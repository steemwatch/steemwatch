package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/steemwatch/steemwatch/cmd/block_processor/tests/data"
	"github.com/steemwatch/steemwatch/pkg/testing/steem"

	"github.com/go-steem/rpc/apis/database"
	"github.com/go-steem/rpc/apis/login"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	tomb "gopkg.in/tomb.v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %# v\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	// Set up the mock RPC endpoint.
	addr := os.Getenv("STEEMD_RPC_ENDPOINT_ADDRESS")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return errors.Wrap(err, "failed to start listening on the given address")
	}

	// Close the listener on signal or error.
	var t tomb.Tomb
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	t.Go(func() error {
		select {
		case <-sigCh:
			signal.Stop(sigCh)
		case <-t.Dying():
		}
		return listener.Close()
	})

	t.Go(func() error {
		server := steem.NewMockRPCServer(logger)
		registerHandlers(server)

		err := http.Serve(listener, server)
		if err != http.ErrServerClosed {
			return errors.Wrap(err, "failed to start serving requests")
		}
		return nil
	})

	return t.Wait()
}

func registerHandlers(server *steem.MockRPCServer) {
	//
	// database_api
	//

	db := database.NumbericAPIID

	server.HandleRPC(db, "get_config",
		func(apiID int, method string, args []*json.RawMessage) (interface{}, error) {
			return map[string]interface{}{
				"STEEMIT_BLOCK_INTERVAL": 1,
			}, nil
		})

	server.HandleRPC(db, "get_dynamic_global_properties",
		func(apiID int, method string, args []*json.RawMessage) (interface{}, error) {
			return map[string]interface{}{
				"last_irreversible_block_num": 1,
			}, nil
		})

	server.HandleRPC(db, "get_ops_in_block",
		func(apiID int, method string, args []*json.RawMessage) (interface{}, error) {
			if len(args) != 2 {
				return nil, steem.ErrInvalidArguments
			}

			// Block 0 is the only allowed block number.
			if string(*args[0]) != "1" {
				return nil, errors.New("invalid block number")
			}

			// Construct the operations.
			// Not too efficient, but whatever.
			ops := make([]string, 0, len(data.Operations))
			for _, op := range data.Operations {
				ops = append(ops, op.FormatOperationObject())
			}

			data := []byte("[" + strings.Join(ops, ",") + "]")
			return json.RawMessage(data), nil
		})

	//
	// login_api
	//

	login := login.NumbericAPIID

	server.HandleRPC(login, "get_api_by_name",
		func(apiID int, method string, args []*json.RawMessage) (interface{}, error) {
			if len(args) != 1 {
				return nil, steem.ErrInvalidArguments
			}

			var apiName string
			if err := json.Unmarshal([]byte(*args[0]), &apiName); err != nil {
				return nil, errors.Wrap(err, "failed to decode args[0]")
			}

			switch apiName {
			case "database_api":
				return 0, nil
			case "login_api":
				return 1, nil
			case "follow_api":
				return 2, nil
			case "network_broadcast_api":
				return 3, nil
			default:
				return 0, errors.Errorf("unknown api name: %v", apiName)
			}
		})
}
