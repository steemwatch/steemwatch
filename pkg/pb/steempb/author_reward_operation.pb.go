// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: author_reward_operation.proto

package steempb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AuthorRewardOperation struct {
}

func (m *AuthorRewardOperation) Reset()         { *m = AuthorRewardOperation{} }
func (m *AuthorRewardOperation) String() string { return proto.CompactTextString(m) }
func (*AuthorRewardOperation) ProtoMessage()    {}
func (*AuthorRewardOperation) Descriptor() ([]byte, []int) {
	return fileDescriptorAuthorRewardOperation, []int{0}
}

func init() {
	proto.RegisterType((*AuthorRewardOperation)(nil), "steemwatch.steem.AuthorRewardOperation")
}

func init() { proto.RegisterFile("author_reward_operation.proto", fileDescriptorAuthorRewardOperation) }

var fileDescriptorAuthorRewardOperation = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x2c, 0x2d, 0xc9,
	0xc8, 0x2f, 0x8a, 0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0x89, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c,
	0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x28, 0x2e, 0x49, 0x4d, 0xcd,
	0x2d, 0x4f, 0x2c, 0x49, 0xce, 0xd0, 0x03, 0x33, 0x95, 0xc4, 0xb9, 0x44, 0x1d, 0xc1, 0x5a, 0x82,
	0xc0, 0x3a, 0xfc, 0x61, 0x1a, 0x9c, 0x38, 0xa3, 0xd8, 0xc1, 0x2a, 0x0a, 0x92, 0x92, 0xd8, 0xc0,
	0x9a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0xd7, 0xff, 0xdd, 0x5d, 0x00, 0x00, 0x00,
}
