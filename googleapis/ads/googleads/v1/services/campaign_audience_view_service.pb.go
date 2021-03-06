// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/campaign_audience_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [CampaignAudienceViewService.GetCampaignAudienceView][google.ads.googleads.v1.services.CampaignAudienceViewService.GetCampaignAudienceView].
type GetCampaignAudienceViewRequest struct {
	// Required. The resource name of the campaign audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignAudienceViewRequest) Reset()         { *m = GetCampaignAudienceViewRequest{} }
func (m *GetCampaignAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignAudienceViewRequest) ProtoMessage()    {}
func (*GetCampaignAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae1606ecaa0347f1, []int{0}
}

func (m *GetCampaignAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetCampaignAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignAudienceViewRequest.Merge(m, src)
}
func (m *GetCampaignAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Size(m)
}
func (m *GetCampaignAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignAudienceViewRequest proto.InternalMessageInfo

func (m *GetCampaignAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignAudienceViewRequest)(nil), "google.ads.googleads.v1.services.GetCampaignAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/campaign_audience_view_service.proto", fileDescriptor_ae1606ecaa0347f1)
}

var fileDescriptor_ae1606ecaa0347f1 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xb1, 0x8a, 0xd5, 0x40,
	0x14, 0x25, 0x11, 0x04, 0x83, 0x36, 0x69, 0x76, 0xcd, 0x8a, 0x86, 0x65, 0x0b, 0x59, 0x70, 0x86,
	0x28, 0xb2, 0x38, 0xa2, 0x38, 0x4f, 0xe4, 0x59, 0xe9, 0xa2, 0x90, 0x62, 0x09, 0x84, 0xd9, 0xe4,
	0x1a, 0x07, 0x92, 0x99, 0x98, 0x49, 0xb2, 0x85, 0xda, 0xd8, 0xf8, 0x01, 0xfe, 0xc1, 0x96, 0x7e,
	0xca, 0xb6, 0x76, 0x56, 0x16, 0xda, 0xf8, 0x09, 0x56, 0x92, 0x9d, 0x99, 0x6c, 0x16, 0x5e, 0xde,
	0xeb, 0x0e, 0xef, 0x9c, 0x7b, 0xee, 0x9d, 0x73, 0x5e, 0xbc, 0x17, 0x85, 0x94, 0x45, 0x09, 0x98,
	0xe5, 0x0a, 0x6b, 0x38, 0xa0, 0x3e, 0xc2, 0x0a, 0x9a, 0x9e, 0x67, 0xa0, 0x70, 0xc6, 0xaa, 0x9a,
	0xf1, 0x42, 0xa4, 0xac, 0xcb, 0x39, 0x88, 0x0c, 0xd2, 0x9e, 0xc3, 0x49, 0x6a, 0x78, 0x54, 0x37,
	0xb2, 0x95, 0x7e, 0xa8, 0x67, 0x11, 0xcb, 0x15, 0x1a, 0x6d, 0x50, 0x1f, 0x21, 0x6b, 0x13, 0x3c,
	0x9d, 0x5b, 0xd4, 0x80, 0x92, 0x5d, 0x33, 0xbf, 0x49, 0x6f, 0x08, 0x6e, 0xd9, 0xf9, 0x9a, 0x63,
	0x26, 0x84, 0x6c, 0x59, 0xcb, 0xa5, 0x50, 0x86, 0xdd, 0x9a, 0xb0, 0x59, 0xc9, 0x41, 0xb4, 0x86,
	0xb8, 0x33, 0x21, 0xde, 0x71, 0x28, 0xf3, 0xf4, 0x18, 0xde, 0xb3, 0x9e, 0xcb, 0xc6, 0x08, 0x6e,
	0x4e, 0x04, 0xf6, 0x14, 0x4d, 0xed, 0x7e, 0xf2, 0x6e, 0x2f, 0xa1, 0x7d, 0x6e, 0xae, 0xa2, 0xe6,
	0xa8, 0x98, 0xc3, 0xc9, 0x1b, 0xf8, 0xd0, 0x81, 0x6a, 0xfd, 0x23, 0xef, 0x86, 0x9d, 0x49, 0x05,
	0xab, 0x60, 0xdb, 0x09, 0x9d, 0xbb, 0xd7, 0x16, 0x0f, 0x7f, 0x51, 0xf7, 0x1f, 0xc5, 0xde, 0xbd,
	0x8b, 0x28, 0x0c, 0xaa, 0xb9, 0x42, 0x99, 0xac, 0xf0, 0x4a, 0xd3, 0xeb, 0xd6, 0xeb, 0x15, 0xab,
	0xe0, 0xfe, 0xa9, 0xeb, 0xed, 0xac, 0x92, 0xbd, 0xd5, 0x89, 0xfa, 0x7f, 0x1c, 0x6f, 0x6b, 0xe6,
	0x3c, 0xff, 0x19, 0xda, 0xd4, 0x07, 0x5a, 0xff, 0xb2, 0xe0, 0x60, 0xd6, 0x61, 0xec, 0x0b, 0xad,
	0x9a, 0xdf, 0x7d, 0xfd, 0x93, 0x5e, 0xce, 0xe4, 0xcb, 0x8f, 0xdf, 0xdf, 0xdc, 0x47, 0xfe, 0xc1,
	0xd0, 0xf5, 0xc7, 0x4b, 0xcc, 0x93, 0xac, 0x53, 0xad, 0xac, 0xa0, 0x51, 0x78, 0x7f, 0x2c, 0x7f,
	0x6a, 0xa6, 0xf0, 0xfe, 0xe7, 0x60, 0xe7, 0x8c, 0x6e, 0xcf, 0xe5, 0xb8, 0xf8, 0xea, 0x7a, 0x7b,
	0x99, 0xac, 0x36, 0x3e, 0x77, 0x11, 0xae, 0x89, 0xf2, 0x70, 0x68, 0xfb, 0xd0, 0x39, 0x7a, 0x69,
	0x5c, 0x0a, 0x59, 0x32, 0x51, 0x20, 0xd9, 0x14, 0xb8, 0x00, 0x71, 0xfe, 0x5f, 0xc0, 0x17, 0x7b,
	0xe7, 0x3f, 0x95, 0xc7, 0x16, 0x9c, 0xba, 0x57, 0x96, 0x94, 0x7e, 0x77, 0xc3, 0xa5, 0x36, 0xa4,
	0xb9, 0x42, 0x1a, 0x0e, 0x28, 0x8e, 0x90, 0x59, 0xac, 0xce, 0xac, 0x24, 0xa1, 0xb9, 0x4a, 0x46,
	0x49, 0x12, 0x47, 0x89, 0x95, 0xfc, 0x75, 0xf7, 0xf4, 0xef, 0x84, 0xd0, 0x5c, 0x11, 0x32, 0x8a,
	0x08, 0x89, 0x23, 0x42, 0xac, 0xec, 0xf8, 0xea, 0xf9, 0x9d, 0x0f, 0xfe, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x62, 0x9a, 0x0f, 0x1f, 0xd1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CampaignAudienceViewServiceClient is the client API for CampaignAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignAudienceViewServiceClient interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error)
}

type campaignAudienceViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCampaignAudienceViewServiceClient(cc grpc.ClientConnInterface) CampaignAudienceViewServiceClient {
	return &campaignAudienceViewServiceClient{cc}
}

func (c *campaignAudienceViewServiceClient) GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error) {
	out := new(resources.CampaignAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAudienceViewServiceServer is the server API for CampaignAudienceViewService service.
type CampaignAudienceViewServiceServer interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error)
}

// UnimplementedCampaignAudienceViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCampaignAudienceViewServiceServer struct {
}

func (*UnimplementedCampaignAudienceViewServiceServer) GetCampaignAudienceView(ctx context.Context, req *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCampaignAudienceView not implemented")
}

func RegisterCampaignAudienceViewServiceServer(s *grpc.Server, srv CampaignAudienceViewServiceServer) {
	s.RegisterService(&_CampaignAudienceViewService_serviceDesc, srv)
}

func _CampaignAudienceViewService_GetCampaignAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/campaign_audience_view_service.proto",
}
