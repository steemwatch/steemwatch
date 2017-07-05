// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer_to_savings_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TransferToSavingsOperation struct {
	From   string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Memo   string `protobuf:"bytes,4,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *TransferToSavingsOperation) Reset()         { *m = TransferToSavingsOperation{} }
func (m *TransferToSavingsOperation) String() string { return proto.CompactTextString(m) }
func (*TransferToSavingsOperation) ProtoMessage()    {}
func (*TransferToSavingsOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorTransferToSavingsOperation, []int{0}
}

func (m *TransferToSavingsOperation) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TransferToSavingsOperation) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *TransferToSavingsOperation) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *TransferToSavingsOperation) GetMemo() string {
	if m != nil {
		return m.Memo
	}
	return ""
}

func init() {
	proto.RegisterType((*TransferToSavingsOperation)(nil), "steemwatch.steem.TransferToSavingsOperation")
}

func init() {
	proto.RegisterFile("transfer_to_savings_operation.proto", fileDescriptorTransferToSavingsOperation)
}

var fileDescriptorTransferToSavingsOperation = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0x4b, 0x2d, 0x8a, 0x2f, 0xc9, 0x8f, 0x2f, 0x4e, 0x2c, 0xcb, 0xcc, 0x4b, 0x2f, 0x8e,
	0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95, 0x72,
	0xb8, 0xa4, 0x42, 0xa0, 0x1a, 0x43, 0xf2, 0x83, 0x21, 0xda, 0xfc, 0x61, 0xba, 0x84, 0x84, 0xb8,
	0x58, 0xd2, 0x8a, 0xf2, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x3e,
	0x2e, 0xa6, 0x92, 0x7c, 0x09, 0x26, 0xb0, 0x08, 0x53, 0x49, 0xbe, 0x90, 0x18, 0x17, 0x5b, 0x62,
	0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x33, 0x58, 0x0c, 0xca, 0x03, 0xe9, 0xcd, 0x4d, 0xcd, 0xcd,
	0x97, 0x60, 0x81, 0xe8, 0x05, 0xb1, 0x9d, 0x38, 0xa3, 0xd8, 0xc1, 0xd6, 0x16, 0x24, 0x25, 0xb1,
	0x81, 0x5d, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x96, 0x81, 0xc9, 0x59, 0xb8, 0x00, 0x00,
	0x00,
}