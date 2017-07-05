// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decline_voting_rights_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeclineVotingRightsOperation struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Decline bool   `protobuf:"varint,2,opt,name=decline,proto3" json:"decline,omitempty"`
}

func (m *DeclineVotingRightsOperation) Reset()         { *m = DeclineVotingRightsOperation{} }
func (m *DeclineVotingRightsOperation) String() string { return proto.CompactTextString(m) }
func (*DeclineVotingRightsOperation) ProtoMessage()    {}
func (*DeclineVotingRightsOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorDeclineVotingRightsOperation, []int{0}
}

func (m *DeclineVotingRightsOperation) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *DeclineVotingRightsOperation) GetDecline() bool {
	if m != nil {
		return m.Decline
	}
	return false
}

func init() {
	proto.RegisterType((*DeclineVotingRightsOperation)(nil), "steemwatch.steem.DeclineVotingRightsOperation")
}

func init() {
	proto.RegisterFile("decline_voting_rights_operation.proto", fileDescriptorDeclineVotingRightsOperation)
}

var fileDescriptorDeclineVotingRightsOperation = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x49, 0x4d, 0xce,
	0xc9, 0xcc, 0x4b, 0x8d, 0x2f, 0xcb, 0x2f, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0xca, 0x4c, 0xcf, 0x28,
	0x29, 0x8e, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33,
	0x95, 0x82, 0xb8, 0x64, 0x5c, 0x20, 0x5a, 0xc3, 0xc0, 0x3a, 0x83, 0xc0, 0x1a, 0xfd, 0x61, 0xfa,
	0x84, 0x24, 0xb8, 0xd8, 0x13, 0x93, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x60, 0x5c, 0x90, 0x0c, 0xd4, 0x52, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x18,
	0xd7, 0x89, 0x33, 0x8a, 0x1d, 0x6c, 0x78, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x5e, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x28, 0x0e, 0xbf, 0xa0, 0x00, 0x00, 0x00,
}