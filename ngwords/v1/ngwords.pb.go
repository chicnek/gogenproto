// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: ngwords/v1/ngwords.proto

package ngwordsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ngwords_v1_ngwords_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ngwords_v1_ngwords_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_ngwords_v1_ngwords_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NgWords []string `protobuf:"bytes,1,rep,name=ng_words,json=ngWords,proto3" json:"ng_words,omitempty"`
	IsPass  bool     `protobuf:"varint,2,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ngwords_v1_ngwords_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ngwords_v1_ngwords_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_ngwords_v1_ngwords_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateResponse) GetNgWords() []string {
	if x != nil {
		return x.NgWords
	}
	return nil
}

func (x *ValidateResponse) GetIsPass() bool {
	if x != nil {
		return x.IsPass
	}
	return false
}

type ConvertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Mark string `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *ConvertRequest) Reset() {
	*x = ConvertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ngwords_v1_ngwords_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertRequest) ProtoMessage() {}

func (x *ConvertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ngwords_v1_ngwords_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertRequest.ProtoReflect.Descriptor instead.
func (*ConvertRequest) Descriptor() ([]byte, []int) {
	return file_ngwords_v1_ngwords_proto_rawDescGZIP(), []int{2}
}

func (x *ConvertRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ConvertRequest) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

type ConvertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ConvertResponse) Reset() {
	*x = ConvertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ngwords_v1_ngwords_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConvertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConvertResponse) ProtoMessage() {}

func (x *ConvertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ngwords_v1_ngwords_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConvertResponse.ProtoReflect.Descriptor instead.
func (*ConvertResponse) Descriptor() ([]byte, []int) {
	return file_ngwords_v1_ngwords_proto_rawDescGZIP(), []int{3}
}

func (x *ConvertResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type BuildRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *BuildRequest) Reset() {
	*x = BuildRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ngwords_v1_ngwords_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRequest) ProtoMessage() {}

func (x *BuildRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ngwords_v1_ngwords_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRequest.ProtoReflect.Descriptor instead.
func (*BuildRequest) Descriptor() ([]byte, []int) {
	return file_ngwords_v1_ngwords_proto_rawDescGZIP(), []int{4}
}

func (x *BuildRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_ngwords_v1_ngwords_proto protoreflect.FileDescriptor

var file_ngwords_v1_ngwords_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x67, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6e, 0x67, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x46, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6e, 0x67, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x6e, 0x67, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73,
	0x22, 0x36, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x23, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a,
	0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32,
	0xd7, 0x01, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1b, 0x2e, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6e,
	0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x05, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x2e, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x69, 0x63, 0x6e, 0x65, 0x6b, 0x2f,
	0x67, 0x6f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x67, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x67, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ngwords_v1_ngwords_proto_rawDescOnce sync.Once
	file_ngwords_v1_ngwords_proto_rawDescData = file_ngwords_v1_ngwords_proto_rawDesc
)

func file_ngwords_v1_ngwords_proto_rawDescGZIP() []byte {
	file_ngwords_v1_ngwords_proto_rawDescOnce.Do(func() {
		file_ngwords_v1_ngwords_proto_rawDescData = protoimpl.X.CompressGZIP(file_ngwords_v1_ngwords_proto_rawDescData)
	})
	return file_ngwords_v1_ngwords_proto_rawDescData
}

var file_ngwords_v1_ngwords_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ngwords_v1_ngwords_proto_goTypes = []interface{}{
	(*ValidateRequest)(nil),  // 0: ngwords.v1.ValidateRequest
	(*ValidateResponse)(nil), // 1: ngwords.v1.ValidateResponse
	(*ConvertRequest)(nil),   // 2: ngwords.v1.ConvertRequest
	(*ConvertResponse)(nil),  // 3: ngwords.v1.ConvertResponse
	(*BuildRequest)(nil),     // 4: ngwords.v1.BuildRequest
	(*emptypb.Empty)(nil),    // 5: google.protobuf.Empty
}
var file_ngwords_v1_ngwords_proto_depIdxs = []int32{
	0, // 0: ngwords.v1.KeywordsService.Validate:input_type -> ngwords.v1.ValidateRequest
	2, // 1: ngwords.v1.KeywordsService.Convert:input_type -> ngwords.v1.ConvertRequest
	4, // 2: ngwords.v1.KeywordsService.Build:input_type -> ngwords.v1.BuildRequest
	1, // 3: ngwords.v1.KeywordsService.Validate:output_type -> ngwords.v1.ValidateResponse
	3, // 4: ngwords.v1.KeywordsService.Convert:output_type -> ngwords.v1.ConvertResponse
	5, // 5: ngwords.v1.KeywordsService.Build:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ngwords_v1_ngwords_proto_init() }
func file_ngwords_v1_ngwords_proto_init() {
	if File_ngwords_v1_ngwords_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ngwords_v1_ngwords_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_ngwords_v1_ngwords_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
		file_ngwords_v1_ngwords_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertRequest); i {
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
		file_ngwords_v1_ngwords_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConvertResponse); i {
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
		file_ngwords_v1_ngwords_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRequest); i {
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
			RawDescriptor: file_ngwords_v1_ngwords_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ngwords_v1_ngwords_proto_goTypes,
		DependencyIndexes: file_ngwords_v1_ngwords_proto_depIdxs,
		MessageInfos:      file_ngwords_v1_ngwords_proto_msgTypes,
	}.Build()
	File_ngwords_v1_ngwords_proto = out.File
	file_ngwords_v1_ngwords_proto_rawDesc = nil
	file_ngwords_v1_ngwords_proto_goTypes = nil
	file_ngwords_v1_ngwords_proto_depIdxs = nil
}
