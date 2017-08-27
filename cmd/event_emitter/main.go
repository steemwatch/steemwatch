package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	raven "github.com/getsentry/raven-go"
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

	logger := zap.New(zapcore.NewTee(cores...)).Named(config.ServiceName)
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

	// Start the service.
	ee := NewEventEmitter(logger, config)
	t.Go(func() error {
		if err := ee.Start(); err != nil {
			return err
		}

		t.Go(func() error {
			<-t.Dying()
			ee.Stop(nil)
			return nil
		})

		return ee.Wait()
	})

	// Wait until all threads are done.
	return t.Wait()
}
