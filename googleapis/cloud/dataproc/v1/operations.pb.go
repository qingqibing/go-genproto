// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/cloud/dataproc/v1/operations.proto
// DO NOT EDIT!

package google_cloud_dataproc_v1 // import "google.golang.org/genproto/googleapis/cloud/dataproc/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/serviceconfig"
import _ "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/golang/protobuf/ptypes/empty"
import google_protobuf3 "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The operation state.
type ClusterOperationStatus_State int32

const (
	// Unused.
	ClusterOperationStatus_UNKNOWN ClusterOperationStatus_State = 0
	// The operation has been created.
	ClusterOperationStatus_PENDING ClusterOperationStatus_State = 1
	// The operation is running.
	ClusterOperationStatus_RUNNING ClusterOperationStatus_State = 2
	// The operation is done; either cancelled or completed.
	ClusterOperationStatus_DONE ClusterOperationStatus_State = 3
)

var ClusterOperationStatus_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "RUNNING",
	3: "DONE",
}
var ClusterOperationStatus_State_value = map[string]int32{
	"UNKNOWN": 0,
	"PENDING": 1,
	"RUNNING": 2,
	"DONE":    3,
}

func (x ClusterOperationStatus_State) String() string {
	return proto.EnumName(ClusterOperationStatus_State_name, int32(x))
}
func (ClusterOperationStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor2, []int{0, 0}
}

// The status of the operation.
type ClusterOperationStatus struct {
	// [Output-only] A message containing the operation state.
	State ClusterOperationStatus_State `protobuf:"varint,1,opt,name=state,enum=google.cloud.dataproc.v1.ClusterOperationStatus_State" json:"state,omitempty"`
	// [Output-only] A message containing the detailed operation state.
	InnerState string `protobuf:"bytes,2,opt,name=inner_state,json=innerState" json:"inner_state,omitempty"`
	// [Output-only]A message containing any operation metadata details.
	Details string `protobuf:"bytes,3,opt,name=details" json:"details,omitempty"`
	// [Output-only] The time this state was entered.
	StateStartTime *google_protobuf3.Timestamp `protobuf:"bytes,4,opt,name=state_start_time,json=stateStartTime" json:"state_start_time,omitempty"`
}

func (m *ClusterOperationStatus) Reset()                    { *m = ClusterOperationStatus{} }
func (m *ClusterOperationStatus) String() string            { return proto.CompactTextString(m) }
func (*ClusterOperationStatus) ProtoMessage()               {}
func (*ClusterOperationStatus) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ClusterOperationStatus) GetStateStartTime() *google_protobuf3.Timestamp {
	if m != nil {
		return m.StateStartTime
	}
	return nil
}

// Metadata describing the operation.
type ClusterOperationMetadata struct {
	// [Output-only] Name of the cluster for the operation.
	ClusterName string `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	// [Output-only] Cluster UUID for the operation.
	ClusterUuid string `protobuf:"bytes,8,opt,name=cluster_uuid,json=clusterUuid" json:"cluster_uuid,omitempty"`
	// [Output-only] Current operation status.
	Status *ClusterOperationStatus `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
	// [Output-only] The previous operation status.
	StatusHistory []*ClusterOperationStatus `protobuf:"bytes,10,rep,name=status_history,json=statusHistory" json:"status_history,omitempty"`
	// [Output-only] The operation type.
	OperationType string `protobuf:"bytes,11,opt,name=operation_type,json=operationType" json:"operation_type,omitempty"`
	// [Output-only] Short description of operation.
	Description string `protobuf:"bytes,12,opt,name=description" json:"description,omitempty"`
}

