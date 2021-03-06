// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/location_view.proto

package resources

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

// A location view summarizes the performance of campaigns by
// Location criteria.
type LocationView struct {
	// Output only. The resource name of the location view.
	// Location view resource names have the form:
	//
	// `customers/{customer_id}/locationViews/{campaign_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationView) Reset()         { *m = LocationView{} }
func (m *LocationView) String() string { return proto.CompactTextString(m) }
func (*LocationView) ProtoMessage()    {}
func (*LocationView) Descriptor() ([]byte, []int) {
	return fileDescriptor_8294fbe2fa5282fe, []int{0}
}

func (m *LocationView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationView.Unmarshal(m, b)
}
func (m *LocationView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationView.Marshal(b, m, deterministic)
}
func (m *LocationView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationView.Merge(m, src)
}
func (m *LocationView) XXX_Size() int {
	return xxx_messageInfo_LocationView.Size(m)
}
func (m *LocationView) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationView.DiscardUnknown(m)
}

var xxx_messageInfo_LocationView proto.InternalMessageInfo

func (m *LocationView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationView)(nil), "google.ads.googleads.v3.resources.LocationView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/location_view.proto", fileDescriptor_8294fbe2fa5282fe)
}

var fileDescriptor_8294fbe2fa5282fe = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x1c, 0xc5, 0x49, 0x0b, 0x82, 0xa1, 0x0e, 0x76, 0xd2, 0x22, 0x68, 0x85, 0xa2, 0x8b, 0x77, 0x60,
	0x70, 0x39, 0xa7, 0xeb, 0x52, 0x10, 0x91, 0xd2, 0x21, 0x83, 0x06, 0xcb, 0x35, 0x39, 0xe3, 0x41,
	0x92, 0x7f, 0xb9, 0x4b, 0xd3, 0xa1, 0xf4, 0xcb, 0x38, 0xfa, 0x31, 0x1c, 0xfd, 0x14, 0xce, 0xfd,
	0x08, 0x0e, 0x22, 0xc9, 0xe5, 0xae, 0x71, 0x51, 0xb7, 0x07, 0xff, 0xdf, 0x7b, 0xf7, 0x78, 0xe7,
	0x5e, 0xc5, 0x00, 0x71, 0xc2, 0x31, 0x8b, 0x14, 0xd6, 0xb2, 0x54, 0x85, 0x87, 0x25, 0x57, 0xb0,
	0x90, 0x21, 0x57, 0x38, 0x81, 0x90, 0xe5, 0x02, 0xb2, 0x69, 0x21, 0xf8, 0x12, 0xcd, 0x25, 0xe4,
	0xd0, 0xed, 0x6b, 0x16, 0xb1, 0x48, 0x21, 0x6b, 0x43, 0x85, 0x87, 0xac, 0xad, 0x77, 0x6c, 0x92,
	0xe7, 0x02, 0x3f, 0x09, 0x9e, 0x44, 0xd3, 0x19, 0x7f, 0x66, 0x85, 0x00, 0xa9, 0x33, 0x7a, 0x87,
	0x0d, 0xc0, 0xd8, 0xea, 0xd3, 0x51, 0xe3, 0xc4, 0xb2, 0x0c, 0xf2, 0xaa, 0x80, 0xd2, 0xd7, 0xd3,
	0x37, 0xc7, 0xed, 0xdc, 0xd6, 0xa5, 0x7c, 0xc1, 0x97, 0xdd, 0x89, 0xbb, 0x67, 0x02, 0xa6, 0x19,
	0x4b, 0xf9, 0x81, 0x73, 0xe2, 0x9c, 0xef, 0x0e, 0x2f, 0x3e, 0x68, 0xfb, 0x93, 0x9e, 0xb9, 0x83,
	0x6d, 0xc3, 0x5a, 0xcd, 0x85, 0x42, 0x21, 0xa4, 0xb8, 0x99, 0x32, 0xe9, 0x98, 0x8c, 0x3b, 0x96,
	0x72, 0xf2, 0xb8, 0xa1, 0x0f, 0xff, 0x74, 0x76, 0x2f, 0xc3, 0x85, 0xca, 0x21, 0xe5, 0x52, 0xe1,
	0x95, 0x91, 0x6b, 0xbb, 0x5b, 0x89, 0x28, 0xbc, 0xfa, 0x31, 0xe3, 0x7a, 0xf8, 0xe5, 0xb8, 0x83,
	0x10, 0x52, 0xf4, 0xe7, 0x90, 0xc3, 0xfd, 0xe6, 0x5b, 0xe3, 0x72, 0x81, 0xb1, 0x73, 0x7f, 0x53,
	0xfb, 0x62, 0x48, 0x58, 0x16, 0x23, 0x90, 0x31, 0x8e, 0x79, 0x56, 0xed, 0x83, 0xb7, 0x55, 0x7f,
	0xf9, 0xd6, 0x6b, 0xab, 0x5e, 0x5a, 0xed, 0x11, 0xa5, 0xaf, 0xad, 0xfe, 0x48, 0x47, 0xd2, 0x48,
	0x21, 0x2d, 0x4b, 0xe5, 0x7b, 0x68, 0x62, 0xc8, 0x77, 0xc3, 0x04, 0x34, 0x52, 0x81, 0x65, 0x02,
	0xdf, 0x0b, 0x2c, 0xb3, 0x69, 0x0d, 0xf4, 0x81, 0x10, 0x1a, 0x29, 0x42, 0x2c, 0x45, 0x88, 0xef,
	0x11, 0x62, 0xb9, 0xd9, 0x4e, 0x55, 0xd6, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x10, 0x0f, 0xb3,
	0x06, 0x82, 0x02, 0x00, 0x00,
}
