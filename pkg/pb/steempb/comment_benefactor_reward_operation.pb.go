// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comment_benefactor_reward_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CommentBenefactorRewardOperation struct {
}

func (m *CommentBenefactorRewardOperation) Reset()         { *m = CommentBenefactorRewardOperation{} }
func (m *CommentBenefactorRewardOperation) String() string { return proto.CompactTextString(m) }
func (*CommentBenefactorRewardOperation) ProtoMessage()    {}
func (*CommentBenefactorRewardOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorCommentBenefactorRewardOperation, []int{0}
}

func init() {
	proto.RegisterType((*CommentBenefactorRewardOperation)(nil), "steemwatch.steem.CommentBenefactorRewardOperation")
}

func init() {
	proto.RegisterFile("comment_benefactor_reward_operation.proto", fileDescriptorCommentBenefactorRewardOperation)
}

var fileDescriptorCommentBenefactorRewardOperation = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0x89, 0x4f, 0x4a, 0xcd, 0x4b, 0x4d, 0x4b, 0x4c, 0x2e, 0xc9, 0x2f, 0x8a, 0x2f,
	0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd, 0x2d, 0x4f, 0x2c, 0x49,
	0xce, 0xd0, 0x03, 0x33, 0x95, 0x94, 0xb8, 0x14, 0x9c, 0x21, 0xda, 0x9d, 0xe0, 0xba, 0x83, 0xc0,
	0x9a, 0xfd, 0x61, 0x7a, 0x9d, 0x38, 0xa3, 0xd8, 0xc1, 0x8a, 0x0b, 0x92, 0x92, 0xd8, 0xc0, 0xe6,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x88, 0x34, 0x70, 0x74, 0x00, 0x00, 0x00,
}