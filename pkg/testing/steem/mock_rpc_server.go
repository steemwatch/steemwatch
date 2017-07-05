package steem

import (
	"context"
	"encoding/json"
	errs "errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/sourcegraph/jsonrpc2"
	jsonrpc2websocket "github.com/sourcegraph/jsonrpc2/websocket"
	"github.com/tchap/zapext/types"
	"github.com/tchap/zapext/zapsentry"
	"go.uber.org/zap"
)

const (
	CodeUnknownMethod = iota
	CodeDecodingError
	CodeHandlerUndefined
	CodeInvalidArguments
	CodeHandlerError
)

var ErrInvalidArguments = errs.New("invalid arguments")

type handlerKey struct {
	apiID  int
	method string
}

// RPCHandlerFunc is a function that is used to handle incoming RPC requests.
type RPCHandlerFunc func(apiID int, method string, args []*json.RawMessage) (interface{}, error)

// MockRPCServer can be used to mock requests to the steemd RPC endpoint.
type MockRPCServer struct {
	logger   *zap.Logger
	upgrader *websocket.Upgrader

	handlers map[handlerKey]RPCHandlerFunc
}

// NewMockRPCServer returns a new steemd RPC mock server.
func NewMockRPCServer(logger *zap.Logger) *MockRPCServer {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &MockRPCServer{
		logger:   logger,
		upgrader: &upgrader,
		handlers: make(map[handlerKey]RPCHandlerFunc),
	}
}

// HandleRPC can be used to register a handler function
// to be used to handle the given combination of API ID and method name.
func (server *MockRPCServer) HandleRPC(
	apiID int,
	method string,
	handler RPCHandlerFunc,
) {
	server.handlers[handlerKey{apiID, method}] = handler
}

// ServeHTTP serves HTTP requests by upgrading to WebSocket
// and serving steemd RPC endpoint requests as specified.
func (server *MockRPCServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.logger.Error(
			"failed to upgrade to WebSocket",
			zap.Object(zapsentry.HTTPRequestKey, &types.HTTPRequest{r}),
			zap.Error(err),
		)
		return
	}

	ctx := r.Context()
	rpcConn := jsonrpc2.NewConn(ctx, jsonrpc2websocket.NewObjectStream(ws), server)
	defer rpcConn.Close()

	select {
	case <-rpcConn.DisconnectNotify():
	case <-ctx.Done():
	}
}

// Handle implements jsonrpc2.Handler.
func (server *MockRPCServer) Handle(ctx context.Context, conn *jsonrpc2.Conn, r *jsonrpc2.Request) {
	rid := r.ID.String()

	server.logger.Debug(
		"JSON-RPC request received",
		zap.String("request_id", rid),
		zap.String("method", r.Method),
		zap.String("params", string(*r.Params)),
	)

	// Parse the request to get Steem RPC API ID, method name and arguments.
	var (
		apiID  int
		method string
		args   []*json.RawMessage
		err    error
	)
	// In case the request method is "call", we have to dive and process request parameters.
	// Otherwise we can access the request object more or less directly.
	if r.Method == "call" {
		apiID, method, args, err = server.unmarshalRequest(r.Params)
	} else {
		apiID = 0
		method = r.Method
		args, err = server.unmarshalArgs(r.Params)
	}
	if err != nil {
		if err := conn.ReplyWithError(ctx, r.ID, &jsonrpc2.Error{
			Code:    CodeDecodingError,
			Message: "failed to decode Steem RPC request object: " + err.Error(),
		}); err != nil {
			server.logger.Error("failed to reply",
				zap.String("request_id", rid), zap.Error(err))
		}
		return
	}

	server.logger.Debug("request parsed",
		zap.Int("api_id", apiID), zap.String("method", method), zap.Reflect("args", args))

	// Get the handle for the (API ID, method name) combination.
	handler, ok := server.handlers[handlerKey{apiID, method}]
	if !ok {
		server.logger.Warn("handler not found",
			zap.Int("api_id", apiID), zap.String("method", method))

		if err := conn.ReplyWithError(ctx, r.ID, &jsonrpc2.Error{
			Code:    CodeHandlerUndefined,
			Message: fmt.Sprintf("mock handler undefined for (API ID = %v, method = %v)", apiID, method),
		}); err != nil {
			server.logger.Error("failed to reply",
				zap.String("request_id", rid), zap.Error(err))
		}
		return
	}

	// Generate the response and send it back.
	response, err := handler(apiID, method, args)

	server.logger.Debug("response generated",
		zap.String("request_id", rid), zap.Reflect("response", response), zap.Error(err))

	if err != nil {
		var code int
		if err == ErrInvalidArguments {
			code = CodeInvalidArguments
		} else {
			code = CodeHandlerError
		}

		if err := conn.ReplyWithError(ctx, r.ID, &jsonrpc2.Error{
			Code:    int64(code),
			Message: "handler returned an error: " + err.Error(),
		}); err != nil {
			server.logger.Error("failed to reply",
				zap.String("request_id", rid), zap.Error(err))
		}
		return
	}

	if err := conn.Reply(ctx, r.ID, response); err != nil {
		server.logger.Error("failed to reply",
			zap.String("request_id", rid), zap.Error(err))
	}
}

func (server *MockRPCServer) unmarshalRequest(
	params *json.RawMessage,
) (int, string, []*json.RawMessage, error) {

	// Unmarshal *json.RawMessage into []*json.RawMessage,
	// i.e. turn params into a raw parameter list.
	var ps []*json.RawMessage
	if err := json.Unmarshal([]byte(*params), &ps); err != nil {
		return 0, "", nil, errors.Wrap(err, "failed to unmarshal request params")
	}

	if n := len(ps); n != 3 {
		return 0, "", nil, errors.Errorf("unexpected param list length: %v", n)
	}

	// Unmarshal the API ID.
	var apiID int
	if err := json.Unmarshal([]byte(*ps[0]), &apiID); err != nil {
		return 0, "", nil, errors.Wrap(err, "failed to unmarshal API ID")
	}

	// Unmarshal the method name.
	var method string
	if err := json.Unmarshal([]byte(*ps[1]), &method); err != nil {
		return 0, "", nil, errors.Wrap(err, "failed to unmarshal method name")
	}

	// Unmarshal the arguments.
	args, err := server.unmarshalArgs(ps[2])
	if err != nil {
		return 0, "", nil, err
	}

	// And we are done!
	return apiID, method, args, nil
}

func (server *MockRPCServer) unmarshalArgs(raw *json.RawMessage) ([]*json.RawMessage, error) {
	var args []*json.RawMessage
	if err := json.Unmarshal([]byte(*raw), &args); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal Steem RPC request arguments")
	}
	return args, nil
}
