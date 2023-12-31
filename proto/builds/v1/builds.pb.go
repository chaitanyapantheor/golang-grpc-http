// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: builds/v1/builds.proto

package buildsv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type AddBuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *AddBuildRequest) Reset() {
	*x = AddBuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builds_v1_builds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBuildRequest) ProtoMessage() {}

func (x *AddBuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builds_v1_builds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBuildRequest.ProtoReflect.Descriptor instead.
func (*AddBuildRequest) Descriptor() ([]byte, []int) {
	return file_builds_v1_builds_proto_rawDescGZIP(), []int{0}
}

func (x *AddBuildRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type AddBuildResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Build *Build `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
}

func (x *AddBuildResponse) Reset() {
	*x = AddBuildResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builds_v1_builds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBuildResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBuildResponse) ProtoMessage() {}

func (x *AddBuildResponse) ProtoReflect() protoreflect.Message {
	mi := &file_builds_v1_builds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBuildResponse.ProtoReflect.Descriptor instead.
func (*AddBuildResponse) Descriptor() ([]byte, []int) {
	return file_builds_v1_builds_proto_rawDescGZIP(), []int{1}
}

func (x *AddBuildResponse) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

type ListBuildsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBuildsRequest) Reset() {
	*x = ListBuildsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builds_v1_builds_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsRequest) ProtoMessage() {}

func (x *ListBuildsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_builds_v1_builds_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsRequest.ProtoReflect.Descriptor instead.
func (*ListBuildsRequest) Descriptor() ([]byte, []int) {
	return file_builds_v1_builds_proto_rawDescGZIP(), []int{2}
}

type ListBuildsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Build *Build `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
}

func (x *ListBuildsResponse) Reset() {
	*x = ListBuildsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builds_v1_builds_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBuildsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBuildsResponse) ProtoMessage() {}

func (x *ListBuildsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_builds_v1_builds_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBuildsResponse.ProtoReflect.Descriptor instead.
func (*ListBuildsResponse) Descriptor() ([]byte, []int) {
	return file_builds_v1_builds_proto_rawDescGZIP(), []int{3}
}

func (x *ListBuildsResponse) GetBuild() *Build {
	if x != nil {
		return x.Build
	}
	return nil
}

type Build struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Build) Reset() {
	*x = Build{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builds_v1_builds_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Build) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Build) ProtoMessage() {}

func (x *Build) ProtoReflect() protoreflect.Message {
	mi := &file_builds_v1_builds_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Build.ProtoReflect.Descriptor instead.
func (*Build) Descriptor() ([]byte, []int) {
	return file_builds_v1_builds_proto_rawDescGZIP(), []int{4}
}

func (x *Build) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Build) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_builds_v1_builds_proto protoreflect.FileDescriptor

var file_builds_v1_builds_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3a, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x2d, 0x0a, 0x05, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x32, 0xcf, 0x02, 0x0a, 0x0c, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x08, 0x41, 0x64,
	0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x1a, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x54, 0x92, 0x41, 0x31, 0x0a, 0x06, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x0b, 0x41, 0x64,
	0x64, 0x20, 0x61, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x1a, 0x1a, 0x41, 0x64, 0x64, 0x20, 0x61,
	0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x62, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0xa2, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x12, 0x1c, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x55, 0x92, 0x41, 0x35, 0x0a, 0x06, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x12, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x1a, 0x1e, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x20, 0x6f, 0x6e, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x62, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x30, 0x01, 0x42, 0x94, 0x02, 0x92, 0x41, 0x68,
	0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x72, 0x5c, 0x0a, 0x23, 0x67, 0x52,
	0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x62, 0x6f, 0x69, 0x6c, 0x65,
	0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x20, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x35, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x70,
	0x61, 0x6e, 0x74, 0x68, 0x65, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x74, 0x61, 0x6e, 0x79, 0x61, 0x70, 0x61, 0x6e, 0x74,
	0x68, 0x65, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_builds_v1_builds_proto_rawDescOnce sync.Once
	file_builds_v1_builds_proto_rawDescData = file_builds_v1_builds_proto_rawDesc
)

func file_builds_v1_builds_proto_rawDescGZIP() []byte {
	file_builds_v1_builds_proto_rawDescOnce.Do(func() {
		file_builds_v1_builds_proto_rawDescData = protoimpl.X.CompressGZIP(file_builds_v1_builds_proto_rawDescData)
	})
	return file_builds_v1_builds_proto_rawDescData
}

var file_builds_v1_builds_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_builds_v1_builds_proto_goTypes = []interface{}{
	(*AddBuildRequest)(nil),    // 0: builds.v1.AddBuildRequest
	(*AddBuildResponse)(nil),   // 1: builds.v1.AddBuildResponse
	(*ListBuildsRequest)(nil),  // 2: builds.v1.ListBuildsRequest
	(*ListBuildsResponse)(nil), // 3: builds.v1.ListBuildsResponse
	(*Build)(nil),              // 4: builds.v1.Build
}
var file_builds_v1_builds_proto_depIdxs = []int32{
	4, // 0: builds.v1.AddBuildResponse.build:type_name -> builds.v1.Build
	4, // 1: builds.v1.ListBuildsResponse.build:type_name -> builds.v1.Build
	0, // 2: builds.v1.BuildService.AddBuild:input_type -> builds.v1.AddBuildRequest
	2, // 3: builds.v1.BuildService.ListBuilds:input_type -> builds.v1.ListBuildsRequest
	1, // 4: builds.v1.BuildService.AddBuild:output_type -> builds.v1.AddBuildResponse
	3, // 5: builds.v1.BuildService.ListBuilds:output_type -> builds.v1.ListBuildsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_builds_v1_builds_proto_init() }
func file_builds_v1_builds_proto_init() {
	if File_builds_v1_builds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_builds_v1_builds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBuildRequest); i {
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
		file_builds_v1_builds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBuildResponse); i {
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
		file_builds_v1_builds_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsRequest); i {
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
		file_builds_v1_builds_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBuildsResponse); i {
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
		file_builds_v1_builds_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Build); i {
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
			RawDescriptor: file_builds_v1_builds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_builds_v1_builds_proto_goTypes,
		DependencyIndexes: file_builds_v1_builds_proto_depIdxs,
		MessageInfos:      file_builds_v1_builds_proto_msgTypes,
	}.Build()
	File_builds_v1_builds_proto = out.File
	file_builds_v1_builds_proto_rawDesc = nil
	file_builds_v1_builds_proto_goTypes = nil
	file_builds_v1_builds_proto_depIdxs = nil
}
