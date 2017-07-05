// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account_witness_vote_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AccountWitnessVoteOperation struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Witness string `protobuf:"bytes,2,opt,name=witness,proto3" json:"witness,omitempty"`
	Approve bool   `protobuf:"varint,3,opt,name=approve,proto3" json:"approve,omitempty"`
}

func (m *AccountWitnessVoteOperation) Reset()         { *m = AccountWitnessVoteOperation{} }
func (m *AccountWitnessVoteOperation) String() string { return proto.CompactTextString(m) }
func (*AccountWitnessVoteOperation) ProtoMessage()    {}
func (*AccountWitnessVoteOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorAccountWitnessVoteOperation, []int{0}
}

func (m *AccountWitnessVoteOperation) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountWitnessVoteOperation) GetWitness() string {
	if m != nil {
		return m.Witness
	}
	return ""
}

func (m *AccountWitnessVoteOperation) GetApprove() bool {
	if m != nil {
		return m.Approve
	}
	return false
}

func init() {
	proto.RegisterType((*AccountWitnessVoteOperation)(nil), "steemwatch.steem.AccountWitnessVoteOperation")
}

func init() {
	proto.RegisterFile("account_witness_vote_operation.proto", fileDescriptorAccountWitnessVoteOperation)
}

var fileDescriptorAccountWitnessVoteOperation = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x89, 0x2f, 0xcf, 0x2c, 0xc9, 0x4b, 0x2d, 0x2e, 0x8e, 0x2f, 0xcb, 0x2f, 0x49,
	0x8d, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95,
	0xb2, 0xb9, 0xa4, 0x1d, 0x21, 0x3a, 0xc3, 0x21, 0x1a, 0xc3, 0xf2, 0x4b, 0x52, 0xfd, 0x61, 0xda,
	0x84, 0x24, 0xb8, 0xd8, 0xa1, 0x06, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x20,
	0x19, 0xa8, 0x55, 0x12, 0x4c, 0x10, 0x19, 0x28, 0x17, 0xac, 0xa7, 0xa0, 0xa0, 0x28, 0xbf, 0x2c,
	0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0xc6, 0x75, 0xe2, 0x8c, 0x62, 0x07, 0xdb, 0x5a,
	0x90, 0x94, 0xc4, 0x06, 0x76, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x0f, 0x55, 0xfe,
	0xb8, 0x00, 0x00, 0x00,
}