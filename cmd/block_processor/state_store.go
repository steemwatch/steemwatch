package main

import (
	"context"
	"io"
)

// StateStore is used to store block processor state.
type StateStore interface {
	// Changes returns a channel that can be used to receive state changes.
	Changes() <-chan *State

	// StoreState stores the given state.
	StoreState(context.Context, *State) error

	// Close() error
	io.Closer
}
