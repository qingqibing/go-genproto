// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/keyword_plan_error.proto

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

// Enum describing possible errors from applying a keyword plan.
type KeywordPlanErrorEnum_KeywordPlanError int32

const (
	// Enum unspecified.
	KeywordPlanErrorEnum_UNSPECIFIED KeywordPlanErrorEnum_KeywordPlanError = 0
	// The received error code is not known in this version.
	KeywordPlanErrorEnum_UNKNOWN KeywordPlanErrorEnum_KeywordPlanError = 1
	// The plan's bid multiplier value is outside the valid range.
	KeywordPlanErrorEnum_BID_MULTIPLIER_OUT_OF_RANGE KeywordPlanErrorEnum_KeywordPlanError = 2
	// The plan's bid value is too high.
	KeywordPlanErrorEnum_BID_TOO_HIGH KeywordPlanErrorEnum_KeywordPlanError = 3
	// The plan's bid value is too low.
	KeywordPlanErrorEnum_BID_TOO_LOW KeywordPlanErrorEnum_KeywordPlanError = 4
	// The plan's cpc bid is not a multiple of the minimum billable unit.
	KeywordPlanErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS KeywordPlanErrorEnum_KeywordPlanError = 5
	// The plan's daily budget value is too low.
	KeywordPlanErrorEnum_DAILY_BUDGET_TOO_LOW KeywordPlanErrorEnum_KeywordPlanError = 6
	// The plan's daily budget is not a multiple of the minimum billable unit.
	KeywordPlanErrorEnum_DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS KeywordPlanErrorEnum_KeywordPlanError = 7
	// The input has an invalid value.
	KeywordPlanErrorEnum_INVALID_VALUE KeywordPlanErrorEnum_KeywordPlanError = 8
	// The plan has no keyword.
	KeywordPlanErrorEnum_KEYWORD_PLAN_HAS_NO_KEYWORDS KeywordPlanErrorEnum_KeywordPlanError = 9
	// The plan is not enabled and API cannot provide mutation, forecast or
	// stats.
	KeywordPlanErrorEnum_KEYWORD_PLAN_NOT_ENABLED KeywordPlanErrorEnum_KeywordPlanError = 10
	// The requested plan cannot be found for providing forecast or stats.
	KeywordPlanErrorEnum_KEYWORD_PLAN_NOT_FOUND KeywordPlanErrorEnum_KeywordPlanError = 11
	// The plan is missing a cpc bid.
	KeywordPlanErrorEnum_MISSING_BID KeywordPlanErrorEnum_KeywordPlanError = 13
	// The plan is missing required forecast_period field.
	KeywordPlanErrorEnum_MISSING_FORECAST_PERIOD KeywordPlanErrorEnum_KeywordPlanError = 14
	// The plan's forecast_period has invalid forecast date range.
	KeywordPlanErrorEnum_INVALID_FORECAST_DATE_RANGE KeywordPlanErrorEnum_KeywordPlanError = 15
	// The plan's name is invalid.
	KeywordPlanErrorEnum_INVALID_NAME KeywordPlanErrorEnum_KeywordPlanError = 16
)

var KeywordPlanErrorEnum_KeywordPlanError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BID_MULTIPLIER_OUT_OF_RANGE",
	3:  "BID_TOO_HIGH",
	4:  "BID_TOO_LOW",
	5:  "BID_TOO_MANY_FRACTIONAL_DIGITS",
	6:  "DAILY_BUDGET_TOO_LOW",
	7:  "DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS",
	8:  "INVALID_VALUE",
	9:  "KEYWORD_PLAN_HAS_NO_KEYWORDS",
	10: "KEYWORD_PLAN_NOT_ENABLED",
	11: "KEYWORD_PLAN_NOT_FOUND",
	13: "MISSING_BID",
	14: "MISSING_FORECAST_PERIOD",
	15: "INVALID_FORECAST_DATE_RANGE",
	16: "INVALID_NAME",
}

