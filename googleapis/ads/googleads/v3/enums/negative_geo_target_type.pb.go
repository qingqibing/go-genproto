// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/negative_geo_target_type.proto

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

// The possible negative geo target types.
type NegativeGeoTargetTypeEnum_NegativeGeoTargetType int32

const (
	// Not specified.
	NegativeGeoTargetTypeEnum_UNSPECIFIED NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 0
	// The value is unknown in this version.
	NegativeGeoTargetTypeEnum_UNKNOWN NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 1
	// Specifies that a user is excluded from seeing the ad if they
	// are in, or show interest in, advertiser's excluded locations.
	NegativeGeoTargetTypeEnum_PRESENCE_OR_INTEREST NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 4
	// Specifies that a user is excluded from seeing the ad if they
	// are in advertiser's excluded locations.
	NegativeGeoTargetTypeEnum_PRESENCE NegativeGeoTargetTypeEnum_NegativeGeoTargetType = 5
)

var NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	4: "PRESENCE_OR_INTEREST",
	5: "PRESENCE",
}

var NegativeGeoTargetTypeEnum_NegativeGeoTargetType_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"PRESENCE_OR_INTEREST": 4,
	"PRESENCE":             5,
}

func (x NegativeGeoTargetTypeEnum_NegativeGeoTargetType) String() string {
	return proto.EnumName(NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name, int32(x))
}

func (NegativeGeoTargetTypeEnum_NegativeGeoTargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_145bd44f7b043ddc, []int{0, 0}
}

// Container for enum describing possible negative geo target types.
type NegativeGeoTargetTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NegativeGeoTargetTypeEnum) Reset()         { *m = NegativeGeoTargetTypeEnum{} }
func (m *NegativeGeoTargetTypeEnum) String() string { return proto.CompactTextString(m) }
func (*NegativeGeoTargetTypeEnum) ProtoMessage()    {}
func (*NegativeGeoTargetTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_145bd44f7b043ddc, []int{0}
}

func (m *NegativeGeoTargetTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Unmarshal(m, b)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Marshal(b, m, deterministic)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NegativeGeoTargetTypeEnum.Merge(m, src)
}
func (m *NegativeGeoTargetTypeEnum) XXX_Size() int {
	return xxx_messageInfo_NegativeGeoTargetTypeEnum.Size(m)
}
func (m *NegativeGeoTargetTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NegativeGeoTargetTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NegativeGeoTargetTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.NegativeGeoTargetTypeEnum_NegativeGeoTargetType", NegativeGeoTargetTypeEnum_NegativeGeoTargetType_name, NegativeGeoTargetTypeEnum_NegativeGeoTargetType_value)
	proto.RegisterType((*NegativeGeoTargetTypeEnum)(nil), "google.ads.googleads.v3.enums.NegativeGeoTargetTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/negative_geo_target_type.proto", fileDescriptor_145bd44f7b043ddc)
}

var fileDescriptor_145bd44f7b043ddc = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x4a, 0xfb, 0x30,
	0x1c, 0xc5, 0x7f, 0xdb, 0xcf, 0x7f, 0x64, 0x82, 0xa5, 0x28, 0xe8, 0x70, 0x17, 0xdb, 0x03, 0xa4,
	0x17, 0xb9, 0x8b, 0xde, 0x74, 0x33, 0x8e, 0x21, 0x64, 0x65, 0xeb, 0x26, 0x48, 0xa5, 0x44, 0x1b,
	0x42, 0x61, 0x4b, 0xca, 0x92, 0x0d, 0xe6, 0xe3, 0x78, 0xe9, 0xa3, 0xf8, 0x28, 0xde, 0xf8, 0x0a,
	0xd2, 0x64, 0xed, 0xd5, 0xf4, 0xa6, 0x1c, 0x7a, 0xbe, 0x9f, 0xc3, 0xc9, 0x01, 0xb7, 0x42, 0x29,
	0xb1, 0xe0, 0x01, 0xcb, 0x74, 0xe0, 0x64, 0xa9, 0x36, 0x28, 0xe0, 0x72, 0xbd, 0xd4, 0x81, 0xe4,
	0x82, 0x99, 0x7c, 0xc3, 0x53, 0xc1, 0x55, 0x6a, 0xd8, 0x4a, 0x70, 0x93, 0x9a, 0x6d, 0xc1, 0x61,
	0xb1, 0x52, 0x46, 0xf9, 0x1d, 0x87, 0x40, 0x96, 0x69, 0x58, 0xd3, 0x70, 0x83, 0xa0, 0xa5, 0xdb,
	0xd7, 0x55, 0x78, 0x91, 0x07, 0x4c, 0x4a, 0x65, 0x98, 0xc9, 0x95, 0xd4, 0x0e, 0xee, 0xbd, 0x81,
	0x2b, 0xba, 0x8b, 0x1f, 0x72, 0x15, 0xdb, 0xf0, 0x78, 0x5b, 0x70, 0x22, 0xd7, 0xcb, 0xde, 0x33,
	0xb8, 0xd8, 0x6b, 0xfa, 0x67, 0xa0, 0x35, 0xa3, 0xd3, 0x88, 0x0c, 0x46, 0xf7, 0x23, 0x72, 0xe7,
	0xfd, 0xf3, 0x5b, 0xe0, 0x78, 0x46, 0x1f, 0xe8, 0xf8, 0x91, 0x7a, 0x0d, 0xff, 0x12, 0x9c, 0x47,
	0x13, 0x32, 0x25, 0x74, 0x40, 0xd2, 0xf1, 0x24, 0x1d, 0xd1, 0x98, 0x4c, 0xc8, 0x34, 0xf6, 0x0e,
	0xfc, 0x53, 0x70, 0x52, 0x39, 0xde, 0x61, 0xff, 0xbb, 0x01, 0xba, 0xaf, 0x6a, 0x09, 0xff, 0xec,
	0xdf, 0x6f, 0xef, 0xad, 0x10, 0x95, 0xed, 0xa3, 0xc6, 0x53, 0x7f, 0x07, 0x0b, 0xb5, 0x60, 0x52,
	0x40, 0xb5, 0x12, 0x81, 0xe0, 0xd2, 0xbe, 0xad, 0x9a, 0xb2, 0xc8, 0xf5, 0x2f, 0xcb, 0xde, 0xd8,
	0xef, 0x7b, 0xf3, 0xff, 0x30, 0x0c, 0x3f, 0x9a, 0x9d, 0xa1, 0x8b, 0x0a, 0x33, 0x0d, 0x9d, 0x2c,
	0xd5, 0x1c, 0xc1, 0x72, 0x0b, 0xfd, 0x59, 0xf9, 0x49, 0x98, 0xe9, 0xa4, 0xf6, 0x93, 0x39, 0x4a,
	0xac, 0xff, 0xd5, 0xec, 0xba, 0x9f, 0x18, 0x87, 0x99, 0xc6, 0xb8, 0xbe, 0xc0, 0x78, 0x8e, 0x30,
	0xb6, 0x37, 0x2f, 0x47, 0xb6, 0x18, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x8c, 0x63, 0x22,
	0xf1, 0x01, 0x00, 0x00,
}