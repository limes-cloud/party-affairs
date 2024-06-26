// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: party_affairs_task_value.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/limes-cloud/resource/api/file/v1"
	v1 "github.com/limes-cloud/user-center/api/user/v1"
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

type TaskValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TaskId    uint32       `protobuf:"varint,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId    uint32       `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Value     string       `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	CreatedAt uint32       `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint32       `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	User      *v1.BaseUser `protobuf:"bytes,7,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *TaskValue) Reset() {
	*x = TaskValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskValue) ProtoMessage() {}

func (x *TaskValue) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskValue.ProtoReflect.Descriptor instead.
func (*TaskValue) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{0}
}

func (x *TaskValue) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskValue) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *TaskValue) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *TaskValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TaskValue) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TaskValue) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TaskValue) GetUser() *v1.BaseUser {
	if x != nil {
		return x.User
	}
	return nil
}

type GetCurTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetCurTaskValueRequest) Reset() {
	*x = GetCurTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurTaskValueRequest) ProtoMessage() {}

func (x *GetCurTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurTaskValueRequest.ProtoReflect.Descriptor instead.
func (*GetCurTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type GetTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetTaskValueRequest) Reset() {
	*x = GetTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskValueRequest) ProtoMessage() {}

func (x *GetTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskValueRequest.ProtoReflect.Descriptor instead.
func (*GetTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{2}
}

func (x *GetTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *GetTaskValueRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PageTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize uint32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TaskId   uint32  `protobuf:"varint,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	UserId   *uint32 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`
}

func (x *PageTaskValueRequest) Reset() {
	*x = PageTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTaskValueRequest) ProtoMessage() {}

func (x *PageTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTaskValueRequest.ProtoReflect.Descriptor instead.
func (*PageTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{3}
}

func (x *PageTaskValueRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageTaskValueRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *PageTaskValueRequest) GetUserId() uint32 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

type PageTaskValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*TaskValue `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageTaskValueReply) Reset() {
	*x = PageTaskValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageTaskValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageTaskValueReply) ProtoMessage() {}

func (x *PageTaskValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageTaskValueReply.ProtoReflect.Descriptor instead.
func (*PageTaskValueReply) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{4}
}

func (x *PageTaskValueReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageTaskValueReply) GetList() []*TaskValue {
	if x != nil {
		return x.List
	}
	return nil
}

type ExportTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ExportTaskValueRequest) Reset() {
	*x = ExportTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportTaskValueRequest) ProtoMessage() {}

func (x *ExportTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportTaskValueRequest.ProtoReflect.Descriptor instead.
func (*ExportTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{5}
}

func (x *ExportTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

type AddTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddTaskValueRequest) Reset() {
	*x = AddTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskValueRequest) ProtoMessage() {}

func (x *AddTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskValueRequest.ProtoReflect.Descriptor instead.
func (*AddTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{6}
}

func (x *AddTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *AddTaskValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UpdateTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32 `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateTaskValueRequest) Reset() {
	*x = UpdateTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskValueRequest) ProtoMessage() {}

func (x *UpdateTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskValueRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateTaskValueRequest) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateTaskValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DeleteTaskValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteTaskValueRequest) Reset() {
	*x = DeleteTaskValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_task_value_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskValueRequest) ProtoMessage() {}

func (x *DeleteTaskValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_task_value_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskValueRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskValueRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_task_value_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTaskValueRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_party_affairs_task_value_proto protoreflect.FileDescriptor

var file_party_affairs_task_value_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0xb0, 0x01, 0x0a, 0x14, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04,
	0x18, 0x32, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x22, 0x56, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02,
	0x20, 0x00, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x59, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x31, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_affairs_task_value_proto_rawDescOnce sync.Once
	file_party_affairs_task_value_proto_rawDescData = file_party_affairs_task_value_proto_rawDesc
)

func file_party_affairs_task_value_proto_rawDescGZIP() []byte {
	file_party_affairs_task_value_proto_rawDescOnce.Do(func() {
		file_party_affairs_task_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_affairs_task_value_proto_rawDescData)
	})
	return file_party_affairs_task_value_proto_rawDescData
}

var file_party_affairs_task_value_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_party_affairs_task_value_proto_goTypes = []interface{}{
	(*TaskValue)(nil),              // 0: admin.TaskValue
	(*GetCurTaskValueRequest)(nil), // 1: admin.GetCurTaskValueRequest
	(*GetTaskValueRequest)(nil),    // 2: admin.GetTaskValueRequest
	(*PageTaskValueRequest)(nil),   // 3: admin.PageTaskValueRequest
	(*PageTaskValueReply)(nil),     // 4: admin.PageTaskValueReply
	(*ExportTaskValueRequest)(nil), // 5: admin.ExportTaskValueRequest
	(*AddTaskValueRequest)(nil),    // 6: admin.AddTaskValueRequest
	(*UpdateTaskValueRequest)(nil), // 7: admin.UpdateTaskValueRequest
	(*DeleteTaskValueRequest)(nil), // 8: admin.DeleteTaskValueRequest
	(*v1.BaseUser)(nil),            // 9: user.BaseUser
}
var file_party_affairs_task_value_proto_depIdxs = []int32{
	9, // 0: admin.TaskValue.user:type_name -> user.BaseUser
	0, // 1: admin.PageTaskValueReply.list:type_name -> admin.TaskValue
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_party_affairs_task_value_proto_init() }
func file_party_affairs_task_value_proto_init() {
	if File_party_affairs_task_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_affairs_task_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskValue); i {
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
		file_party_affairs_task_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageTaskValueReply); i {
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
		file_party_affairs_task_value_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskValueRequest); i {
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
		file_party_affairs_task_value_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskValueRequest); i {
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
	file_party_affairs_task_value_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_affairs_task_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_affairs_task_value_proto_goTypes,
		DependencyIndexes: file_party_affairs_task_value_proto_depIdxs,
		MessageInfos:      file_party_affairs_task_value_proto_msgTypes,
	}.Build()
	File_party_affairs_task_value_proto = out.File
	file_party_affairs_task_value_proto_rawDesc = nil
	file_party_affairs_task_value_proto_goTypes = nil
	file_party_affairs_task_value_proto_depIdxs = nil
}
