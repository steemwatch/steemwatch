package main

import (
	"context"

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

func (ee *EventEmitter) Start() error {
	sc, err := stan.Connect(
		ee.config.STANClusterID,
		ee.config.STANClientID,
		stan.NatsURL(ee.config.STANURL),
		stan.ConnectWait(ee.config.STANConnectWait),
		stan.PubAckWait(ee.config.STANPubAckWait),
	)
	if err != nil {
		return ee.nuked(
			err, "failed to connect to STAN",
			zap.String("url", ee.config.STANURL),
			zap.String("cluster_id", ee.config.STANClusterID),
			zap.String("client_id", ee.config.STANClientID),
			zap.Error(err),
		)
	}

	ee.logger.Info(
		"connected to STAN",
		zap.String("url", ee.config.STANURL),
		zap.String("cluster_id", ee.config.STANClusterID),
		zap.String("client_id", ee.config.STANClientID),
	)

	if _, err := sc.QueueSubscribe(
		ee.config.STANInputSubject,
		ee.config.ServiceName,
		ee.handleMessage,
		stan.DurableName(ee.config.ServiceName),
	); err != nil {
		return ee.nuked(
			err, "failed to subscribe to the input subject",
			zap.String("subject", ee.config.STANInputSubject),
		)
	}

	ee.sc = sc
	return nil
}

func (ee *EventEmitter) Stop(ctx context.Context) error {
	ee.t.Kill(nil)
	return nil
}

func (ee *EventEmitter) Wait() error {
	err := ee.t.Wait()
	ee.logger.Info("terminated")
	return err
}

func (ee *EventEmitter) loop() error {
	for {
	}
}

func (ee *EventEmitter) nuked(
	err error,
	msg string,
	fields ...zapcore.Field,
) error {
	if err == nil {
		return nil
	}

	ee.logger.Error(msg, append(fields, zap.Error(err))...)
	return errors.Wrap(err, msg)
}
