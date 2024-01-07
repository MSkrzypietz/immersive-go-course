// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: prober/prober.proto

package prober

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

// The request message
type ProbeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Count    int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ProbeRequest) Reset() {
	*x = ProbeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prober_prober_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeRequest) ProtoMessage() {}

func (x *ProbeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prober_prober_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeRequest.ProtoReflect.Descriptor instead.
func (*ProbeRequest) Descriptor() ([]byte, []int) {
	return file_prober_prober_proto_rawDescGZIP(), []int{0}
}

func (x *ProbeRequest) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *ProbeRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// The response message containing the result
type ProbeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvgLatencyMsecs float32 `protobuf:"fixed32,1,opt,name=avg_latency_msecs,json=avgLatencyMsecs,proto3" json:"avg_latency_msecs,omitempty"`
	SuccessCount    int32   `protobuf:"varint,2,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
	ErrorCount      int32   `protobuf:"varint,3,opt,name=error_count,json=errorCount,proto3" json:"error_count,omitempty"`
}

func (x *ProbeReply) Reset() {
	*x = ProbeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prober_prober_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProbeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProbeReply) ProtoMessage() {}

func (x *ProbeReply) ProtoReflect() protoreflect.Message {
	mi := &file_prober_prober_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProbeReply.ProtoReflect.Descriptor instead.
func (*ProbeReply) Descriptor() ([]byte, []int) {
	return file_prober_prober_proto_rawDescGZIP(), []int{1}
}

func (x *ProbeReply) GetAvgLatencyMsecs() float32 {
	if x != nil {
		return x.AvgLatencyMsecs
	}
	return 0
}

func (x *ProbeReply) GetSuccessCount() int32 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

func (x *ProbeReply) GetErrorCount() int32 {
	if x != nil {
		return x.ErrorCount
	}
	return 0
}

var File_prober_prober_proto protoreflect.FileDescriptor

var file_prober_prober_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x22, 0x40, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x7e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a,
	0x11, 0x61, 0x76, 0x67, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x6d, 0x73, 0x65,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x61, 0x76, 0x67, 0x4c, 0x61, 0x74,
	0x65, 0x6e, 0x63, 0x79, 0x4d, 0x73, 0x65, 0x63, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0x40, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x6f, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x6f, 0x64, 0x65, 0x59, 0x6f, 0x75, 0x72, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x69,
	0x6d, 0x6d, 0x65, 0x72, 0x73, 0x69, 0x76, 0x65, 0x2d, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prober_prober_proto_rawDescOnce sync.Once
	file_prober_prober_proto_rawDescData = file_prober_prober_proto_rawDesc
)

func file_prober_prober_proto_rawDescGZIP() []byte {
	file_prober_prober_proto_rawDescOnce.Do(func() {
		file_prober_prober_proto_rawDescData = protoimpl.X.CompressGZIP(file_prober_prober_proto_rawDescData)
	})
	return file_prober_prober_proto_rawDescData
}

var file_prober_prober_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_prober_prober_proto_goTypes = []interface{}{
	(*ProbeRequest)(nil), // 0: prober.ProbeRequest
	(*ProbeReply)(nil),   // 1: prober.ProbeReply
}
var file_prober_prober_proto_depIdxs = []int32{
	0, // 0: prober.Prober.DoProbes:input_type -> prober.ProbeRequest
	1, // 1: prober.Prober.DoProbes:output_type -> prober.ProbeReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_prober_prober_proto_init() }
func file_prober_prober_proto_init() {
	if File_prober_prober_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prober_prober_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeRequest); i {
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
		file_prober_prober_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProbeReply); i {
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
			RawDescriptor: file_prober_prober_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prober_prober_proto_goTypes,
		DependencyIndexes: file_prober_prober_proto_depIdxs,
		MessageInfos:      file_prober_prober_proto_msgTypes,
	}.Build()
	File_prober_prober_proto = out.File
	file_prober_prober_proto_rawDesc = nil
	file_prober_prober_proto_goTypes = nil
	file_prober_prober_proto_depIdxs = nil
}