var KeywordPlanErrorEnum_KeywordPlanError_value = map[string]int32{
	"UNSPECIFIED":                             0,
	"UNKNOWN":                                 1,
	"BID_MULTIPLIER_OUT_OF_RANGE":             2,
	"BID_TOO_HIGH":                            3,
	"BID_TOO_LOW":                             4,
	"BID_TOO_MANY_FRACTIONAL_DIGITS":          5,
	"DAILY_BUDGET_TOO_LOW":                    6,
	"DAILY_BUDGET_TOO_MANY_FRACTIONAL_DIGITS": 7,
	"INVALID_VALUE":                           8,
	"KEYWORD_PLAN_HAS_NO_KEYWORDS":            9,
	"KEYWORD_PLAN_NOT_ENABLED":                10,
	"KEYWORD_PLAN_NOT_FOUND":                  11,
	"MISSING_BID":                             13,
	"MISSING_FORECAST_PERIOD":                 14,
	"INVALID_FORECAST_DATE_RANGE":             15,
	"INVALID_NAME":                            16,
}

func (x KeywordPlanErrorEnum_KeywordPlanError) String() string {
	return proto.EnumName(KeywordPlanErrorEnum_KeywordPlanError_name, int32(x))
}

func (KeywordPlanErrorEnum_KeywordPlanError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_93fd9cdbc08c3d07, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// resource (keyword plan, keyword plan campaign, keyword plan ad group or
// keyword plan keyword) or KeywordPlanService RPC.
type KeywordPlanErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanErrorEnum) Reset()         { *m = KeywordPlanErrorEnum{} }
func (m *KeywordPlanErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanErrorEnum) ProtoMessage()    {}
func (*KeywordPlanErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_93fd9cdbc08c3d07, []int{0}
}

func (m *KeywordPlanErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanErrorEnum.Merge(m, src)
}
func (m *KeywordPlanErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanErrorEnum.Size(m)
}
func (m *KeywordPlanErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.KeywordPlanErrorEnum_KeywordPlanError", KeywordPlanErrorEnum_KeywordPlanError_name, KeywordPlanErrorEnum_KeywordPlanError_value)
	proto.RegisterType((*KeywordPlanErrorEnum)(nil), "google.ads.googleads.v3.errors.KeywordPlanErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/keyword_plan_error.proto", fileDescriptor_93fd9cdbc08c3d07)
}

