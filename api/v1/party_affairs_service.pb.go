// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: party_affairs_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_party_affairs_service_proto protoreflect.FileDescriptor

var file_party_affairs_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x21, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x6e,
	0x65, 0x77, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72,
	0x73, 0x5f, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73,
	0x5f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x5f, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc2, 0x20, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x4e, 0x65, 0x77, 0x73, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x57, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x51, 0x5a, 0x28, 0x12, 0x26, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x25,
	0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x7a, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x22, 0x25, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x12, 0x80, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01, 0x2a, 0x1a, 0x25, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x79, 0x12, 0x7d, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65,
	0x77, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x20, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x2a, 0x25, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x12, 0x98, 0x01, 0x0a, 0x0f, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x51, 0x5a, 0x28, 0x12,
	0x26, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d,
	0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x8b,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x4e, 0x65, 0x77, 0x73, 0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x5a, 0x27, 0x12,
	0x25, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x70, 0x0a, 0x0e,
	0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x76,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x01, 0x2a,
	0x1a, 0x24, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x73, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66,
	0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x65, 0x77, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0xaf, 0x01, 0x0a, 0x13,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x59, 0x5a, 0x2c, 0x12, 0x2a, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x79, 0x12, 0x29, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x86, 0x01,
	0x0a, 0x13, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x8c, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x1a, 0x29, 0x2f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x79, 0x12, 0x89, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79,
	0x12, 0x24, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x31,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x2a, 0x29, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x12, 0xac, 0x01, 0x0a, 0x13, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x5f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x59, 0x5a, 0x2c, 0x12, 0x2a, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66,
	0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x9f, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x5d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x57, 0x5a, 0x2b, 0x12, 0x29, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x28, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x33, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61,
	0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x1a, 0x28, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x7f, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x2a, 0x28, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x42, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x6c, 0x6c, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x5a, 0x22, 0x12, 0x20, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x1f, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x73,
	0x12, 0x67, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x6a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66,
	0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x47, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x5a, 0x20, 0x12, 0x1e, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1d, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x74, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x5a,
	0x1f, 0x12, 0x1d, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72,
	0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x1c, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73,
	0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x61,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61,
	0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x12, 0x67, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12,
	0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x1a, 0x1c, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x64, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1e, 0x2a, 0x1c, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69,
	0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x12, 0x74, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1b, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x25, 0x12, 0x23, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72,
	0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x51, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4b, 0x5a, 0x25, 0x12,
	0x23, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x22, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66,
	0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x72, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d,
	0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x78, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01,
	0x2a, 0x1a, 0x23, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72,
	0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x74, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x2a, 0x22, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x2d, 0x61, 0x66, 0x66, 0x61, 0x69, 0x72, 0x73, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0x5a, 0x07,
	0x2e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_party_affairs_service_proto_goTypes = []interface{}{
	(*emptypb.Empty)(nil),                 // 0: google.protobuf.Empty
	(*AddNewsClassifyRequest)(nil),        // 1: admin.AddNewsClassifyRequest
	(*UpdateNewsClassifyRequest)(nil),     // 2: admin.UpdateNewsClassifyRequest
	(*DeleteNewsClassifyRequest)(nil),     // 3: admin.DeleteNewsClassifyRequest
	(*PageNewsRequest)(nil),               // 4: admin.PageNewsRequest
	(*GetNewsRequest)(nil),                // 5: admin.GetNewsRequest
	(*AddNewsRequest)(nil),                // 6: admin.AddNewsRequest
	(*UpdateNewsRequest)(nil),             // 7: admin.UpdateNewsRequest
	(*DeleteNewsRequest)(nil),             // 8: admin.DeleteNewsRequest
	(*AddResourceClassifyRequest)(nil),    // 9: admin.AddResourceClassifyRequest
	(*UpdateResourceClassifyRequest)(nil), // 10: admin.UpdateResourceClassifyRequest
	(*DeleteResourceClassifyRequest)(nil), // 11: admin.DeleteResourceClassifyRequest
	(*PageResourceRequest)(nil),           // 12: admin.PageResourceRequest
	(*GetResourceRequest)(nil),            // 13: admin.GetResourceRequest
	(*AddResourceRequest)(nil),            // 14: admin.AddResourceRequest
	(*UpdateResourceRequest)(nil),         // 15: admin.UpdateResourceRequest
	(*DeleteResourceRequest)(nil),         // 16: admin.DeleteResourceRequest
	(*AddBannerRequest)(nil),              // 17: admin.AddBannerRequest
	(*UpdateBannerRequest)(nil),           // 18: admin.UpdateBannerRequest
	(*DeleteBannerRequest)(nil),           // 19: admin.DeleteBannerRequest
	(*PageTaskRequest)(nil),               // 20: admin.PageTaskRequest
	(*GetTaskRequest)(nil),                // 21: admin.GetTaskRequest
	(*AddTaskRequest)(nil),                // 22: admin.AddTaskRequest
	(*UpdateTaskRequest)(nil),             // 23: admin.UpdateTaskRequest
	(*DeleteTaskRequest)(nil),             // 24: admin.DeleteTaskRequest
	(*PageTaskValueRequest)(nil),          // 25: admin.PageTaskValueRequest
	(*GetTaskValueRequest)(nil),           // 26: admin.GetTaskValueRequest
	(*AddTaskValueRequest)(nil),           // 27: admin.AddTaskValueRequest
	(*UpdateTaskValueRequest)(nil),        // 28: admin.UpdateTaskValueRequest
	(*DeleteTaskValueRequest)(nil),        // 29: admin.DeleteTaskValueRequest
	(*AllNewsClassifyReply)(nil),          // 30: admin.AllNewsClassifyReply
	(*PageNewsReply)(nil),                 // 31: admin.PageNewsReply
	(*News)(nil),                          // 32: admin.News
	(*AllResourceClassifyReply)(nil),      // 33: admin.AllResourceClassifyReply
	(*PageResourceReply)(nil),             // 34: admin.PageResourceReply
	(*Resource)(nil),                      // 35: admin.Resource
	(*AllBannerReply)(nil),                // 36: admin.AllBannerReply
	(*PageTaskReply)(nil),                 // 37: admin.PageTaskReply
	(*Task)(nil),                          // 38: admin.Task
	(*PageTaskValueReply)(nil),            // 39: admin.PageTaskValueReply
	(*TaskValue)(nil),                     // 40: admin.TaskValue
}
var file_party_affairs_service_proto_depIdxs = []int32{
	0,  // 0: admin.Service.AllNewsClassify:input_type -> google.protobuf.Empty
	1,  // 1: admin.Service.AddNewsClassify:input_type -> admin.AddNewsClassifyRequest
	2,  // 2: admin.Service.UpdateNewsClassify:input_type -> admin.UpdateNewsClassifyRequest
	3,  // 3: admin.Service.DeleteNewsClassify:input_type -> admin.DeleteNewsClassifyRequest
	4,  // 4: admin.Service.PageNewsContent:input_type -> admin.PageNewsRequest
	5,  // 5: admin.Service.GetNewsContent:input_type -> admin.GetNewsRequest
	6,  // 6: admin.Service.AddNewsContent:input_type -> admin.AddNewsRequest
	7,  // 7: admin.Service.UpdateNewsContent:input_type -> admin.UpdateNewsRequest
	8,  // 8: admin.Service.DeleteNewsContent:input_type -> admin.DeleteNewsRequest
	0,  // 9: admin.Service.AllResourceClassify:input_type -> google.protobuf.Empty
	9,  // 10: admin.Service.AddResourceClassify:input_type -> admin.AddResourceClassifyRequest
	10, // 11: admin.Service.UpdateResourceClassify:input_type -> admin.UpdateResourceClassifyRequest
	11, // 12: admin.Service.DeleteResourceClassify:input_type -> admin.DeleteResourceClassifyRequest
	12, // 13: admin.Service.PageResourceContent:input_type -> admin.PageResourceRequest
	13, // 14: admin.Service.GetResourceContent:input_type -> admin.GetResourceRequest
	14, // 15: admin.Service.AddResourceContent:input_type -> admin.AddResourceRequest
	15, // 16: admin.Service.UpdateResourceContent:input_type -> admin.UpdateResourceRequest
	16, // 17: admin.Service.DeleteResourceContent:input_type -> admin.DeleteResourceRequest
	0,  // 18: admin.Service.AllBanner:input_type -> google.protobuf.Empty
	17, // 19: admin.Service.AddBanner:input_type -> admin.AddBannerRequest
	18, // 20: admin.Service.UpdateBanner:input_type -> admin.UpdateBannerRequest
	19, // 21: admin.Service.DeleteBanner:input_type -> admin.DeleteBannerRequest
	20, // 22: admin.Service.PageTask:input_type -> admin.PageTaskRequest
	21, // 23: admin.Service.GetTask:input_type -> admin.GetTaskRequest
	22, // 24: admin.Service.AddTask:input_type -> admin.AddTaskRequest
	23, // 25: admin.Service.UpdateTask:input_type -> admin.UpdateTaskRequest
	24, // 26: admin.Service.DeleteTask:input_type -> admin.DeleteTaskRequest
	25, // 27: admin.Service.PageTaskValue:input_type -> admin.PageTaskValueRequest
	26, // 28: admin.Service.GetTaskValue:input_type -> admin.GetTaskValueRequest
	27, // 29: admin.Service.AddTaskValue:input_type -> admin.AddTaskValueRequest
	28, // 30: admin.Service.UpdateTaskValue:input_type -> admin.UpdateTaskValueRequest
	29, // 31: admin.Service.DeleteTaskValue:input_type -> admin.DeleteTaskValueRequest
	30, // 32: admin.Service.AllNewsClassify:output_type -> admin.AllNewsClassifyReply
	0,  // 33: admin.Service.AddNewsClassify:output_type -> google.protobuf.Empty
	0,  // 34: admin.Service.UpdateNewsClassify:output_type -> google.protobuf.Empty
	0,  // 35: admin.Service.DeleteNewsClassify:output_type -> google.protobuf.Empty
	31, // 36: admin.Service.PageNewsContent:output_type -> admin.PageNewsReply
	32, // 37: admin.Service.GetNewsContent:output_type -> admin.News
	0,  // 38: admin.Service.AddNewsContent:output_type -> google.protobuf.Empty
	0,  // 39: admin.Service.UpdateNewsContent:output_type -> google.protobuf.Empty
	0,  // 40: admin.Service.DeleteNewsContent:output_type -> google.protobuf.Empty
	33, // 41: admin.Service.AllResourceClassify:output_type -> admin.AllResourceClassifyReply
	0,  // 42: admin.Service.AddResourceClassify:output_type -> google.protobuf.Empty
	0,  // 43: admin.Service.UpdateResourceClassify:output_type -> google.protobuf.Empty
	0,  // 44: admin.Service.DeleteResourceClassify:output_type -> google.protobuf.Empty
	34, // 45: admin.Service.PageResourceContent:output_type -> admin.PageResourceReply
	35, // 46: admin.Service.GetResourceContent:output_type -> admin.Resource
	0,  // 47: admin.Service.AddResourceContent:output_type -> google.protobuf.Empty
	0,  // 48: admin.Service.UpdateResourceContent:output_type -> google.protobuf.Empty
	0,  // 49: admin.Service.DeleteResourceContent:output_type -> google.protobuf.Empty
	36, // 50: admin.Service.AllBanner:output_type -> admin.AllBannerReply
	0,  // 51: admin.Service.AddBanner:output_type -> google.protobuf.Empty
	0,  // 52: admin.Service.UpdateBanner:output_type -> google.protobuf.Empty
	0,  // 53: admin.Service.DeleteBanner:output_type -> google.protobuf.Empty
	37, // 54: admin.Service.PageTask:output_type -> admin.PageTaskReply
	38, // 55: admin.Service.GetTask:output_type -> admin.Task
	0,  // 56: admin.Service.AddTask:output_type -> google.protobuf.Empty
	0,  // 57: admin.Service.UpdateTask:output_type -> google.protobuf.Empty
	0,  // 58: admin.Service.DeleteTask:output_type -> google.protobuf.Empty
	39, // 59: admin.Service.PageTaskValue:output_type -> admin.PageTaskValueReply
	40, // 60: admin.Service.GetTaskValue:output_type -> admin.TaskValue
	0,  // 61: admin.Service.AddTaskValue:output_type -> google.protobuf.Empty
	0,  // 62: admin.Service.UpdateTaskValue:output_type -> google.protobuf.Empty
	0,  // 63: admin.Service.DeleteTaskValue:output_type -> google.protobuf.Empty
	32, // [32:64] is the sub-list for method output_type
	0,  // [0:32] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_party_affairs_service_proto_init() }
func file_party_affairs_service_proto_init() {
	if File_party_affairs_service_proto != nil {
		return
	}
	file_party_affairs_news_classify_proto_init()
	file_party_affairs_news_proto_init()
	file_party_affairs_resource_classify_proto_init()
	file_party_affairs_resource_proto_init()
	file_party_affairs_banner_proto_init()
	file_party_affairs_task_proto_init()
	file_party_affairs_task_value_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_party_affairs_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_affairs_service_proto_goTypes,
		DependencyIndexes: file_party_affairs_service_proto_depIdxs,
	}.Build()
	File_party_affairs_service_proto = out.File
	file_party_affairs_service_proto_rawDesc = nil
	file_party_affairs_service_proto_goTypes = nil
	file_party_affairs_service_proto_depIdxs = nil
}
