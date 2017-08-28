package main

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/steemwatch/steemwatch/pkg/pb/steempb"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	tomb "gopkg.in/tomb.v2"
)

// EventEmitter implements event_emitter service.
type EventEmitter struct {
	// Logger
	logger *zap.Logger

	// Config
	config *Config

	// Management
	t *tomb.Tomb
}

// NewEventEmitter returns a new event_emitter service instance.
func NewEventEmitter(
	logger *zap.Logger,
	config *Config,
) *EventEmitter {
	return &EventEmitter{
		logger: logger,
		config: config,
		t:      &tomb.Tomb{},
	}
}

func (srv *EventEmitter) Start() error {
	srv.t.Go(func() error {
		srv.logger.Info("started")
		defer srv.logger.Info("terminated")

		sc, err := stan.Connect(
			srv.config.STANClusterID,
			srv.config.STANClientID,
			stan.NatsURL(srv.config.STANURL),
			stan.ConnectWait(srv.config.STANConnectWait),
			stan.PubAckWait(srv.config.STANPubAckWait),
		)
		if err != nil {
			return srv.nuked(
				err, "failed to connect to STAN",
				zap.String("url", srv.config.STANURL),
				zap.String("cluster_id", srv.config.STANClusterID),
				zap.String("client_id", srv.config.STANClientID),
				zap.Error(err),
			)
		}

		srv.logger.Info(
			"connected to STAN",
			zap.String("url", srv.config.STANURL),
			zap.String("cluster_id", srv.config.STANClusterID),
			zap.String("client_id", srv.config.STANClientID),
		)

		if _, err := sc.QueueSubscribe(
			srv.config.STANInputSubject,
			srv.config.ServiceName,
			srv.handleMessage,
			stan.DurableName(srv.config.ServiceName),
		); err != nil {
			return srv.nuked(
				err, "failed to subscribe to the input subject",
				zap.String("subject", srv.config.STANInputSubject),
			)
		}

		// Wait for the termination signal.
		<-srv.t.Dying()

		if err := sc.Close(); err != nil {
			srv.nuked(err, "failed to close STAN connection")
		}
		return nil
	})
	return nil
}

func (srv *EventEmitter) Stop(ctx context.Context) error {
	srv.t.Kill(nil)
	return nil
}

func (srv *EventEmitter) Wait() error {
	return srv.t.Wait()
}

func (srv *EventEmitter) handleMessage(msg *stan.Msg) {
	var ops steempb.BlockOperations
	if err := proto.Unmarshal(msg.Data, &ops); err != nil {
		srv.t.Kill(srv.nuked(err, "failed to unmarshal message data"))
		return
	}

	for _, op := range ops.GetOperations() {
		data, err := proto.Marshal(harvest(op.GetOperation()))
		if err != nil {
			srv.t.Kill(srv.nuked(err, "failed to marshal event object"))
		}
	}

	if err := msg.Ack(); err != nil {
		srv.t.Kill(srv.nuked(err, "failed to ACK input message"))
	}
}

func (srv *EventEmitter) nuked(
	err error,
	msg string,
	fields ...zapcore.Field,
) error {
	if err == nil {
		return nil
	}

	srv.logger.Error(msg, append(fields, zap.Error(err))...)
	return errors.Wrap(err, msg)
}
