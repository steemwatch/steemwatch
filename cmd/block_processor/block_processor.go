package main

import (
	"context"
	"time"

	"github.com/go-steem/rpc"
	"github.com/go-steem/rpc/types"
	"github.com/gogo/protobuf/proto"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/pkg/errors"
	"github.com/steemwatch/steemwatch/pkg/pb/steempb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	tomb "gopkg.in/tomb.v2"
)

// StateLoadTimeout is the timeout to wait for to get the state loaded.
// When the deadline is exceeded, the state is reset.
const StateLoadTimeout = 5 * time.Second

// BlockProcessor implements the block processing logic.
type BlockProcessor struct {
	// Logger
	logger *zap.Logger

	// State management
	store StateStore
	state *State

	// Input - steemd RPC client
	client *rpc.Client

	// Output - STAN
	sc      stan.Conn
	subject string

	t *tomb.Tomb
}

// NewBlockProcessor is the block processor constructor.
func NewBlockProcessor(
	logger *zap.Logger,
	store StateStore,
	client *rpc.Client,
	sc stan.Conn,
	subject string,
) *BlockProcessor {
	return &BlockProcessor{
		logger:  logger,
		store:   store,
		client:  client,
		sc:      sc,
		subject: subject,
		t:       &tomb.Tomb{},
	}
}

// Start starts the block processor in the background, i.e. it is non-blocking.
func (bp *BlockProcessor) Start() error {
	inputCh := make(chan []*types.OperationObject)
	outputCh := make(chan []*steempb.OperationObject)
	commitCh := make(chan uint32)

	bp.t.Go(func() error {
		return bp.fetcher(bp.logger.Named("fetcher"), inputCh)
	})

	bp.t.Go(func() error {
		return bp.convertor(bp.logger.Named("convertor"), inputCh, outputCh)
	})

	bp.t.Go(func() error {
		return bp.publisher(bp.logger.Named("publisher"), outputCh, commitCh)
	})

	bp.t.Go(func() error {
		return bp.committed(bp.logger.Named("committer"), commitCh)
	})

	bp.logger.Info("started")
	return nil
}

// Stop terminates the block processor.
// The shutdown process can be interrupted by the context being closed.
func (bp *BlockProcessor) Stop(ctx context.Context) error {
	bp.t.Kill(nil)
	return nil
}

// Wait blocks until block processor is terminated.
func (bp *BlockProcessor) Wait() error {
	err := bp.t.Wait()
	bp.logger.Info("terminated")
	return err
}

// fetcher is a background goroutine that handles fetching of Steem blocks
// and committing block numbers that has already been processed.
func (bp *BlockProcessor) fetcher(
	logger *zap.Logger,
	outputCh chan<- []*types.OperationObject,
) error {
	logger.Info("started")
	defer logger.Info("terminated")

	defer close(outputCh)

	// Get the blockchain config.
	config, err := bp.client.Database.GetConfig()
	if err != nil {
		return bp.nuked(logger, err, "failed to get blockchain configuration")
	}

	// Keep fetching blocks until interrupted.
	//
	// Tick every STEEMIT_BLOCK_INTERVAL seconds.
	blockTicker := time.NewTicker(time.Duration(config.SteemitBlockInterval) * time.Second)
	defer blockTicker.Stop()
	var blockTickCh <-chan time.Time

	// State is set once received over the store changes channel.
	var nextBlockNumber uint32
	stateLoadTimeoutCh := time.After(StateLoadTimeout)

	for {
		select {
		// Receiving on the state change channel enables the processing loop.
		case state := <-bp.store.Changes():
			bp.logger.Info("state loaded", zap.Reflect("state", state))
			nextBlockNumber = state.LastProcessedBlockNumber + 1
			blockTickCh = blockTicker.C

		case <-stateLoadTimeoutCh:
			bp.logger.Info("state not loaded, starting with the last block")
			blockTickCh = blockTicker.C
			stateLoadTimeoutCh = nil

		// This is the main block processing case.
		case <-blockTickCh:
			// Get current properties.
			props, err := bp.client.Database.GetDynamicGlobalProperties()
			if err != nil {
				return bp.nuked(logger, err, "failed to get dynamic global properties")
			}

			// Process new blocks.
			if nextBlockNumber == 0 {
				nextBlockNumber = props.LastIrreversibleBlockNum
			}

			for props.LastIrreversibleBlockNum >= nextBlockNumber {
				// Fetch the block operations from the database.
				if ce := logger.Check(zapcore.DebugLevel, "fetching block ..."); ce != nil {
					ce.Write(zap.Uint32("block_number", nextBlockNumber))
				}

				operations, err := bp.client.Database.GetOpsInBlock(nextBlockNumber, false)
				if err != nil {
					return bp.nuked(
						logger, err, "failed to fetch block operations",
						zap.Uint32("block_number", nextBlockNumber),
					)
				}

				// Pass the block to the next component.
				select {
				case outputCh <- operations:
				case <-bp.t.Dying():
					return nil
				}

				// Block fetched, increment the next block number.
				nextBlockNumber++
			}

		case <-bp.t.Dying():
			return nil
		}
	}
}

