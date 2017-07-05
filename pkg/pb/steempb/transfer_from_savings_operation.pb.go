// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer_from_savings_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TransferFromSavingsOperation struct {
	From      string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	RequestID uint32 `protobuf:"varint,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	To        string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Amount    string `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Memo      string `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *TransferFromSavingsOperation) Reset()         { *m = TransferFromSavingsOperation{} }
func (m *TransferFromSavingsOperation) String() string { return proto.CompactTextString(m) }
func (*TransferFromSavingsOperation) ProtoMessage()    {}
func (*TransferFromSavingsOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferFromSavingsOperation, []int{0}
}

func (m *TransferFromSavingsOperation) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransferFromSavingsOperation) GetRequestID() uint32 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *TransferFromSavingsOperation) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransferFromSavingsOperation) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransferFromSavingsOperation) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterType((*TransferFromSavingsOperation)(nil), "steemwatch.steem.TransferFromSavingsOperation")
}

func init() {
	proto.RegisterFile("transfer_from_savings_operation.proto", fileDescriptorTransferFromSavingsOperation)
}

var fileDescriptorTransferFromSavingsOperation = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xb5, 0x56, 0x36, 0x50, 0x91, 0x1c, 0x24, 0x88, 0x60, 0x11, 0x84, 0x1e, 0x74,
	0x7b, 0xf0, 0x0d, 0x8a, 0x08, 0x3d, 0x09, 0xab, 0x27, 0x2f, 0x21, 0x69, 0xb3, 0xd9, 0x1c, 0x92,
	0x59, 0x93, 0x89, 0xbe, 0x8c, 0xcf, 0xe6, 0xc1, 0x27, 0x91, 0x9d, 0xac, 0xb7, 0xef, 0x1f, 0x26,
	0xf9, 0xfe, 0x61, 0x77, 0x18, 0x55, 0x48, 0xbd, 0x89, 0xb2, 0x8f, 0xe0, 0x65, 0x52, 0x9f, 0x2e,
	0xd8, 0x24, 0x61, 0x34, 0x51, 0xa1, 0x83, 0xd0, 0x8e, 0x11, 0x10, 0xf8, 0x45, 0x42, 0x63, 0xfc,
	0x97, 0xc2, 0xc3, 0xd0, 0x12, 0x5e, 0x3d, 0x58, 0x87, 0x43, 0xd6, 0xed, 0x01, 0xfc, 0xd6, 0x82,
	0x85, 0x2d, 0x2d, 0xea, 0xdc, 0x53, 0xa2, 0x40, 0x54, 0x3e, 0xb8, 0xfd, 0xae, 0xd8, 0xf5, 0xdb,
	0xac, 0x7a, 0x8e, 0xe0, 0x5f, 0x8b, 0xe8, 0xe5, 0xdf, 0xc3, 0x39, 0x5b, 0x4c, 0x0d, 0x44, 0xb5,
	0xae, 0x36, 0x4d, 0x47, 0xcc, 0xef, 0x19, 0x8b, 0xe6, 0x23, 0x9b, 0x84, 0xd2, 0x1d, 0x45, 0xbd,
	0xae, 0x36, 0xab, 0xdd, 0xea, 0xf7, 0xe7, 0xa6, 0xe9, 0xca, 0x74, 0xff, 0xd4, 0x35, 0xf3, 0xc2,
	0xfe, 0xc8, 0xcf, 0x59, 0x8d, 0x20, 0x4e, 0xe8, 0x7d, 0x8d, 0xc0, 0x2f, 0xd9, 0x52, 0x79, 0xc8,
	0x01, 0xc5, 0x82, 0x66, 0x73, 0x9a, 0x4c, 0xde, 0x78, 0x10, 0xa7, 0xc5, 0x34, 0xf1, 0xae, 0x79,
	0x3f, 0xa3, 0xb3, 0x46, 0xad, 0x97, 0x54, 0xf8, 0xf1, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x02, 0x4d,
	0x1b, 0x86, 0x1a, 0x01, 0x00, 0x00,
}
