package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	raven "github.com/getsentry/raven-go"
	"github.com/go-steem/rpc"
	"github.com/go-steem/rpc/transports/websocket"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
	"github.com/tchap/zapext/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	tomb "gopkg.in/tomb.v2"
)

// Version contains the service version and is set at build time.
var Version = ""

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %# v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load configuration from the environment.
	config, err := LoadConfig()
	if err != nil {
		return err
	}

	// Set up logging.
	var cores []zapcore.Core

	var logLevel zap.AtomicLevel
	if err := logLevel.UnmarshalText([]byte(config.LogLevel)); err != nil {
		return errors.Wrap(err, "failed to unmarshal log level")
	}

	// Console logging.
	cores = append(cores, zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		os.Stderr,
		logLevel,
	))

	// Sentry logging.
	if dsn := config.SentryDSN; dsn != "" {
		client, err := raven.NewWithTags(dsn, map[string]string{
			"service_name":        config.ServiceName,
			"service_instance_id": config.ServiceInstanceID,
		})
		if err != nil {
			return errors.Wrap(err, "failed to instantiate Sentry client")
		}
		defer client.Wait()

		client.SetRelease(Version)

		cores = append(cores, zapsentry.NewCore(
			logLevel,
			client,
			zapsentry.SetStackTracePackagePrefixes([]string{"github.com/steemwatch"}),
			zapsentry.SetWaitEnabler(zapcore.PanicLevel),
		))
	}

	logger := zap.New(zapcore.NewTee(cores...))
	defer logger.Sync()

	// Start catching signals.
	var t tomb.Tomb
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	t.Go(func() error {
		select {
		case <-sigCh:
			logger.Info("signal received")
			signal.Stop(sigCh)
			t.Kill(nil)
		case <-t.Dying():
		}
		return nil
	})

	// Prepare a STAN connection factory.
	connectToSTAN := func(clientID string) (stan.Conn, error) {
		conn, err := stan.Connect(
			config.STANClusterID,
			clientID,
			stan.NatsURL(config.STANURL),
			stan.ConnectWait(config.STANConnectWait),
			stan.PubAckWait(config.STANPubAckWait),
		)
		if err != nil {
			errMsg := "failed to connect to STAN"

			logger.Error(
				errMsg,
				zap.String("url", config.STANURL),
				zap.String("cluster_id", config.STANClusterID),
				zap.String("client_id", config.STANClientID),
				zap.Error(err),
			)

			return nil, errors.Wrap(err, errMsg)
		}

		logger.Info(
			"connected to STAN",
			zap.String("url", config.STANURL),
			zap.String("cluster_id", config.STANClusterID),
			zap.String("client_id", clientID),
		)
		return conn, nil
	}

	// Set up the state store.
	storeConn, err := connectToSTAN(config.STANClientID + "_state")
	if err != nil {
		return err
	}
	defer storeConn.Close()

	store, err := NewSTANStateStore(logger, storeConn, config.STANStateSubject)
	if err != nil {
		logger.Error("failed to set up the state store", zap.Error(err))
		return err
	}
	defer store.Close()

	// Prepare a steemd RPC client.
	monitorCh := make(chan interface{})
	t.Go(func() error {
		for {
			select {
			case event, ok := <-monitorCh:
				if !ok {
					return nil
				}

				state := fmt.Sprintf("%v", event)
				logger.Info("WebSocket state changed", zap.String("state", state))

			case <-t.Dying():
				return nil
			}
		}
	})

	transport, err := websocket.NewTransport(
		config.SteemdRPCEndpointAddress,
		websocket.SetAutoReconnectEnabled(true),
		websocket.SetAutoReconnectMaxDelay(config.SteemdMaxReconnectDelay),
		websocket.SetDialTimeout(config.SteemdDialTimeout),
		websocket.SetMonitor(monitorCh),
		websocket.SetReadTimeout(config.SteemdReadTimeout),
		websocket.SetWriteTimeout(config.SteemdWriteTimeout),
	)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Steem RPC transport")
	}

	client, err := rpc.NewClient(transport)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Steem RPC client")
	}
	defer client.Close()

	// Prepare a STAN connection.
	conn, err := connectToSTAN(config.STANClientID)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Initialize the block processor and keep processing until interrupted.
	blockProcessor := NewBlockProcessor(logger, store, client, conn, config.STANOutputSubject)
	t.Go(func() error {
		if err := blockProcessor.Start(); err != nil {
			return err
		}

		t.Go(func() error {
			<-t.Dying()
			blockProcessor.Stop(nil)
			return nil
		})

		return blockProcessor.Wait()
	})

	// Wait until all threads are done.
	return t.Wait()
}
