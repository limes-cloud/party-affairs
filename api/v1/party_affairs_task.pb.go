// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: party_affairs_task.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/limes-cloud/resource/api/v1"
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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Desc      string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	IsUpdate  bool   `protobuf:"varint,4,opt,name=is_update,json=isUpdate,proto3" json:"is_update,omitempty"`
	Start     uint32 `protobuf:"varint,5,opt,name=start,proto3" json:"start,omitempty"`
	End       uint32 `protobuf:"varint,6,opt,name=end,proto3" json:"end,omitempty"`
	Config    string `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	CreatedAt uint32 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint32 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Task) GetIsUpdate() bool {
	if x != nil {
		return x.IsUpdate
	}
	return false
}

func (x *Task) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Task) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Task) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Task) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Task) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{1}
}

func (x *GetTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PageTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize  uint32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Title     *string `protobuf:"bytes,3,opt,name=title,proto3,oneof" json:"title,omitempty"`
	NotFinish *bool   `protobuf:"varint,4,opt,name=not_finish,json=notFinish,proto3,oneof" json:"not_finish,omitempty"`
}

func (x *PageTaskRequest) Reset() {
	*x = PageTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTaskRequest) ProtoMessage() {}

func (x *PageTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTaskRequest.ProtoReflect.Descriptor instead.
func (*PageTaskRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{2}
}

func (x *PageTaskRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageTaskRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageTaskRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *PageTaskRequest) GetNotFinish() bool {
	if x != nil && x.NotFinish != nil {
		return *x.NotFinish
	}
	return false
}

type PageTaskReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32  `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*Task `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageTaskReply) Reset() {
	*x = PageTaskReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTaskReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTaskReply) ProtoMessage() {}

func (x *PageTaskReply) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTaskReply.ProtoReflect.Descriptor instead.
func (*PageTaskReply) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{3}
}

func (x *PageTaskReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageTaskReply) GetList() []*Task {
	if x != nil {
		return x.List
	}
	return nil
}

type AddTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc     string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	IsUpdate *bool  `protobuf:"varint,3,opt,name=is_update,json=isUpdate,proto3,oneof" json:"is_update,omitempty"`
	Start    uint32 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	End      uint32 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	Config   string `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *AddTaskRequest) Reset() {
	*x = AddTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskRequest) ProtoMessage() {}

func (x *AddTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskRequest.ProtoReflect.Descriptor instead.
func (*AddTaskRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{4}
}

func (x *AddTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddTaskRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *AddTaskRequest) GetIsUpdate() bool {
	if x != nil && x.IsUpdate != nil {
		return *x.IsUpdate
	}
	return false
}

func (x *AddTaskRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *AddTaskRequest) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *AddTaskRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc     string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	IsUpdate *bool  `protobuf:"varint,3,opt,name=is_update,json=isUpdate,proto3,oneof" json:"is_update,omitempty"`
	Start    uint32 `protobuf:"varint,4,opt,name=start,proto3" json:"start,omitempty"`
	End      uint32 `protobuf:"varint,5,opt,name=end,proto3" json:"end,omitempty"`
	Config   string `protobuf:"bytes,6,opt,name=config,proto3" json:"config,omitempty"`
	Id       uint32 `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTaskRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTaskRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UpdateTaskRequest) GetIsUpdate() bool {
	if x != nil && x.IsUpdate != nil {
		return *x.IsUpdate
	}
	return false
}

func (x *UpdateTaskRequest) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *UpdateTaskRequest) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *UpdateTaskRequest) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *UpdateTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTaskRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_party_affairs_task_proto protoreflect.FileDescriptor

var file_party_affairs_task_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xae, 0x01, 0x0a,
	0x0f, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x32, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x22, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x22, 0x46, 0x0a,
	0x0d, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xd7, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x01, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64,
	0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22,
	0xf3, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x20, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x01, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1f, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x69, 0x73, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_affairs_task_proto_rawDescOnce sync.Once
	file_party_affairs_task_proto_rawDescData = file_party_affairs_task_proto_rawDesc
)

func file_party_affairs_task_proto_rawDescGZIP() []byte {
	file_party_affairs_task_proto_rawDescOnce.Do(func() {
		file_party_affairs_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_affairs_task_proto_rawDescData)
	})
	return file_party_affairs_task_proto_rawDescData
}

var file_party_affairs_task_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_party_affairs_task_proto_goTypes = []interface{}{
	(*Task)(nil),              // 0: admin.Task
	(*GetTaskRequest)(nil),    // 1: admin.GetTaskRequest
	(*PageTaskRequest)(nil),   // 2: admin.PageTaskRequest
	(*PageTaskReply)(nil),     // 3: admin.PageTaskReply
	(*AddTaskRequest)(nil),    // 4: admin.AddTaskRequest
	(*UpdateTaskRequest)(nil), // 5: admin.UpdateTaskRequest
	(*DeleteTaskRequest)(nil), // 6: admin.DeleteTaskRequest
}
var file_party_affairs_task_proto_depIdxs = []int32{
	0, // 0: admin.PageTaskReply.list:type_name -> admin.Task
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_affairs_task_proto_init() }
func file_party_affairs_task_proto_init() {
	if File_party_affairs_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_affairs_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_party_affairs_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRequest); i {
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
		file_party_affairs_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTaskRequest); i {
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
		file_party_affairs_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTaskReply); i {
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
		file_party_affairs_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskRequest); i {
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
		file_party_affairs_task_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRequest); i {
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
		file_party_affairs_task_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRequest); i {
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
	file_party_affairs_task_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_party_affairs_task_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_party_affairs_task_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_affairs_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_affairs_task_proto_goTypes,
		DependencyIndexes: file_party_affairs_task_proto_depIdxs,
		MessageInfos:      file_party_affairs_task_proto_msgTypes,
	}.Build()
	File_party_affairs_task_proto = out.File
	file_party_affairs_task_proto_rawDesc = nil
	file_party_affairs_task_proto_goTypes = nil
	file_party_affairs_task_proto_depIdxs = nil
}
