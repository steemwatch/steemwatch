package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-steem/rpc"
	"github.com/go-steem/rpc/transports/websocket"
	"github.com/kr/pretty"
	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:  "watch_operations",
		Usage: "watch block appearing on the blockchain",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "rpc-endpoint",
				Value: "wss://gtg.steem.house:8090",
				Usage: "steemd RPC endpoint to connect to",
			},
			&cli.IntFlag{
				Name:  "block-num-from",
				Value: -1,
				Usage: "block number to start with (-1 = the latest)",
			},
			&cli.BoolFlag{
				Name:  "raw",
				Value: false,
				Usage: "print raw RPC responses",
			},
			&cli.BoolFlag{
				Name:  "monitor",
				Value: false,
				Usage: "monitor the WebSocket connection",
			},
		},
		Action: run,
	}

	app.Run(os.Args)
}

func run(ctx *cli.Context) error {
	var (
		endpoint     = ctx.String("rpc-endpoint")
		blockNumFrom = ctx.Int("block-num-from")
		raw          = ctx.Bool("raw")
	)

	var opts []websocket.Option
	if ctx.Bool("monitor") {
		monitorCh := make(chan interface{}, 1)
		go func() {
			for e := range monitorCh {
				fmt.Fprintf(os.Stderr, "WebSocket: %v\n", e)
			}
		}()

		opts = append(opts, websocket.SetMonitor(monitorCh))
	}

	transport, err := websocket.NewTransport([]string{endpoint}, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Steem RPC transport")
	}

	client, err := rpc.NewClient(transport)
	if err != nil {
		return errors.Wrap(err, "failed to initialize Steem RPC client")
	}
	defer client.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	dyingCh := make(chan struct{})
	go func() {
		<-sigCh
		client.Close()
		close(dyingCh)
		signal.Stop(sigCh)
	}()

	// Get config.
	config, err := client.Database.GetConfig()
	if err != nil {
		return errors.Wrap(err, "failed to get steemd database config")
	}

	// Get the next block number.
	props, err := client.Database.GetDynamicGlobalProperties()
	if err != nil {
		return errors.Wrap(err, "failed to get dynamic global properties")
	}

	var next uint32
	if blockNumFrom == -1 {
		next = props.LastIrreversibleBlockNum + 1
	} else {
		next = uint32(blockNumFrom)
	}

	// Keep processing blocks until interrupted.
	for {
		// Get current properties.
		props, err := client.Database.GetDynamicGlobalProperties()
		if err != nil {
			return errors.Wrap(err, "failed to get dynamic global properties")
		}

		// Process new blocks.
		for props.LastIrreversibleBlockNum >= next {
			// Fetch the block from the database.
			if raw {
				ops, err := client.Database.GetOpsInBlockRaw(next, false)
				if err != nil {
					return errors.Wrap(err, "failed to get block")
				}

				var v []interface{}
				if err := json.Unmarshal([]byte(*ops), &v); err != nil {
					return errors.Wrap(err, "failed to unmarshal a block")
				}

				blockIndent, err := json.MarshalIndent(v, "", "  ")
				if err != nil {
					return errors.Wrap(err, "failed to format raw block")
				}

				fmt.Println()
				fmt.Println("==============================")
				fmt.Println("BLOCK", next)
				fmt.Println("==============================")
				fmt.Println(string(blockIndent))
			} else {
				ops, err := client.Database.GetOpsInBlock(next, false)
				if err != nil {
					return errors.Wrap(err, "failed to get block operations")
				}

				for _, op := range ops {
					fmt.Printf(
						"[BLOCK %v] [OP %v] %# v\n",
						op.BlockNumber, op.Operation.Type(), pretty.Formatter(op.Operation.Data()))
				}
			}

			next++
		}

		// Wait for STEEMIT_BLOCK_INTERVAL seconds before the next iteration.
		// In case a signal is received, exit immediately.
		select {
		case <-time.After(time.Duration(config.SteemitBlockInterval) * time.Second):
		case <-dyingCh:
			return nil
		}
	}
}
