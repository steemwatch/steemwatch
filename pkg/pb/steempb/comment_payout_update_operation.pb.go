// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comment_payout_update_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommentPayoutUpdateOperation struct {
}

func (m *CommentPayoutUpdateOperation) Reset()         { *m = CommentPayoutUpdateOperation{} }
func (m *CommentPayoutUpdateOperation) String() string { return proto.CompactTextString(m) }
func (*CommentPayoutUpdateOperation) ProtoMessage()    {}
func (*CommentPayoutUpdateOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorCommentPayoutUpdateOperation, []int{0}
}

func init() {
	proto.RegisterType((*CommentPayoutUpdateOperation)(nil), "steemwatch.steem.CommentPayoutUpdateOperation")
}

func init() {
	proto.RegisterFile("comment_payout_update_operation.proto", fileDescriptorCommentPayoutUpdateOperation)
}

var fileDescriptorCommentPayoutUpdateOperation = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0x89, 0x2f, 0x48, 0xac, 0xcc, 0x2f, 0x2d, 0x89, 0x2f, 0x2d, 0x48, 0x49, 0x2c,
	0x49, 0x8d, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33,
	0x95, 0xe4, 0xb8, 0x64, 0x9c, 0x21, 0x5a, 0x03, 0xc0, 0x3a, 0x43, 0xc1, 0x1a, 0xfd, 0x61, 0xfa,
	0x9c, 0x38, 0xa3, 0xd8, 0xc1, 0x0a, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0x66, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xf4, 0x76, 0x6e, 0x0c, 0x6c, 0x00, 0x00, 0x00,
}
