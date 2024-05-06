// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: party_affairs_resource.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/limes-cloud/resource/api/file/v1"
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

type ResourceClassify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight    *uint32 `protobuf:"varint,3,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	CreatedAt uint32  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint32  `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *ResourceClassify) Reset() {
	*x = ResourceClassify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceClassify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceClassify) ProtoMessage() {}

func (x *ResourceClassify) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceClassify.ProtoReflect.Descriptor instead.
func (*ResourceClassify) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceClassify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResourceClassify) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceClassify) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *ResourceClassify) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ResourceClassify) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type AllResourceClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ResourceClassify `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AllResourceClassifyReply) Reset() {
	*x = AllResourceClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllResourceClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllResourceClassifyReply) ProtoMessage() {}

func (x *AllResourceClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllResourceClassifyReply.ProtoReflect.Descriptor instead.
func (*AllResourceClassifyReply) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{1}
}

func (x *AllResourceClassifyReply) GetList() []*ResourceClassify {
	if x != nil {
		return x.List
	}
	return nil
}

type AddResourceClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *AddResourceClassifyRequest) Reset() {
	*x = AddResourceClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResourceClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResourceClassifyRequest) ProtoMessage() {}

func (x *AddResourceClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResourceClassifyRequest.ProtoReflect.Descriptor instead.
func (*AddResourceClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{2}
}

func (x *AddResourceClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddResourceClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type UpdateResourceClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *UpdateResourceClassifyRequest) Reset() {
	*x = UpdateResourceClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceClassifyRequest) ProtoMessage() {}

func (x *UpdateResourceClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceClassifyRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResourceClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateResourceClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResourceClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type DeleteResourceClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteResourceClassifyRequest) Reset() {
	*x = DeleteResourceClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceClassifyRequest) ProtoMessage() {}

func (x *DeleteResourceClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceClassifyRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteResourceClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ResourceContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title            string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Desc             string            `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Url              string            `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	DownloadCount    uint32            `protobuf:"varint,6,opt,name=download_count,json=downloadCount,proto3" json:"download_count,omitempty"`
	ClassifyId       uint32            `protobuf:"varint,7,opt,name=classify_id,json=classifyId,proto3" json:"classify_id,omitempty"`
	CreatedAt        uint32            `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        uint32            `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ResourceClassify *ResourceClassify `protobuf:"bytes,10,opt,name=resource_classify,json=resourceClassify,proto3" json:"resource_classify,omitempty"`
	Resource         *v1.File          `protobuf:"bytes,13,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ResourceContent) Reset() {
	*x = ResourceContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceContent) ProtoMessage() {}

func (x *ResourceContent) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceContent.ProtoReflect.Descriptor instead.
func (*ResourceContent) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceContent) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResourceContent) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ResourceContent) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ResourceContent) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ResourceContent) GetDownloadCount() uint32 {
	if x != nil {
		return x.DownloadCount
	}
	return 0
}

func (x *ResourceContent) GetClassifyId() uint32 {
	if x != nil {
		return x.ClassifyId
	}
	return 0
}

func (x *ResourceContent) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ResourceContent) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *ResourceContent) GetResourceClassify() *ResourceClassify {
	if x != nil {
		return x.ResourceClassify
	}
	return nil
}

func (x *ResourceContent) GetResource() *v1.File {
	if x != nil {
		return x.Resource
	}
	return nil
}

type GetResourceContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetResourceContentRequest) Reset() {
	*x = GetResourceContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourceContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceContentRequest) ProtoMessage() {}

func (x *GetResourceContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceContentRequest.ProtoReflect.Descriptor instead.
func (*GetResourceContentRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{6}
}

func (x *GetResourceContentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PageResourceContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       uint32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   uint32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	ClassifyId *uint32 `protobuf:"varint,3,opt,name=classify_id,json=classifyId,proto3,oneof" json:"classify_id,omitempty"`
	Title      *string `protobuf:"bytes,4,opt,name=title,proto3,oneof" json:"title,omitempty"`
}

func (x *PageResourceContentRequest) Reset() {
	*x = PageResourceContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResourceContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResourceContentRequest) ProtoMessage() {}

func (x *PageResourceContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResourceContentRequest.ProtoReflect.Descriptor instead.
func (*PageResourceContentRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{7}
}

func (x *PageResourceContentRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageResourceContentRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageResourceContentRequest) GetClassifyId() uint32 {
	if x != nil && x.ClassifyId != nil {
		return *x.ClassifyId
	}
	return 0
}

func (x *PageResourceContentRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type PageResourceContentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32             `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ResourceContent `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PageResourceContentReply) Reset() {
	*x = PageResourceContentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageResourceContentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageResourceContentReply) ProtoMessage() {}

func (x *PageResourceContentReply) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageResourceContentReply.ProtoReflect.Descriptor instead.
func (*PageResourceContentReply) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{8}
}

func (x *PageResourceContentReply) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageResourceContentReply) GetList() []*ResourceContent {
	if x != nil {
		return x.List
	}
	return nil
}

type AddResourceContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc       string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Url        string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	ClassifyId uint32 `protobuf:"varint,4,opt,name=classify_id,json=classifyId,proto3" json:"classify_id,omitempty"`
}

