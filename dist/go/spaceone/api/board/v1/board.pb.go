//A Board is a bulletin-board-type resource for posting notices and announcements in Cloudforet.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/board/v1/board.proto

package v1

import (
	v1 "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v1"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

//	{
//	  "name": "notice",
//	  "categories": ["admin", "developer", "devops"],
//	  "tags": {"a": "b"}
//	}
type CreateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Categories []string `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateBoardRequest) Reset() {
	*x = CreateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBoardRequest) ProtoMessage() {}

func (x *CreateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBoardRequest.ProtoReflect.Descriptor instead.
func (*CreateBoardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBoardRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *CreateBoardRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

//	{
//	   "board_id": "board-123456789012",
//	   "name": "system notice",
//	   "tags": {"b": "c"}
//	}
type UpdateBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,3,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateBoardRequest) Reset() {
	*x = UpdateBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBoardRequest) ProtoMessage() {}

func (x *UpdateBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBoardRequest.ProtoReflect.Descriptor instead.
func (*UpdateBoardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *UpdateBoardRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBoardRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

//	{
//	   "board_id": "board-123456789012",
//	   "categories": ["Developer", "SRE", "Devops"]
//	}
type SetBoardCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	// +optional
	Categories []string `protobuf:"bytes,2,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *SetBoardCategoriesRequest) Reset() {
	*x = SetBoardCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBoardCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBoardCategoriesRequest) ProtoMessage() {}

func (x *SetBoardCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBoardCategoriesRequest.ProtoReflect.Descriptor instead.
func (*SetBoardCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{2}
}

func (x *SetBoardCategoriesRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *SetBoardCategoriesRequest) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

//	{
//	   "board_id": "board-123456789012"
//	}
type BoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
}

func (x *BoardRequest) Reset() {
	*x = BoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardRequest) ProtoMessage() {}

func (x *BoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardRequest.ProtoReflect.Descriptor instead.
func (*BoardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{3}
}

func (x *BoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

//	{
//	   "board_id": "board-123456789012"
//	}
type GetBoardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	// +optional
	Only []string `protobuf:"bytes,3,rep,name=only,proto3" json:"only,omitempty"`
}

func (x *GetBoardRequest) Reset() {
	*x = GetBoardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBoardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBoardRequest) ProtoMessage() {}

func (x *GetBoardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBoardRequest.ProtoReflect.Descriptor instead.
func (*GetBoardRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{4}
}

func (x *GetBoardRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *GetBoardRequest) GetOnly() []string {
	if x != nil {
		return x.Only
	}
	return nil
}

//	{
//	   "query": {}
//	}
type BoardQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	BoardId string `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Query *v1.Query `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *BoardQuery) Reset() {
	*x = BoardQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardQuery) ProtoMessage() {}

func (x *BoardQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardQuery.ProtoReflect.Descriptor instead.
func (*BoardQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{5}
}

func (x *BoardQuery) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *BoardQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoardQuery) GetQuery() *v1.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

type BoardStatQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query *v1.StatisticsQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *BoardStatQuery) Reset() {
	*x = BoardStatQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardStatQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardStatQuery) ProtoMessage() {}

func (x *BoardStatQuery) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardStatQuery.ProtoReflect.Descriptor instead.
func (*BoardStatQuery) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{6}
}

func (x *BoardStatQuery) GetQuery() *v1.StatisticsQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

//	{
//	    "board_id": "board-123456789012",
//	    "name": "notice",
//	    "categories": [
//	        "admin",
//	        "developer",
//	        "devops"
//	    ],
//	    "tags": {
//	        "a": "b"
//	    },
//	    "created_at": "2022-01-01T06:47:27.759Z"
//	}
type BoardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BoardId    string          `protobuf:"bytes,1,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name       string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Categories []string        `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty"`
	Tags       *_struct.Struct `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt  string          `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *BoardInfo) Reset() {
	*x = BoardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardInfo) ProtoMessage() {}

func (x *BoardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardInfo.ProtoReflect.Descriptor instead.
func (*BoardInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{7}
}

func (x *BoardInfo) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *BoardInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoardInfo) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

func (x *BoardInfo) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *BoardInfo) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

