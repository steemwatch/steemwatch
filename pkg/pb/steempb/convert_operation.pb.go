// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: convert_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConvertOperation struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	RequestID uint32 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *ConvertOperation) Reset()                    { *m = ConvertOperation{} }
func (m *ConvertOperation) String() string            { return proto.CompactTextString(m) }
func (*ConvertOperation) ProtoMessage()               {}
func (*ConvertOperation) Descriptor() ([]byte, []int) { return fileDescriptorConvertOperation, []int{0} }

func (m *ConvertOperation) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ConvertOperation) GetRequestID() uint32 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *ConvertOperation) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*ConvertOperation)(nil), "steemwatch.steem.ConvertOperation")
}

func init() { proto.RegisterFile("convert_operation.proto", fileDescriptorConvertOperation) }

var fileDescriptorConvertOperation = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0xcf, 0x2b,
	0x4b, 0x2d, 0x2a, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0,
	0x03, 0x33, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3,
	0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x0a, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x18,
	0xa0, 0x94, 0xc7, 0x25, 0xe0, 0x0c, 0x31, 0xdb, 0x1f, 0x66, 0xb4, 0x90, 0x08, 0x17, 0x6b, 0x7e,
	0x79, 0x5e, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0xa4, 0xc3, 0xc5,
	0x55, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x12, 0x9f, 0x99, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0xeb, 0xc4, 0xfb, 0xe8, 0x9e, 0x3c, 0x67, 0x10, 0x44, 0xd4, 0xd3, 0x25, 0x88, 0x13, 0xaa, 0xc0,
	0x33, 0x45, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x19, 0x6c, 0x08,
	0x94, 0xe7, 0xc4, 0x19, 0xc5, 0x0e, 0x76, 0x67, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x05, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x27, 0x0b, 0x1d, 0xdd, 0x00, 0x00, 0x00,
}
