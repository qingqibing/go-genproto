// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/genomics/v1/datasets.proto
// DO NOT EDIT!

package google_genomics_v1 // import "google.golang.org/genproto/googleapis/genomics/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import google_iam_v11 "google.golang.org/genproto/googleapis/iam/v1"
import google_iam_v1 "google.golang.org/genproto/googleapis/iam/v1"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import google_protobuf2 "google.golang.org/genproto/protobuf"
import google_protobuf6 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A Dataset is a collection of genomic data.
//
// For more genomics resource definitions, see [Fundamentals of Google
// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
type Dataset struct {
	// The server-generated dataset ID, unique across all datasets.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// The Google Developers Console project ID that this dataset belongs to.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	// The dataset name.
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// The time this dataset was created, in seconds from the epoch.
	CreateTime *google_protobuf6.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
}

func (m *Dataset) Reset()                    { *m = Dataset{} }
func (m *Dataset) String() string            { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()               {}
func (*Dataset) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Dataset) GetCreateTime() *google_protobuf6.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

// The dataset list request.
type ListDatasetsRequest struct {
	// Required. The project to list datasets for.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
	// The maximum number of results to return in a single page. If unspecified,
	// defaults to 50. The maximum value is 1024.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// To get the next page of results, set this parameter to the value of
	// `nextPageToken` from the previous response.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *ListDatasetsRequest) Reset()                    { *m = ListDatasetsRequest{} }
func (m *ListDatasetsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDatasetsRequest) ProtoMessage()               {}
func (*ListDatasetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// The dataset list response.
type ListDatasetsResponse struct {
	// The list of matching Datasets.
	Datasets []*Dataset `protobuf:"bytes,1,rep,name=datasets" json:"datasets,omitempty"`
	// The continuation token, which is used to page through large result sets.
	// Provide this value in a subsequent request to return the next page of
	// results. This field will be empty if there aren't any additional results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *ListDatasetsResponse) Reset()                    { *m = ListDatasetsResponse{} }
func (m *ListDatasetsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDatasetsResponse) ProtoMessage()               {}
func (*ListDatasetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ListDatasetsResponse) GetDatasets() []*Dataset {
	if m != nil {
		return m.Datasets
	}
	return nil
}

type CreateDatasetRequest struct {
	// The dataset to be created. Must contain projectId and name.
	Dataset *Dataset `protobuf:"bytes,1,opt,name=dataset" json:"dataset,omitempty"`
}

func (m *CreateDatasetRequest) Reset()                    { *m = CreateDatasetRequest{} }
func (m *CreateDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateDatasetRequest) ProtoMessage()               {}
func (*CreateDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CreateDatasetRequest) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

type UpdateDatasetRequest struct {
	// The ID of the dataset to be updated.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
	// The new dataset data.
	Dataset *Dataset `protobuf:"bytes,2,opt,name=dataset" json:"dataset,omitempty"`
	// An optional mask specifying which fields to update. At this time, the only
	// mutable field is [name][google.genomics.v1.Dataset.name]. The only
	// acceptable value is "name". If unspecified, all mutable fields will be
	// updated.
	UpdateMask *google_protobuf2.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask" json:"update_mask,omitempty"`
}

func (m *UpdateDatasetRequest) Reset()                    { *m = UpdateDatasetRequest{} }
func (m *UpdateDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateDatasetRequest) ProtoMessage()               {}
func (*UpdateDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *UpdateDatasetRequest) GetDataset() *Dataset {
	if m != nil {
		return m.Dataset
	}
	return nil
}

func (m *UpdateDatasetRequest) GetUpdateMask() *google_protobuf2.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type DeleteDatasetRequest struct {
	// The ID of the dataset to be deleted.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
}

func (m *DeleteDatasetRequest) Reset()                    { *m = DeleteDatasetRequest{} }
func (m *DeleteDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDatasetRequest) ProtoMessage()               {}
func (*DeleteDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type UndeleteDatasetRequest struct {
	// The ID of the dataset to be undeleted.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
}

func (m *UndeleteDatasetRequest) Reset()                    { *m = UndeleteDatasetRequest{} }
func (m *UndeleteDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*UndeleteDatasetRequest) ProtoMessage()               {}
func (*UndeleteDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type GetDatasetRequest struct {
	// The ID of the dataset.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId" json:"dataset_id,omitempty"`
}

func (m *GetDatasetRequest) Reset()                    { *m = GetDatasetRequest{} }
func (m *GetDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetDatasetRequest) ProtoMessage()               {}
func (*GetDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func init() {
	proto.RegisterType((*Dataset)(nil), "google.genomics.v1.Dataset")
	proto.RegisterType((*ListDatasetsRequest)(nil), "google.genomics.v1.ListDatasetsRequest")
	proto.RegisterType((*ListDatasetsResponse)(nil), "google.genomics.v1.ListDatasetsResponse")
	proto.RegisterType((*CreateDatasetRequest)(nil), "google.genomics.v1.CreateDatasetRequest")
	proto.RegisterType((*UpdateDatasetRequest)(nil), "google.genomics.v1.UpdateDatasetRequest")
	proto.RegisterType((*DeleteDatasetRequest)(nil), "google.genomics.v1.DeleteDatasetRequest")
	proto.RegisterType((*UndeleteDatasetRequest)(nil), "google.genomics.v1.UndeleteDatasetRequest")
	proto.RegisterType((*GetDatasetRequest)(nil), "google.genomics.v1.GetDatasetRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for DatasetServiceV1 service

type DatasetServiceV1Client interface {
	// Lists datasets within a project.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error)
	// Creates a new dataset.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Gets a dataset by ID.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Updates a dataset.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// This method supports patch semantics.
	UpdateDataset(ctx context.Context, in *UpdateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Deletes a dataset and all of its contents (all read group sets,
	// reference sets, variant sets, call sets, annotation sets, etc.)
	// This is reversible (up to one week after the deletion) via
	// the
	// [datasets.undelete][google.genomics.v1.DatasetServiceV1.UndeleteDataset]
	// operation.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	// Undeletes a dataset by restoring a dataset which was deleted via this API.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// This operation is only possible for a week after the deletion occurred.
	UndeleteDataset(ctx context.Context, in *UndeleteDatasetRequest, opts ...grpc.CallOption) (*Dataset, error)
	// Sets the access control policy on the specified dataset. Replaces any
	// existing policy.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// See <a href="/iam/docs/managing-policies#setting_a_policy">Setting a
	// Policy</a> for more information.
	SetIamPolicy(ctx context.Context, in *google_iam_v11.SetIamPolicyRequest, opts ...grpc.CallOption) (*google_iam_v1.Policy, error)
	// Gets the access control policy for the dataset. This is empty if the
	// policy or resource does not exist.
	//
	// See <a href="/iam/docs/managing-policies#getting_a_policy">Getting a
	// Policy</a> for more information.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	GetIamPolicy(ctx context.Context, in *google_iam_v11.GetIamPolicyRequest, opts ...grpc.CallOption) (*google_iam_v1.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// See <a href="/iam/docs/managing-policies#testing_permissions">Testing
	// Permissions</a> for more information.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	TestIamPermissions(ctx context.Context, in *google_iam_v11.TestIamPermissionsRequest, opts ...grpc.CallOption) (*google_iam_v11.TestIamPermissionsResponse, error)
}

type datasetServiceV1Client struct {
	cc *grpc.ClientConn
}

func NewDatasetServiceV1Client(cc *grpc.ClientConn) DatasetServiceV1Client {
	return &datasetServiceV1Client{cc}
}

func (c *datasetServiceV1Client) ListDatasets(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error) {
	out := new(ListDatasetsResponse)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/ListDatasets", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/CreateDataset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/GetDataset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) UpdateDataset(ctx context.Context, in *UpdateDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/UpdateDataset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/DeleteDataset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) UndeleteDataset(ctx context.Context, in *UndeleteDatasetRequest, opts ...grpc.CallOption) (*Dataset, error) {
	out := new(Dataset)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/UndeleteDataset", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) SetIamPolicy(ctx context.Context, in *google_iam_v11.SetIamPolicyRequest, opts ...grpc.CallOption) (*google_iam_v1.Policy, error) {
	out := new(google_iam_v1.Policy)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/SetIamPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) GetIamPolicy(ctx context.Context, in *google_iam_v11.GetIamPolicyRequest, opts ...grpc.CallOption) (*google_iam_v1.Policy, error) {
	out := new(google_iam_v1.Policy)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/GetIamPolicy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceV1Client) TestIamPermissions(ctx context.Context, in *google_iam_v11.TestIamPermissionsRequest, opts ...grpc.CallOption) (*google_iam_v11.TestIamPermissionsResponse, error) {
	out := new(google_iam_v11.TestIamPermissionsResponse)
	err := grpc.Invoke(ctx, "/google.genomics.v1.DatasetServiceV1/TestIamPermissions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DatasetServiceV1 service

type DatasetServiceV1Server interface {
	// Lists datasets within a project.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	ListDatasets(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error)
	// Creates a new dataset.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	CreateDataset(context.Context, *CreateDatasetRequest) (*Dataset, error)
	// Gets a dataset by ID.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	GetDataset(context.Context, *GetDatasetRequest) (*Dataset, error)
	// Updates a dataset.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// This method supports patch semantics.
	UpdateDataset(context.Context, *UpdateDatasetRequest) (*Dataset, error)
	// Deletes a dataset and all of its contents (all read group sets,
	// reference sets, variant sets, call sets, annotation sets, etc.)
	// This is reversible (up to one week after the deletion) via
	// the
	// [datasets.undelete][google.genomics.v1.DatasetServiceV1.UndeleteDataset]
	// operation.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	DeleteDataset(context.Context, *DeleteDatasetRequest) (*google_protobuf1.Empty, error)
	// Undeletes a dataset by restoring a dataset which was deleted via this API.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// This operation is only possible for a week after the deletion occurred.
	UndeleteDataset(context.Context, *UndeleteDatasetRequest) (*Dataset, error)
	// Sets the access control policy on the specified dataset. Replaces any
	// existing policy.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	//
	// See <a href="/iam/docs/managing-policies#setting_a_policy">Setting a
	// Policy</a> for more information.
	SetIamPolicy(context.Context, *google_iam_v11.SetIamPolicyRequest) (*google_iam_v1.Policy, error)
	// Gets the access control policy for the dataset. This is empty if the
	// policy or resource does not exist.
	//
	// See <a href="/iam/docs/managing-policies#getting_a_policy">Getting a
	// Policy</a> for more information.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	GetIamPolicy(context.Context, *google_iam_v11.GetIamPolicyRequest) (*google_iam_v1.Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// See <a href="/iam/docs/managing-policies#testing_permissions">Testing
	// Permissions</a> for more information.
	//
	// For the definitions of datasets and other genomics resources, see
	// [Fundamentals of Google
	// Genomics](https://cloud.google.com/genomics/fundamentals-of-google-genomics)
	TestIamPermissions(context.Context, *google_iam_v11.TestIamPermissionsRequest) (*google_iam_v11.TestIamPermissionsResponse, error)
}

func RegisterDatasetServiceV1Server(s *grpc.Server, srv DatasetServiceV1Server) {
	s.RegisterService(&_DatasetServiceV1_serviceDesc, srv)
}

func _DatasetServiceV1_ListDatasets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).ListDatasets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/ListDatasets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).ListDatasets(ctx, req.(*ListDatasetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).CreateDataset(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_GetDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).GetDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/GetDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).GetDataset(ctx, req.(*GetDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_UpdateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).UpdateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/UpdateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).UpdateDataset(ctx, req.(*UpdateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_DeleteDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).DeleteDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/DeleteDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).DeleteDataset(ctx, req.(*DeleteDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_UndeleteDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UndeleteDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).UndeleteDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/UndeleteDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).UndeleteDataset(ctx, req.(*UndeleteDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_iam_v11.SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).SetIamPolicy(ctx, req.(*google_iam_v11.SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_iam_v11.GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).GetIamPolicy(ctx, req.(*google_iam_v11.GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetServiceV1_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_iam_v11.TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceV1Server).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.genomics.v1.DatasetServiceV1/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceV1Server).TestIamPermissions(ctx, req.(*google_iam_v11.TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DatasetServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.genomics.v1.DatasetServiceV1",
	HandlerType: (*DatasetServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDatasets",
			Handler:    _DatasetServiceV1_ListDatasets_Handler,
		},
		{
			MethodName: "CreateDataset",
			Handler:    _DatasetServiceV1_CreateDataset_Handler,
		},
		{
			MethodName: "GetDataset",
			Handler:    _DatasetServiceV1_GetDataset_Handler,
		},
		{
			MethodName: "UpdateDataset",
			Handler:    _DatasetServiceV1_UpdateDataset_Handler,
		},
		{
			MethodName: "DeleteDataset",
			Handler:    _DatasetServiceV1_DeleteDataset_Handler,
		},
		{
			MethodName: "UndeleteDataset",
			Handler:    _DatasetServiceV1_UndeleteDataset_Handler,
		},
		{
			MethodName: "SetIamPolicy",
			Handler:    _DatasetServiceV1_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _DatasetServiceV1_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _DatasetServiceV1_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/genomics/v1/datasets.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 804 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xce, 0x16, 0x14, 0x7a, 0xa0, 0xa0, 0x63, 0x45, 0x2c, 0x21, 0xd1, 0x89, 0x02, 0x56, 0xdc,
	0x4d, 0x8b, 0x84, 0x08, 0x21, 0x31, 0x88, 0x12, 0x12, 0x49, 0x9a, 0x02, 0xde, 0x36, 0xcb, 0x76,
	0x58, 0x57, 0xba, 0x3f, 0x76, 0xa6, 0x20, 0x22, 0x37, 0xdc, 0x79, 0xed, 0x03, 0x98, 0x78, 0xe7,
	0xf3, 0xf8, 0x0a, 0x3e, 0x84, 0x97, 0xce, 0xcc, 0xce, 0xb6, 0xdb, 0x76, 0x59, 0x8a, 0xf1, 0x66,
	0x19, 0xce, 0x7c, 0xe7, 0x7c, 0xdf, 0x99, 0xf3, 0xed, 0x6c, 0xe1, 0xa5, 0xed, 0xfb, 0x76, 0x83,
	0xe8, 0xb6, 0xdf, 0x30, 0x3d, 0x5b, 0xf7, 0x9b, 0xb6, 0x61, 0x13, 0x2f, 0x68, 0xfa, 0xcc, 0x37,
	0xc2, 0x2d, 0x33, 0x70, 0xa8, 0x88, 0xf9, 0xae, 0x63, 0x51, 0xe3, 0xb8, 0x64, 0xd4, 0x4d, 0x66,
	0x52, 0xc2, 0xa8, 0x2e, 0x51, 0x08, 0x45, 0x15, 0x14, 0x44, 0x3f, 0x2e, 0x15, 0xb6, 0x07, 0xab,
	0xca, 0x1f, 0x06, 0x25, 0xcd, 0x63, 0xc7, 0x22, 0x96, 0xef, 0x1d, 0x3a, 0xb6, 0x61, 0x7a, 0x9e,
	0xcf, 0x4c, 0xe6, 0xf8, 0x9e, 0x2a, 0x5f, 0x58, 0x1f, 0xac, 0x94, 0x63, 0xba, 0x42, 0x1b, 0xff,
	0x53, 0x0b, 0xfc, 0x86, 0x63, 0x9d, 0xaa, 0xf4, 0x17, 0xd7, 0x4a, 0xef, 0x4a, 0x5d, 0xb2, 0x1d,
	0xf6, 0xbe, 0x75, 0xa0, 0x5b, 0xbe, 0x6b, 0x84, 0xe9, 0x86, 0xdc, 0x38, 0x68, 0x1d, 0x1a, 0x01,
	0x3b, 0x0d, 0x08, 0x35, 0x88, 0xcb, 0x17, 0xe1, 0x53, 0x25, 0x3d, 0x4f, 0xe1, 0x6b, 0x67, 0x1f,
	0x3a, 0xa4, 0x51, 0xaf, 0xb9, 0x26, 0x3d, 0x52, 0x59, 0x6b, 0x57, 0x53, 0x31, 0xc7, 0x25, 0x94,
	0x99, 0x6e, 0xd0, 0x59, 0x85, 0xc9, 0xf8, 0xab, 0x06, 0x23, 0x9b, 0xe1, 0x4c, 0xd0, 0x04, 0x64,
	0x9c, 0xfa, 0xb4, 0xf6, 0x40, 0x5b, 0xc8, 0x56, 0xf9, 0x0a, 0xcd, 0x02, 0x70, 0xd0, 0x07, 0x62,
	0xb1, 0x1a, 0x8f, 0x67, 0x64, 0x3c, 0xab, 0x22, 0xdb, 0x75, 0x84, 0x60, 0xd8, 0x33, 0x5d, 0x32,
	0x3d, 0x24, 0x37, 0xe4, 0x1a, 0xad, 0xc1, 0x98, 0xd5, 0x24, 0x26, 0x23, 0x35, 0x41, 0x34, 0x3d,
	0xcc, 0xb7, 0xc6, 0xca, 0x05, 0x5d, 0xf5, 0x15, 0xc9, 0xd2, 0xf7, 0x22, 0x15, 0x55, 0x08, 0xe1,
	0x22, 0x80, 0x03, 0xb8, 0xf3, 0xd6, 0xa1, 0x4c, 0xc9, 0xa1, 0x55, 0xf2, 0xb1, 0xc5, 0x51, 0x3d,
	0x32, 0xb4, 0x5e, 0x19, 0x33, 0x90, 0x0d, 0x4c, 0x9b, 0xd4, 0xa8, 0xf3, 0x99, 0x48, 0x91, 0x37,
	0xaa, 0xa3, 0x22, 0xb0, 0xcb, 0xff, 0x97, 0xb9, 0x62, 0x93, 0xf9, 0x47, 0xc4, 0x53, 0x4a, 0x25,
	0x7c, 0x4f, 0x04, 0xf0, 0x09, 0xe4, 0xbb, 0x19, 0x69, 0xc0, 0xcd, 0x43, 0xd0, 0x0a, 0x8c, 0x46,
	0x46, 0xe5, 0x84, 0x43, 0xbc, 0x87, 0x19, 0xbd, 0xdf, 0xa9, 0xba, 0xca, 0xab, 0xb6, 0xc1, 0x68,
	0x0e, 0x26, 0x3d, 0xf2, 0x89, 0xd5, 0x62, 0xa4, 0xe1, 0xb9, 0xe5, 0x44, 0xb8, 0xd2, 0x26, 0xde,
	0x81, 0xfc, 0x2b, 0xd9, 0x78, 0x54, 0x42, 0xf5, 0xba, 0x0c, 0x23, 0xaa, 0x96, 0x6c, 0xf4, 0x0a,
	0xde, 0x08, 0x8b, 0x7f, 0x6a, 0x90, 0xdf, 0x0f, 0xea, 0xfd, 0xf5, 0x78, 0xff, 0x0a, 0x13, 0x3b,
	0x3b, 0x15, 0xe1, 0x67, 0x17, 0xa3, 0xcb, 0x0c, 0x4e, 0x27, 0xa6, 0xdc, 0x92, 0x6c, 0xd2, 0x86,
	0xf2, 0x58, 0x93, 0xa6, 0xfc, 0x46, 0x38, 0x75, 0x87, 0x23, 0xaa, 0x10, 0xc2, 0xc5, 0x1a, 0x2f,
	0x43, 0x7e, 0x93, 0x34, 0xc8, 0x35, 0xa5, 0xe2, 0x15, 0x98, 0xda, 0xf7, 0xea, 0xff, 0x90, 0x58,
	0x86, 0xdb, 0x5b, 0x84, 0x5d, 0x2b, 0xa7, 0xfc, 0x3d, 0x0b, 0xb7, 0x54, 0xc6, 0x6e, 0x78, 0xc5,
	0xbc, 0x2b, 0xa1, 0x13, 0x18, 0x8f, 0x9b, 0x05, 0xcd, 0x27, 0x9d, 0x55, 0x82, 0x81, 0x0b, 0x0b,
	0x57, 0x03, 0x43, 0xdf, 0xe1, 0xfc, 0xc5, 0xaf, 0xdf, 0xdf, 0x32, 0x13, 0x68, 0x3c, 0x7e, 0x55,
	0xa2, 0x16, 0xe4, 0xba, 0xcc, 0x82, 0x12, 0x0b, 0x26, 0xf9, 0xa9, 0x90, 0x36, 0x4f, 0x3c, 0x2b,
	0xd9, 0xee, 0xe1, 0x2e, 0xb6, 0xd5, 0xf6, 0x94, 0x29, 0x40, 0xe7, 0xe0, 0xd0, 0xe3, 0xa4, 0x4a,
	0x7d, 0x07, 0x9b, 0x4e, 0xf8, 0x50, 0x12, 0xce, 0xa0, 0xfb, 0x71, 0x42, 0xe3, 0xac, 0x33, 0x89,
	0x73, 0x74, 0xa1, 0x41, 0xae, 0xcb, 0xc9, 0xc9, 0xcd, 0x26, 0x99, 0x3d, 0x9d, 0xbb, 0x28, 0xb9,
	0x1f, 0x95, 0x2f, 0xe7, 0xee, 0x74, 0xce, 0x20, 0xd7, 0x65, 0xd1, 0x64, 0x0d, 0x49, 0x2e, 0x2e,
	0x4c, 0xf5, 0xbd, 0x05, 0xaf, 0xc5, 0x05, 0x1f, 0xb5, 0x5e, 0x4c, 0x69, 0x9d, 0x5f, 0xc5, 0x93,
	0x3d, 0x16, 0x47, 0xc5, 0xc4, 0xe6, 0x13, 0xdf, 0x83, 0xf4, 0xf6, 0x9f, 0x49, 0xfe, 0x79, 0x8c,
	0x2f, 0x6f, 0xbf, 0xa5, 0xca, 0xae, 0x6a, 0x45, 0xf4, 0x05, 0xc6, 0x77, 0xf9, 0x9b, 0x60, 0xba,
	0x15, 0xf9, 0x51, 0x43, 0x38, 0xaa, 0xcd, 0xbf, 0x75, 0xa2, 0x6c, 0x7c, 0x33, 0xe2, 0xbf, 0xdb,
	0x83, 0x09, 0x77, 0x71, 0x49, 0x32, 0x3f, 0xc5, 0x73, 0x82, 0xf9, 0xac, 0x49, 0xa8, 0xdf, 0x6a,
	0x5a, 0x64, 0xbd, 0xad, 0xa1, 0x78, 0xbe, 0x4a, 0x63, 0xd5, 0x14, 0xfb, 0x56, 0x1a, 0xfb, 0xd6,
	0x7f, 0x65, 0xb7, 0x7b, 0xd8, 0x7f, 0x68, 0x80, 0xf6, 0x78, 0x49, 0x11, 0x24, 0x4d, 0xd7, 0xa1,
	0x54, 0xfc, 0xa2, 0xe8, 0x78, 0x40, 0x11, 0xf4, 0x43, 0x22, 0x29, 0x4f, 0x06, 0x40, 0xaa, 0x17,
	0x7e, 0x45, 0xca, 0x2b, 0xe1, 0xc5, 0xcb, 0xe5, 0xb1, 0xbe, 0x6c, 0x2e, 0x72, 0x63, 0x11, 0xa6,
	0xf8, 0xf7, 0x3e, 0x61, 0xe2, 0x1b, 0xb9, 0xe8, 0x56, 0xa9, 0x08, 0x07, 0x56, 0xb4, 0x3f, 0x9a,
	0x76, 0x70, 0x53, 0xba, 0x71, 0xe9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x65, 0x28, 0x36,
	0xb1, 0x09, 0x00, 0x00,
}