var fileDescriptor_93fd9cdbc08c3d07 = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x69, 0x0b, 0x1b, 0xb8, 0x8c, 0x19, 0x6b, 0xc0, 0xb4, 0x55, 0x05, 0xf5, 0xc2, 0x01,
	0x29, 0x39, 0xe4, 0x80, 0x14, 0x4e, 0x4e, 0xed, 0xa6, 0x56, 0x53, 0x3b, 0xca, 0x9f, 0x56, 0x45,
	0x95, 0xac, 0x40, 0xaa, 0xa8, 0xa2, 0x8b, 0xab, 0xa4, 0x0c, 0xf1, 0x75, 0x38, 0x72, 0xe1, 0x6b,
	0x20, 0x3e, 0x0a, 0x27, 0x3e, 0x02, 0x4a, 0xdd, 0x44, 0x62, 0xd3, 0x38, 0xe5, 0xd5, 0xf3, 0xfe,
	0x9e, 0x27, 0x8e, 0xf3, 0x80, 0xb7, 0x99, 0x52, 0xd9, 0x66, 0x65, 0x26, 0x69, 0x69, 0xea, 0xb1,
	0x9a, 0xae, 0x2d, 0x73, 0x55, 0x14, 0xaa, 0x28, 0xcd, 0x4f, 0xab, 0xaf, 0x5f, 0x54, 0x91, 0xca,
	0xed, 0x26, 0xc9, 0xe5, 0x5e, 0x33, 0xb6, 0x85, 0xda, 0x29, 0xd4, 0xd7, 0xb4, 0x91, 0xa4, 0xa5,
	0xd1, 0x18, 0x8d, 0x6b, 0xcb, 0xd0, 0xc6, 0x8b, 0x5e, 0x1d, 0xbc, 0x5d, 0x9b, 0x49, 0x9e, 0xab,
	0x5d, 0xb2, 0x5b, 0xab, 0xbc, 0xd4, 0xee, 0xc1, 0xcf, 0x0e, 0x38, 0x9b, 0xe8, 0x68, 0x7f, 0x93,
	0xe4, 0xb4, 0xf2, 0xd0, 0xfc, 0xf3, 0xd5, 0xe0, 0x47, 0x07, 0xc0, 0x9b, 0x0b, 0x74, 0x0a, 0xba,
	0x31, 0x0f, 0x7d, 0x3a, 0x64, 0x23, 0x46, 0x09, 0xbc, 0x87, 0xba, 0xe0, 0x38, 0xe6, 0x13, 0x2e,
	0xe6, 0x1c, 0xb6, 0xd0, 0x4b, 0x70, 0xe9, 0x30, 0x22, 0xa7, 0xb1, 0x17, 0x31, 0xdf, 0x63, 0x34,
	0x90, 0x22, 0x8e, 0xa4, 0x18, 0xc9, 0x00, 0x73, 0x97, 0xc2, 0x36, 0x82, 0xe0, 0x71, 0x05, 0x44,
	0x42, 0xc8, 0x31, 0x73, 0xc7, 0xb0, 0x53, 0x05, 0xd6, 0x8a, 0x27, 0xe6, 0xf0, 0x3e, 0x1a, 0x80,
	0x7e, 0x2d, 0x4c, 0x31, 0x5f, 0xc8, 0x51, 0x80, 0x87, 0x11, 0x13, 0x1c, 0x7b, 0x92, 0x30, 0x97,
	0x45, 0x21, 0x7c, 0x80, 0xce, 0xc1, 0x19, 0xc1, 0xcc, 0x5b, 0x48, 0x27, 0x26, 0x2e, 0x8d, 0x1a,
	0xf7, 0x11, 0x7a, 0x03, 0x5e, 0xdf, 0xda, 0xdc, 0x11, 0x73, 0x8c, 0x9e, 0x82, 0x13, 0xc6, 0x67,
	0xd8, 0x63, 0x44, 0xce, 0xb0, 0x17, 0x53, 0xf8, 0x10, 0xbd, 0x02, 0xbd, 0x09, 0x5d, 0xcc, 0x45,
	0x40, 0xa4, 0xef, 0x61, 0x2e, 0xc7, 0x38, 0x94, 0x5c, 0xc8, 0x83, 0x16, 0xc2, 0x47, 0xa8, 0x07,
	0xce, 0xff, 0x21, 0xb8, 0x88, 0x24, 0xe5, 0xd8, 0xf1, 0x28, 0x81, 0x00, 0x5d, 0x80, 0xe7, 0xb7,
	0xb6, 0x23, 0x11, 0x73, 0x02, 0xbb, 0xd5, 0xa7, 0x4e, 0x59, 0x18, 0x32, 0xee, 0x4a, 0x87, 0x11,
	0x78, 0x82, 0x2e, 0xc1, 0x8b, 0x5a, 0x18, 0x89, 0x80, 0x0e, 0x71, 0x18, 0x49, 0x9f, 0x06, 0x4c,
	0x10, 0xf8, 0xa4, 0xba, 0xcb, 0xfa, 0x70, 0xcd, 0x92, 0xe0, 0x88, 0x1e, 0xee, 0xf2, 0xb4, 0xba,
	0xcb, 0x1a, 0xe0, 0x78, 0x4a, 0x21, 0x74, 0xfe, 0xb4, 0xc0, 0xe0, 0xa3, 0xba, 0x32, 0xfe, 0xdf,
	0x07, 0xe7, 0xd9, 0xcd, 0xbf, 0xea, 0x57, 0x45, 0xf0, 0x5b, 0xef, 0xc9, 0xc1, 0x98, 0xa9, 0x4d,
	0x92, 0x67, 0x86, 0x2a, 0x32, 0x33, 0x5b, 0xe5, 0xfb, 0x9a, 0xd4, 0x8d, 0xdc, 0xae, 0xcb, 0xbb,
	0x0a, 0xfa, 0x4e, 0x3f, 0xbe, 0xb5, 0x3b, 0x2e, 0xc6, 0xdf, 0xdb, 0x7d, 0x57, 0x87, 0xe1, 0xb4,
	0x34, 0xf4, 0x58, 0x4d, 0x33, 0xcb, 0xd8, 0xbf, 0xb2, 0xfc, 0x55, 0x03, 0x4b, 0x9c, 0x96, 0xcb,
	0x06, 0x58, 0xce, 0xac, 0xa5, 0x06, 0x7e, 0xb7, 0x07, 0x5a, 0xb5, 0x6d, 0x9c, 0x96, 0xb6, 0xdd,
	0x20, 0xb6, 0x3d, 0xb3, 0x6c, 0x5b, 0x43, 0x1f, 0x8e, 0xf6, 0xa7, 0xb3, 0xfe, 0x06, 0x00, 0x00,
	0xff, 0xff, 0xec, 0x38, 0x0a, 0xed, 0x3d, 0x03, 0x00, 0x00,
}
