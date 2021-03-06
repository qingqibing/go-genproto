// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/asset_error.proto

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

// Enum describing possible asset errors.
type AssetErrorEnum_AssetError int32

const (
	// Enum unspecified.
	AssetErrorEnum_UNSPECIFIED AssetErrorEnum_AssetError = 0
	// The received error code is not known in this version.
	AssetErrorEnum_UNKNOWN AssetErrorEnum_AssetError = 1
	// The customer is not whitelisted for this asset type.
	AssetErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 2
	// Assets are duplicated across operations.
	AssetErrorEnum_DUPLICATE_ASSET AssetErrorEnum_AssetError = 3
	// The asset name is duplicated, either across operations or with an
	// existing asset.
	AssetErrorEnum_DUPLICATE_ASSET_NAME AssetErrorEnum_AssetError = 4
	// The Asset.asset_data oneof is empty.
	AssetErrorEnum_ASSET_DATA_IS_MISSING AssetErrorEnum_AssetError = 5
	// The asset has a name which is different from an existing duplicate that
	// represents the same content.
	AssetErrorEnum_CANNOT_MODIFY_ASSET_NAME AssetErrorEnum_AssetError = 6
)

var AssetErrorEnum_AssetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE",
	3: "DUPLICATE_ASSET",
	4: "DUPLICATE_ASSET_NAME",
	5: "ASSET_DATA_IS_MISSING",
	6: "CANNOT_MODIFY_ASSET_NAME",
}

var AssetErrorEnum_AssetError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE": 2,
	"DUPLICATE_ASSET":                         3,
	"DUPLICATE_ASSET_NAME":                    4,
	"ASSET_DATA_IS_MISSING":                   5,
	"CANNOT_MODIFY_ASSET_NAME":                6,
}

func (x AssetErrorEnum_AssetError) String() string {
	return proto.EnumName(AssetErrorEnum_AssetError_name, int32(x))
}

func (AssetErrorEnum_AssetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a6381eb5f1114281, []int{0, 0}
}

// Container for enum describing possible asset errors.
type AssetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetErrorEnum) Reset()         { *m = AssetErrorEnum{} }
func (m *AssetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AssetErrorEnum) ProtoMessage()    {}
func (*AssetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6381eb5f1114281, []int{0}
}

func (m *AssetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetErrorEnum.Unmarshal(m, b)
}
func (m *AssetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetErrorEnum.Marshal(b, m, deterministic)
}
func (m *AssetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetErrorEnum.Merge(m, src)
}
func (m *AssetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AssetErrorEnum.Size(m)
}
func (m *AssetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AssetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.AssetErrorEnum_AssetError", AssetErrorEnum_AssetError_name, AssetErrorEnum_AssetError_value)
	proto.RegisterType((*AssetErrorEnum)(nil), "google.ads.googleads.v3.errors.AssetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/asset_error.proto", fileDescriptor_a6381eb5f1114281)
}

