// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/group_placement_view_service.proto

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

// Request message for [GroupPlacementViewService.GetGroupPlacementView][google.ads.googleads.v1.services.GroupPlacementViewService.GetGroupPlacementView].
type GetGroupPlacementViewRequest struct {
	// Required. The resource name of the Group Placement view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGroupPlacementViewRequest) Reset()         { *m = GetGroupPlacementViewRequest{} }
func (m *GetGroupPlacementViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetGroupPlacementViewRequest) ProtoMessage()    {}
func (*GetGroupPlacementViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_722a806b0135cef4, []int{0}
}

func (m *GetGroupPlacementViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Unmarshal(m, b)
}
func (m *GetGroupPlacementViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Marshal(b, m, deterministic)
}
func (m *GetGroupPlacementViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGroupPlacementViewRequest.Merge(m, src)
}
func (m *GetGroupPlacementViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetGroupPlacementViewRequest.Size(m)
}
func (m *GetGroupPlacementViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGroupPlacementViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGroupPlacementViewRequest proto.InternalMessageInfo

func (m *GetGroupPlacementViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetGroupPlacementViewRequest)(nil), "google.ads.googleads.v1.services.GetGroupPlacementViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/group_placement_view_service.proto", fileDescriptor_722a806b0135cef4)
}

var fileDescriptor_722a806b0135cef4 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6b, 0xd5, 0x40,
	0x14, 0xe5, 0x45, 0x10, 0x0c, 0xba, 0x09, 0x88, 0x6d, 0x2c, 0xfa, 0x28, 0x5d, 0x48, 0xc5, 0x19,
	0x62, 0x29, 0xc2, 0xf8, 0x01, 0xf3, 0x5c, 0xc4, 0x8d, 0xf2, 0x50, 0x08, 0x22, 0x81, 0x30, 0x4d,
	0xae, 0x71, 0x20, 0x99, 0x89, 0x33, 0x93, 0x54, 0x10, 0x37, 0x82, 0xbf, 0xc0, 0x8d, 0x6b, 0x97,
	0xfe, 0x94, 0x6e, 0xdd, 0x09, 0x82, 0x0b, 0x57, 0xfe, 0x04, 0x57, 0x92, 0x4e, 0x26, 0x7d, 0x8f,
	0x36, 0x76, 0x77, 0xc8, 0x39, 0xf7, 0xdc, 0x3b, 0xe7, 0x10, 0xff, 0x71, 0x29, 0x65, 0x59, 0x01,
	0x66, 0x85, 0xc6, 0x16, 0xf6, 0xa8, 0x8b, 0xb0, 0x06, 0xd5, 0xf1, 0x1c, 0x34, 0x2e, 0x95, 0x6c,
	0x9b, 0xac, 0xa9, 0x58, 0x0e, 0x35, 0x08, 0x93, 0x75, 0x1c, 0x0e, 0xb3, 0x81, 0x45, 0x8d, 0x92,
	0x46, 0x06, 0x73, 0x3b, 0x89, 0x58, 0xa1, 0xd1, 0x68, 0x82, 0xba, 0x08, 0x39, 0x93, 0xf0, 0xc1,
	0xd4, 0x1a, 0x05, 0x5a, 0xb6, 0x6a, 0x6a, 0x8f, 0xf5, 0x0f, 0xb7, 0xdc, 0x74, 0xc3, 0x31, 0x13,
	0x42, 0x1a, 0x66, 0xb8, 0x14, 0x7a, 0x60, 0xaf, 0xad, 0xb0, 0x79, 0xc5, 0x41, 0x98, 0x81, 0xb8,
	0xb9, 0x42, 0xbc, 0xe6, 0x50, 0x15, 0xd9, 0x01, 0xbc, 0x61, 0x1d, 0x97, 0x6a, 0x10, 0x6c, 0xae,
	0x08, 0xdc, 0x21, 0x96, 0xda, 0x7e, 0xe7, 0x6f, 0xc5, 0x60, 0xe2, 0xfe, 0xa6, 0xa5, 0x3b, 0x29,
	0xe1, 0x70, 0xf8, 0x1c, 0xde, 0xb6, 0xa0, 0x4d, 0xf0, 0xd2, 0xbf, 0xe2, 0x26, 0x32, 0xc1, 0x6a,
	0xd8, 0x98, 0xcd, 0x67, 0xb7, 0x2e, 0x2d, 0xf6, 0x7e, 0x51, 0xef, 0x2f, 0xbd, 0xe3, 0xdf, 0x3e,
	0x89, 0x61, 0x40, 0x0d, 0xd7, 0x28, 0x97, 0x35, 0x3e, 0xc3, 0xf2, 0xb2, 0x73, 0x7a, 0xc6, 0x6a,
	0xb8, 0xfb, 0xc5, 0xf3, 0x37, 0x4f, 0x8b, 0x5e, 0xd8, 0x24, 0x83, 0x9f, 0x33, 0xff, 0xea, 0x99,
	0x87, 0x05, 0x8f, 0xd0, 0x79, 0x2d, 0xa0, 0xff, 0xbd, 0x28, 0xdc, 0x9f, 0x9c, 0x1f, 0x3b, 0x42,
	0xa7, 0xa7, 0xb7, 0x9f, 0xfe, 0xa0, 0xeb, 0x49, 0x7c, 0xfc, 0xfe, 0xfb, 0xb3, 0x77, 0x2f, 0xd8,
	0xef, 0xdb, 0x7d, 0xbf, 0xc6, 0x3c, 0xcc, 0x5b, 0x6d, 0x64, 0x0d, 0x4a, 0xe3, 0x5d, 0x5b, 0xf7,
	0x9a, 0x95, 0xc6, 0xbb, 0x1f, 0xc2, 0xeb, 0x47, 0x74, 0x63, 0x2a, 0xbb, 0xc5, 0x27, 0xcf, 0xdf,
	0xc9, 0x65, 0x7d, 0xee, 0x43, 0x17, 0x37, 0x26, 0x03, 0x5c, 0xf6, 0xed, 0x2e, 0x67, 0xaf, 0x9e,
	0x0c, 0x1e, 0xa5, 0xac, 0x98, 0x28, 0x91, 0x54, 0x25, 0x2e, 0x41, 0x1c, 0x77, 0x8f, 0x4f, 0xb6,
	0x4e, 0xff, 0x16, 0xf7, 0x1d, 0xf8, 0xea, 0x5d, 0x88, 0x29, 0xfd, 0xe6, 0xcd, 0x63, 0x6b, 0x48,
	0x0b, 0x8d, 0x2c, 0xec, 0x51, 0x12, 0xa1, 0x61, 0xb1, 0x3e, 0x72, 0x92, 0x94, 0x16, 0x3a, 0x1d,
	0x25, 0x69, 0x12, 0xa5, 0x4e, 0xf2, 0xc7, 0xdb, 0xb1, 0xdf, 0x09, 0xa1, 0x85, 0x26, 0x64, 0x14,
	0x11, 0x92, 0x44, 0x84, 0x38, 0xd9, 0xc1, 0xc5, 0xe3, 0x3b, 0xf7, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x4d, 0x4d, 0x74, 0xc4, 0xbd, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GroupPlacementViewServiceClient is the client API for GroupPlacementViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GroupPlacementViewServiceClient interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error)
}

type groupPlacementViewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupPlacementViewServiceClient(cc grpc.ClientConnInterface) GroupPlacementViewServiceClient {
	return &groupPlacementViewServiceClient{cc}
}

func (c *groupPlacementViewServiceClient) GetGroupPlacementView(ctx context.Context, in *GetGroupPlacementViewRequest, opts ...grpc.CallOption) (*resources.GroupPlacementView, error) {
	out := new(resources.GroupPlacementView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.GroupPlacementViewService/GetGroupPlacementView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupPlacementViewServiceServer is the server API for GroupPlacementViewService service.
type GroupPlacementViewServiceServer interface {
	// Returns the requested Group Placement view in full detail.
	GetGroupPlacementView(context.Context, *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error)
}

// UnimplementedGroupPlacementViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGroupPlacementViewServiceServer struct {
}

func (*UnimplementedGroupPlacementViewServiceServer) GetGroupPlacementView(ctx context.Context, req *GetGroupPlacementViewRequest) (*resources.GroupPlacementView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupPlacementView not implemented")
}

func RegisterGroupPlacementViewServiceServer(s *grpc.Server, srv GroupPlacementViewServiceServer) {
	s.RegisterService(&_GroupPlacementViewService_serviceDesc, srv)
}

func _GroupPlacementViewService_GetGroupPlacementView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupPlacementViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.GroupPlacementViewService/GetGroupPlacementView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupPlacementViewServiceServer).GetGroupPlacementView(ctx, req.(*GetGroupPlacementViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupPlacementViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.GroupPlacementViewService",
	HandlerType: (*GroupPlacementViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupPlacementView",
			Handler:    _GroupPlacementViewService_GetGroupPlacementView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/group_placement_view_service.proto",
}
