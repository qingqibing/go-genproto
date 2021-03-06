// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/remarketing_action.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A remarketing action. A snippet of JavaScript code that will collect the
// product id and the type of page people visited (product page, shopping cart
// page, purchase page, general site visit) on an advertiser's website.
type RemarketingAction struct {
	// Immutable. The resource name of the remarketing action.
	// Remarketing action resource names have the form:
	//
	// `customers/{customer_id}/remarketingActions/{remarketing_action_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. Id of the remarketing action.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the remarketing action.
	//
	// This field is required and should not be empty when creating new
	// remarketing actions.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. The snippets used for tracking remarketing actions.
	TagSnippets          []*common.TagSnippet `protobuf:"bytes,4,rep,name=tag_snippets,json=tagSnippets,proto3" json:"tag_snippets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RemarketingAction) Reset()         { *m = RemarketingAction{} }
func (m *RemarketingAction) String() string { return proto.CompactTextString(m) }
func (*RemarketingAction) ProtoMessage()    {}
func (*RemarketingAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_62ce490472884744, []int{0}
}

func (m *RemarketingAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemarketingAction.Unmarshal(m, b)
}
func (m *RemarketingAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemarketingAction.Marshal(b, m, deterministic)
}
func (m *RemarketingAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemarketingAction.Merge(m, src)
}
func (m *RemarketingAction) XXX_Size() int {
	return xxx_messageInfo_RemarketingAction.Size(m)
}
func (m *RemarketingAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RemarketingAction.DiscardUnknown(m)
}

var xxx_messageInfo_RemarketingAction proto.InternalMessageInfo

func (m *RemarketingAction) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *RemarketingAction) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *RemarketingAction) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RemarketingAction) GetTagSnippets() []*common.TagSnippet {
	if m != nil {
		return m.TagSnippets
	}
	return nil
}

func init() {
	proto.RegisterType((*RemarketingAction)(nil), "google.ads.googleads.v1.resources.RemarketingAction")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/remarketing_action.proto", fileDescriptor_62ce490472884744)
}

var fileDescriptor_62ce490472884744 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8a, 0x13, 0x31,
	0x1c, 0xc7, 0xe9, 0xcc, 0x2a, 0x38, 0x5d, 0x0f, 0xce, 0x41, 0xea, 0xba, 0x68, 0x57, 0x58, 0x28,
	0x8b, 0x24, 0x9d, 0x2a, 0x1e, 0xa2, 0x97, 0xf4, 0xb2, 0xe8, 0x41, 0x96, 0x59, 0xa9, 0x20, 0x85,
	0x92, 0xce, 0x64, 0x63, 0xb4, 0x49, 0x86, 0x24, 0xad, 0x87, 0x65, 0x4f, 0x3e, 0x82, 0x6f, 0xe0,
	0xd1, 0x47, 0xf1, 0x29, 0xf6, 0xbc, 0x8f, 0xe0, 0x49, 0x3a, 0x99, 0xa4, 0x85, 0xa1, 0xca, 0xde,
	0xbe, 0xed, 0xef, 0xfb, 0xf9, 0xfd, 0xcb, 0x6f, 0x12, 0xc4, 0x94, 0x62, 0x0b, 0x0a, 0x49, 0x69,
	0xa0, 0x93, 0x6b, 0xb5, 0xca, 0xa0, 0xa6, 0x46, 0x2d, 0x75, 0x41, 0x0d, 0xd4, 0x54, 0x10, 0xfd,
	0x95, 0x5a, 0x2e, 0xd9, 0x8c, 0x14, 0x96, 0x2b, 0x09, 0x2a, 0xad, 0xac, 0x4a, 0x8f, 0x1c, 0x00,
	0x48, 0x69, 0x40, 0x60, 0xc1, 0x2a, 0x03, 0x81, 0x3d, 0x18, 0xee, 0x4a, 0x5f, 0x28, 0x21, 0x94,
	0x84, 0x96, 0xb0, 0x99, 0x91, 0xbc, 0xaa, 0xa8, 0x75, 0x49, 0x0f, 0x9e, 0x7a, 0xa2, 0xe2, 0xf0,
	0x82, 0xd3, 0x45, 0x39, 0x9b, 0xd3, 0xcf, 0x64, 0xc5, 0x95, 0x6e, 0x0c, 0x8f, 0xb6, 0x0c, 0xbe,
	0x50, 0x13, 0x7a, 0xd2, 0x84, 0xea, 0x5f, 0xf3, 0xe5, 0x05, 0xfc, 0xa6, 0x49, 0x55, 0x51, 0x6d,
	0x9a, 0xf8, 0xe1, 0x16, 0x4a, 0xa4, 0x54, 0x96, 0xac, 0xa7, 0x69, 0xa2, 0xcf, 0x7e, 0xc4, 0xc9,
	0x83, 0x7c, 0x33, 0x2b, 0xae, 0x47, 0x4d, 0x3f, 0x26, 0xf7, 0x7d, 0x95, 0x99, 0x24, 0x82, 0xf6,
	0x3a, 0xfd, 0xce, 0xe0, 0xde, 0x78, 0x74, 0x8d, 0xef, 0xfc, 0xc1, 0xcf, 0x93, 0x93, 0xcd, 0xe0,
	0x8d, 0xaa, 0xb8, 0x01, 0x85, 0x12, 0xb0, 0x95, 0x2a, 0xdf, 0xf7, 0x89, 0xde, 0x13, 0x41, 0xd3,
	0x61, 0x12, 0xf1, 0xb2, 0x17, 0xf5, 0x3b, 0x83, 0xee, 0xe8, 0x71, 0x03, 0x03, 0xdf, 0x39, 0x78,
	0x2b, 0xed, 0xab, 0x97, 0x13, 0xb2, 0x58, 0xd2, 0x71, 0x7c, 0x8d, 0xe3, 0x3c, 0xe2, 0x65, 0x3a,
	0x4c, 0xf6, 0xea, 0x0e, 0xe2, 0x9a, 0x39, 0x6c, 0x31, 0xe7, 0x56, 0x73, 0xc9, 0x6a, 0x28, 0xaf,
	0x9d, 0x69, 0x9e, 0xec, 0x6f, 0x6d, 0xd8, 0xf4, 0xf6, 0xfa, 0xf1, 0xa0, 0x3b, 0x3a, 0x01, 0xbb,
	0x1e, 0xce, 0xbd, 0x0a, 0xf8, 0x40, 0xd8, 0xb9, 0x43, 0x5c, 0xf1, 0xae, 0x0d, 0x7f, 0x18, 0x24,
	0x6e, 0xf0, 0x97, 0xdb, 0x8c, 0x9d, 0xbe, 0x29, 0x96, 0xc6, 0x2a, 0x41, 0xb5, 0x81, 0x97, 0x5e,
	0x5e, 0x6d, 0x5f, 0x95, 0xf3, 0x19, 0x78, 0xd9, 0xbe, 0xb4, 0xab, 0xf1, 0xf7, 0x28, 0x39, 0x2e,
	0x94, 0x00, 0xff, 0xbd, 0xb5, 0xf1, 0xc3, 0x56, 0xe9, 0xb3, 0xf5, 0x66, 0xce, 0x3a, 0x9f, 0xde,
	0x35, 0x30, 0x53, 0x0b, 0x22, 0x19, 0x50, 0x9a, 0x41, 0x46, 0x65, 0xbd, 0x37, 0xb8, 0x69, 0xff,
	0x1f, 0xdf, 0xc0, 0xeb, 0xa0, 0x7e, 0x46, 0xf1, 0x29, 0xc6, 0xbf, 0xa2, 0xa3, 0x53, 0x97, 0x12,
	0x97, 0x06, 0x38, 0xb9, 0x56, 0x93, 0x0c, 0xe4, 0xde, 0xf9, 0xdb, 0x7b, 0xa6, 0xb8, 0x34, 0xd3,
	0xe0, 0x99, 0x4e, 0xb2, 0x69, 0xf0, 0xdc, 0x44, 0xc7, 0x2e, 0x80, 0x10, 0x2e, 0x0d, 0x42, 0xc1,
	0x85, 0xd0, 0x24, 0x43, 0x28, 0xf8, 0xe6, 0x77, 0xeb, 0x66, 0x5f, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0xe3, 0x5b, 0xe7, 0xaf, 0x03, 0x00, 0x00,
}
