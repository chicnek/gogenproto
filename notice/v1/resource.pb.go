// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: notice/v1/resource.proto

package notice_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Notice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject   string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Disabled  bool   `protobuf:"varint,97,opt,name=disabled,proto3" json:"disabled,omitempty"`
	CreatedAt string `protobuf:"bytes,98,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,99,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Notice) Reset() {
	*x = Notice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notice_v1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notice) ProtoMessage() {}

func (x *Notice) ProtoReflect() protoreflect.Message {
	mi := &file_notice_v1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notice.ProtoReflect.Descriptor instead.
func (*Notice) Descriptor() ([]byte, []int) {
	return file_notice_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Notice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Notice) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Notice) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notice) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Notice) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Notice) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_notice_v1_resource_proto protoreflect.FileDescriptor

var file_notice_v1_resource_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6e, 0x6f, 0x74, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x61, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x62, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x63,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x69, 0x63, 0x6e, 0x65, 0x6b, 0x2f, 0x67, 0x6f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notice_v1_resource_proto_rawDescOnce sync.Once
	file_notice_v1_resource_proto_rawDescData = file_notice_v1_resource_proto_rawDesc
)

func file_notice_v1_resource_proto_rawDescGZIP() []byte {
	file_notice_v1_resource_proto_rawDescOnce.Do(func() {
		file_notice_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_notice_v1_resource_proto_rawDescData)
	})
	return file_notice_v1_resource_proto_rawDescData
}

var file_notice_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_notice_v1_resource_proto_goTypes = []interface{}{
	(*Notice)(nil), // 0: notice.v1.Notice
}
var file_notice_v1_resource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notice_v1_resource_proto_init() }
func file_notice_v1_resource_proto_init() {
	if File_notice_v1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notice_v1_resource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notice); i {
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
			RawDescriptor: file_notice_v1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notice_v1_resource_proto_goTypes,
		DependencyIndexes: file_notice_v1_resource_proto_depIdxs,
		MessageInfos:      file_notice_v1_resource_proto_msgTypes,
	}.Build()
	File_notice_v1_resource_proto = out.File
	file_notice_v1_resource_proto_rawDesc = nil
	file_notice_v1_resource_proto_goTypes = nil
	file_notice_v1_resource_proto_depIdxs = nil
}
