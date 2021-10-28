// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/tensorflow_server.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Defines the configuration of a single TensorFlow server.
type ServerDef struct {
	// The cluster of which this server is a member.
	Cluster *ClusterDef `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// The name of the job of which this server is a member.
	//
	// NOTE(mrry): The `cluster` field must contain a `JobDef` with a `name` field
	// that matches this name.
	JobName string `protobuf:"bytes,2,opt,name=job_name,json=jobName,proto3" json:"job_name,omitempty"`
	// The task index of this server in its job.
	//
	// NOTE: The `cluster` field must contain a `JobDef` with a matching `name`
	// and a mapping in its `tasks` field for this index.
	TaskIndex int32 `protobuf:"varint,3,opt,name=task_index,json=taskIndex,proto3" json:"task_index,omitempty"`
	// The default configuration for sessions that run on this server.
	DefaultSessionConfig *ConfigProto `protobuf:"bytes,4,opt,name=default_session_config,json=defaultSessionConfig,proto3" json:"default_session_config,omitempty"`
	// The protocol to be used by this server.
	//
	// Acceptable values include: "grpc", "grpc+verbs".
	Protocol             string   `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerDef) Reset()         { *m = ServerDef{} }
func (m *ServerDef) String() string { return proto.CompactTextString(m) }
func (*ServerDef) ProtoMessage()    {}
func (*ServerDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f0f8cbd85b669e4, []int{0}
}

func (m *ServerDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerDef.Unmarshal(m, b)
}
func (m *ServerDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerDef.Marshal(b, m, deterministic)
}
func (m *ServerDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerDef.Merge(m, src)
}
func (m *ServerDef) XXX_Size() int {
	return xxx_messageInfo_ServerDef.Size(m)
}
func (m *ServerDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerDef.DiscardUnknown(m)
}

var xxx_messageInfo_ServerDef proto.InternalMessageInfo

func (m *ServerDef) GetCluster() *ClusterDef {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ServerDef) GetJobName() string {
	if m != nil {
		return m.JobName
	}
	return ""
}

func (m *ServerDef) GetTaskIndex() int32 {
	if m != nil {
		return m.TaskIndex
	}
	return 0
}

func (m *ServerDef) GetDefaultSessionConfig() *ConfigProto {
	if m != nil {
		return m.DefaultSessionConfig
	}
	return nil
}

func (m *ServerDef) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerDef)(nil), "tensorflow.ServerDef")
}

func init() {
	proto.RegisterFile("tensorflow/core/protobuf/tensorflow_server.proto", fileDescriptor_7f0f8cbd85b669e4)
}

var fileDescriptor_7f0f8cbd85b669e4 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0xe5, 0xf7, 0xa5, 0xb4, 0x35, 0x4c, 0x16, 0x2a, 0x21, 0x12, 0x52, 0x84, 0x04, 0xca,
	0x94, 0x54, 0xb0, 0x32, 0x95, 0x2e, 0x0c, 0xa0, 0x2a, 0xdd, 0x58, 0xa2, 0xfc, 0x39, 0x07, 0x97,
	0xc4, 0x87, 0x6c, 0x07, 0xf8, 0xbc, 0x7c, 0x0a, 0x46, 0x14, 0xbb, 0x25, 0xed, 0x90, 0x2d, 0x77,
	0xbf, 0x7b, 0x9e, 0xbb, 0x27, 0xa6, 0x73, 0x03, 0x52, 0xa3, 0xe2, 0x35, 0x7e, 0xc6, 0x05, 0x2a,
	0x88, 0xdf, 0x15, 0x1a, 0xcc, 0x5b, 0x1e, 0xf7, 0x20, 0xd5, 0xa0, 0x3e, 0x40, 0x45, 0x16, 0x31,
	0xda, 0x03, 0xff, 0x7a, 0x50, 0x5d, 0xa0, 0xe4, 0xa2, 0x72, 0x12, 0xff, 0x66, 0x78, 0xac, 0x6e,
	0xb5, 0xd9, 0x59, 0x5f, 0x7d, 0x13, 0x3a, 0x5d, 0xdb, 0x5d, 0x4b, 0xe0, 0x6c, 0x4e, 0xc7, 0x5b,
	0xec, 0x91, 0x80, 0x84, 0x27, 0xb7, 0xb3, 0xa8, 0xf7, 0x89, 0x1e, 0x1c, 0x5a, 0x02, 0x4f, 0x76,
	0x63, 0xec, 0x82, 0x4e, 0x36, 0x98, 0xa7, 0x32, 0x6b, 0xc0, 0xfb, 0x17, 0x90, 0x70, 0x9a, 0x8c,
	0x37, 0x98, 0x3f, 0x67, 0x0d, 0xb0, 0x4b, 0x4a, 0x4d, 0xa6, 0xdf, 0x52, 0x21, 0x4b, 0xf8, 0xf2,
	0xfe, 0x07, 0x24, 0x1c, 0x25, 0xd3, 0xae, 0xf3, 0xd8, 0x35, 0xd8, 0x13, 0x9d, 0x95, 0xc0, 0xb3,
	0xb6, 0x36, 0xa9, 0x06, 0xad, 0x05, 0xca, 0xd4, 0x25, 0xf0, 0x8e, 0xec, 0xea, 0xf3, 0x83, 0xd5,
	0x96, 0xac, 0xba, 0x93, 0x93, 0xb3, 0xad, 0x6c, 0xed, 0x54, 0x0e, 0x31, 0x9f, 0x4e, 0x6c, 0xa2,
	0x02, 0x6b, 0x6f, 0x64, 0x0f, 0xf9, 0xab, 0x17, 0x0d, 0xf5, 0x51, 0x55, 0xfb, 0x7e, 0xa5, 0xd0,
	0x46, 0xb5, 0xd2, 0x88, 0x06, 0x16, 0xa7, 0x2e, 0xbf, 0x35, 0xd7, 0x2b, 0xf2, 0x72, 0x5f, 0x09,
	0xf3, 0xda, 0xe6, 0x51, 0x81, 0xcd, 0xde, 0x8b, 0x0c, 0x7c, 0x56, 0x78, 0xf8, 0x7b, 0x7f, 0x08,
	0xc9, 0x8f, 0x6d, 0x71, 0xf7, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xa7, 0x8e, 0xe0, 0xe9, 0x01,
	0x00, 0x00,
}
