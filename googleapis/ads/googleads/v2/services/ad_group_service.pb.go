// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Request message for [AdGroupService.GetAdGroup][google.ads.googleads.v2.services.AdGroupService.GetAdGroup].
type GetAdGroupRequest struct {
	// Required. The resource name of the ad group to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupRequest) Reset()         { *m = GetAdGroupRequest{} }
func (m *GetAdGroupRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupRequest) ProtoMessage()    {}
func (*GetAdGroupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2551d73f06de7113, []int{0}
}

func (m *GetAdGroupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupRequest.Unmarshal(m, b)
}
func (m *GetAdGroupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupRequest.Merge(m, src)
}
func (m *GetAdGroupRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupRequest.Size(m)
}
func (m *GetAdGroupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupRequest proto.InternalMessageInfo

func (m *GetAdGroupRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupService.MutateAdGroups][google.ads.googleads.v2.services.AdGroupService.MutateAdGroups].
type MutateAdGroupsRequest struct {
	// Required. The ID of the customer whose ad groups are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The list of operations to perform on individual ad groups.
	Operations []*AdGroupOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupsRequest) Reset()         { *m = MutateAdGroupsRequest{} }
func (m *MutateAdGroupsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupsRequest) ProtoMessage()    {}
func (*MutateAdGroupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2551d73f06de7113, []int{1}
}

func (m *MutateAdGroupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupsRequest.Merge(m, src)
}
func (m *MutateAdGroupsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupsRequest.Size(m)
}
func (m *MutateAdGroupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupsRequest proto.InternalMessageInfo

func (m *MutateAdGroupsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupsRequest) GetOperations() []*AdGroupOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group.
type AdGroupOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupOperation_Create
	//	*AdGroupOperation_Update
	//	*AdGroupOperation_Remove
	Operation            isAdGroupOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AdGroupOperation) Reset()         { *m = AdGroupOperation{} }
func (m *AdGroupOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupOperation) ProtoMessage()    {}
func (*AdGroupOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2551d73f06de7113, []int{2}
}

func (m *AdGroupOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupOperation.Unmarshal(m, b)
}
func (m *AdGroupOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupOperation.Merge(m, src)
}
func (m *AdGroupOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupOperation.Size(m)
}
func (m *AdGroupOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupOperation proto.InternalMessageInfo

func (m *AdGroupOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdGroupOperation_Operation interface {
	isAdGroupOperation_Operation()
}

type AdGroupOperation_Create struct {
	Create *resources.AdGroup `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupOperation_Update struct {
	Update *resources.AdGroup `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupOperation_Create) isAdGroupOperation_Operation() {}

func (*AdGroupOperation_Update) isAdGroupOperation_Operation() {}

func (*AdGroupOperation_Remove) isAdGroupOperation_Operation() {}

func (m *AdGroupOperation) GetOperation() isAdGroupOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupOperation) GetCreate() *resources.AdGroup {
	if x, ok := m.GetOperation().(*AdGroupOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupOperation) GetUpdate() *resources.AdGroup {
	if x, ok := m.GetOperation().(*AdGroupOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupOperation_Create)(nil),
		(*AdGroupOperation_Update)(nil),
		(*AdGroupOperation_Remove)(nil),
	}
}

// Response message for an ad group mutate.
type MutateAdGroupsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MutateAdGroupsResponse) Reset()         { *m = MutateAdGroupsResponse{} }
func (m *MutateAdGroupsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupsResponse) ProtoMessage()    {}
func (*MutateAdGroupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2551d73f06de7113, []int{3}
}

func (m *MutateAdGroupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupsResponse.Merge(m, src)
}
func (m *MutateAdGroupsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupsResponse.Size(m)
}
func (m *MutateAdGroupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupsResponse proto.InternalMessageInfo

func (m *MutateAdGroupsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupsResponse) GetResults() []*MutateAdGroupResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad group mutate.
type MutateAdGroupResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupResult) Reset()         { *m = MutateAdGroupResult{} }
func (m *MutateAdGroupResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupResult) ProtoMessage()    {}
func (*MutateAdGroupResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2551d73f06de7113, []int{4}
}

func (m *MutateAdGroupResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupResult.Unmarshal(m, b)
}
func (m *MutateAdGroupResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupResult.Merge(m, src)
}
func (m *MutateAdGroupResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupResult.Size(m)
}
func (m *MutateAdGroupResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupResult proto.InternalMessageInfo

func (m *MutateAdGroupResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupRequest")
	proto.RegisterType((*MutateAdGroupsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupsRequest")
	proto.RegisterType((*AdGroupOperation)(nil), "google.ads.googleads.v2.services.AdGroupOperation")
	proto.RegisterType((*MutateAdGroupsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupsResponse")
	proto.RegisterType((*MutateAdGroupResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_service.proto", fileDescriptor_2551d73f06de7113)
}

var fileDescriptor_2551d73f06de7113 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x4e, 0xdb, 0x4a,
	0x14, 0xbe, 0x76, 0xae, 0xb8, 0x97, 0x09, 0x70, 0x6f, 0x07, 0x41, 0xd3, 0xb4, 0x52, 0x23, 0x17,
	0xa9, 0x51, 0x44, 0xed, 0xca, 0xa8, 0xa2, 0x32, 0x62, 0xe1, 0xa8, 0x25, 0x74, 0x41, 0xa1, 0x46,
	0xa2, 0x52, 0x15, 0xc9, 0x9a, 0xd8, 0x43, 0x6a, 0x61, 0x7b, 0xdc, 0x19, 0x3b, 0x12, 0x42, 0x6c,
	0x78, 0x85, 0xbe, 0x41, 0x97, 0xdd, 0xb1, 0xed, 0x23, 0xb0, 0xad, 0xd4, 0x05, 0x2b, 0x16, 0xac,
	0xfa, 0x08, 0xac, 0x2a, 0x7b, 0x66, 0xf2, 0x57, 0xa2, 0x88, 0xee, 0x4e, 0xe6, 0x7c, 0xdf, 0x77,
	0xce, 0x99, 0x6f, 0x8e, 0x03, 0xd6, 0xbb, 0x84, 0x74, 0x43, 0x6c, 0x20, 0x9f, 0x19, 0x3c, 0xcc,
	0xa3, 0x9e, 0x69, 0x30, 0x4c, 0x7b, 0x81, 0x87, 0x99, 0x81, 0x7c, 0xb7, 0x4b, 0x49, 0x96, 0xb8,
	0xe2, 0x44, 0x4f, 0x28, 0x49, 0x09, 0xac, 0x71, 0xb4, 0x8e, 0x7c, 0xa6, 0xf7, 0x89, 0x7a, 0xcf,
	0xd4, 0x25, 0xb1, 0xfa, 0x7c, 0x92, 0x34, 0xc5, 0x8c, 0x64, 0x74, 0x58, 0x9b, 0x6b, 0x56, 0x1f,
	0x49, 0x46, 0x12, 0x18, 0x28, 0x8e, 0x49, 0x8a, 0xd2, 0x80, 0xc4, 0x4c, 0x64, 0xef, 0x0f, 0x65,
	0xbd, 0x30, 0xc0, 0x71, 0x2a, 0x12, 0x8f, 0x87, 0x12, 0x87, 0x01, 0x0e, 0x7d, 0xb7, 0x83, 0x3f,
	0xa2, 0x5e, 0x40, 0xa8, 0x00, 0x3c, 0x18, 0x02, 0xc8, 0xe2, 0x22, 0x25, 0xc6, 0x30, 0x8a, 0x5f,
	0x9d, 0xec, 0x50, 0x08, 0x44, 0x88, 0x1d, 0x8d, 0x95, 0xa5, 0x89, 0x67, 0xb0, 0x14, 0xa5, 0x99,
	0xe8, 0x47, 0xeb, 0x80, 0x7b, 0x2d, 0x9c, 0xda, 0x7e, 0x2b, 0x9f, 0xc0, 0xc1, 0x9f, 0x32, 0xcc,
	0x52, 0xb8, 0x03, 0xe6, 0x65, 0x05, 0x37, 0x46, 0x11, 0xae, 0x28, 0x35, 0xa5, 0x3e, 0xdb, 0xac,
	0x5f, 0xd9, 0xea, 0x8d, 0xad, 0x81, 0xda, 0xe0, 0xaa, 0x44, 0x94, 0x04, 0x4c, 0xf7, 0x48, 0x64,
	0x48, 0x9d, 0x39, 0x49, 0x7f, 0x8b, 0x22, 0xac, 0x5d, 0x2b, 0x60, 0x69, 0x27, 0x4b, 0x51, 0x8a,
	0x45, 0x9e, 0xc9, 0x42, 0x2b, 0xa0, 0xec, 0x65, 0x2c, 0x25, 0x11, 0xa6, 0x6e, 0xe0, 0x8b, 0x32,
	0xa5, 0x2b, 0x5b, 0x75, 0x80, 0x3c, 0x7f, 0xe3, 0xc3, 0xf7, 0x00, 0x90, 0x04, 0x53, 0x7e, 0x8f,
	0x15, 0xb5, 0x56, 0xaa, 0x97, 0x4d, 0x53, 0x9f, 0x66, 0x9d, 0x2e, 0x8a, 0xed, 0x4a, 0xaa, 0x10,
	0x1e, 0x48, 0xc1, 0xa7, 0xe0, 0xbf, 0x04, 0xd1, 0x34, 0x40, 0xa1, 0x7b, 0x88, 0x82, 0x30, 0xa3,
	0xb8, 0x52, 0xaa, 0x29, 0xf5, 0x7f, 0x9d, 0x05, 0x71, 0xbc, 0xc5, 0x4f, 0xe1, 0x13, 0x30, 0xdf,
	0x43, 0x61, 0xe0, 0xa3, 0x14, 0xbb, 0x24, 0x0e, 0x8f, 0x2b, 0x7f, 0x17, 0xb0, 0x39, 0x79, 0xb8,
	0x1b, 0x87, 0xc7, 0xda, 0x99, 0x0a, 0xfe, 0x1f, 0xaf, 0x09, 0x37, 0x40, 0x39, 0x4b, 0x0a, 0x5e,
	0xee, 0x46, 0xc1, 0x2b, 0x9b, 0x55, 0xd9, 0xbc, 0x34, 0x4c, 0xdf, 0xca, 0x0d, 0xdb, 0x41, 0xec,
	0xc8, 0x01, 0x1c, 0x9e, 0xc7, 0xf0, 0x15, 0x98, 0xf1, 0x28, 0x46, 0x29, 0x37, 0xa0, 0x6c, 0x36,
	0x26, 0x0e, 0xdd, 0x7f, 0x8d, 0x72, 0xea, 0xed, 0xbf, 0x1c, 0xc1, 0xcd, 0x55, 0xb8, 0x66, 0x45,
	0xfd, 0x13, 0x15, 0xce, 0x85, 0x15, 0x30, 0x43, 0x71, 0x44, 0x7a, 0xfc, 0x8a, 0x66, 0xf3, 0x0c,
	0xff, 0xdd, 0x2c, 0x83, 0xd9, 0xfe, 0x9d, 0x6a, 0xe7, 0x0a, 0x58, 0x1e, 0xf7, 0x9a, 0x25, 0x24,
	0x66, 0x18, 0x6e, 0x81, 0xa5, 0xb1, 0xdb, 0x76, 0x31, 0xa5, 0x84, 0x16, 0x82, 0x65, 0x13, 0xca,
	0xb6, 0x68, 0xe2, 0xe9, 0xfb, 0xc5, 0x1b, 0x75, 0x16, 0x47, 0x7d, 0x78, 0x9d, 0xc3, 0xe1, 0x2e,
	0xf8, 0x87, 0x62, 0x96, 0x85, 0xa9, 0x7c, 0x0b, 0x2f, 0xa6, 0xbf, 0x85, 0x91, 0x96, 0x9c, 0x82,
	0xed, 0x48, 0x15, 0xcd, 0x02, 0x8b, 0xb7, 0xe4, 0x73, 0xd3, 0x6f, 0xd9, 0x82, 0xd1, 0xb7, 0x6d,
	0x7e, 0x2b, 0x81, 0x05, 0x41, 0xdb, 0xe7, 0xc5, 0xe0, 0xb9, 0x02, 0xc0, 0x60, 0xa7, 0xe0, 0xda,
	0xf4, 0xee, 0x7e, 0xdb, 0xc0, 0xea, 0x1d, 0x3c, 0xd2, 0x9a, 0x97, 0xf6, 0x68, 0xa3, 0x67, 0xdf,
	0xaf, 0x3f, 0xab, 0xab, 0xb0, 0x91, 0x7f, 0xa6, 0x4e, 0x46, 0x32, 0x9b, 0x72, 0xab, 0x98, 0xd1,
	0x30, 0x90, 0x30, 0xc8, 0x68, 0x9c, 0xc2, 0x1f, 0x0a, 0x58, 0x18, 0xb5, 0x0d, 0xae, 0xdf, 0xf1,
	0x56, 0xe5, 0x52, 0x57, 0x5f, 0xde, 0x9d, 0xc8, 0x5f, 0x88, 0xf6, 0xee, 0xd2, 0x5e, 0x1e, 0xfa,
	0x1e, 0xac, 0x0e, 0x56, 0xb5, 0x18, 0xc9, 0xd4, 0x9e, 0xe5, 0x23, 0x0d, 0x66, 0x38, 0x19, 0x02,
	0x6f, 0x36, 0x4e, 0xfb, 0x13, 0x59, 0x51, 0xa1, 0x6f, 0x29, 0x8d, 0xea, 0xc3, 0x0b, 0xbb, 0x32,
	0xe9, 0x73, 0xd5, 0xbc, 0x51, 0xc0, 0x8a, 0x47, 0xa2, 0xa9, 0xfd, 0x36, 0x17, 0x47, 0x2d, 0xde,
	0xcb, 0xd7, 0x76, 0x4f, 0xf9, 0xb0, 0x2d, 0x88, 0x5d, 0x12, 0xa2, 0xb8, 0xab, 0x13, 0xda, 0x35,
	0xba, 0x38, 0x2e, 0x96, 0xda, 0x18, 0x94, 0x9a, 0xfc, 0xb7, 0xb4, 0x21, 0x83, 0x2f, 0x6a, 0xa9,
	0x65, 0xdb, 0x5f, 0xd5, 0x5a, 0x8b, 0x0b, 0xda, 0x3e, 0xd3, 0x79, 0x98, 0x47, 0x07, 0xa6, 0x2e,
	0x0a, 0xb3, 0x0b, 0x09, 0x69, 0xdb, 0x3e, 0x6b, 0xf7, 0x21, 0xed, 0x03, 0xb3, 0x2d, 0x21, 0x3f,
	0xd5, 0x15, 0x7e, 0x6e, 0x59, 0xb6, 0xcf, 0x2c, 0xab, 0x0f, 0xb2, 0xac, 0x03, 0xd3, 0xb2, 0x24,
	0xac, 0x33, 0x53, 0xf4, 0xb9, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x46, 0xfe, 0x07, 0x6b, 0x3d,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdGroupServiceClient is the client API for AdGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupServiceClient interface {
	// Returns the requested ad group in full detail.
	GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error)
}

type adGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdGroupServiceClient(cc grpc.ClientConnInterface) AdGroupServiceClient {
	return &adGroupServiceClient{cc}
}

func (c *adGroupServiceClient) GetAdGroup(ctx context.Context, in *GetAdGroupRequest, opts ...grpc.CallOption) (*resources.AdGroup, error) {
	out := new(resources.AdGroup)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupService/GetAdGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupServiceClient) MutateAdGroups(ctx context.Context, in *MutateAdGroupsRequest, opts ...grpc.CallOption) (*MutateAdGroupsResponse, error) {
	out := new(MutateAdGroupsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupService/MutateAdGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupServiceServer is the server API for AdGroupService service.
type AdGroupServiceServer interface {
	// Returns the requested ad group in full detail.
	GetAdGroup(context.Context, *GetAdGroupRequest) (*resources.AdGroup, error)
	// Creates, updates, or removes ad groups. Operation statuses are returned.
	MutateAdGroups(context.Context, *MutateAdGroupsRequest) (*MutateAdGroupsResponse, error)
}

// UnimplementedAdGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupServiceServer struct {
}

func (*UnimplementedAdGroupServiceServer) GetAdGroup(ctx context.Context, req *GetAdGroupRequest) (*resources.AdGroup, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroup not implemented")
}
func (*UnimplementedAdGroupServiceServer) MutateAdGroups(ctx context.Context, req *MutateAdGroupsRequest) (*MutateAdGroupsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroups not implemented")
}

func RegisterAdGroupServiceServer(s *grpc.Server, srv AdGroupServiceServer) {
	s.RegisterService(&_AdGroupService_serviceDesc, srv)
}

func _AdGroupService_GetAdGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupService/GetAdGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).GetAdGroup(ctx, req.(*GetAdGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupService_MutateAdGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupService/MutateAdGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupServiceServer).MutateAdGroups(ctx, req.(*MutateAdGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupService",
	HandlerType: (*AdGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroup",
			Handler:    _AdGroupService_GetAdGroup_Handler,
		},
		{
			MethodName: "MutateAdGroups",
			Handler:    _AdGroupService_MutateAdGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_service.proto",
}
