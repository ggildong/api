// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: spaceone/api/identity/v1/endpoint.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EndpointInfo_EndpointState int32

const (
	EndpointInfo_NONE     EndpointInfo_EndpointState = 0
	EndpointInfo_ACTIVE   EndpointInfo_EndpointState = 1
	EndpointInfo_INACTIVE EndpointInfo_EndpointState = 2
)

// Enum value maps for EndpointInfo_EndpointState.
var (
	EndpointInfo_EndpointState_name = map[int32]string{
		0: "NONE",
		1: "ACTIVE",
		2: "INACTIVE",
	}
	EndpointInfo_EndpointState_value = map[string]int32{
		"NONE":     0,
		"ACTIVE":   1,
		"INACTIVE": 2,
	}
)

func (x EndpointInfo_EndpointState) Enum() *EndpointInfo_EndpointState {
	p := new(EndpointInfo_EndpointState)
	*p = x
	return p
}

func (x EndpointInfo_EndpointState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EndpointInfo_EndpointState) Descriptor() protoreflect.EnumDescriptor {
	return file_spaceone_api_identity_v1_endpoint_proto_enumTypes[0].Descriptor()
}

func (EndpointInfo_EndpointState) Type() protoreflect.EnumType {
	return &file_spaceone_api_identity_v1_endpoint_proto_enumTypes[0]
}

func (x EndpointInfo_EndpointState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EndpointInfo_EndpointState.Descriptor instead.
func (EndpointInfo_EndpointState) EnumDescriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_endpoint_proto_rawDescGZIP(), []int{0, 0}
}

type EndpointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service  string                     `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Name     string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Endpoint string                     `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	State    EndpointInfo_EndpointState `protobuf:"varint,4,opt,name=state,proto3,enum=spaceone.api.identity.v1.EndpointInfo_EndpointState" json:"state,omitempty"`
	Version  string                     `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *EndpointInfo) Reset() {
	*x = EndpointInfo{}
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointInfo) ProtoMessage() {}

func (x *EndpointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointInfo.ProtoReflect.Descriptor instead.
func (*EndpointInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *EndpointInfo) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EndpointInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EndpointInfo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *EndpointInfo) GetState() EndpointInfo_EndpointState {
	if x != nil {
		return x.State
	}
	return EndpointInfo_NONE
}

func (x *EndpointInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type EndpointQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Query *v1.Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// +optional
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	// example: public | internal
	// +optional
	EndpointType string `protobuf:"bytes,3,opt,name=endpoint_type,json=endpointType,proto3" json:"endpoint_type,omitempty"`
	DomainId     string `protobuf:"bytes,6,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *EndpointQuery) Reset() {
	*x = EndpointQuery{}
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointQuery) ProtoMessage() {}

func (x *EndpointQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointQuery.ProtoReflect.Descriptor instead.
func (*EndpointQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *EndpointQuery) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *EndpointQuery) GetEndpointType() string {
	if x != nil {
		return x.EndpointType
	}
	return ""
}

func (x *EndpointQuery) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

type EndpointsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*EndpointInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32           `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *EndpointsInfo) Reset() {
	*x = EndpointsInfo{}
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EndpointsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointsInfo) ProtoMessage() {}

func (x *EndpointsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v1_endpoint_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointsInfo.ProtoReflect.Descriptor instead.
func (*EndpointsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v1_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *EndpointsInfo) GetResults() []*EndpointInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *EndpointsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_identity_v1_endpoint_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v1_endpoint_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4a,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x22, 0x9e, 0x01, 0x0a, 0x0d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x72, 0x0a, 0x0d, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x07, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x8b,
	0x01, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x7f, 0x0a, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x27, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a,
	0x22, 0x1a, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v1_endpoint_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v1_endpoint_proto_rawDescData = file_spaceone_api_identity_v1_endpoint_proto_rawDesc
)

func file_spaceone_api_identity_v1_endpoint_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v1_endpoint_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v1_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v1_endpoint_proto_rawDescData)
	})
	return file_spaceone_api_identity_v1_endpoint_proto_rawDescData
}

var file_spaceone_api_identity_v1_endpoint_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spaceone_api_identity_v1_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spaceone_api_identity_v1_endpoint_proto_goTypes = []any{
	(EndpointInfo_EndpointState)(0), // 0: spaceone.api.identity.v1.EndpointInfo.EndpointState
	(*EndpointInfo)(nil),            // 1: spaceone.api.identity.v1.EndpointInfo
	(*EndpointQuery)(nil),           // 2: spaceone.api.identity.v1.EndpointQuery
	(*EndpointsInfo)(nil),           // 3: spaceone.api.identity.v1.EndpointsInfo
	(*v1.Query)(nil),                // 4: spaceone.api.core.v1.Query
}
var file_spaceone_api_identity_v1_endpoint_proto_depIdxs = []int32{
	0, // 0: spaceone.api.identity.v1.EndpointInfo.state:type_name -> spaceone.api.identity.v1.EndpointInfo.EndpointState
	4, // 1: spaceone.api.identity.v1.EndpointQuery.query:type_name -> spaceone.api.core.v1.Query
	1, // 2: spaceone.api.identity.v1.EndpointsInfo.results:type_name -> spaceone.api.identity.v1.EndpointInfo
	2, // 3: spaceone.api.identity.v1.Endpoint.list:input_type -> spaceone.api.identity.v1.EndpointQuery
	3, // 4: spaceone.api.identity.v1.Endpoint.list:output_type -> spaceone.api.identity.v1.EndpointsInfo
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v1_endpoint_proto_init() }
func file_spaceone_api_identity_v1_endpoint_proto_init() {
	if File_spaceone_api_identity_v1_endpoint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_identity_v1_endpoint_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v1_endpoint_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v1_endpoint_proto_depIdxs,
		EnumInfos:         file_spaceone_api_identity_v1_endpoint_proto_enumTypes,
		MessageInfos:      file_spaceone_api_identity_v1_endpoint_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v1_endpoint_proto = out.File
	file_spaceone_api_identity_v1_endpoint_proto_rawDesc = nil
	file_spaceone_api_identity_v1_endpoint_proto_goTypes = nil
	file_spaceone_api_identity_v1_endpoint_proto_depIdxs = nil
}
