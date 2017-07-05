// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: limit_order_cancel_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LimitOrderCancelOperation struct {
	Owner   string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	OrderID uint32 `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (m *LimitOrderCancelOperation) Reset()         { *m = LimitOrderCancelOperation{} }
func (m *LimitOrderCancelOperation) String() string { return proto.CompactTextString(m) }
func (*LimitOrderCancelOperation) ProtoMessage()    {}
func (*LimitOrderCancelOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorLimitOrderCancelOperation, []int{0}
}

func (m *LimitOrderCancelOperation) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *LimitOrderCancelOperation) GetOrderID() uint32 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

func init() {
	proto.RegisterType((*LimitOrderCancelOperation)(nil), "steemwatch.steem.LimitOrderCancelOperation")
}

func init() {
	proto.RegisterFile("limit_order_cancel_operation.proto", fileDescriptorLimitOrderCancelOperation)
}

var fileDescriptorLimitOrderCancelOperation = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xc9, 0xcc, 0xcd,
	0x2c, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x8a, 0x4f, 0x4e, 0xcc, 0x4b, 0x4e, 0xcd, 0x89, 0xcf,
	0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0xa5, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0x0a, 0x93, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x18, 0xa0, 0x14, 0xc9, 0x25, 0xe9,
	0x03, 0xb2, 0xc6, 0x1f, 0x64, 0x8b, 0x33, 0xd8, 0x12, 0x7f, 0x98, 0x1d, 0x42, 0x22, 0x5c, 0xac,
	0xf9, 0xe5, 0x79, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x1a,
	0x17, 0x07, 0xc4, 0x4d, 0x99, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc, 0x4e, 0xdc, 0x8f, 0xee,
	0xc9, 0xb3, 0x83, 0x4d, 0xf0, 0x74, 0x09, 0x62, 0x07, 0x4b, 0x7a, 0xa6, 0x38, 0x71, 0x46, 0xb1,
	0x83, 0x9d, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xb6, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x12,
	0x8b, 0xb9, 0xb8, 0xd3, 0x00, 0x00, 0x00,
}