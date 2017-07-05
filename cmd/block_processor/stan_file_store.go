package main

import (
	"context"
	"encoding/json"

	stan "github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type stanStateStore struct {
	logger *zap.Logger

	sc      stan.Conn
	subject string
	sub     stan.Subscription

	changes chan *State
	closed  chan struct{}
}

// NewSTANStateStore returns a state store that stores state in a STAN channel.
func NewSTANStateStore(logger *zap.Logger, sc stan.Conn, subject string) (StateStore, error) {
	s := &stanStateStore{
		logger:  logger,
		sc:      sc,
		subject: subject,
		changes: make(chan *State, 1),
		closed:  make(chan struct{}),
	}

	var err error
	s.sub, err = sc.Subscribe(subject, func(msg *stan.Msg) {
		// Decode message data.
		var state State
		if err := json.Unmarshal(msg.Data, &state); err != nil {
			s.logError(
				err, "state store: failed to decode message data",
				zap.String("data", string(msg.Data)),
			)
			return
		}

		// Pass the received state on.
		s.changes <- &state

		// Unsubscribe.
		s.ensureNotSubscribed()
	}, stan.StartWithLastReceived())
	if err != nil {
		return nil, s.logError(err, "state store: failed to subscribe")
	}

	return s, nil
}

func (s *stanStateStore) Changes() <-chan *State {
	return s.changes
}

func (s *stanStateStore) StoreState(ctx context.Context, state *State) error {
	// Make sure we are not subscribed.
	s.ensureNotSubscribed()

	// Encode the state object.
	data, err := json.Marshal(state)
	if err != nil {
		return s.logError(err, "state store: failed to encode state")
	}

	// Publish it to the channel.
	return s.logError(s.sc.Publish(s.subject, data), "state store: failed to publish state")
}

func (s *stanStateStore) Close() error {
	return s.logError(s.sc.Close(), "state store: failed to close")
}

func (s *stanStateStore) logError(err error, msg string, fields ...zapcore.Field) error {
	if err == nil {
		return nil
	}

	s.logger.Error(msg, append(fields, zap.Error(err))...)
	return errors.Wrap(err, msg)
}

func (s *stanStateStore) ensureNotSubscribed() {
	if s.sub.IsValid() {
		if err := s.sub.Unsubscribe(); err != nil {
			s.logger.Error("failed to unsubscribe", zap.Error(err))
		}
	}
}
