package data

import (
	"fmt"

	"github.com/steemwatch/steemwatch/pkg/pb/steempb"

	"github.com/golang/protobuf/ptypes/timestamp"
)

const OperationObjectFormat = `
{
	"block": 1,
	"trx_id": "6af0811250a1a28ad12e67417edaf6ee8c849c27",
    "trx_in_block": 1,
    "op": ["%v", %v],
    "op_in_trx": 1,
	"virtual_op": 1,
	"timestamp": "2017-07-25T20:00:00"
}`

type Operation struct {
	Name string
	Body string
	Op   *steempb.Operation
}

func (op *Operation) FormatOperationObject() string {
	return fmt.Sprintf(OperationObjectFormat, op.Name, op.Body)
}

func (op *Operation) OperationObject() *steempb.OperationObject {
	return &steempb.OperationObject{
		BlockNumber:            1,
		TransactionID:          "6af0811250a1a28ad12e67417edaf6ee8c849c27",
		TransactionInBlock:     1,
		Operation:              op.Op,
		OperationInTransaction: 1,
		VirtualOperation:       1,
		Timestamp:              &timestamp.Timestamp{Seconds: 1501012800},
	}
}

var Operations = []Operation{
	// account_update
	{
		"account_update",
		`
		{
       		"account": "somebody",
        	"json_metadata": "{\"profile\":{}}",
			"memo_key": "STM8K77gmLQ6tauRHusxr7E5mpzqvGArsy2vmSijHpiPy4EY2B9hX",
			"owner": {
          		"account_auths": [
					["streemian", 1]
				],
          		"key_auths": [
            		["STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7", 1]
          		],
          		"weight_threshold": 1
			},
			"active": {
          		"account_auths": [
					["streemian", 1]
				],
          		"key_auths": [
            		["STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7", 1]
          		],
          		"weight_threshold": 1
        	},
        	"posting": {
          		"account_auths": [
					["streemian", 1]
				],
          		"key_auths": [
            		["STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7", 1]
          		],
          		"weight_threshold": 1
			}
      	}
		`,
		&steempb.Operation{Operation: &steempb.Operation_AccountUpdate{
			AccountUpdate: &steempb.AccountUpdateOperation{
				Account: "somebody",
				Owner: &steempb.Authority{
					WeightThreshold: 1,
					AccountAuths: map[string]int64{
						"streemian": 1,
					},
					KeyAuths: map[string]int64{
						"STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7": 1,
					},
				},
				Active: &steempb.Authority{
					WeightThreshold: 1,
					AccountAuths: map[string]int64{
						"streemian": 1,
					},
					KeyAuths: map[string]int64{
						"STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7": 1,
					},
				},
				Posting: &steempb.Authority{
					WeightThreshold: 1,
					AccountAuths: map[string]int64{
						"streemian": 1,
					},
					KeyAuths: map[string]int64{
						"STM6PNBPPVzTJLcoFQudzgRP7MsWDuKSomek8TbPHsPNzMaxfxaL7": 1,
					},
				},
				MemoKey:      "STM8K77gmLQ6tauRHusxr7E5mpzqvGArsy2vmSijHpiPy4EY2B9hX",
				JSONMetadata: "{\"profile\":{}}",
			},
		}},
	},
	// account_witness_vote
	{
		"account_witness_vote",
		`
		{
    		"account": "somebody",
        	"approve": true,
        	"witness": "netuoso"
      	}
		`,
		&steempb.Operation{Operation: &steempb.Operation_AccountWitnessVote{
			AccountWitnessVote: &steempb.AccountWitnessVoteOperation{
				Account: "somebody",
				Approve: true,
				Witness: "netuoso",
			},
		}},
	},
	// comment
	{
		"comment",
		`
		{
        	"author": "somebody",
        	"body": "Look at the sky. We are not alone.",
        	"json_metadata": "{\"app\":\"chainbb/0.3\",\"format\":\"markdown+html\",\"tags\":[\"chainbb\"]}",
        	"parent_author": "someparentauthor",
        	"parent_permlink": "chainbb",
        	"permlink": "thoughts",
        	"title": "Thoughts"
      	}
		`,
		&steempb.Operation{Operation: &steempb.Operation_Comment{
			Comment: &steempb.CommentOperation{
				ParentAuthor:   "someparentauthor",
				ParentPermlink: "chainbb",
				Author:         "somebody",
				Permlink:       "thoughts",
				Title:          "Thoughts",
				Body:           "Look at the sky. We are not alone.",
				JSONMetadata:   "{\"app\":\"chainbb/0.3\",\"format\":\"markdown+html\",\"tags\":[\"chainbb\"]}",
			},
		}},
	},
	// vote
	{
		"vote",
		`
		{
        	"author": "thegodbox",
        	"permlink": "ready-for-hot-dogs-and-tacos-tacotuesday",
        	"voter": "rogeer",
        	"weight": 100
		}
		`,
		&steempb.Operation{Operation: &steempb.Operation_Vote{
			Vote: &steempb.VoteOperation{
				Author:   "thegodbox",
				Permlink: "ready-for-hot-dogs-and-tacos-tacotuesday",
				Voter:    "rogeer",
				Weight:   100,
			},
		}},
	},
	// custom_json
	{
		"custom_json",
		`
		{
        	"id": "follow",
        	"json": "[\"follow\",{\"follower\":\"mardah.resonance\",\"following\":\"organik\",\"what\":[]}]",
        	"required_auths": [],
        	"required_posting_auths": [
          		"mardah.resonance"
			]
		}
		`,
		&steempb.Operation{Operation: &steempb.Operation_CustomJSON{
			CustomJSON: &steempb.CustomJSONOperation{
				RequiredAuths:        []string{},
				RequiredPostingAuths: []string{"mardah.resonance"},
				ID:                   "follow",
				JSON:                 "[\"follow\",{\"follower\":\"mardah.resonance\",\"following\":\"organik\",\"what\":[]}]",
			},
		}},
	},
	// transfer
	{
		"transfer",
		`
		{
        	"amount": "0.001 SBD",
        	"from": "ruse",
        	"memo": "somememo",
        	"to": "minnowsupport"
      	}
		`,
		&steempb.Operation{Operation: &steempb.Operation_Transfer{
			Transfer: &steempb.TransferOperation{
				Amount: "0.001 SBD",
				From:   "ruse",
				To:     "minnowsupport",
				Memo:   "somememo",
			},
		}},
	},
}
