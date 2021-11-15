// Defines the text format for including per-op API definition and
// overrides for client language op code generators.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: tensorflow/core/framework/api_def.proto

package api_def_go_proto

import (
	attr_value_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/attr_value_go_proto"
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

type ApiDef_Visibility int32

const (
	// Normally this is "VISIBLE" unless you are inheriting a
	// different value from another ApiDef.
	ApiDef_DEFAULT_VISIBILITY ApiDef_Visibility = 0
	// Publicly visible in the API.
	ApiDef_VISIBLE ApiDef_Visibility = 1
	// Do not include this op in the generated API. If visibility is
	// set to 'SKIP', other fields are ignored for this op.
	ApiDef_SKIP ApiDef_Visibility = 2
	// Hide this op by putting it into an internal namespace (or whatever
	// is appropriate in the target language).
	ApiDef_HIDDEN ApiDef_Visibility = 3
)

// Enum value maps for ApiDef_Visibility.
var (
	ApiDef_Visibility_name = map[int32]string{
		0: "DEFAULT_VISIBILITY",
		1: "VISIBLE",
		2: "SKIP",
		3: "HIDDEN",
	}
	ApiDef_Visibility_value = map[string]int32{
		"DEFAULT_VISIBILITY": 0,
		"VISIBLE":            1,
		"SKIP":               2,
		"HIDDEN":             3,
	}
)

func (x ApiDef_Visibility) Enum() *ApiDef_Visibility {
	p := new(ApiDef_Visibility)
	*p = x
	return p
}

func (x ApiDef_Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiDef_Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_framework_api_def_proto_enumTypes[0].Descriptor()
}

func (ApiDef_Visibility) Type() protoreflect.EnumType {
	return &file_tensorflow_core_framework_api_def_proto_enumTypes[0]
}

func (x ApiDef_Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiDef_Visibility.Descriptor instead.
func (ApiDef_Visibility) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{0, 0}
}

// Used to specify and override the default API & behavior in the
// generated code for client languages, from what you would get from
// the OpDef alone. There will be a set of ApiDefs that are common
// to all client languages, and another set per client language.
// The per-client-language ApiDefs will inherit values from the
// common ApiDefs which it can either replace or modify.
//
// We separate the API definition from the OpDef so we can evolve the
// API while remaining backwards compatible when interpreting old
// graphs.  Overrides go in an "api_def.pbtxt" file with a text-format
// ApiDefs message.
//
// WARNING: Be *very* careful changing the API for any existing op --
// you can change the semantics of existing code.  These changes may
// need to wait until a major release of TensorFlow to avoid breaking
// our compatibility promises.
type ApiDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the op (in the OpDef) to specify the API for.
	GraphOpName string `protobuf:"bytes,1,opt,name=graph_op_name,json=graphOpName,proto3" json:"graph_op_name,omitempty"`
	// If this op is deprecated, set deprecation message to the message
	// that should be logged when this op is used.
	// The message should indicate alternative op to use, if any.
	DeprecationMessage string `protobuf:"bytes,12,opt,name=deprecation_message,json=deprecationMessage,proto3" json:"deprecation_message,omitempty"`
	// Major version when the op will be deleted. For e.g. set this
	// value to 2 if op API should be removed in TensorFlow 2.0 and
	// deprecated in versions before that.
	DeprecationVersion int32              `protobuf:"varint,13,opt,name=deprecation_version,json=deprecationVersion,proto3" json:"deprecation_version,omitempty"`
	Visibility         ApiDef_Visibility  `protobuf:"varint,2,opt,name=visibility,proto3,enum=tensorflow.ApiDef_Visibility" json:"visibility,omitempty"`
	Endpoint           []*ApiDef_Endpoint `protobuf:"bytes,3,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	InArg              []*ApiDef_Arg      `protobuf:"bytes,4,rep,name=in_arg,json=inArg,proto3" json:"in_arg,omitempty"`
	OutArg             []*ApiDef_Arg      `protobuf:"bytes,5,rep,name=out_arg,json=outArg,proto3" json:"out_arg,omitempty"`
	// List of original in_arg names to specify new argument order.
	// Length of arg_order should be either empty to keep current order
	// or match size of in_arg.
	ArgOrder []string       `protobuf:"bytes,11,rep,name=arg_order,json=argOrder,proto3" json:"arg_order,omitempty"`
	Attr     []*ApiDef_Attr `protobuf:"bytes,6,rep,name=attr,proto3" json:"attr,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// Modify an existing/inherited description by adding text to the beginning
	// or end.
	DescriptionPrefix string `protobuf:"bytes,9,opt,name=description_prefix,json=descriptionPrefix,proto3" json:"description_prefix,omitempty"`
	DescriptionSuffix string `protobuf:"bytes,10,opt,name=description_suffix,json=descriptionSuffix,proto3" json:"description_suffix,omitempty"`
}

func (x *ApiDef) Reset() {
	*x = ApiDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDef) ProtoMessage() {}

