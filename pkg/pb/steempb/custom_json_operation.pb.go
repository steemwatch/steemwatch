// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: custom_json_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CustomJSONOperation struct {
	RequiredAuths        []string `protobuf:"bytes,1,rep,name=required_auths,json=requiredAuths" json:"required_auths,omitempty"`
	RequiredPostingAuths []string `protobuf:"bytes,2,rep,name=required_posting_auths,json=requiredPostingAuths" json:"required_posting_auths,omitempty"`
	ID                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	JSON                 string   `protobuf:"bytes,4,opt,name=json,proto3" json:"json,omitempty"`
}

func (m *CustomJSONOperation) Reset()         { *m = CustomJSONOperation{} }
func (m *CustomJSONOperation) String() string { return proto.CompactTextString(m) }
func (*CustomJSONOperation) ProtoMessage()    {}
func (*CustomJSONOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorCustomJsonOperation, []int{0}
}

func (m *CustomJSONOperation) GetRequiredAuths() []string {
	if m != nil {
		return m.RequiredAuths
	}
	return nil
}

func (m *CustomJSONOperation) GetRequiredPostingAuths() []string {
	if m != nil {
		return m.RequiredPostingAuths
	}
	return nil
}

func (m *CustomJSONOperation) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *CustomJSONOperation) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

func init() {
	proto.RegisterType((*CustomJSONOperation)(nil), "steemwatch.steem.CustomJSONOperation")
}

func init() { proto.RegisterFile("custom_json_operation.proto", fileDescriptorCustomJsonOperation) }

var fileDescriptorCustomJsonOperation = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2e, 0x2d, 0x2e,
	0xc9, 0xcf, 0x8d, 0xcf, 0x2a, 0xce, 0xcf, 0x8b, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f,
	0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0xa5, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0x0a, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30,
	0x07, 0xcc, 0x82, 0x18, 0xa0, 0xb4, 0x82, 0x91, 0x4b, 0xd8, 0x19, 0x6c, 0x81, 0x57, 0xb0, 0xbf,
	0x9f, 0x3f, 0xcc, 0x78, 0x21, 0x55, 0x2e, 0xbe, 0xa2, 0xd4, 0xc2, 0xd2, 0xcc, 0xa2, 0xd4, 0x94,
	0xf8, 0xc4, 0xd2, 0x92, 0x8c, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x5e, 0x98, 0xa8,
	0x23, 0x48, 0x50, 0xc8, 0x84, 0x4b, 0x0c, 0xae, 0xac, 0x20, 0xbf, 0xb8, 0x24, 0x33, 0x2f, 0x1d,
	0xaa, 0x9c, 0x09, 0xac, 0x5c, 0x04, 0x26, 0x1b, 0x00, 0x91, 0x84, 0xe8, 0x12, 0xe3, 0x62, 0xca,
	0x4c, 0x91, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x74, 0x62, 0x7b, 0x74, 0x4f, 0x9e, 0xc9, 0xd3, 0x25,
	0x88, 0x29, 0x33, 0x45, 0x48, 0x86, 0x8b, 0x05, 0xe4, 0x4b, 0x09, 0x16, 0xb0, 0x0c, 0xc7, 0xa3,
	0x7b, 0xf2, 0x2c, 0x20, 0x57, 0x05, 0x81, 0x45, 0x9d, 0x38, 0xa3, 0xd8, 0xc1, 0x5e, 0x2c, 0x48,
	0x4a, 0x62, 0x03, 0x3b, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xd8, 0x0e, 0xc4, 0x1c,
	0x01, 0x00, 0x00,
}
