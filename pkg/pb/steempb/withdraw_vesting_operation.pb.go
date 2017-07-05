// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: withdraw_vesting_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WithdrawVestingOperation struct {
	Account       string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	VestingShares string `protobuf:"bytes,2,opt,name=vesting_shares,json=vestingShares,proto3" json:"vesting_shares,omitempty"`
}

func (m *WithdrawVestingOperation) Reset()         { *m = WithdrawVestingOperation{} }
func (m *WithdrawVestingOperation) String() string { return proto.CompactTextString(m) }
func (*WithdrawVestingOperation) ProtoMessage()    {}
func (*WithdrawVestingOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorWithdrawVestingOperation, []int{0}
}

func (m *WithdrawVestingOperation) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *WithdrawVestingOperation) GetVestingShares() string {
	if m != nil {
		return m.VestingShares
	}
	return ""
}

func init() {
	proto.RegisterType((*WithdrawVestingOperation)(nil), "steemwatch.steem.WithdrawVestingOperation")
}

func init() {
	proto.RegisterFile("withdraw_vesting_operation.proto", fileDescriptorWithdrawVestingOperation)
}

var fileDescriptorWithdrawVestingOperation = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0xcf, 0x2c, 0xc9,
	0x48, 0x29, 0x4a, 0x2c, 0x8f, 0x2f, 0x4b, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x8f, 0xcf, 0x2f, 0x48,
	0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e,
	0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95, 0xa2, 0xb9, 0x24, 0xc2,
	0xa1, 0xba, 0xc2, 0x20, 0x9a, 0xfc, 0x61, 0x7a, 0x84, 0x24, 0xb8, 0xd8, 0x13, 0x93, 0x93, 0xf3,
	0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x55, 0x2e, 0x3e,
	0x98, 0x15, 0xc5, 0x19, 0x89, 0x45, 0xa9, 0xc5, 0x12, 0x4c, 0x60, 0x05, 0xbc, 0x50, 0xd1, 0x60,
	0xb0, 0xa0, 0x13, 0x67, 0x14, 0x3b, 0xd8, 0x96, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x03, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0xb4, 0x8d, 0x81, 0xa4, 0x00, 0x00, 0x00,
}