//	{
//	     "results": [
//	         {
//	             "board_id": "board-123456789012",
//	             "name": "dev-notice",
//	             "categories": [
//	                 "flower",
//	                 "school",
//	                 "spaceone"
//	             ],
//	             "tags": {
//	                 "b": "c"
//	             },
//	             "created_at": "2022-01-01T05:16:08.549Z"
//	         },
//	         {
//	             "board_id": "board-987654321098",
//	             "name": "notice",
//	             "tags": {
//	                 "a": "b"
//	             },
//	             "created_at": "2022-01-01T05:24:19.758Z"
//	         }
//	     ],
//	     "total_count": 2
//	}
type BoardsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results    []*BoardInfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TotalCount int32        `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
}

func (x *BoardsInfo) Reset() {
	*x = BoardsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_board_v1_board_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardsInfo) ProtoMessage() {}

func (x *BoardsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_board_v1_board_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardsInfo.ProtoReflect.Descriptor instead.
func (*BoardsInfo) Descriptor() ([]byte, []int) {
	return file_spaceone_api_board_v1_board_proto_rawDescGZIP(), []int{8}
}

func (x *BoardsInfo) GetResults() []*BoardInfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *BoardsInfo) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_spaceone_api_board_v1_board_proto protoreflect.FileDescriptor

var file_spaceone_api_board_v1_board_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x70, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x56,
	0x0a, 0x19, 0x53, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0c, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x22, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6f,
	0x6e, 0x6c, 0x79, 0x22, 0x6e, 0x0a, 0x0a, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x4d, 0x0a, 0x0e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x22, 0xa6, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x2b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x69, 0x0a, 0x0a, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc0, 0x06, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x78, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a,
	0x01, 0x2a, 0x22, 0x16, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x78, 0x0a, 0x06, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x0e, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x30, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x73, 0x65, 0x74, 0x2d, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x23, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x6f, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f,
	0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x65,
	0x74, 0x12, 0x6d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x21, 0x2e, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0x67, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19,
	0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72,
	0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_board_v1_board_proto_rawDescOnce sync.Once
	file_spaceone_api_board_v1_board_proto_rawDescData = file_spaceone_api_board_v1_board_proto_rawDesc
)

func file_spaceone_api_board_v1_board_proto_rawDescGZIP() []byte {
	file_spaceone_api_board_v1_board_proto_rawDescOnce.Do(func() {
		file_spaceone_api_board_v1_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_board_v1_board_proto_rawDescData)
	})
	return file_spaceone_api_board_v1_board_proto_rawDescData
}

var file_spaceone_api_board_v1_board_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_spaceone_api_board_v1_board_proto_goTypes = []interface{}{
	(*CreateBoardRequest)(nil),        // 0: spaceone.api.board.v1.CreateBoardRequest
	(*UpdateBoardRequest)(nil),        // 1: spaceone.api.board.v1.UpdateBoardRequest
	(*SetBoardCategoriesRequest)(nil), // 2: spaceone.api.board.v1.SetBoardCategoriesRequest
	(*BoardRequest)(nil),              // 3: spaceone.api.board.v1.BoardRequest
	(*GetBoardRequest)(nil),           // 4: spaceone.api.board.v1.GetBoardRequest
	(*BoardQuery)(nil),                // 5: spaceone.api.board.v1.BoardQuery
	(*BoardStatQuery)(nil),            // 6: spaceone.api.board.v1.BoardStatQuery
	(*BoardInfo)(nil),                 // 7: spaceone.api.board.v1.BoardInfo
	(*BoardsInfo)(nil),                // 8: spaceone.api.board.v1.BoardsInfo
	(*_struct.Struct)(nil),            // 9: google.protobuf.Struct
	(*v1.Query)(nil),                  // 10: spaceone.api.core.v1.Query
	(*v1.StatisticsQuery)(nil),        // 11: spaceone.api.core.v1.StatisticsQuery
	(*empty.Empty)(nil),               // 12: google.protobuf.Empty
}
var file_spaceone_api_board_v1_board_proto_depIdxs = []int32{
	9,  // 0: spaceone.api.board.v1.CreateBoardRequest.tags:type_name -> google.protobuf.Struct
	9,  // 1: spaceone.api.board.v1.UpdateBoardRequest.tags:type_name -> google.protobuf.Struct
	10, // 2: spaceone.api.board.v1.BoardQuery.query:type_name -> spaceone.api.core.v1.Query
	11, // 3: spaceone.api.board.v1.BoardStatQuery.query:type_name -> spaceone.api.core.v1.StatisticsQuery
	9,  // 4: spaceone.api.board.v1.BoardInfo.tags:type_name -> google.protobuf.Struct
	7,  // 5: spaceone.api.board.v1.BoardsInfo.results:type_name -> spaceone.api.board.v1.BoardInfo
	0,  // 6: spaceone.api.board.v1.Board.create:input_type -> spaceone.api.board.v1.CreateBoardRequest
	1,  // 7: spaceone.api.board.v1.Board.update:input_type -> spaceone.api.board.v1.UpdateBoardRequest
	2,  // 8: spaceone.api.board.v1.Board.set_categories:input_type -> spaceone.api.board.v1.SetBoardCategoriesRequest
	3,  // 9: spaceone.api.board.v1.Board.delete:input_type -> spaceone.api.board.v1.BoardRequest
	4,  // 10: spaceone.api.board.v1.Board.get:input_type -> spaceone.api.board.v1.GetBoardRequest
	5,  // 11: spaceone.api.board.v1.Board.list:input_type -> spaceone.api.board.v1.BoardQuery
	6,  // 12: spaceone.api.board.v1.Board.stat:input_type -> spaceone.api.board.v1.BoardStatQuery
	7,  // 13: spaceone.api.board.v1.Board.create:output_type -> spaceone.api.board.v1.BoardInfo
	7,  // 14: spaceone.api.board.v1.Board.update:output_type -> spaceone.api.board.v1.BoardInfo
	7,  // 15: spaceone.api.board.v1.Board.set_categories:output_type -> spaceone.api.board.v1.BoardInfo
	12, // 16: spaceone.api.board.v1.Board.delete:output_type -> google.protobuf.Empty
	7,  // 17: spaceone.api.board.v1.Board.get:output_type -> spaceone.api.board.v1.BoardInfo
	8,  // 18: spaceone.api.board.v1.Board.list:output_type -> spaceone.api.board.v1.BoardsInfo
	9,  // 19: spaceone.api.board.v1.Board.stat:output_type -> google.protobuf.Struct
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_spaceone_api_board_v1_board_proto_init() }
func file_spaceone_api_board_v1_board_proto_init() {
	if File_spaceone_api_board_v1_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_board_v1_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBoardRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBoardRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBoardCategoriesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBoardRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardStatQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spaceone_api_board_v1_board_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardsInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spaceone_api_board_v1_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_board_v1_board_proto_goTypes,
		DependencyIndexes: file_spaceone_api_board_v1_board_proto_depIdxs,
		MessageInfos:      file_spaceone_api_board_v1_board_proto_msgTypes,
	}.Build()
	File_spaceone_api_board_v1_board_proto = out.File
	file_spaceone_api_board_v1_board_proto_rawDesc = nil
	file_spaceone_api_board_v1_board_proto_goTypes = nil
	file_spaceone_api_board_v1_board_proto_depIdxs = nil
}