package main

import (
	"github.com/steemwatch/steemwatch/pkg/pb/steempb"

	"github.com/go-steem/rpc/types"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func convert(op *types.OperationObject) *steempb.OperationObject {
	return &steempb.OperationObject{
		BlockNumber:            op.BlockNumber,
		TransactionID:          op.TransactionID,
		TransactionInBlock:     op.TransactionInBlock,
		Operation:              convertOperation(op.Operation),
		OperationInTransaction: uint32(op.OperationInTransaction),
		VirtualOperation:       op.VirtualOperation,
		Timestamp: &timestamp.Timestamp{
			Seconds: op.Timestamp.Unix(),
			Nanos:   int32(op.Timestamp.Nanosecond()),
		},
	}
}

func convertOperation(op types.Operation) *steempb.Operation {
	switch op := op.Data().(type) {
	case *types.AccountUpdateOperation:
		var (
			owner   *steempb.Authority
			active  *steempb.Authority
			posting *steempb.Authority
		)
		if op.Owner != nil {
			owner = &steempb.Authority{
				WeightThreshold: op.Owner.WeightThreshold,
				AccountAuths:    op.Owner.AccountAuths,
				KeyAuths:        op.Owner.KeyAuths,
			}
		}
		if op.Active != nil {
			active = &steempb.Authority{
				WeightThreshold: op.Active.WeightThreshold,
				AccountAuths:    op.Active.AccountAuths,
				KeyAuths:        op.Active.KeyAuths,
			}
		}
		if op.Posting != nil {
			posting = &steempb.Authority{
				WeightThreshold: op.Posting.WeightThreshold,
				AccountAuths:    op.Posting.AccountAuths,
				KeyAuths:        op.Posting.KeyAuths,
			}
		}

		return &steempb.Operation{Operation: &steempb.Operation_AccountUpdate{
			AccountUpdate: &steempb.AccountUpdateOperation{
				Account:      op.Account,
				Owner:        owner,
				Active:       active,
				Posting:      posting,
				MemoKey:      op.MemoKey,
				JSONMetadata: op.JsonMetadata,
			},
		}}

	case *types.AccountWitnessVoteOperation:
		return &steempb.Operation{Operation: &steempb.Operation_AccountWitnessVote{
			AccountWitnessVote: (*steempb.AccountWitnessVoteOperation)(op),
		}}

	case *types.CommentOperation:
		return &steempb.Operation{Operation: &steempb.Operation_Comment{
			Comment: &steempb.CommentOperation{
				ParentAuthor:   op.ParentAuthor,
				ParentPermlink: op.ParentPermlink,
				Author:         op.Author,
				Permlink:       op.Permlink,
				Title:          op.Title,
				Body:           op.Body,
				JSONMetadata:   op.JsonMetadata,
			},
		}}

	case *types.VoteOperation:
		return &steempb.Operation{Operation: &steempb.Operation_Vote{
			Vote: &steempb.VoteOperation{
				Author:   op.Author,
				Permlink: op.Permlink,
				Voter:    op.Voter,
				Weight:   int32(op.Weight),
			},
		}}

	case *types.CustomJSONOperation:
		return &steempb.Operation{Operation: &steempb.Operation_CustomJSON{
			CustomJSON: (*steempb.CustomJSONOperation)(op),
		}}

	case *types.TransferOperation:
		return &steempb.Operation{Operation: &steempb.Operation_Transfer{
			Transfer: (*steempb.TransferOperation)(op),
		}}
	}

	return nil
}
