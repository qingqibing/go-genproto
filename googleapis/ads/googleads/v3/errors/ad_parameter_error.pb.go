// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/ad_parameter_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Enum describing possible ad parameter errors.
type AdParameterErrorEnum_AdParameterError int32

const (
	// Enum unspecified.
	AdParameterErrorEnum_UNSPECIFIED AdParameterErrorEnum_AdParameterError = 0
	// The received error code is not known in this version.
	AdParameterErrorEnum_UNKNOWN AdParameterErrorEnum_AdParameterError = 1
	// The ad group criterion must be a keyword criterion.
	AdParameterErrorEnum_AD_GROUP_CRITERION_MUST_BE_KEYWORD AdParameterErrorEnum_AdParameterError = 2
	// The insertion text is invalid.
	AdParameterErrorEnum_INVALID_INSERTION_TEXT_FORMAT AdParameterErrorEnum_AdParameterError = 3
)

var AdParameterErrorEnum_AdParameterError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_CRITERION_MUST_BE_KEYWORD",
	3: "INVALID_INSERTION_TEXT_FORMAT",
}

var AdParameterErrorEnum_AdParameterError_value = map[string]int32{
	"UNSPECIFIED":                        0,
	"UNKNOWN":                            1,
	"AD_GROUP_CRITERION_MUST_BE_KEYWORD": 2,
	"INVALID_INSERTION_TEXT_FORMAT":      3,
}

func (x AdParameterErrorEnum_AdParameterError) String() string {
	return proto.EnumName(AdParameterErrorEnum_AdParameterError_name, int32(x))
}

func (AdParameterErrorEnum_AdParameterError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cdb54cc8afb74021, []int{0, 0}
}

// Container for enum describing possible ad parameter errors.
type AdParameterErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdParameterErrorEnum) Reset()         { *m = AdParameterErrorEnum{} }
func (m *AdParameterErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdParameterErrorEnum) ProtoMessage()    {}
func (*AdParameterErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdb54cc8afb74021, []int{0}
}

func (m *AdParameterErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdParameterErrorEnum.Unmarshal(m, b)
}
func (m *AdParameterErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdParameterErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdParameterErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdParameterErrorEnum.Merge(m, src)
}
func (m *AdParameterErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdParameterErrorEnum.Size(m)
}
func (m *AdParameterErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdParameterErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdParameterErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.AdParameterErrorEnum_AdParameterError", AdParameterErrorEnum_AdParameterError_name, AdParameterErrorEnum_AdParameterError_value)
	proto.RegisterType((*AdParameterErrorEnum)(nil), "google.ads.googleads.v3.errors.AdParameterErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/ad_parameter_error.proto", fileDescriptor_cdb54cc8afb74021)
}

var fileDescriptor_cdb54cc8afb74021 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xa4, 0xf0, 0x0b, 0xd3, 0x85, 0x21, 0xe8, 0x46, 0xb4, 0x60, 0x16, 0x2e, 0x27,
	0x8b, 0x2c, 0x84, 0x71, 0x35, 0x6d, 0xa6, 0x25, 0xd4, 0x26, 0x21, 0x4d, 0x52, 0x95, 0xc0, 0x30,
	0x9a, 0x10, 0x0a, 0x6d, 0x26, 0xcc, 0xc4, 0x6e, 0x7c, 0x0c, 0xdf, 0xc0, 0xa5, 0x8f, 0xe2, 0xa3,
	0xb8, 0xf2, 0x11, 0x24, 0x19, 0xd3, 0x45, 0x41, 0x57, 0x73, 0xb8, 0x7c, 0xe7, 0xcc, 0x3d, 0x17,
	0x5c, 0x97, 0x9c, 0x97, 0x9b, 0xc2, 0x66, 0xb9, 0xb4, 0x95, 0x6c, 0xd5, 0xce, 0xb1, 0x0b, 0x21,
	0xb8, 0x90, 0x36, 0xcb, 0x69, 0xcd, 0x04, 0xdb, 0x16, 0x4d, 0x21, 0x68, 0x37, 0x83, 0xb5, 0xe0,
	0x0d, 0x37, 0x47, 0x8a, 0x86, 0x2c, 0x97, 0x70, 0x6f, 0x84, 0x3b, 0x07, 0x2a, 0xe3, 0xd9, 0x79,
	0x1f, 0x5c, 0xaf, 0x6d, 0x56, 0x55, 0xbc, 0x61, 0xcd, 0x9a, 0x57, 0x52, 0xb9, 0xad, 0x57, 0x0d,
	0x9c, 0xe0, 0x3c, 0xec, 0x93, 0x49, 0xeb, 0x21, 0xd5, 0xf3, 0xd6, 0x7a, 0x01, 0xc6, 0xe1, 0xdc,
	0x3c, 0x06, 0xc3, 0xc4, 0x5f, 0x86, 0x64, 0xe2, 0x4d, 0x3d, 0xe2, 0x1a, 0xff, 0xcc, 0x21, 0x38,
	0x4a, 0xfc, 0xb9, 0x1f, 0xac, 0x7c, 0x43, 0x33, 0xaf, 0x80, 0x85, 0x5d, 0x3a, 0x8b, 0x82, 0x24,
	0xa4, 0x93, 0xc8, 0x8b, 0x49, 0xe4, 0x05, 0x3e, 0x5d, 0x24, 0xcb, 0x98, 0x8e, 0x09, 0x9d, 0x93,
	0xfb, 0x55, 0x10, 0xb9, 0x86, 0x6e, 0x5e, 0x82, 0x0b, 0xcf, 0x4f, 0xf1, 0xad, 0xe7, 0x52, 0xcf,
	0x5f, 0x92, 0x28, 0x6e, 0xb1, 0x98, 0xdc, 0xc5, 0x74, 0x1a, 0x44, 0x0b, 0x1c, 0x1b, 0x83, 0xf1,
	0x97, 0x06, 0xac, 0x27, 0xbe, 0x85, 0x7f, 0x57, 0x1b, 0x9f, 0x1e, 0x6e, 0x18, 0xb6, 0x9d, 0x42,
	0xed, 0xc1, 0xfd, 0x31, 0x96, 0x7c, 0xc3, 0xaa, 0x12, 0x72, 0x51, 0xda, 0x65, 0x51, 0x75, 0x8d,
	0xfb, 0xe3, 0xd6, 0x6b, 0xf9, 0xdb, 0xad, 0x6f, 0xd4, 0xf3, 0xa6, 0x0f, 0x66, 0x18, 0xbf, 0xeb,
	0xa3, 0x99, 0x0a, 0xc3, 0xb9, 0x84, 0x4a, 0xb6, 0x2a, 0x75, 0x60, 0xf7, 0xa5, 0xfc, 0xe8, 0x81,
	0x0c, 0xe7, 0x32, 0xdb, 0x03, 0x59, 0xea, 0x64, 0x0a, 0xf8, 0xd4, 0x2d, 0x35, 0x45, 0x08, 0xe7,
	0x12, 0xa1, 0x3d, 0x82, 0x50, 0xea, 0x20, 0xa4, 0xa0, 0xc7, 0xff, 0xdd, 0x76, 0xce, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x4a, 0xd8, 0x47, 0x08, 0x02, 0x00, 0x00,
}
