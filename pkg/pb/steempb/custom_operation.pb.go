// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: custom_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CustomOperation struct {
}

func (m *CustomOperation) Reset()                    { *m = CustomOperation{} }
func (m *CustomOperation) String() string            { return proto.CompactTextString(m) }
func (*CustomOperation) ProtoMessage()               {}
func (*CustomOperation) Descriptor() ([]byte, []int) { return fileDescriptorCustomOperation, []int{0} }

func init() {
	proto.RegisterType((*CustomOperation)(nil), "steemwatch.steem.CustomOperation")
}

func init() { proto.RegisterFile("custom_operation.proto", fileDescriptorCustomOperation) }

var fileDescriptorCustomOperation = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2e, 0x2d, 0x2e,
	0xc9, 0xcf, 0x8d, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03,
	0x33, 0x95, 0x04, 0xb9, 0xf8, 0x9d, 0xc1, 0x6a, 0xfd, 0x61, 0x4a, 0x9d, 0x38, 0xa3, 0xd8, 0xc1,
	0x72, 0x05, 0x49, 0x49, 0x6c, 0x60, 0x6d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x67,
	0xec, 0x92, 0x50, 0x00, 0x00, 0x00,
}