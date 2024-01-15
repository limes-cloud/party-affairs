// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: party_affairs_news_classify.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type NewsClassify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight    *uint32 `protobuf:"varint,3,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
	CreatedAt uint32  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt uint32  `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *NewsClassify) Reset() {
	*x = NewsClassify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_news_classify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsClassify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsClassify) ProtoMessage() {}

func (x *NewsClassify) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_news_classify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsClassify.ProtoReflect.Descriptor instead.
func (*NewsClassify) Descriptor() ([]byte, []int) {
	return file_party_affairs_news_classify_proto_rawDescGZIP(), []int{0}
}

func (x *NewsClassify) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewsClassify) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewsClassify) GetWeight() uint32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

func (x *NewsClassify) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NewsClassify) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type AllNewsClassifyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*NewsClassify `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AllNewsClassifyReply) Reset() {
	*x = AllNewsClassifyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_news_classify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllNewsClassifyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllNewsClassifyReply) ProtoMessage() {}

func (x *AllNewsClassifyReply) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_news_classify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllNewsClassifyReply.ProtoReflect.Descriptor instead.
func (*AllNewsClassifyReply) Descriptor() ([]byte, []int) {
	return file_party_affairs_news_classify_proto_rawDescGZIP(), []int{1}
}

func (x *AllNewsClassifyReply) GetList() []*NewsClassify {
	if x != nil {
		return x.List
	}
	return nil
}

type AddNewsClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *AddNewsClassifyRequest) Reset() {
	*x = AddNewsClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_news_classify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddNewsClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddNewsClassifyRequest) ProtoMessage() {}

func (x *AddNewsClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_news_classify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddNewsClassifyRequest.ProtoReflect.Descriptor instead.
func (*AddNewsClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_news_classify_proto_rawDescGZIP(), []int{2}
}

func (x *AddNewsClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddNewsClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type UpdateNewsClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Weight uint32 `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *UpdateNewsClassifyRequest) Reset() {
	*x = UpdateNewsClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_news_classify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNewsClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNewsClassifyRequest) ProtoMessage() {}

func (x *UpdateNewsClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_news_classify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNewsClassifyRequest.ProtoReflect.Descriptor instead.
func (*UpdateNewsClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_news_classify_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNewsClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateNewsClassifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNewsClassifyRequest) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type DeleteNewsClassifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteNewsClassifyRequest) Reset() {
	*x = DeleteNewsClassifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_affairs_news_classify_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNewsClassifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNewsClassifyRequest) ProtoMessage() {}

func (x *DeleteNewsClassifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_affairs_news_classify_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNewsClassifyRequest.ProtoReflect.Descriptor instead.
func (*DeleteNewsClassifyRequest) Descriptor() ([]byte, []int) {
	return file_party_affairs_news_classify_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNewsClassifyRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_party_affairs_news_classify_proto protoreflect.FileDescriptor

var file_party_affairs_news_classify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x6e, 0x65, 0x77, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3f,
	0x0a, 0x14, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4e, 0x65, 0x77,
	0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x56, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52,
	0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x72, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x77, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x28, 0x00, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x34, 0x0a, 0x19, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_affairs_news_classify_proto_rawDescOnce sync.Once
	file_party_affairs_news_classify_proto_rawDescData = file_party_affairs_news_classify_proto_rawDesc
)

func file_party_affairs_news_classify_proto_rawDescGZIP() []byte {
	file_party_affairs_news_classify_proto_rawDescOnce.Do(func() {
		file_party_affairs_news_classify_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_affairs_news_classify_proto_rawDescData)
	})
	return file_party_affairs_news_classify_proto_rawDescData
}

var file_party_affairs_news_classify_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_party_affairs_news_classify_proto_goTypes = []interface{}{
	(*NewsClassify)(nil),              // 0: admin.NewsClassify
	(*AllNewsClassifyReply)(nil),      // 1: admin.AllNewsClassifyReply
	(*AddNewsClassifyRequest)(nil),    // 2: admin.AddNewsClassifyRequest
	(*UpdateNewsClassifyRequest)(nil), // 3: admin.UpdateNewsClassifyRequest
	(*DeleteNewsClassifyRequest)(nil), // 4: admin.DeleteNewsClassifyRequest
}
var file_party_affairs_news_classify_proto_depIdxs = []int32{
	0, // 0: admin.AllNewsClassifyReply.list:type_name -> admin.NewsClassify
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_affairs_news_classify_proto_init() }
func file_party_affairs_news_classify_proto_init() {
	if File_party_affairs_news_classify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_affairs_news_classify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsClassify); i {
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
		file_party_affairs_news_classify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllNewsClassifyReply); i {
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
		file_party_affairs_news_classify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddNewsClassifyRequest); i {
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
		file_party_affairs_news_classify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNewsClassifyRequest); i {
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
		file_party_affairs_news_classify_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNewsClassifyRequest); i {
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
	file_party_affairs_news_classify_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_affairs_news_classify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_party_affairs_news_classify_proto_goTypes,
		DependencyIndexes: file_party_affairs_news_classify_proto_depIdxs,
		MessageInfos:      file_party_affairs_news_classify_proto_msgTypes,
	}.Build()
	File_party_affairs_news_classify_proto = out.File
	file_party_affairs_news_classify_proto_rawDesc = nil
	file_party_affairs_news_classify_proto_goTypes = nil
	file_party_affairs_news_classify_proto_depIdxs = nil
}
