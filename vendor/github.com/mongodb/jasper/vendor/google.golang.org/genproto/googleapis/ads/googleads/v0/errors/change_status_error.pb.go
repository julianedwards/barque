// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/change_status_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible change status errors.
type ChangeStatusErrorEnum_ChangeStatusError int32

const (
	// Enum unspecified.
	ChangeStatusErrorEnum_UNSPECIFIED ChangeStatusErrorEnum_ChangeStatusError = 0
	// The received error code is not known in this version.
	ChangeStatusErrorEnum_UNKNOWN ChangeStatusErrorEnum_ChangeStatusError = 1
	// The requested start date is too old.
	ChangeStatusErrorEnum_START_DATE_TOO_OLD ChangeStatusErrorEnum_ChangeStatusError = 3
)

var ChangeStatusErrorEnum_ChangeStatusError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	3: "START_DATE_TOO_OLD",
}
var ChangeStatusErrorEnum_ChangeStatusError_value = map[string]int32{
	"UNSPECIFIED":        0,
	"UNKNOWN":            1,
	"START_DATE_TOO_OLD": 3,
}

func (x ChangeStatusErrorEnum_ChangeStatusError) String() string {
	return proto.EnumName(ChangeStatusErrorEnum_ChangeStatusError_name, int32(x))
}
func (ChangeStatusErrorEnum_ChangeStatusError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_change_status_error_6bf40ad350498990, []int{0, 0}
}

// Container for enum describing possible change status errors.
type ChangeStatusErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeStatusErrorEnum) Reset()         { *m = ChangeStatusErrorEnum{} }
func (m *ChangeStatusErrorEnum) String() string { return proto.CompactTextString(m) }
func (*ChangeStatusErrorEnum) ProtoMessage()    {}
func (*ChangeStatusErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_change_status_error_6bf40ad350498990, []int{0}
}
func (m *ChangeStatusErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatusErrorEnum.Unmarshal(m, b)
}
func (m *ChangeStatusErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatusErrorEnum.Marshal(b, m, deterministic)
}
func (dst *ChangeStatusErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatusErrorEnum.Merge(dst, src)
}
func (m *ChangeStatusErrorEnum) XXX_Size() int {
	return xxx_messageInfo_ChangeStatusErrorEnum.Size(m)
}
func (m *ChangeStatusErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatusErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatusErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChangeStatusErrorEnum)(nil), "google.ads.googleads.v0.errors.ChangeStatusErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.ChangeStatusErrorEnum_ChangeStatusError", ChangeStatusErrorEnum_ChangeStatusError_name, ChangeStatusErrorEnum_ChangeStatusError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/change_status_error.proto", fileDescriptor_change_status_error_6bf40ad350498990)
}

var fileDescriptor_change_status_error_6bf40ad350498990 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0xc7, 0x6d, 0x07, 0x0a, 0xd9, 0xc1, 0x5a, 0x70, 0xc7, 0x1d, 0xfa, 0x00, 0x69, 0xc1, 0x8b,
	0xc4, 0x53, 0xb6, 0xd6, 0x51, 0x94, 0xb6, 0xd8, 0xae, 0x82, 0x14, 0x4a, 0xb6, 0x94, 0x28, 0x6c,
	0xcd, 0xc8, 0xaf, 0xdb, 0x03, 0x79, 0xf4, 0x51, 0x7c, 0x14, 0x6f, 0xbe, 0x81, 0x34, 0x71, 0xbd,
	0x0c, 0x3d, 0xe5, 0xcb, 0x37, 0x9f, 0xef, 0xef, 0x1f, 0xba, 0x15, 0x52, 0x8a, 0x4d, 0xe3, 0x33,
	0x0e, 0xbe, 0x91, 0xbd, 0x3a, 0x04, 0x7e, 0xa3, 0x94, 0x54, 0xe0, 0xaf, 0x5f, 0x59, 0x2b, 0x9a,
	0x1a, 0x3a, 0xd6, 0xed, 0xa1, 0xd6, 0x26, 0xde, 0x29, 0xd9, 0x49, 0x77, 0x6a, 0x70, 0xcc, 0x38,
	0xe0, 0x21, 0x89, 0x0f, 0x01, 0x36, 0x49, 0x6f, 0x85, 0xae, 0xe7, 0x3a, 0x9c, 0xeb, 0x6c, 0xd4,
	0xbb, 0x51, 0xbb, 0xdf, 0x7a, 0x31, 0xba, 0x3a, 0xf9, 0x70, 0x2f, 0xd1, 0x78, 0x99, 0xe4, 0x59,
	0x34, 0x8f, 0xef, 0xe3, 0x28, 0x74, 0xce, 0xdc, 0x31, 0xba, 0x58, 0x26, 0x0f, 0x49, 0xfa, 0x9c,
	0x38, 0x96, 0x3b, 0x41, 0x6e, 0x5e, 0xd0, 0xa7, 0xa2, 0x0e, 0x69, 0x11, 0xd5, 0x45, 0x9a, 0xd6,
	0xe9, 0x63, 0xe8, 0x8c, 0x66, 0xdf, 0x16, 0xf2, 0xd6, 0x72, 0x8b, 0xff, 0x1f, 0x65, 0x36, 0x39,
	0xe9, 0x97, 0xf5, 0x2b, 0x64, 0xd6, 0x4b, 0xf8, 0x9b, 0x14, 0x72, 0xc3, 0x5a, 0x81, 0xa5, 0x12,
	0xbe, 0x68, 0x5a, 0xbd, 0xe0, 0xf1, 0x1c, 0xbb, 0x37, 0xf8, 0xeb, 0x3a, 0x77, 0xe6, 0x79, 0xb7,
	0x47, 0x0b, 0x4a, 0x3f, 0xec, 0xe9, 0xc2, 0x14, 0xa3, 0x1c, 0xb0, 0x91, 0xbd, 0x2a, 0x03, 0xac,
	0x5b, 0xc2, 0xe7, 0x11, 0xa8, 0x28, 0x87, 0x6a, 0x00, 0xaa, 0x32, 0xa8, 0x0c, 0xf0, 0x65, 0x7b,
	0xc6, 0x25, 0x84, 0x72, 0x20, 0x64, 0x40, 0x08, 0x29, 0x03, 0x42, 0x0c, 0xb4, 0x3a, 0xd7, 0xd3,
	0xdd, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x80, 0x1c, 0xf9, 0xba, 0x01, 0x00, 0x00,
}