func (x *ApiDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDef.ProtoReflect.Descriptor instead.
func (*ApiDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{0}
}

func (x *ApiDef) GetGraphOpName() string {
	if x != nil {
		return x.GraphOpName
	}
	return ""
}

func (x *ApiDef) GetDeprecationMessage() string {
	if x != nil {
		return x.DeprecationMessage
	}
	return ""
}

func (x *ApiDef) GetDeprecationVersion() int32 {
	if x != nil {
		return x.DeprecationVersion
	}
	return 0
}

func (x *ApiDef) GetVisibility() ApiDef_Visibility {
	if x != nil {
		return x.Visibility
	}
	return ApiDef_DEFAULT_VISIBILITY
}

func (x *ApiDef) GetEndpoint() []*ApiDef_Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *ApiDef) GetInArg() []*ApiDef_Arg {
	if x != nil {
		return x.InArg
	}
	return nil
}

func (x *ApiDef) GetOutArg() []*ApiDef_Arg {
	if x != nil {
		return x.OutArg
	}
	return nil
}

func (x *ApiDef) GetArgOrder() []string {
	if x != nil {
		return x.ArgOrder
	}
	return nil
}

func (x *ApiDef) GetAttr() []*ApiDef_Attr {
	if x != nil {
		return x.Attr
	}
	return nil
}

func (x *ApiDef) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *ApiDef) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiDef) GetDescriptionPrefix() string {
	if x != nil {
		return x.DescriptionPrefix
	}
	return ""
}

func (x *ApiDef) GetDescriptionSuffix() string {
	if x != nil {
		return x.DescriptionSuffix
	}
	return ""
}

type ApiDefs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op []*ApiDef `protobuf:"bytes,1,rep,name=op,proto3" json:"op,omitempty"`
}

func (x *ApiDefs) Reset() {
	*x = ApiDefs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDefs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDefs) ProtoMessage() {}

func (x *ApiDefs) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDefs.ProtoReflect.Descriptor instead.
func (*ApiDefs) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{1}
}

func (x *ApiDefs) GetOp() []*ApiDef {
	if x != nil {
		return x.Op
	}
	return nil
}

// If you specify any endpoint, this will replace all of the
// inherited endpoints.  The first endpoint should be the
// "canonical" endpoint, and should not be deprecated (unless all
// endpoints are deprecated).
type ApiDef_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name should be either like "CamelCaseName" or
	// "Package.CamelCaseName". Client-language-specific ApiDefs may
	// use a snake_case convention instead of CamelCase.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Set if this endpoint is deprecated. If set to true, a message suggesting
	// to use a non-deprecated endpoint instead will be printed. If all
	// endpoints are deprecated, set deprecation_message in ApiDef instead.
	Deprecated bool `protobuf:"varint,3,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	// Major version when an endpoint will be deleted. For e.g. set this
	// value to 2 if endpoint should be removed in TensorFlow 2.0 and
	// deprecated in versions before that.
	DeprecationVersion int32 `protobuf:"varint,4,opt,name=deprecation_version,json=deprecationVersion,proto3" json:"deprecation_version,omitempty"`
}

func (x *ApiDef_Endpoint) Reset() {
	*x = ApiDef_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDef_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDef_Endpoint) ProtoMessage() {}

func (x *ApiDef_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDef_Endpoint.ProtoReflect.Descriptor instead.
func (*ApiDef_Endpoint) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ApiDef_Endpoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiDef_Endpoint) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

func (x *ApiDef_Endpoint) GetDeprecationVersion() int32 {
	if x != nil {
		return x.DeprecationVersion
	}
	return 0
}

type ApiDef_Arg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this arg in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Note: this will replace any inherited arg doc. There is no
	// current way of modifying arg descriptions (other than replacing
	// them entirely) as can be done with op descriptions.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ApiDef_Arg) Reset() {
	*x = ApiDef_Arg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDef_Arg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDef_Arg) ProtoMessage() {}

func (x *ApiDef_Arg) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDef_Arg.ProtoReflect.Descriptor instead.
func (*ApiDef_Arg) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ApiDef_Arg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiDef_Arg) GetRenameTo() string {
	if x != nil {
		return x.RenameTo
	}
	return ""
}

func (x *ApiDef_Arg) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type ApiDef_Attr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this attr in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Specify a new default value to use for this attr.  This default
	// will be used when creating new graphs, as opposed to the
	// default in the OpDef, which will be used when interpreting old
	// GraphDefs.
	DefaultValue *attr_value_go_proto.AttrValue `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Note: this will replace any inherited attr doc, there is no current
	// way of modifying attr descriptions as can be done with op descriptions.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ApiDef_Attr) Reset() {
	*x = ApiDef_Attr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDef_Attr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDef_Attr) ProtoMessage() {}

func (x *ApiDef_Attr) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_framework_api_def_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDef_Attr.ProtoReflect.Descriptor instead.
func (*ApiDef_Attr) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_framework_api_def_proto_rawDescGZIP(), []int{0, 2}
}

func (x *ApiDef_Attr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ApiDef_Attr) GetRenameTo() string {
	if x != nil {
		return x.RenameTo
	}
	return ""
}

