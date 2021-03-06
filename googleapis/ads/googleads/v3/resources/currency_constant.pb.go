// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/currency_constant.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A currency constant.
type CurrencyConstant struct {
	// Output only. The resource name of the currency constant.
	// Currency constant resource names have the form:
	//
	// `currencyConstants/{currency_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. ISO 4217 three-letter currency code, e.g. "USD"
	Code *wrappers.StringValue `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// Output only. Full English name of the currency.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Standard symbol for describing this currency, e.g. '$' for US Dollars.
	Symbol *wrappers.StringValue `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Output only. The billable unit for this currency. Billed amounts should be multiples of
	// this value.
	BillableUnitMicros   *wrappers.Int64Value `protobuf:"bytes,5,opt,name=billable_unit_micros,json=billableUnitMicros,proto3" json:"billable_unit_micros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CurrencyConstant) Reset()         { *m = CurrencyConstant{} }
func (m *CurrencyConstant) String() string { return proto.CompactTextString(m) }
func (*CurrencyConstant) ProtoMessage()    {}
func (*CurrencyConstant) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f14d3d22f3007ad, []int{0}
}

func (m *CurrencyConstant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyConstant.Unmarshal(m, b)
}
func (m *CurrencyConstant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyConstant.Marshal(b, m, deterministic)
}
func (m *CurrencyConstant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyConstant.Merge(m, src)
}
func (m *CurrencyConstant) XXX_Size() int {
	return xxx_messageInfo_CurrencyConstant.Size(m)
}
func (m *CurrencyConstant) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyConstant.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyConstant proto.InternalMessageInfo

func (m *CurrencyConstant) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CurrencyConstant) GetCode() *wrappers.StringValue {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *CurrencyConstant) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CurrencyConstant) GetSymbol() *wrappers.StringValue {
	if m != nil {
		return m.Symbol
	}
	return nil
}

func (m *CurrencyConstant) GetBillableUnitMicros() *wrappers.Int64Value {
	if m != nil {
		return m.BillableUnitMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*CurrencyConstant)(nil), "google.ads.googleads.v3.resources.CurrencyConstant")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/currency_constant.proto", fileDescriptor_6f14d3d22f3007ad)
}

var fileDescriptor_6f14d3d22f3007ad = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x6a, 0xd5, 0x30,
	0x1c, 0xc6, 0x39, 0xed, 0x1c, 0x58, 0x15, 0xa4, 0x28, 0x1c, 0xe7, 0xd0, 0x4d, 0x38, 0x30, 0x11,
	0x12, 0xb4, 0x2a, 0x2c, 0x5e, 0xe5, 0xec, 0x62, 0x28, 0x28, 0xa3, 0xb2, 0x5e, 0xc8, 0x81, 0x92,
	0xa6, 0x59, 0x0d, 0xb4, 0xf9, 0x97, 0x24, 0x3d, 0x32, 0xc4, 0x1b, 0x1f, 0xc5, 0x4b, 0x1f, 0xc2,
	0x07, 0xf0, 0x29, 0x76, 0xbd, 0x47, 0xf0, 0x4a, 0xda, 0xa6, 0xdd, 0xe1, 0x0c, 0x74, 0xbb, 0xfb,
	0xda, 0xff, 0xf7, 0xfb, 0xf2, 0x25, 0xfc, 0x83, 0xfd, 0x02, 0xa0, 0x28, 0x05, 0x66, 0xb9, 0xc1,
	0xbd, 0x6c, 0xd5, 0x32, 0xc2, 0x5a, 0x18, 0x68, 0x34, 0x17, 0x06, 0xf3, 0x46, 0x6b, 0xa1, 0xf8,
	0x69, 0xca, 0x41, 0x19, 0xcb, 0x94, 0x45, 0xb5, 0x06, 0x0b, 0xe1, 0x6e, 0xef, 0x47, 0x2c, 0x37,
	0x68, 0x44, 0xd1, 0x32, 0x42, 0x23, 0xba, 0xf5, 0x78, 0x48, 0xaf, 0x25, 0x3e, 0x91, 0xa2, 0xcc,
	0xd3, 0x4c, 0x7c, 0x66, 0x4b, 0x09, 0xba, 0xcf, 0xd8, 0x7a, 0xb0, 0x62, 0x18, 0x30, 0x37, 0x7a,
	0xe4, 0x46, 0xdd, 0x57, 0xd6, 0x9c, 0xe0, 0x2f, 0x9a, 0xd5, 0xb5, 0xd0, 0xc6, 0xcd, 0xb7, 0x57,
	0x50, 0xa6, 0x14, 0x58, 0x66, 0x25, 0x28, 0x37, 0x7d, 0xf2, 0xcb, 0x0f, 0xee, 0x1e, 0xb8, 0xe2,
	0x07, 0xae, 0x77, 0x98, 0x04, 0x77, 0x86, 0x43, 0x52, 0xc5, 0x2a, 0x31, 0x9d, 0xec, 0x4c, 0xf6,
	0x6e, 0xce, 0x9f, 0x9f, 0x51, 0xff, 0x0f, 0x7d, 0x16, 0x3c, 0xbd, 0xb8, 0x85, 0x53, 0xb5, 0x34,
	0x88, 0x43, 0x85, 0xd7, 0x93, 0xe2, 0xdb, 0x43, 0xce, 0x07, 0x56, 0x89, 0xf0, 0x55, 0xb0, 0xc1,
	0x21, 0x17, 0x53, 0x6f, 0x67, 0xb2, 0x77, 0xeb, 0xc5, 0xb6, 0xa3, 0xd1, 0xd0, 0x1c, 0x7d, 0xb4,
	0x5a, 0xaa, 0x22, 0x61, 0x65, 0x23, 0xe6, 0xfe, 0x19, 0xf5, 0xe3, 0xce, 0xde, 0x62, 0x5d, 0x0b,
	0xff, 0xca, 0x58, 0x6b, 0x0f, 0xf7, 0x83, 0x4d, 0x73, 0x5a, 0x65, 0x50, 0x4e, 0x37, 0xae, 0x0a,
	0x3a, 0x20, 0x8c, 0x83, 0x7b, 0x99, 0x2c, 0x4b, 0x96, 0x95, 0x22, 0x6d, 0x94, 0xb4, 0x69, 0x25,
	0xb9, 0x06, 0x33, 0xbd, 0xd1, 0x05, 0x3d, 0xbc, 0x14, 0xf4, 0x56, 0xd9, 0xd7, 0x2f, 0x57, 0x72,
	0xc2, 0x81, 0x3e, 0x56, 0xd2, 0xbe, 0xef, 0x58, 0x72, 0x7c, 0x4e, 0xe3, 0x6b, 0x3c, 0x5d, 0x38,
	0xe3, 0x6b, 0x7f, 0x0c, 0xfe, 0x7a, 0x69, 0xc5, 0xbe, 0xcd, 0xbf, 0x7b, 0xc1, 0x8c, 0x43, 0x85,
	0xfe, 0xbb, 0x64, 0xf3, 0xfb, 0xeb, 0x47, 0x1c, 0xb5, 0xf5, 0x8f, 0x26, 0x9f, 0xde, 0x39, 0xb6,
	0x80, 0x92, 0xa9, 0x02, 0x81, 0x2e, 0x70, 0x21, 0x54, 0x77, 0x39, 0x7c, 0xd1, 0xf2, 0x1f, 0xab,
	0xff, 0x66, 0x54, 0x3f, 0x3c, 0xff, 0x90, 0xd2, 0x9f, 0xde, 0xee, 0x61, 0x1f, 0x49, 0x73, 0x83,
	0x7a, 0xd9, 0xaa, 0x24, 0x42, 0xf1, 0xe0, 0xfc, 0x3d, 0x78, 0x16, 0x34, 0x37, 0x8b, 0xd1, 0xb3,
	0x48, 0xa2, 0xc5, 0xe8, 0x39, 0xf7, 0x66, 0xfd, 0x80, 0x10, 0x9a, 0x1b, 0x42, 0x46, 0x17, 0x21,
	0x49, 0x44, 0xc8, 0xe8, 0xcb, 0x36, 0xbb, 0xb2, 0xd1, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47,
	0x0c, 0x03, 0xe9, 0xa6, 0x03, 0x00, 0x00,
}
