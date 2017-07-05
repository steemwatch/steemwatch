# Block Processor

Block Processor, or `block_processor`, takes care of fetching block operations
from the blockchain and passing them on into NATS Streaming, encoded
using Protocol Buffers.

When run for the first time, `block_processor` starts with the last block on
the blockchain. Then it periodically stores the last processed block number
in a dedicated NATS Streaming channel. When started again later, it loads
the last processed block number from the channel and keeps processing blocks
where it left off previously.