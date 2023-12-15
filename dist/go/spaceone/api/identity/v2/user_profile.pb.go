// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: spaceone/api/identity/v2/user_profile.proto

package v2

import (
	_ "github.com/cloudforet-io/api/dist/go/spaceone/api/core/v2"
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

type UpdateUserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	// +optional
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// +optional
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// +optional
	Language string `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	// +optional
	Timezone string `protobuf:"bytes,5,opt,name=timezone,proto3" json:"timezone,omitempty"`
	// +optional
	Tags *_struct.Struct `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateUserProfileRequest) Reset() {
	*x = UpdateUserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserProfileRequest) ProtoMessage() {}

func (x *UpdateUserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserProfileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserProfileRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *UpdateUserProfileRequest) GetTags() *_struct.Struct {
	if x != nil {
		return x.Tags
	}
	return nil
}

type VerifyEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// +optional
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *VerifyEmailRequest) Reset() {
	*x = VerifyEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyEmailRequest) ProtoMessage() {}

func (x *VerifyEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyEmailRequest.ProtoReflect.Descriptor instead.
func (*VerifyEmailRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ConfirmEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyCode string `protobuf:"bytes,1,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
}

func (x *ConfirmEmailRequest) Reset() {
	*x = ConfirmEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmEmailRequest) ProtoMessage() {}

func (x *ConfirmEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmEmailRequest.ProtoReflect.Descriptor instead.
func (*ConfirmEmailRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{2}
}

func (x *ConfirmEmailRequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

//	{
//	 "user_id": "example@cloudforet.com",
//	 "mfa_type": "EMAIL",
//	 "options": {"email": "wonny@cloudforet.com"},
//	 "domain_id": "domain-a1b2c3d4e5f6"
//	}
type EnableMFARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EMAIL
	MfaType string `protobuf:"bytes,1,opt,name=mfa_type,json=mfaType,proto3" json:"mfa_type,omitempty"`
	// If mfa_type is EMAIL, email is required in options. options will be saved in mfa's options field.
	Options *_struct.Struct `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *EnableMFARequest) Reset() {
	*x = EnableMFARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableMFARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableMFARequest) ProtoMessage() {}

func (x *EnableMFARequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableMFARequest.ProtoReflect.Descriptor instead.
func (*EnableMFARequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{3}
}

func (x *EnableMFARequest) GetMfaType() string {
	if x != nil {
		return x.MfaType
	}
	return ""
}

func (x *EnableMFARequest) GetOptions() *_struct.Struct {
	if x != nil {
		return x.Options
	}
	return nil
}

//	{
//	 "user_id": "example@cloudforet.com",
//	 "force": false,
//	 "domain_id": "domain-a1b2c3d4e5f6"
//	}
type DisableMFARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DisableMFARequest) Reset() {
	*x = DisableMFARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableMFARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableMFARequest) ProtoMessage() {}

func (x *DisableMFARequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableMFARequest.ProtoReflect.Descriptor instead.
func (*DisableMFARequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{4}
}

//	{
//	 "user_id": "example@cloudforet",
//	 "verify_code": "123456",
//	 "domain_id": "domain-a1b2c3d4e5f6"
//	}
type ConfirmMFARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerifyCode string `protobuf:"bytes,1,opt,name=verify_code,json=verifyCode,proto3" json:"verify_code,omitempty"`
}

func (x *ConfirmMFARequest) Reset() {
	*x = ConfirmMFARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmMFARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmMFARequest) ProtoMessage() {}

func (x *ConfirmMFARequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmMFARequest.ProtoReflect.Descriptor instead.
func (*ConfirmMFARequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{5}
}

func (x *ConfirmMFARequest) GetVerifyCode() string {
	if x != nil {
		return x.VerifyCode
	}
	return ""
}

type UserProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserProfileRequest) Reset() {
	*x = UserProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfileRequest) ProtoMessage() {}

func (x *UserProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfileRequest.ProtoReflect.Descriptor instead.
func (*UserProfileRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{6}
}

type UserPasswordResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DomainId string `protobuf:"bytes,21,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *UserPasswordResetRequest) Reset() {
	*x = UserPasswordResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPasswordResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPasswordResetRequest) ProtoMessage() {}

func (x *UserPasswordResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spaceone_api_identity_v2_user_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPasswordResetRequest.ProtoReflect.Descriptor instead.
func (*UserPasswordResetRequest) Descriptor() ([]byte, []int) {
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP(), []int{7}
}

func (x *UserPasswordResetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserPasswordResetRequest) GetDomainId() string {
	if x != nil {
		return x.DomainId
	}
	return ""
}

var File_spaceone_api_identity_v2_user_profile_proto protoreflect.FileDescriptor

var file_spaceone_api_identity_v2_user_profile_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x28, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x2a, 0x0a, 0x12, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x36, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a,
	0x10, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x66, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x66, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4d,
	0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x50, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x49, 0x64, 0x32, 0xa8, 0x0a, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x06, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01, 0x2a,
	0x22, 0x20, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x96, 0x01, 0x0a,
	0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2d,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x2d,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x8b, 0x01, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x33,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x8d, 0x01, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d,
	0x66, 0x61, 0x12, 0x2a, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x2d,
	0x6d, 0x66, 0x61, 0x12, 0x90, 0x01, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6d, 0x66, 0x61, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22,
	0x25, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x2d, 0x6d, 0x66, 0x61, 0x12, 0x90, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x5f, 0x6d, 0x66, 0x61, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x4d, 0x46, 0x41, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a,
	0x01, 0x2a, 0x22, 0x25, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x2d, 0x6d, 0x66, 0x61, 0x12, 0x81, 0x01, 0x0a, 0x03, 0x67, 0x65,
	0x74, 0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d,
	0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x9d, 0x01,
	0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x12, 0x2c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x67,
	0x65, 0x74, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x66, 0x6f, 0x72, 0x65, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x73, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6f, 0x6e, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spaceone_api_identity_v2_user_profile_proto_rawDescOnce sync.Once
	file_spaceone_api_identity_v2_user_profile_proto_rawDescData = file_spaceone_api_identity_v2_user_profile_proto_rawDesc
)

func file_spaceone_api_identity_v2_user_profile_proto_rawDescGZIP() []byte {
	file_spaceone_api_identity_v2_user_profile_proto_rawDescOnce.Do(func() {
		file_spaceone_api_identity_v2_user_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_spaceone_api_identity_v2_user_profile_proto_rawDescData)
	})
	return file_spaceone_api_identity_v2_user_profile_proto_rawDescData
}

var file_spaceone_api_identity_v2_user_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_spaceone_api_identity_v2_user_profile_proto_goTypes = []interface{}{
	(*UpdateUserProfileRequest)(nil), // 0: spaceone.api.identity.v2.UpdateUserProfileRequest
	(*VerifyEmailRequest)(nil),       // 1: spaceone.api.identity.v2.VerifyEmailRequest
	(*ConfirmEmailRequest)(nil),      // 2: spaceone.api.identity.v2.ConfirmEmailRequest
	(*EnableMFARequest)(nil),         // 3: spaceone.api.identity.v2.EnableMFARequest
	(*DisableMFARequest)(nil),        // 4: spaceone.api.identity.v2.DisableMFARequest
	(*ConfirmMFARequest)(nil),        // 5: spaceone.api.identity.v2.ConfirmMFARequest
	(*UserProfileRequest)(nil),       // 6: spaceone.api.identity.v2.UserProfileRequest
	(*UserPasswordResetRequest)(nil), // 7: spaceone.api.identity.v2.UserPasswordResetRequest
	(*_struct.Struct)(nil),           // 8: google.protobuf.Struct
	(*UserInfo)(nil),                 // 9: spaceone.api.identity.v2.UserInfo
	(*empty.Empty)(nil),              // 10: google.protobuf.Empty
	(*WorkspacesInfo)(nil),           // 11: spaceone.api.identity.v2.WorkspacesInfo
}
var file_spaceone_api_identity_v2_user_profile_proto_depIdxs = []int32{
	8,  // 0: spaceone.api.identity.v2.UpdateUserProfileRequest.tags:type_name -> google.protobuf.Struct
	8,  // 1: spaceone.api.identity.v2.EnableMFARequest.options:type_name -> google.protobuf.Struct
	0,  // 2: spaceone.api.identity.v2.UserProfile.update:input_type -> spaceone.api.identity.v2.UpdateUserProfileRequest
	1,  // 3: spaceone.api.identity.v2.UserProfile.verify_email:input_type -> spaceone.api.identity.v2.VerifyEmailRequest
	2,  // 4: spaceone.api.identity.v2.UserProfile.confirm_email:input_type -> spaceone.api.identity.v2.ConfirmEmailRequest
	6,  // 5: spaceone.api.identity.v2.UserProfile.reset_password:input_type -> spaceone.api.identity.v2.UserProfileRequest
	3,  // 6: spaceone.api.identity.v2.UserProfile.enable_mfa:input_type -> spaceone.api.identity.v2.EnableMFARequest
	4,  // 7: spaceone.api.identity.v2.UserProfile.disable_mfa:input_type -> spaceone.api.identity.v2.DisableMFARequest
	5,  // 8: spaceone.api.identity.v2.UserProfile.confirm_mfa:input_type -> spaceone.api.identity.v2.ConfirmMFARequest
	6,  // 9: spaceone.api.identity.v2.UserProfile.get:input_type -> spaceone.api.identity.v2.UserProfileRequest
	6,  // 10: spaceone.api.identity.v2.UserProfile.get_workspaces:input_type -> spaceone.api.identity.v2.UserProfileRequest
	9,  // 11: spaceone.api.identity.v2.UserProfile.update:output_type -> spaceone.api.identity.v2.UserInfo
	10, // 12: spaceone.api.identity.v2.UserProfile.verify_email:output_type -> google.protobuf.Empty
	9,  // 13: spaceone.api.identity.v2.UserProfile.confirm_email:output_type -> spaceone.api.identity.v2.UserInfo
	10, // 14: spaceone.api.identity.v2.UserProfile.reset_password:output_type -> google.protobuf.Empty
	9,  // 15: spaceone.api.identity.v2.UserProfile.enable_mfa:output_type -> spaceone.api.identity.v2.UserInfo
	9,  // 16: spaceone.api.identity.v2.UserProfile.disable_mfa:output_type -> spaceone.api.identity.v2.UserInfo
	9,  // 17: spaceone.api.identity.v2.UserProfile.confirm_mfa:output_type -> spaceone.api.identity.v2.UserInfo
	9,  // 18: spaceone.api.identity.v2.UserProfile.get:output_type -> spaceone.api.identity.v2.UserInfo
	11, // 19: spaceone.api.identity.v2.UserProfile.get_workspaces:output_type -> spaceone.api.identity.v2.WorkspacesInfo
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_spaceone_api_identity_v2_user_profile_proto_init() }
func file_spaceone_api_identity_v2_user_profile_proto_init() {
	if File_spaceone_api_identity_v2_user_profile_proto != nil {
		return
	}
	file_spaceone_api_identity_v2_workspace_proto_init()
	file_spaceone_api_identity_v2_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserProfileRequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyEmailRequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmEmailRequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableMFARequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableMFARequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmMFARequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfileRequest); i {
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
		file_spaceone_api_identity_v2_user_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPasswordResetRequest); i {
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
			RawDescriptor: file_spaceone_api_identity_v2_user_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spaceone_api_identity_v2_user_profile_proto_goTypes,
		DependencyIndexes: file_spaceone_api_identity_v2_user_profile_proto_depIdxs,
		MessageInfos:      file_spaceone_api_identity_v2_user_profile_proto_msgTypes,
	}.Build()
	File_spaceone_api_identity_v2_user_profile_proto = out.File
	file_spaceone_api_identity_v2_user_profile_proto_rawDesc = nil
	file_spaceone_api_identity_v2_user_profile_proto_goTypes = nil
	file_spaceone_api_identity_v2_user_profile_proto_depIdxs = nil
}