func (bp *BlockProcessor) convertor(
	logger *zap.Logger,
	inputCh <-chan []*types.OperationObject,
	outputCh chan<- []*steempb.OperationObject,
) error {
	logger.Info("started")
	defer logger.Info("terminated")

	defer close(outputCh)

	for inputOps := range inputCh {
		outputOps := make([]*steempb.OperationObject, 0, len(inputOps))
		for _, op := range inputOps {
			if out := convert(op); out != nil {
				outputOps = append(outputOps, out)
			}
		}

		outputCh <- outputOps
	}

	return nil
}

func (bp *BlockProcessor) publisher(
	logger *zap.Logger,
	inputCh <-chan []*steempb.OperationObject,
	commitCh chan<- uint32,
) error {
	logger.Info("started")
	defer logger.Info("terminated")

	defer close(commitCh)

	for ops := range inputCh {
		// No operations harvested, send 0 to the committer.
		// That signals to increment the last block number and commit.
		if len(ops) == 0 {
			commitCh <- 0
			continue
		}

		out := steempb.BlockOperations{Operations: ops}
		data, err := proto.Marshal(&out)
		if err != nil {
			return bp.nuked(logger, err, "failed to marshal output operations")
		}

		if err := bp.sc.Publish(bp.subject, data); err != nil {
			return bp.nuked(logger, err, "failed to publish output message")
		}

		blockNumber := ops[0].BlockNumber
		if ce := logger.Check(zapcore.DebugLevel, "operations published"); ce != nil {
			ce.Write(
				zap.Uint32("block_number", blockNumber),
				zap.Int("operation_count", len(ops)),
			)
		}

		commitCh <- blockNumber
	}

	return nil
}

func (bp *BlockProcessor) committed(
	logger *zap.Logger,
	commitCh <-chan uint32,
) error {
	logger.Info("started")
	defer logger.Info("terminated")

	var state State
	ctx := bp.t.Context(nil)

	for blockNumber := range commitCh {
		// Store the new next block number.
		//
		// Receiving 0 means that we have to increment the previous block number.
		// In case there is no previous block number encountered, we do nothing.
		//
		// When the block number is non-zero, we simply store it.
		if blockNumber == 0 {
			if state.LastProcessedBlockNumber == 0 {
				continue
			}
			state.LastProcessedBlockNumber++
		} else {
			state.LastProcessedBlockNumber = blockNumber
		}

		if err := bp.store.StoreState(ctx, &state); err != nil {
			return bp.nuked(logger, err, "failed to store state")
		}

		if ce := logger.Check(zapcore.DebugLevel, "state stored"); ce != nil {
			ce.Write(zap.Reflect("state", &state))
		}
	}

	return nil
}

func (bp *BlockProcessor) nuked(
	logger *zap.Logger,
	err error,
	msg string,
	fields ...zapcore.Field,
) error {
	if err == nil {
		return nil
	}

	logger.Error(msg, append(fields, zap.Error(err))...)
	return errors.Wrap(err, msg)
}