var fileDescriptor_a6381eb5f1114281 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x8b, 0xd4, 0x30,
	0x14, 0xc7, 0x6d, 0x57, 0x57, 0xc8, 0x82, 0x53, 0xa2, 0xc2, 0x2a, 0xcb, 0x1e, 0x7a, 0xf1, 0x20,
	0xa4, 0x42, 0x6f, 0xf1, 0x94, 0x6d, 0x33, 0x63, 0x70, 0x9b, 0x96, 0x49, 0x3a, 0xc3, 0x48, 0x21,
	0x54, 0x5b, 0xca, 0xc0, 0x4c, 0x33, 0x34, 0x75, 0x3e, 0x90, 0x47, 0xbf, 0x84, 0x77, 0x8f, 0x7e,
	0x8c, 0xf9, 0x14, 0xd2, 0xc6, 0x99, 0x11, 0xc1, 0x3d, 0xe5, 0xff, 0x5e, 0x7e, 0xff, 0xf7, 0x92,
	0xf7, 0xc0, 0xbb, 0x46, 0xeb, 0x66, 0x53, 0x07, 0x65, 0x65, 0x02, 0x2b, 0x07, 0xb5, 0x0f, 0x83,
	0xba, 0xeb, 0x74, 0x67, 0x82, 0xd2, 0x98, 0xba, 0x57, 0x63, 0x80, 0x76, 0x9d, 0xee, 0x35, 0xbc,
	0xb5, 0x18, 0x2a, 0x2b, 0x83, 0x4e, 0x0e, 0xb4, 0x0f, 0x91, 0x75, 0xbc, 0xbe, 0x39, 0x56, 0xdc,
	0xad, 0x83, 0xb2, 0x6d, 0x75, 0x5f, 0xf6, 0x6b, 0xdd, 0x1a, 0xeb, 0xf6, 0x7f, 0x39, 0xe0, 0x19,
	0x19, 0x6a, 0xd2, 0x81, 0xa6, 0xed, 0xd7, 0xad, 0xff, 0xc3, 0x01, 0xe0, 0x9c, 0x82, 0x13, 0x70,
	0x95, 0x73, 0x91, 0xd1, 0x88, 0x4d, 0x19, 0x8d, 0xbd, 0x47, 0xf0, 0x0a, 0x3c, 0xcd, 0xf9, 0x47,
	0x9e, 0x2e, 0xb9, 0xe7, 0xc0, 0xb7, 0xe0, 0x4d, 0x94, 0x0b, 0x99, 0x26, 0x74, 0xae, 0x78, 0x2a,
	0xd5, 0xf2, 0x03, 0x93, 0xf4, 0x9e, 0x09, 0x49, 0x63, 0x35, 0x4d, 0xe7, 0x8a, 0x08, 0x41, 0xa5,
	0x92, 0xab, 0x8c, 0x7a, 0x2e, 0x7c, 0x0e, 0x26, 0x71, 0x9e, 0xdd, 0xb3, 0x88, 0x48, 0x6a, 0x6f,
	0xbc, 0x0b, 0x78, 0x0d, 0x5e, 0xfc, 0x93, 0x54, 0x9c, 0x24, 0xd4, 0x7b, 0x0c, 0x5f, 0x81, 0x97,
	0x36, 0x8e, 0x89, 0x24, 0x8a, 0x09, 0x95, 0x30, 0x21, 0x18, 0x9f, 0x79, 0x4f, 0xe0, 0x0d, 0xb8,
	0x8e, 0x08, 0x1f, 0x1a, 0x26, 0x69, 0xcc, 0xa6, 0xab, 0xbf, 0x8d, 0x97, 0x77, 0x07, 0x07, 0xf8,
	0x5f, 0xf4, 0x16, 0x3d, 0x3c, 0x99, 0xbb, 0xc9, 0xf9, 0x97, 0xd9, 0x30, 0x8c, 0xcc, 0xf9, 0x14,
	0xff, 0xb1, 0x34, 0x7a, 0x53, 0xb6, 0x0d, 0xd2, 0x5d, 0x13, 0x34, 0x75, 0x3b, 0x8e, 0xea, 0xb8,
	0x8e, 0xdd, 0xda, 0xfc, 0x6f, 0x3b, 0xef, 0xed, 0xf1, 0xcd, 0xbd, 0x98, 0x11, 0xf2, 0xdd, 0xbd,
	0x9d, 0xd9, 0x62, 0xa4, 0x32, 0xc8, 0xca, 0x41, 0x2d, 0x42, 0x34, 0xb6, 0x34, 0x3f, 0x8f, 0x40,
	0x41, 0x2a, 0x53, 0x9c, 0x80, 0x62, 0x11, 0x16, 0x16, 0x38, 0xb8, 0xbe, 0xcd, 0x62, 0x4c, 0x2a,
	0x83, 0xf1, 0x09, 0xc1, 0x78, 0x11, 0x62, 0x6c, 0xa1, 0xcf, 0x97, 0xe3, 0xeb, 0xc2, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xed, 0x6e, 0xea, 0x21, 0x3a, 0x02, 0x00, 0x00,
}