func (x *AddResourceContentRequest) Reset() {
	*x = AddResourceContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResourceContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResourceContentRequest) ProtoMessage() {}

func (x *AddResourceContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResourceContentRequest.ProtoReflect.Descriptor instead.
func (*AddResourceContentRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{9}
}

func (x *AddResourceContentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddResourceContentRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *AddResourceContentRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *AddResourceContentRequest) GetClassifyId() uint32 {
	if x != nil {
		return x.ClassifyId
	}
	return 0
}

type UpdateResourceContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc       string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Url        string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	ClassifyId uint32 `protobuf:"varint,4,opt,name=classify_id,json=classifyId,proto3" json:"classify_id,omitempty"`
	Id         uint32 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateResourceContentRequest) Reset() {
	*x = UpdateResourceContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResourceContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResourceContentRequest) ProtoMessage() {}

func (x *UpdateResourceContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResourceContentRequest.ProtoReflect.Descriptor instead.
func (*UpdateResourceContentRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateResourceContentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateResourceContentRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UpdateResourceContentRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *UpdateResourceContentRequest) GetClassifyId() uint32 {
	if x != nil {
		return x.ClassifyId
	}
	return 0
}

func (x *UpdateResourceContentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResourceContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteResourceContentRequest) Reset() {
	*x = DeleteResourceContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_resource_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResourceContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResourceContentRequest) ProtoMessage() {}

func (x *DeleteResourceContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_resource_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResourceContentRequest.ProtoReflect.Descriptor instead.
func (*DeleteResourceContentRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_resource_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteResourceContentRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_party_affairs_resource_proto protoreflect.FileDescriptor

var file_party_affairs_resource_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x47, 0x0a, 0x18, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x5a, 0x0a, 0x1a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x28, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x76, 0x0a, 0x1d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x38, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd5, 0x02,
	0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x34, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0xbc, 0x01, 0x0a, 0x1a,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a,
	0x04, 0x18, 0x32, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x24, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x5f, 0x69, 0x64,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x5c, 0x0a, 0x18, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2a, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x19, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a,
	0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x28, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x37, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_affairs_resource_proto_rawDescOnce sync.Once
	file_party_affairs_resource_proto_rawDescData = file_party_affairs_resource_proto_rawDesc
)

func file_party_affairs_resource_proto_rawDescGZIP() []byte {
	file_party_affairs_resource_proto_rawDescOnce.Do(func() {
		file_party_affairs_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_affairs_resource_proto_rawDescData)
	})
	return file_party_affairs_resource_proto_rawDescData
}

var file_party_affairs_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_party_affairs_resource_proto_goTypes = []interface{}{
	(*ResourceClassify)(nil),              // 0: admin.ResourceClassify
	(*AllResourceClassifyReply)(nil),      // 1: admin.AllResourceClassifyReply
	(*AddResourceClassifyRequest)(nil),    // 2: admin.AddResourceClassifyRequest
	(*UpdateResourceClassifyRequest)(nil), // 3: admin.UpdateResourceClassifyRequest
	(*DeleteResourceClassifyRequest)(nil), // 4: admin.DeleteResourceClassifyRequest
	(*ResourceContent)(nil),               // 5: admin.ResourceContent
	(*GetResourceContentRequest)(nil),     // 6: admin.GetResourceContentRequest
	(*PageResourceContentRequest)(nil),    // 7: admin.PageResourceContentRequest
	(*PageResourceContentReply)(nil),      // 8: admin.PageResourceContentReply
	(*AddResourceContentRequest)(nil),     // 9: admin.AddResourceContentRequest
	(*UpdateResourceContentRequest)(nil),  // 10: admin.UpdateResourceContentRequest
	(*DeleteResourceContentRequest)(nil),  // 11: admin.DeleteResourceContentRequest
	(*v1.File)(nil),                       // 12: resource.File
}
var file_party_affairs_resource_proto_depIdxs = []int32{
	0,  // 0: admin.AllResourceClassifyReply.list:type_name -> admin.ResourceClassify
	0,  // 1: admin.ResourceContent.resource_classify:type_name -> admin.ResourceClassify
	12, // 2: admin.ResourceContent.resource:type_name -> resource.File
	5,  // 3: admin.PageResourceContentReply.list:type_name -> admin.ResourceContent
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_party_affairs_resource_proto_init() }
func file_party_affairs_resource_proto_init() {
	if File_party_affairs_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_affairs_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceClassify); i {
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
		file_party_affairs_resource_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllResourceClassifyReply); i {
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
		file_party_affairs_resource_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResourceClassifyRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceClassifyRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceClassifyRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceContent); i {
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
		file_party_affairs_resource_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourceContentRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResourceContentRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageResourceContentReply); i {
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
		file_party_affairs_resource_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResourceContentRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResourceContentRequest); i {
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
		file_party_affairs_resource_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResourceContentRequest); i {
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
	file_party_affairs_resource_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_party_affairs_resource_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_affairs_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_affairs_resource_proto_goTypes,
		DependencyIndexes: file_party_affairs_resource_proto_depIdxs,
		MessageInfos:      file_party_affairs_resource_proto_msgTypes,
	}.Build()
	File_party_affairs_resource_proto = out.File
	file_party_affairs_resource_proto_rawDesc = nil
	file_party_affairs_resource_proto_goTypes = nil
	file_party_affairs_resource_proto_depIdxs = nil
}
