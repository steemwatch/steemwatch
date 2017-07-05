// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: set_withdraw_vesting_route_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SetWithdrawVestingRouteOperation struct {
	FromAccount string `protobuf:"bytes,1,opt,name=from_account,json=fromAccount,proto3" json:"from_account,omitempty"`
	ToAccount   string `protobuf:"bytes,2,opt,name=to_account,json=toAccount,proto3" json:"to_account,omitempty"`
	Percent     uint32 `protobuf:"varint,3,opt,name=percent,proto3" json:"percent,omitempty"`
	AutoVest    bool   `protobuf:"varint,4,opt,name=auto_vest,json=autoVest,proto3" json:"auto_vest,omitempty"`
}

func (m *SetWithdrawVestingRouteOperation) Reset()         { *m = SetWithdrawVestingRouteOperation{} }
func (m *SetWithdrawVestingRouteOperation) String() string { return proto.CompactTextString(m) }
func (*SetWithdrawVestingRouteOperation) ProtoMessage()    {}
func (*SetWithdrawVestingRouteOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorSetWithdrawVestingRouteOperation, []int{0}
}

func (m *SetWithdrawVestingRouteOperation) GetFromAccount() string {
	if m != nil {
		return m.FromAccount
	}
	return ""
}

func (m *SetWithdrawVestingRouteOperation) GetToAccount() string {
	if m != nil {
		return m.ToAccount
	}
	return ""
}

func (m *SetWithdrawVestingRouteOperation) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *SetWithdrawVestingRouteOperation) GetAutoVest() bool {
	if m != nil {
		return m.AutoVest
	}
	return false
}

func init() {
	proto.RegisterType((*SetWithdrawVestingRouteOperation)(nil), "steemwatch.steem.SetWithdrawVestingRouteOperation")
}

func init() {
	proto.RegisterFile("set_withdraw_vesting_route_operation.proto", fileDescriptorSetWithdrawVestingRouteOperation)
}

var fileDescriptorSetWithdrawVestingRouteOperation = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0x8a, 0x6d, 0xa2, 0x82, 0xe4, 0x14, 0x10, 0x21, 0x7a, 0x2a, 0x1e, 0x7a, 0xf1,
	0x09, 0xf4, 0x05, 0x84, 0x08, 0x0a, 0x5e, 0x42, 0x1a, 0x67, 0xb7, 0x3d, 0x34, 0x13, 0xd2, 0xe9,
	0xf6, 0x61, 0xf6, 0x65, 0x97, 0xa4, 0xdb, 0xbd, 0xcd, 0xfc, 0xf3, 0x7f, 0xf0, 0x8d, 0x78, 0x9b,
	0x80, 0xec, 0x32, 0x50, 0xff, 0x9f, 0xdc, 0x62, 0x0f, 0x30, 0xd1, 0x10, 0xf6, 0x36, 0xe1, 0x4c,
	0x60, 0x31, 0x42, 0x72, 0x34, 0x60, 0x68, 0x63, 0x42, 0x42, 0xf9, 0x38, 0x11, 0xc0, 0xb8, 0x38,
	0xf2, 0x7d, 0x5b, 0xc6, 0xd7, 0x23, 0x13, 0xfa, 0x1b, 0xe8, 0xf7, 0xcc, 0xff, 0xac, 0xb8, 0xc9,
	0xf4, 0xd7, 0x06, 0xcb, 0x17, 0x71, 0xbf, 0x4b, 0x38, 0x5a, 0xe7, 0x3d, 0xce, 0x81, 0x14, 0xd3,
	0xac, 0xe1, 0xe6, 0x2e, 0x67, 0x1f, 0x6b, 0x24, 0x9f, 0x85, 0x20, 0xbc, 0x14, 0xae, 0x4a, 0x81,
	0x13, 0x6e, 0x67, 0x25, 0xaa, 0x08, 0xc9, 0x43, 0x20, 0x75, 0xad, 0x59, 0xf3, 0x60, 0xb6, 0x55,
	0x3e, 0x09, 0xee, 0x66, 0xc2, 0x22, 0xae, 0x6e, 0x34, 0x6b, 0x6a, 0x53, 0xe7, 0x20, 0x9b, 0x7c,
	0xf2, 0xbf, 0xaa, 0x68, 0xc6, 0xae, 0xbb, 0x2d, 0x1f, 0xbc, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe9, 0xa3, 0xe9, 0x82, 0xef, 0x00, 0x00, 0x00,
}
