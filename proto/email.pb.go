// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v4.25.3
// source: proto/email.proto

package email

import (
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

type AddEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AddEmailRequest) Reset() {
	*x = AddEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmailRequest) ProtoMessage() {}

func (x *AddEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmailRequest.ProtoReflect.Descriptor instead.
func (*AddEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{0}
}

func (x *AddEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEmailResponse) Reset() {
	*x = AddEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEmailResponse) ProtoMessage() {}

func (x *AddEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEmailResponse.ProtoReflect.Descriptor instead.
func (*AddEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{1}
}

func (x *AddEmailResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetEmailRequest) Reset() {
	*x = GetEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailRequest) ProtoMessage() {}

func (x *GetEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailRequest.ProtoReflect.Descriptor instead.
func (*GetEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *EmailEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *GetEmailResponse) Reset() {
	*x = GetEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailResponse) ProtoMessage() {}

func (x *GetEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailResponse.ProtoReflect.Descriptor instead.
func (*GetEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmailResponse) GetEntry() *EmailEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *EmailEntry `protobuf:"bytes,1,opt,name=entry,proto3" json:"entry,omitempty"`
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEmailRequest) GetEntry() *EmailEntry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type UpdateEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateEmailResponse) Reset() {
	*x = UpdateEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailResponse) ProtoMessage() {}

func (x *UpdateEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DeleteEmailRequest) Reset() {
	*x = DeleteEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmailRequest) ProtoMessage() {}

func (x *DeleteEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmailRequest.ProtoReflect.Descriptor instead.
func (*DeleteEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DeleteEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteEmailResponse) Reset() {
	*x = DeleteEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEmailResponse) ProtoMessage() {}

func (x *DeleteEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEmailResponse.ProtoReflect.Descriptor instead.
func (*DeleteEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEmailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetEmailBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetEmailBatchRequest) Reset() {
	*x = GetEmailBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailBatchRequest) ProtoMessage() {}

func (x *GetEmailBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailBatchRequest.ProtoReflect.Descriptor instead.
func (*GetEmailBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{8}
}

func (x *GetEmailBatchRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEmailBatchRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetEmailBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*EmailEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GetEmailBatchResponse) Reset() {
	*x = GetEmailBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailBatchResponse) ProtoMessage() {}

func (x *GetEmailBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailBatchResponse.ProtoReflect.Descriptor instead.
func (*GetEmailBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{9}
}

func (x *GetEmailBatchResponse) GetEntries() []*EmailEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type EmailEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	ConfirmedAt int64  `protobuf:"varint,3,opt,name=confirmed_at,json=confirmedAt,proto3" json:"confirmed_at,omitempty"`
	OptOut      bool   `protobuf:"varint,4,opt,name=opt_out,json=optOut,proto3" json:"opt_out,omitempty"`
}

func (x *EmailEntry) Reset() {
	*x = EmailEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_email_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailEntry) ProtoMessage() {}

func (x *EmailEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_email_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailEntry.ProtoReflect.Descriptor instead.
func (*EmailEntry) Descriptor() ([]byte, []int) {
	return file_proto_email_proto_rawDescGZIP(), []int{10}
}

func (x *EmailEntry) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmailEntry) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailEntry) GetConfirmedAt() int64 {
	if x != nil {
		return x.ConfirmedAt
	}
	return 0
}

func (x *EmailEntry) GetOptOut() bool {
	if x != nil {
		return x.OptOut
	}
	return false
}

var File_proto_email_proto protoreflect.FileDescriptor

var file_proto_email_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3d, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x2f, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x6e, 0x0a, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x5f, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x4f,
	0x75, 0x74, 0x32, 0xe0, 0x02, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x16, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x41, 0x64, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x19, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x6e, 0x65, 0x73, 0x69, 0x6f, 0x75, 0x73, 0x2f, 0x6d,
	0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_email_proto_rawDescOnce sync.Once
	file_proto_email_proto_rawDescData = file_proto_email_proto_rawDesc
)

func file_proto_email_proto_rawDescGZIP() []byte {
	file_proto_email_proto_rawDescOnce.Do(func() {
		file_proto_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_email_proto_rawDescData)
	})
	return file_proto_email_proto_rawDescData
}

var file_proto_email_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_email_proto_goTypes = []interface{}{
	(*AddEmailRequest)(nil),       // 0: email.AddEmailRequest
	(*AddEmailResponse)(nil),      // 1: email.AddEmailResponse
	(*GetEmailRequest)(nil),       // 2: email.GetEmailRequest
	(*GetEmailResponse)(nil),      // 3: email.GetEmailResponse
	(*UpdateEmailRequest)(nil),    // 4: email.UpdateEmailRequest
	(*UpdateEmailResponse)(nil),   // 5: email.UpdateEmailResponse
	(*DeleteEmailRequest)(nil),    // 6: email.DeleteEmailRequest
	(*DeleteEmailResponse)(nil),   // 7: email.DeleteEmailResponse
	(*GetEmailBatchRequest)(nil),  // 8: email.GetEmailBatchRequest
	(*GetEmailBatchResponse)(nil), // 9: email.GetEmailBatchResponse
	(*EmailEntry)(nil),            // 10: email.EmailEntry
}
var file_proto_email_proto_depIdxs = []int32{
	10, // 0: email.GetEmailResponse.entry:type_name -> email.EmailEntry
	10, // 1: email.UpdateEmailRequest.entry:type_name -> email.EmailEntry
	10, // 2: email.GetEmailBatchResponse.entries:type_name -> email.EmailEntry
	0,  // 3: email.EmailService.AddEmail:input_type -> email.AddEmailRequest
	2,  // 4: email.EmailService.GetEmail:input_type -> email.GetEmailRequest
	4,  // 5: email.EmailService.UpdateEmail:input_type -> email.UpdateEmailRequest
	6,  // 6: email.EmailService.DeleteEmail:input_type -> email.DeleteEmailRequest
	8,  // 7: email.EmailService.GetEmailBatch:input_type -> email.GetEmailBatchRequest
	1,  // 8: email.EmailService.AddEmail:output_type -> email.AddEmailResponse
	3,  // 9: email.EmailService.GetEmail:output_type -> email.GetEmailResponse
	5,  // 10: email.EmailService.UpdateEmail:output_type -> email.UpdateEmailResponse
	7,  // 11: email.EmailService.DeleteEmail:output_type -> email.DeleteEmailResponse
	9,  // 12: email.EmailService.GetEmailBatch:output_type -> email.GetEmailBatchResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_email_proto_init() }
func file_proto_email_proto_init() {
	if File_proto_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEmailRequest); i {
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
		file_proto_email_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEmailResponse); i {
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
		file_proto_email_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailRequest); i {
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
		file_proto_email_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailResponse); i {
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
		file_proto_email_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequest); i {
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
		file_proto_email_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailResponse); i {
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
		file_proto_email_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmailRequest); i {
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
		file_proto_email_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEmailResponse); i {
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
		file_proto_email_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailBatchRequest); i {
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
		file_proto_email_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailBatchResponse); i {
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
		file_proto_email_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailEntry); i {
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
			RawDescriptor: file_proto_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_email_proto_goTypes,
		DependencyIndexes: file_proto_email_proto_depIdxs,
		MessageInfos:      file_proto_email_proto_msgTypes,
	}.Build()
	File_proto_email_proto = out.File
	file_proto_email_proto_rawDesc = nil
	file_proto_email_proto_goTypes = nil
	file_proto_email_proto_depIdxs = nil
}
