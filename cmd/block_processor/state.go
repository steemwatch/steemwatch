package main

// State represents the block processor state.
type State struct {
	LastProcessedBlockNumber uint32 `json:"last_processed_block_number"`
}
