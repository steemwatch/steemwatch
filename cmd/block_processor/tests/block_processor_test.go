package blockprocessor_test

import (
	"os"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/steemwatch/steemwatch/cmd/block_processor/tests/data"
	"github.com/steemwatch/steemwatch/pkg/pb/steempb"

	"github.com/gogo/protobuf/proto"
	"github.com/nats-io/go-nats-streaming"
)

func TestBlockProcessor(t *testing.T) {
	// Start processing output messages.
	var (
		stanURL   = os.Getenv("STAN_URL")
		clusterID = os.Getenv("STAN_CLUSTER_ID")
		clientID  = os.Getenv("STAN_CLIENT_ID")
		subject   = os.Getenv("STAN_OUTPUT_SUBJECT")
	)

	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(stanURL))
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()

	msgCh := make(chan *stan.Msg, 1)
	if _, err := conn.Subscribe(subject, func(msg *stan.Msg) {
		msgCh <- msg
	}, stan.DeliverAllAvailable()); err != nil {
		t.Error(err)
		return
	}

	msg := <-msgCh
	var blockOps steempb.BlockOperations
	if err := proto.Unmarshal(msg.Data, &blockOps); err != nil {
		t.Error(err)
		return
	}

	// Assemble the expected block operations object.
	var expected steempb.BlockOperations
	expected.Operations = make([]*steempb.OperationObject, 0, len(data.Operations))
	for _, op := range data.Operations {
		expected.Operations = append(expected.Operations, op.OperationObject())
	}

	// Compare.
	if diff := pretty.Diff(expected, blockOps); len(diff) != 0 {
		t.Errorf("unexpected message payload received: \n%v", strings.Join(diff, "\n"))
	}
}
