// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: tensorflow/core/protobuf/status.proto

package for_core_protos_go_proto

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

// If included as a payload, this message flags the Status to be a "derived"
// Status. Used by StatusGroup to ignore certain Statuses when reporting
// errors to end users.
type DerivedStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DerivedStatus) Reset() {
	*x = DerivedStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DerivedStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DerivedStatus) ProtoMessage() {}

func (x *DerivedStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DerivedStatus.ProtoReflect.Descriptor instead.
func (*DerivedStatus) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_status_proto_rawDescGZIP(), []int{0}
}

var File_tensorflow_core_protobuf_status_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_status_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x22, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x57, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_status_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_status_proto_rawDescData = file_tensorflow_core_protobuf_status_proto_rawDesc
)

func file_tensorflow_core_protobuf_status_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_status_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_status_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_status_proto_rawDescData
}

var file_tensorflow_core_protobuf_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_status_proto_goTypes = []interface{}{
	(*DerivedStatus)(nil), // 0: tensorflow.DerivedStatus
}
var file_tensorflow_core_protobuf_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_status_proto_init() }
func file_tensorflow_core_protobuf_status_proto_init() {
	if File_tensorflow_core_protobuf_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DerivedStatus); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_status_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_status_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_status_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_status_proto = out.File
	file_tensorflow_core_protobuf_status_proto_rawDesc = nil
	file_tensorflow_core_protobuf_status_proto_goTypes = nil
	file_tensorflow_core_protobuf_status_proto_depIdxs = nil
}