func (m *ClusterOperationMetadata) Reset()                    { *m = ClusterOperationMetadata{} }
func (m *ClusterOperationMetadata) String() string            { return proto.CompactTextString(m) }
func (*ClusterOperationMetadata) ProtoMessage()               {}
func (*ClusterOperationMetadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ClusterOperationMetadata) GetStatus() *ClusterOperationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterOperationMetadata) GetStatusHistory() []*ClusterOperationStatus {
	if m != nil {
		return m.StatusHistory
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterOperationStatus)(nil), "google.cloud.dataproc.v1.ClusterOperationStatus")
	proto.RegisterType((*ClusterOperationMetadata)(nil), "google.cloud.dataproc.v1.ClusterOperationMetadata")
	proto.RegisterEnum("google.cloud.dataproc.v1.ClusterOperationStatus_State", ClusterOperationStatus_State_name, ClusterOperationStatus_State_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/cloud/dataproc/v1/operations.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0xdd, 0x6a, 0xdb, 0x30,
	0x18, 0x9d, 0x93, 0xb6, 0x69, 0xe5, 0x36, 0x0b, 0xba, 0x18, 0x22, 0x0c, 0x96, 0x05, 0x06, 0xb9,
	0xb2, 0x96, 0x16, 0xc6, 0x60, 0x77, 0x59, 0xca, 0x52, 0xb6, 0x39, 0x21, 0x6d, 0xe8, 0x65, 0x50,
	0x6c, 0x55, 0x15, 0xd8, 0x92, 0x90, 0xe4, 0x40, 0x1e, 0x67, 0xef, 0xb3, 0x87, 0x1a, 0x92, 0xe2,
	0x2c, 0xfb, 0x29, 0xeb, 0x7a, 0x63, 0xeb, 0xfb, 0x39, 0xe7, 0x3b, 0xe7, 0x93, 0xc0, 0x84, 0x49,
	0xc9, 0x0a, 0x9a, 0x30, 0x59, 0x10, 0xc1, 0x12, 0xa9, 0x19, 0x66, 0x54, 0x28, 0x2d, 0xad, 0xc4,
	0xa1, 0x44, 0x14, 0x37, 0x38, 0x2b, 0x64, 0x95, 0xe3, 0x9c, 0x58, 0xa2, 0xb4, 0xcc, 0xf0, 0x7a,
	0x88, 0xa5, 0xa2, 0x9a, 0x58, 0x2e, 0x85, 0x49, 0x7c, 0x37, 0x44, 0x5b, 0x26, 0xdf, 0x9a, 0xd4,
	0xad, 0xc9, 0x7a, 0xd8, 0xbd, 0x7a, 0xdc, 0x0c, 0xa2, 0x38, 0x36, 0x54, 0xaf, 0x79, 0x46, 0x33,
	0x29, 0xee, 0x38, 0xc3, 0x44, 0x08, 0x69, 0xf7, 0x87, 0x74, 0x47, 0x8f, 0xa3, 0x2a, 0xa4, 0x60,
	0xba, 0x12, 0x82, 0x0b, 0xf6, 0x87, 0xd0, 0xee, 0x05, 0xe3, 0xf6, 0xbe, 0x5a, 0x25, 0x99, 0x2c,
	0x71, 0xe0, 0xc1, 0xbe, 0xb0, 0xaa, 0xee, 0xb0, 0xb2, 0x1b, 0x45, 0x0d, 0xa6, 0xa5, 0xb2, 0x9b,
	0xf0, 0xdd, 0x82, 0x3e, 0xfc, 0x1b, 0x64, 0x79, 0x49, 0x8d, 0x25, 0xa5, 0xfa, 0x79, 0x0a, 0xe0,
	0xfe, 0xb7, 0x06, 0x78, 0xf1, 0xb1, 0xa8, 0x8c, 0xa5, 0x7a, 0x5a, 0xab, 0xb9, 0xb6, 0xc4, 0x56,
	0x06, 0x7e, 0x01, 0x87, 0xc6, 0x12, 0x4b, 0x51, 0xd4, 0x8b, 0x06, 0xed, 0xf3, 0x77, 0xc9, 0x43,
	0x5b, 0x4c, 0xfe, 0x4e, 0x90, 0xb8, 0x1f, 0x9d, 0x07, 0x12, 0xf8, 0x0a, 0xc4, 0x5c, 0x08, 0xaa,
	0x97, 0x81, 0xb3, 0xd1, 0x8b, 0x06, 0x27, 0x73, 0xe0, 0x53, 0xbe, 0x0f, 0x22, 0xd0, 0xca, 0xa9,
	0x25, 0xbc, 0x30, 0xa8, 0xe9, 0x8b, 0x75, 0x08, 0xc7, 0xa0, 0xe3, 0x41, 0x0e, 0xaa, 0xed, 0xd2,
	0x59, 0x40, 0x07, 0xbd, 0x68, 0x10, 0x9f, 0x77, 0x6b, 0x4d, 0xb5, 0xe1, 0xe4, 0xa6, 0xf6, 0x37,
	0x6f, 0x7b, 0xcc, 0xb5, 0x83, 0xb8, 0x64, 0xff, 0x3d, 0x38, 0x0c, 0x83, 0x62, 0xd0, 0x5a, 0xa4,
	0x9f, 0xd3, 0xe9, 0x6d, 0xda, 0x79, 0xe6, 0x82, 0xd9, 0x65, 0x3a, 0xbe, 0x4a, 0x3f, 0x75, 0x22,
	0x17, 0xcc, 0x17, 0x69, 0xea, 0x82, 0x06, 0x3c, 0x06, 0x07, 0xe3, 0x69, 0x7a, 0xd9, 0x69, 0xf6,
	0xbf, 0x37, 0x00, 0xfa, 0xdd, 0xe2, 0x57, 0x6a, 0x89, 0x5b, 0x01, 0x7c, 0x0d, 0x4e, 0xb3, 0x50,
	0x5b, 0x0a, 0x52, 0x52, 0xd4, 0xf2, 0xda, 0xe3, 0x6d, 0x2e, 0x25, 0x25, 0xdd, 0x6f, 0xa9, 0x2a,
	0x9e, 0xa3, 0xe3, 0x5f, 0x5a, 0x16, 0x15, 0xcf, 0xe1, 0x04, 0x1c, 0x19, 0xbf, 0x34, 0x74, 0xe2,
	0x8d, 0xbd, 0xfd, 0xdf, 0x65, 0xcf, 0xb7, 0x78, 0x78, 0x0b, 0xda, 0xe1, 0xb4, 0xbc, 0xe7, 0xc6,
	0x4a, 0xbd, 0x41, 0xa0, 0xd7, 0x7c, 0x12, 0xe3, 0x59, 0xe0, 0x99, 0x04, 0x1a, 0xf8, 0x06, 0xb4,
	0x77, 0xef, 0x75, 0xe9, 0x9e, 0x15, 0x8a, 0xbd, 0x8f, 0xb3, 0x5d, 0xf6, 0x66, 0xa3, 0x28, 0xec,
	0x81, 0x38, 0xa7, 0x26, 0xd3, 0x5c, 0xb9, 0x14, 0x3a, 0x0d, 0x5e, 0xf7, 0x52, 0xa3, 0x21, 0x78,
	0x99, 0xc9, 0xf2, 0x41, 0x39, 0xa3, 0xe7, 0x3b, 0x21, 0x66, 0xe6, 0xae, 0x75, 0x16, 0xad, 0x8e,
	0xfc, 0xfd, 0x5e, 0xfc, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x53, 0x8b, 0x67, 0x13, 0x04, 0x00,
	0x00,
}