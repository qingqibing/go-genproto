// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/frequency_cap_level.proto

package enums

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

// The level on which the cap is to be applied (e.g ad group ad, ad group).
// Cap is applied to all the resources of this level.
type FrequencyCapLevelEnum_FrequencyCapLevel int32

const (
	// Not specified.
	FrequencyCapLevelEnum_UNSPECIFIED FrequencyCapLevelEnum_FrequencyCapLevel = 0
	// Used for return value only. Represents value unknown in this version.
	FrequencyCapLevelEnum_UNKNOWN FrequencyCapLevelEnum_FrequencyCapLevel = 1
	// The cap is applied at the ad group ad level.
	FrequencyCapLevelEnum_AD_GROUP_AD FrequencyCapLevelEnum_FrequencyCapLevel = 2
	// The cap is applied at the ad group level.
	FrequencyCapLevelEnum_AD_GROUP FrequencyCapLevelEnum_FrequencyCapLevel = 3
	// The cap is applied at the campaign level.
	FrequencyCapLevelEnum_CAMPAIGN FrequencyCapLevelEnum_FrequencyCapLevel = 4
)

var FrequencyCapLevelEnum_FrequencyCapLevel_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_AD",
	3: "AD_GROUP",
	4: "CAMPAIGN",
}

var FrequencyCapLevelEnum_FrequencyCapLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"AD_GROUP_AD": 2,
	"AD_GROUP":    3,
	"CAMPAIGN":    4,
}

func (x FrequencyCapLevelEnum_FrequencyCapLevel) String() string {
	return proto.EnumName(FrequencyCapLevelEnum_FrequencyCapLevel_name, int32(x))
}

func (FrequencyCapLevelEnum_FrequencyCapLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7591f0d13b525d9e, []int{0, 0}
}

// Container for enum describing the level on which the cap is to be applied.
type FrequencyCapLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrequencyCapLevelEnum) Reset()         { *m = FrequencyCapLevelEnum{} }
func (m *FrequencyCapLevelEnum) String() string { return proto.CompactTextString(m) }
func (*FrequencyCapLevelEnum) ProtoMessage()    {}
func (*FrequencyCapLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7591f0d13b525d9e, []int{0}
}

func (m *FrequencyCapLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequencyCapLevelEnum.Unmarshal(m, b)
}
func (m *FrequencyCapLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequencyCapLevelEnum.Marshal(b, m, deterministic)
}
func (m *FrequencyCapLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequencyCapLevelEnum.Merge(m, src)
}
func (m *FrequencyCapLevelEnum) XXX_Size() int {
	return xxx_messageInfo_FrequencyCapLevelEnum.Size(m)
}
func (m *FrequencyCapLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequencyCapLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FrequencyCapLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FrequencyCapLevelEnum_FrequencyCapLevel", FrequencyCapLevelEnum_FrequencyCapLevel_name, FrequencyCapLevelEnum_FrequencyCapLevel_value)
	proto.RegisterType((*FrequencyCapLevelEnum)(nil), "google.ads.googleads.v3.enums.FrequencyCapLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/frequency_cap_level.proto", fileDescriptor_7591f0d13b525d9e)
}

var fileDescriptor_7591f0d13b525d9e = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x1c, 0x76, 0x9d, 0xa8, 0x64, 0x82, 0xb5, 0xa0, 0x07, 0x71, 0x87, 0xed, 0x01, 0x92, 0x43, 0x0e,
	0x42, 0x3c, 0x65, 0xff, 0xca, 0x50, 0xbb, 0xa2, 0x6c, 0x82, 0x14, 0x4b, 0x5c, 0x63, 0x18, 0x74,
	0x49, 0x5d, 0xb6, 0x8a, 0xaf, 0xe3, 0xd1, 0x47, 0xf1, 0x51, 0xf4, 0x25, 0x24, 0x89, 0xed, 0x65,
	0xe8, 0xa5, 0x7c, 0xfd, 0x7d, 0x7f, 0xf8, 0xf2, 0x81, 0x0b, 0xa1, 0x94, 0xc8, 0x39, 0x62, 0x99,
	0x46, 0x0e, 0x1a, 0x54, 0x62, 0xc4, 0xe5, 0x66, 0xa9, 0xd1, 0xf3, 0x8a, 0xbf, 0x6c, 0xb8, 0x9c,
	0xbf, 0xa5, 0x73, 0x56, 0xa4, 0x39, 0x2f, 0x79, 0x0e, 0x8b, 0x95, 0x5a, 0xab, 0xa0, 0xed, 0xd4,
	0x90, 0x65, 0x1a, 0xd6, 0x46, 0x58, 0x62, 0x68, 0x8d, 0x67, 0xe7, 0x55, 0x6e, 0xb1, 0x40, 0x4c,
	0x4a, 0xb5, 0x66, 0xeb, 0x85, 0x92, 0xda, 0x99, 0xbb, 0xaf, 0xe0, 0x64, 0x54, 0x25, 0xf7, 0x59,
	0x71, 0x6d, 0x72, 0x87, 0x72, 0xb3, 0xec, 0x3e, 0x82, 0xe3, 0x2d, 0x22, 0x38, 0x02, 0xad, 0x69,
	0x74, 0x17, 0x0f, 0xfb, 0xe3, 0xd1, 0x78, 0x38, 0xf0, 0x77, 0x82, 0x16, 0xd8, 0x9f, 0x46, 0x57,
	0xd1, 0xe4, 0x3e, 0xf2, 0x1b, 0x86, 0xa5, 0x83, 0x34, 0xbc, 0x9d, 0x4c, 0xe3, 0x94, 0x0e, 0x7c,
	0x2f, 0x38, 0x04, 0x07, 0xd5, 0xc1, 0x6f, 0x9a, 0xbf, 0x3e, 0xbd, 0x89, 0xe9, 0x38, 0x8c, 0xfc,
	0xdd, 0xde, 0x77, 0x03, 0x74, 0xe6, 0x6a, 0x09, 0xff, 0x2d, 0xdf, 0x3b, 0xdd, 0xea, 0x10, 0x9b,
	0xda, 0x71, 0xe3, 0xa1, 0xf7, 0x6b, 0x14, 0x2a, 0x67, 0x52, 0x40, 0xb5, 0x12, 0x48, 0x70, 0x69,
	0x1f, 0x55, 0xcd, 0x57, 0x2c, 0xf4, 0x1f, 0x6b, 0x5e, 0xda, 0xef, 0xbb, 0xd7, 0x0c, 0x29, 0xfd,
	0xf0, 0xda, 0xa1, 0x8b, 0xa2, 0x99, 0x86, 0x0e, 0x1a, 0x34, 0xc3, 0xd0, 0x0c, 0xa1, 0x3f, 0x2b,
	0x3e, 0xa1, 0x99, 0x4e, 0x6a, 0x3e, 0x99, 0xe1, 0xc4, 0xf2, 0x5f, 0x5e, 0xc7, 0x1d, 0x09, 0xa1,
	0x99, 0x26, 0xa4, 0x56, 0x10, 0x32, 0xc3, 0x84, 0x58, 0xcd, 0xd3, 0x9e, 0x2d, 0x86, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xc2, 0x8b, 0x44, 0x67, 0xe5, 0x01, 0x00, 0x00,
}
