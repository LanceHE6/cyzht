// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.19.4
// source: file_server.proto

package file_server

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

// 文件上传请求消息结构体
type UploadAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileContent []byte `protobuf:"bytes,2,opt,name=file_content,json=fileContent,proto3" json:"file_content,omitempty"`
	FileName    string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileType    string `protobuf:"bytes,4,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
}

func (x *UploadAvatarRequest) Reset() {
	*x = UploadAvatarRequest{}
	mi := &file_file_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarRequest) ProtoMessage() {}

func (x *UploadAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarRequest.ProtoReflect.Descriptor instead.
func (*UploadAvatarRequest) Descriptor() ([]byte, []int) {
	return file_file_server_proto_rawDescGZIP(), []int{0}
}

func (x *UploadAvatarRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadAvatarRequest) GetFileContent() []byte {
	if x != nil {
		return x.FileContent
	}
	return nil
}

func (x *UploadAvatarRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadAvatarRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

// 文件上传响应消息结构体，包含返回的文件ID
type UploadAvatarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *UploadAvatarResponse) Reset() {
	*x = UploadAvatarResponse{}
	mi := &file_file_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadAvatarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadAvatarResponse) ProtoMessage() {}

func (x *UploadAvatarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadAvatarResponse.ProtoReflect.Descriptor instead.
func (*UploadAvatarResponse) Descriptor() ([]byte, []int) {
	return file_file_server_proto_rawDescGZIP(), []int{1}
}

func (x *UploadAvatarResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type GetAvatarUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAvatarUrlRequest) Reset() {
	*x = GetAvatarUrlRequest{}
	mi := &file_file_server_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvatarUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarUrlRequest) ProtoMessage() {}

func (x *GetAvatarUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_file_server_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarUrlRequest.ProtoReflect.Descriptor instead.
func (*GetAvatarUrlRequest) Descriptor() ([]byte, []int) {
	return file_file_server_proto_rawDescGZIP(), []int{2}
}

func (x *GetAvatarUrlRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAvatarUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *GetAvatarUrlResponse) Reset() {
	*x = GetAvatarUrlResponse{}
	mi := &file_file_server_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAvatarUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAvatarUrlResponse) ProtoMessage() {}

func (x *GetAvatarUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_file_server_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAvatarUrlResponse.ProtoReflect.Descriptor instead.
func (*GetAvatarUrlResponse) Descriptor() ([]byte, []int) {
	return file_file_server_proto_rawDescGZIP(), []int{3}
}

func (x *GetAvatarUrlResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

var File_file_server_proto protoreflect.FileDescriptor

var file_file_server_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x22, 0x82, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x31, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x31, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55,
	0x72, 0x6c, 0x32, 0xbb, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x2e, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_file_server_proto_rawDescOnce sync.Once
	file_file_server_proto_rawDescData = file_file_server_proto_rawDesc
)

func file_file_server_proto_rawDescGZIP() []byte {
	file_file_server_proto_rawDescOnce.Do(func() {
		file_file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_server_proto_rawDescData)
	})
	return file_file_server_proto_rawDescData
}

var file_file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_file_server_proto_goTypes = []any{
	(*UploadAvatarRequest)(nil),  // 0: file_server.UploadAvatarRequest
	(*UploadAvatarResponse)(nil), // 1: file_server.UploadAvatarResponse
	(*GetAvatarUrlRequest)(nil),  // 2: file_server.GetAvatarUrlRequest
	(*GetAvatarUrlResponse)(nil), // 3: file_server.GetAvatarUrlResponse
}
var file_file_server_proto_depIdxs = []int32{
	0, // 0: file_server.FileService.UploadAvatar:input_type -> file_server.UploadAvatarRequest
	2, // 1: file_server.FileService.GetAvatarUrl:input_type -> file_server.GetAvatarUrlRequest
	1, // 2: file_server.FileService.UploadAvatar:output_type -> file_server.UploadAvatarResponse
	3, // 3: file_server.FileService.GetAvatarUrl:output_type -> file_server.GetAvatarUrlResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_file_server_proto_init() }
func file_file_server_proto_init() {
	if File_file_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_file_server_proto_goTypes,
		DependencyIndexes: file_file_server_proto_depIdxs,
		MessageInfos:      file_file_server_proto_msgTypes,
	}.Build()
	File_file_server_proto = out.File
	file_file_server_proto_rawDesc = nil
	file_file_server_proto_goTypes = nil
	file_file_server_proto_depIdxs = nil
}
