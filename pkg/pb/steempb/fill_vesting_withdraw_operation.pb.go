// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fill_vesting_withdraw_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FillVestingWithdrawOperation struct {
}

func (m *FillVestingWithdrawOperation) Reset()         { *m = FillVestingWithdrawOperation{} }
func (m *FillVestingWithdrawOperation) String() string { return proto.CompactTextString(m) }
func (*FillVestingWithdrawOperation) ProtoMessage()    {}
func (*FillVestingWithdrawOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorFillVestingWithdrawOperation, []int{0}
}

func init() {
	proto.RegisterType((*FillVestingWithdrawOperation)(nil), "steemwatch.steem.FillVestingWithdrawOperation")
}

func init() {
	proto.RegisterFile("fill_vesting_withdraw_operation.proto", fileDescriptorFillVestingWithdrawOperation)
}

var fileDescriptorFillVestingWithdrawOperation = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcb, 0xcc, 0xc9,
	0x89, 0x2f, 0x4b, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a,
	0x2c, 0x8f, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33,
	0x95, 0xe4, 0xb8, 0x64, 0xdc, 0x32, 0x73, 0x72, 0xc2, 0x20, 0x3a, 0xc3, 0xa1, 0x1a, 0xfd, 0x61,
	0xfa, 0x9c, 0x38, 0xa3, 0xd8, 0xc1, 0x0a, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x66, 0x18, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc1, 0x60, 0xb6, 0x78, 0x6c, 0x00, 0x00, 0x00,
}