func (x *ApiDef_Attr) GetDefaultValue() *attr_value_go_proto.AttrValue {
	if x != nil {
		return x.DefaultValue
	}
	return nil
}

func (x *ApiDef_Attr) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_tensorflow_core_framework_api_def_proto protoreflect.FileDescriptor

var file_tensorflow_core_framework_api_def_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x5f,
	0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x61, 0x74, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf6, 0x07, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x12, 0x22, 0x0a, 0x0d,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x4f, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64,
	0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x69, 0x6e,
	0x5f, 0x61, 0x72, 0x67, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x2e, 0x41,
	0x72, 0x67, 0x52, 0x05, 0x69, 0x6e, 0x41, 0x72, 0x67, 0x12, 0x2f, 0x0a, 0x07, 0x6f, 0x75, 0x74,
	0x5f, 0x61, 0x72, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x2e, 0x41,
	0x72, 0x67, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x41, 0x72, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72,
	0x67, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x72, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04, 0x61, 0x74, 0x74, 0x72, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x2d, 0x0a, 0x12, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x1a, 0x6f,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2f,
	0x0a, 0x13, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x64, 0x65, 0x70,
	0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a,
	0x58, 0x0a, 0x03, 0x41, 0x72, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x95, 0x01, 0x0a, 0x04, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x54, 0x6f, 0x12, 0x3a, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x47, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x12, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x53, 0x49, 0x42,
	0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x03, 0x22, 0x2d, 0x0a, 0x07, 0x41, 0x70,
	0x69, 0x44, 0x65, 0x66, 0x73, 0x12, 0x22, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x41,
	0x70, 0x69, 0x44, 0x65, 0x66, 0x52, 0x02, 0x6f, 0x70, 0x42, 0x7d, 0x0a, 0x18, 0x6f, 0x72, 0x67,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x0c, 0x41, 0x70, 0x69, 0x44, 0x65, 0x66, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x64, 0x65, 0x66, 0x5f, 0x67, 0x6f, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_framework_api_def_proto_rawDescOnce sync.Once
	file_tensorflow_core_framework_api_def_proto_rawDescData = file_tensorflow_core_framework_api_def_proto_rawDesc
)

func file_tensorflow_core_framework_api_def_proto_rawDescGZIP() []byte {
	file_tensorflow_core_framework_api_def_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_framework_api_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_framework_api_def_proto_rawDescData)
	})
	return file_tensorflow_core_framework_api_def_proto_rawDescData
}

var file_tensorflow_core_framework_api_def_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_framework_api_def_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_core_framework_api_def_proto_goTypes = []interface{}{
	(ApiDef_Visibility)(0),                // 0: tensorflow.ApiDef.Visibility
	(*ApiDef)(nil),                        // 1: tensorflow.ApiDef
	(*ApiDefs)(nil),                       // 2: tensorflow.ApiDefs
	(*ApiDef_Endpoint)(nil),               // 3: tensorflow.ApiDef.Endpoint
	(*ApiDef_Arg)(nil),                    // 4: tensorflow.ApiDef.Arg
	(*ApiDef_Attr)(nil),                   // 5: tensorflow.ApiDef.Attr
	(*attr_value_go_proto.AttrValue)(nil), // 6: tensorflow.AttrValue
}
var file_tensorflow_core_framework_api_def_proto_depIdxs = []int32{
	0, // 0: tensorflow.ApiDef.visibility:type_name -> tensorflow.ApiDef.Visibility
	3, // 1: tensorflow.ApiDef.endpoint:type_name -> tensorflow.ApiDef.Endpoint
	4, // 2: tensorflow.ApiDef.in_arg:type_name -> tensorflow.ApiDef.Arg
	4, // 3: tensorflow.ApiDef.out_arg:type_name -> tensorflow.ApiDef.Arg
	5, // 4: tensorflow.ApiDef.attr:type_name -> tensorflow.ApiDef.Attr
	1, // 5: tensorflow.ApiDefs.op:type_name -> tensorflow.ApiDef
	6, // 6: tensorflow.ApiDef.Attr.default_value:type_name -> tensorflow.AttrValue
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_tensorflow_core_framework_api_def_proto_init() }
func file_tensorflow_core_framework_api_def_proto_init() {
	if File_tensorflow_core_framework_api_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_framework_api_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDef); i {
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
		file_tensorflow_core_framework_api_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDefs); i {
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
		file_tensorflow_core_framework_api_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDef_Endpoint); i {
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
		file_tensorflow_core_framework_api_def_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDef_Arg); i {
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
		file_tensorflow_core_framework_api_def_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDef_Attr); i {
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
			RawDescriptor: file_tensorflow_core_framework_api_def_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_framework_api_def_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_framework_api_def_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_framework_api_def_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_framework_api_def_proto_msgTypes,
	}.Build()
	File_tensorflow_core_framework_api_def_proto = out.File
	file_tensorflow_core_framework_api_def_proto_rawDesc = nil
	file_tensorflow_core_framework_api_def_proto_goTypes = nil
	file_tensorflow_core_framework_api_def_proto_depIdxs = nil
}
