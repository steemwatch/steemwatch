// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: recover_account_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RecoverAccountOperation struct {
}

func (m *RecoverAccountOperation) Reset()         { *m = RecoverAccountOperation{} }
func (m *RecoverAccountOperation) String() string { return proto.CompactTextString(m) }
func (*RecoverAccountOperation) ProtoMessage()    {}
func (*RecoverAccountOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorRecoverAccountOperation, []int{0}
}

func init() {
	proto.RegisterType((*RecoverAccountOperation)(nil), "steemwatch.steem.RecoverAccountOperation")
}

func init() {
	proto.RegisterFile("recover_account_operation.proto", fileDescriptorRecoverAccountOperation)
}

var fileDescriptorRecoverAccountOperation = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4a, 0x4d, 0xce,
	0x2f, 0x4b, 0x2d, 0x8a, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0xcf, 0x2f, 0x48, 0x2d,
	0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49,
	0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95, 0x24, 0xb9, 0xc4, 0x83, 0x20,
	0x9a, 0x1c, 0x21, 0x7a, 0xfc, 0x61, 0x5a, 0x9c, 0x38, 0xa3, 0xd8, 0xc1, 0x6a, 0x0a, 0x92, 0x92,
	0xd8, 0xc0, 0xda, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x68, 0xdb, 0xce, 0x61, 0x00,
	0x00, 0x00,
}
