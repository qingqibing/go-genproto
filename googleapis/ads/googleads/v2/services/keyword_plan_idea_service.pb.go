// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/keyword_plan_idea_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v2.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{0}
}

func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(m, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{1}
}

func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (m *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(m, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{2}
}

func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (m *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(m, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{3}
}

func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (m *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(m, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{4}
}

func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(m, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cad9891eddcc8ab, []int{5}
}

func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(m, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v2.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v2.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v2.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v2.services.GenerateKeywordIdeaResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/keyword_plan_idea_service.proto", fileDescriptor_5cad9891eddcc8ab)
}

var fileDescriptor_5cad9891eddcc8ab = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0xe3, 0x46,
	0x14, 0xaf, 0x1d, 0x4a, 0xc2, 0xa4, 0xaa, 0xc4, 0x08, 0xb5, 0x29, 0xa0, 0x12, 0x45, 0x1c, 0x28,
	0x52, 0xed, 0xca, 0x1c, 0x4a, 0x4d, 0x23, 0xd5, 0xa9, 0xda, 0x04, 0x55, 0x45, 0xc8, 0x94, 0x1c,
	0xaa, 0x48, 0xd6, 0x60, 0x3f, 0x5c, 0x2b, 0xce, 0x4c, 0x3a, 0x33, 0x0e, 0xfb, 0x47, 0x5c, 0xb8,
	0xee, 0x71, 0xbf, 0xc1, 0x1e, 0xf7, 0x3b, 0xec, 0x17, 0xe0, 0x8a, 0xf6, 0xbe, 0x87, 0x3d, 0xed,
	0xa7, 0x58, 0xd9, 0x1e, 0x27, 0x01, 0x92, 0x0d, 0xcb, 0xed, 0xf9, 0xcd, 0xef, 0xfd, 0xde, 0x6f,
	0xde, 0x7b, 0xf3, 0x8c, 0x7e, 0x0b, 0x19, 0x0b, 0x63, 0x30, 0x49, 0x20, 0xcc, 0xdc, 0x4c, 0xad,
	0x91, 0x65, 0x0a, 0xe0, 0xa3, 0xc8, 0x07, 0x61, 0xf6, 0xe1, 0xe9, 0x05, 0xe3, 0x81, 0x37, 0x8c,
	0x09, 0xf5, 0xa2, 0x00, 0x88, 0xa7, 0x8e, 0x8c, 0x21, 0x67, 0x92, 0xe1, 0x7a, 0x1e, 0x66, 0x90,
	0x40, 0x18, 0x63, 0x06, 0x63, 0x64, 0x19, 0x05, 0xc3, 0xfa, 0xfe, 0xbc, 0x1c, 0x3e, 0x1b, 0x0c,
	0x18, 0xbd, 0x9d, 0x21, 0xf7, 0xe5, 0xdc, 0xf3, 0x23, 0x81, 0x26, 0x83, 0x3b, 0xd2, 0x28, 0xc8,
	0x0b, 0xc6, 0xfb, 0x2a, 0x72, 0xb3, 0x88, 0x1c, 0x46, 0x26, 0xa1, 0x94, 0x49, 0x22, 0x23, 0x46,
	0x85, 0x3a, 0xdd, 0x9a, 0x3a, 0x3d, 0x8f, 0x20, 0x0e, 0xbc, 0x33, 0xf8, 0x8f, 0x8c, 0x22, 0xc6,
	0x15, 0xe0, 0x7b, 0x05, 0xc8, 0xbe, 0xce, 0x92, 0x73, 0xf3, 0x82, 0x93, 0xe1, 0x10, 0x78, 0x41,
	0xf0, 0xed, 0x14, 0x81, 0x1f, 0x47, 0x40, 0x65, 0x7e, 0xd0, 0x78, 0xbb, 0x84, 0x36, 0xda, 0x40,
	0x81, 0x13, 0x09, 0x7f, 0xe5, 0xf2, 0x0e, 0x03, 0x20, 0xc2, 0x85, 0xff, 0x13, 0x10, 0x12, 0x6f,
	0xa1, 0xaa, 0x9f, 0x08, 0xc9, 0x06, 0xc0, 0xbd, 0x28, 0xa8, 0x69, 0x75, 0x6d, 0x67, 0xc5, 0x45,
	0x85, 0xeb, 0x30, 0xc0, 0x4d, 0x54, 0x89, 0x09, 0x0d, 0x13, 0x12, 0x42, 0xad, 0x5c, 0xd7, 0x76,
	0xaa, 0xd6, 0xa6, 0x2a, 0xab, 0x51, 0x88, 0x31, 0x4e, 0x24, 0x8f, 0x68, 0xd8, 0x25, 0x71, 0x02,
	0xad, 0xd2, 0x3b, 0x47, 0x77, 0xc7, 0x21, 0xf8, 0x08, 0xad, 0x85, 0xc0, 0x3c, 0x49, 0x78, 0x08,
	0xd2, 0xf3, 0x19, 0x15, 0x92, 0x50, 0x29, 0x6a, 0x95, 0x7a, 0x69, 0x11, 0x95, 0x8b, 0x43, 0x60,
	0xff, 0x64, 0x81, 0xbf, 0x17, 0x71, 0xf8, 0x19, 0x5a, 0x9b, 0x55, 0xe5, 0xda, 0x4a, 0x5d, 0xdb,
	0xf9, 0xda, 0xea, 0x18, 0xf3, 0x9a, 0x9f, 0x35, 0xc8, 0x50, 0x15, 0x38, 0x8e, 0x09, 0x3d, 0xca,
	0x03, 0xff, 0xa0, 0xc9, 0x60, 0x86, 0xdb, 0xc5, 0xfd, 0x7b, 0x3e, 0x7c, 0x3e, 0xc9, 0x4d, 0x68,
	0xe0, 0x25, 0x3c, 0xf6, 0x04, 0x40, 0x50, 0xd3, 0xb3, 0xb2, 0xec, 0x19, 0x8b, 0x06, 0xaf, 0xc8,
	0xe3, 0xd0, 0xe0, 0x94, 0xc7, 0x27, 0x00, 0x41, 0xe7, 0x0b, 0x77, 0xb5, 0x7f, 0xd7, 0x89, 0x5d,
	0xf4, 0x55, 0x91, 0x27, 0xe3, 0x2f, 0x65, 0xfc, 0x3f, 0x3e, 0x98, 0x5f, 0x31, 0x57, 0xfb, 0x93,
	0x4f, 0xfc, 0x27, 0xaa, 0x8c, 0xf5, 0x7e, 0x99, 0xf1, 0xfd, 0xb0, 0x98, 0x6f, 0xa2, 0xb2, 0x9c,
	0xe4, 0x66, 0x6b, 0x19, 0x2d, 0xa5, 0x1c, 0x8d, 0x4b, 0xb4, 0x7a, 0xef, 0x36, 0xd8, 0x40, 0xa5,
	0x84, 0xc7, 0xd9, 0x10, 0x2d, 0xea, 0x6d, 0x0a, 0xc4, 0xfb, 0xa8, 0xa2, 0x34, 0x8a, 0x9a, 0xfe,
	0x80, 0x81, 0x18, 0xa3, 0x1b, 0x6d, 0x54, 0x9d, 0xba, 0xec, 0x2d, 0x22, 0xed, 0xb3, 0x88, 0x7e,
	0x41, 0xe5, 0x47, 0xaa, 0x6f, 0xc8, 0x99, 0x2f, 0xcb, 0x05, 0x31, 0x64, 0x54, 0x00, 0x3e, 0x45,
	0x65, 0x0e, 0x22, 0x89, 0x65, 0x21, 0xe9, 0x60, 0x71, 0xc1, 0x67, 0xf3, 0x25, 0xb1, 0x74, 0x0b,
	0xae, 0xc6, 0x1b, 0x0d, 0x7d, 0x37, 0x17, 0x86, 0x7f, 0x42, 0x4b, 0x12, 0x9e, 0x48, 0x35, 0x92,
	0x9f, 0xbe, 0x44, 0x86, 0xc4, 0x74, 0x32, 0xd4, 0xd9, 0x32, 0x1d, 0x80, 0xe4, 0x91, 0x2f, 0xd4,
	0xd0, 0xfd, 0x3a, 0x57, 0xb3, 0xda, 0x8b, 0x53, 0x4f, 0xa7, 0x13, 0x09, 0xc9, 0x78, 0xe4, 0x93,
	0xf8, 0xef, 0x9c, 0x63, 0xfc, 0x88, 0x52, 0x81, 0xca, 0x67, 0xbd, 0xd0, 0xd1, 0x37, 0x53, 0x41,
	0xe9, 0xd1, 0x49, 0x7e, 0x7d, 0x7c, 0xa3, 0xa1, 0xb5, 0x59, 0xbb, 0x0a, 0x37, 0x1f, 0x55, 0xb9,
	0x62, 0xc7, 0xad, 0x37, 0x1f, 0x5b, 0xf8, 0xac, 0x91, 0x8d, 0xe6, 0xd5, 0xcd, 0xfb, 0x97, 0xfa,
	0xcf, 0x0d, 0x2b, 0xfb, 0x45, 0xa8, 0xcd, 0x28, 0xcc, 0xe7, 0x53, 0x7b, 0xb3, 0xb9, 0x7b, 0x69,
	0x87, 0x33, 0x14, 0xd8, 0xda, 0xee, 0xfa, 0xc6, 0xb5, 0x53, 0x9b, 0x24, 0x55, 0xd6, 0x30, 0x12,
	0x69, 0x05, 0x5b, 0x57, 0x3a, 0xda, 0xf6, 0xd9, 0x60, 0xa1, 0xc0, 0xd6, 0xc6, 0xec, 0x9a, 0x1d,
	0xa7, 0x8d, 0x3d, 0xd6, 0xfe, 0xed, 0x28, 0x82, 0x90, 0xa5, 0x9b, 0xd7, 0x60, 0x3c, 0x34, 0x43,
	0xa0, 0x59, 0xdb, 0xcd, 0x49, 0xca, 0xf9, 0x7f, 0xd5, 0x83, 0xc2, 0x78, 0xa5, 0x97, 0xda, 0x8e,
	0xf3, 0x5a, 0xaf, 0xb7, 0x73, 0x42, 0x27, 0x10, 0x46, 0x6e, 0xa6, 0x56, 0xd7, 0x32, 0x54, 0x62,
	0x71, 0x5d, 0x40, 0x7a, 0x4e, 0x20, 0x7a, 0x63, 0x48, 0xaf, 0x6b, 0xf5, 0x0a, 0xc8, 0x07, 0x7d,
	0x3b, 0xf7, 0xdb, 0xb6, 0x13, 0x08, 0xdb, 0x1e, 0x83, 0x6c, 0xbb, 0x6b, 0xd9, 0x76, 0x01, 0x3b,
	0x5b, 0xce, 0x74, 0xee, 0x7d, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x94, 0xb6, 0xdd, 0x31, 0xfc, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanIdeaServiceClient(cc grpc.ClientConnInterface) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

// UnimplementedKeywordPlanIdeaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanIdeaServiceServer struct {
}

func (*UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordIdeas(ctx context.Context, req *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordIdeas not implemented")
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/keyword_plan_idea_service.proto",
}
