// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/conversion_action_status.proto

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

// Possible statuses of a conversion action.
type ConversionActionStatusEnum_ConversionActionStatus int32

const (
	// Not specified.
	ConversionActionStatusEnum_UNSPECIFIED ConversionActionStatusEnum_ConversionActionStatus = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionStatusEnum_UNKNOWN ConversionActionStatusEnum_ConversionActionStatus = 1
	// Conversions will be recorded.
	ConversionActionStatusEnum_ENABLED ConversionActionStatusEnum_ConversionActionStatus = 2
	// Conversions will not be recorded.
	ConversionActionStatusEnum_REMOVED ConversionActionStatusEnum_ConversionActionStatus = 3
	// Conversions will not be recorded and the conversion action will not
	// appear in the UI.
	ConversionActionStatusEnum_HIDDEN ConversionActionStatusEnum_ConversionActionStatus = 4
)

var ConversionActionStatusEnum_ConversionActionStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
	4: "HIDDEN",
}

var ConversionActionStatusEnum_ConversionActionStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
	"HIDDEN":      4,
}

func (x ConversionActionStatusEnum_ConversionActionStatus) String() string {
	return proto.EnumName(ConversionActionStatusEnum_ConversionActionStatus_name, int32(x))
}

func (ConversionActionStatusEnum_ConversionActionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8fa0bbddeff25dcf, []int{0, 0}
}

// Container for enum describing possible statuses of a conversion action.
type ConversionActionStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionStatusEnum) Reset()         { *m = ConversionActionStatusEnum{} }
func (m *ConversionActionStatusEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionStatusEnum) ProtoMessage()    {}
func (*ConversionActionStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa0bbddeff25dcf, []int{0}
}

func (m *ConversionActionStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionStatusEnum.Unmarshal(m, b)
}
func (m *ConversionActionStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionStatusEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionStatusEnum.Merge(m, src)
}
func (m *ConversionActionStatusEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionStatusEnum.Size(m)
}
func (m *ConversionActionStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ConversionActionStatusEnum_ConversionActionStatus", ConversionActionStatusEnum_ConversionActionStatus_name, ConversionActionStatusEnum_ConversionActionStatus_value)
	proto.RegisterType((*ConversionActionStatusEnum)(nil), "google.ads.googleads.v3.enums.ConversionActionStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/conversion_action_status.proto", fileDescriptor_8fa0bbddeff25dcf)
}

var fileDescriptor_8fa0bbddeff25dcf = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0xa4, 0x29, 0x2a, 0x92, 0x3b, 0x10, 0x65, 0x60, 0x28, 0x74, 0x68, 0x3f, 0xc0, 0x1e, 0xbc,
	0x19, 0x16, 0xa7, 0x31, 0xa5, 0x02, 0xdc, 0x8a, 0xaa, 0x41, 0x42, 0x91, 0x90, 0x69, 0x22, 0xab,
	0x52, 0x6b, 0x57, 0x71, 0x92, 0x81, 0xcf, 0x61, 0xe4, 0x53, 0xf8, 0x14, 0x26, 0x3e, 0x01, 0xd9,
	0x69, 0x32, 0x15, 0x16, 0xfb, 0x9e, 0xef, 0xdd, 0xe9, 0x7c, 0xe0, 0x46, 0x6a, 0x2d, 0xb7, 0x19,
	0x12, 0xa9, 0x41, 0x35, 0xb4, 0xa8, 0xc2, 0x28, 0x53, 0xe5, 0xce, 0xa0, 0xb5, 0x56, 0x55, 0x96,
	0x9b, 0x8d, 0x56, 0xaf, 0x62, 0x5d, 0xd8, 0xcb, 0x14, 0xa2, 0x28, 0x0d, 0xdc, 0xe7, 0xba, 0xd0,
	0xc1, 0xb0, 0x96, 0x40, 0x91, 0x1a, 0xd8, 0xaa, 0x61, 0x85, 0xa1, 0x53, 0x0f, 0xae, 0x1a, 0xf3,
	0xfd, 0x06, 0x09, 0xa5, 0x74, 0x21, 0xac, 0xc5, 0x41, 0x3c, 0x7e, 0x07, 0x83, 0x49, 0x6b, 0x4f,
	0x9d, 0xfb, 0xd2, 0x99, 0x33, 0x55, 0xee, 0xc6, 0x09, 0xb8, 0x38, 0xce, 0x06, 0xe7, 0xa0, 0xbf,
	0xe2, 0xcb, 0x05, 0x9b, 0xcc, 0x6e, 0x67, 0x2c, 0xf2, 0x4f, 0x82, 0x3e, 0x38, 0x5b, 0xf1, 0x7b,
	0x3e, 0x7f, 0xe6, 0x7e, 0xc7, 0x0e, 0x8c, 0xd3, 0xf0, 0x81, 0x45, 0xbe, 0x67, 0x87, 0x27, 0xf6,
	0x38, 0x8f, 0x59, 0xe4, 0x77, 0x03, 0x00, 0x7a, 0x77, 0xb3, 0x28, 0x62, 0xdc, 0x3f, 0x0d, 0x7f,
	0x3a, 0x60, 0xb4, 0xd6, 0x3b, 0xf8, 0x6f, 0xfe, 0xf0, 0xf2, 0x78, 0x82, 0x85, 0x8d, 0xbf, 0xe8,
	0xbc, 0x84, 0x07, 0xb5, 0xd4, 0x5b, 0xa1, 0x24, 0xd4, 0xb9, 0x44, 0x32, 0x53, 0xee, 0x73, 0x4d,
	0x97, 0xfb, 0x8d, 0xf9, 0xa3, 0xda, 0x6b, 0x77, 0x7e, 0x78, 0xdd, 0x29, 0xa5, 0x9f, 0xde, 0x70,
	0x5a, 0x5b, 0xd1, 0xd4, 0xc0, 0x1a, 0x5a, 0x14, 0x63, 0x68, 0xbb, 0x30, 0x5f, 0x0d, 0x9f, 0xd0,
	0xd4, 0x24, 0x2d, 0x9f, 0xc4, 0x38, 0x71, 0xfc, 0xb7, 0x37, 0xaa, 0x1f, 0x09, 0xa1, 0xa9, 0x21,
	0xa4, 0xdd, 0x20, 0x24, 0xc6, 0x84, 0xb8, 0x9d, 0xb7, 0x9e, 0x0b, 0x86, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x8b, 0xdd, 0x69, 0xf2, 0x01, 0x00, 0x00,
}
