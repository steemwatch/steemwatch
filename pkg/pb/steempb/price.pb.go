// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: price.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Price struct {
	Base  string `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Quote string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (m *Price) Reset()                    { *m = Price{} }
func (m *Price) String() string            { return proto.CompactTextString(m) }
func (*Price) ProtoMessage()               {}
func (*Price) Descriptor() ([]byte, []int) { return fileDescriptorPrice, []int{0} }

func (m *Price) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Price) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func init() {
	proto.RegisterType((*Price)(nil), "steemwatch.steem.Price")
}

func init() { proto.RegisterFile("price.proto", fileDescriptorPrice) }

var fileDescriptorPrice = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x4c,
	0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f,
	0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95, 0x0c, 0xb9, 0x58, 0x03, 0x40, 0x0a, 0x84, 0x84, 0xb8,
	0x58, 0x92, 0x12, 0x8b, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11,
	0x2e, 0xd6, 0xc2, 0xd2, 0xfc, 0x92, 0x54, 0x09, 0x26, 0xb0, 0x20, 0x84, 0xe3, 0xc4, 0x19, 0xc5,
	0x0e, 0xd6, 0x5b, 0x90, 0x94, 0xc4, 0x06, 0x36, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x89, 0x49, 0x0f, 0x65, 0x00, 0x00, 0x00,
}