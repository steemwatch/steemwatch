// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: return_vesting_delegation_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ReturnVestingDelegationOperation struct {
}

func (m *ReturnVestingDelegationOperation) Reset()         { *m = ReturnVestingDelegationOperation{} }
func (m *ReturnVestingDelegationOperation) String() string { return proto.CompactTextString(m) }
func (*ReturnVestingDelegationOperation) ProtoMessage()    {}
func (*ReturnVestingDelegationOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorReturnVestingDelegationOperation, []int{0}
}

func init() {
	proto.RegisterType((*ReturnVestingDelegationOperation)(nil), "steemwatch.steem.ReturnVestingDelegationOperation")
}

func init() {
	proto.RegisterFile("return_vesting_delegation_operation.proto", fileDescriptorReturnVestingDelegationOperation)
}

var fileDescriptorReturnVestingDelegationOperation = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x4a, 0x2d, 0x29,
	0x2d, 0xca, 0x8b, 0x2f, 0x4b, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0x8f, 0x4f, 0x49, 0xcd, 0x49, 0x4d,
	0x4f, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf, 0x2f, 0x48, 0x2d, 0x02, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x04, 0x8a, 0x4b, 0x52, 0x53, 0x73, 0xcb, 0x13, 0x4b, 0x92, 0x33, 0xf4, 0xc0,
	0x4c, 0x25, 0x25, 0x2e, 0x85, 0x20, 0xb0, 0xf6, 0x30, 0x88, 0x6e, 0x17, 0xb8, 0x66, 0x7f, 0x98,
	0x5e, 0x27, 0xce, 0x28, 0x76, 0xb0, 0xe2, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x39, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb6, 0x7c, 0xee, 0xeb, 0x74, 0x00, 0x00, 0x00,
}
