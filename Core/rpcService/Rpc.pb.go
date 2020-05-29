// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: Rpc.proto

package rpcService

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []byte `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
}

func (x *RpcRequest) Reset() {
	*x = RpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRequest) ProtoMessage() {}

func (x *RpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcRequest.ProtoReflect.Descriptor instead.
func (*RpcRequest) Descriptor() ([]byte, []int) {
	return file_Rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RpcRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

type RpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response []byte `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *RpcResponse) Reset() {
	*x = RpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponse) ProtoMessage() {}

func (x *RpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RpcResponse.ProtoReflect.Descriptor instead.
func (*RpcResponse) Descriptor() ([]byte, []int) {
	return file_Rpc_proto_rawDescGZIP(), []int{1}
}

func (x *RpcResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_Rpc_proto protoreflect.FileDescriptor

var file_Rpc_proto_rawDesc = []byte{
	0x0a, 0x09, 0x52, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0b, 0x52,
	0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x47, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x70, 0x63, 0x12, 0x12, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x52, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x52, 0x70, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Rpc_proto_rawDescOnce sync.Once
	file_Rpc_proto_rawDescData = file_Rpc_proto_rawDesc
)

func file_Rpc_proto_rawDescGZIP() []byte {
	file_Rpc_proto_rawDescOnce.Do(func() {
		file_Rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_Rpc_proto_rawDescData)
	})
	return file_Rpc_proto_rawDescData
}

var file_Rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Rpc_proto_goTypes = []interface{}{
	(*RpcRequest)(nil),  // 0: Models.RpcRequest
	(*RpcResponse)(nil), // 1: Models.RpcResponse
}
var file_Rpc_proto_depIdxs = []int32{
	0, // 0: Models.RpcService.GetProdListRpc:input_type -> Models.RpcRequest
	1, // 1: Models.RpcService.GetProdListRpc:output_type -> Models.RpcResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Rpc_proto_init() }
func file_Rpc_proto_init() {
	if File_Rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcRequest); i {
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
		file_Rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RpcResponse); i {
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
			RawDescriptor: file_Rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Rpc_proto_goTypes,
		DependencyIndexes: file_Rpc_proto_depIdxs,
		MessageInfos:      file_Rpc_proto_msgTypes,
	}.Build()
	File_Rpc_proto = out.File
	file_Rpc_proto_rawDesc = nil
	file_Rpc_proto_goTypes = nil
	file_Rpc_proto_depIdxs = nil
}
