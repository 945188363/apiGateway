package Domain

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"sync"
)

type Message struct {
	Code int
	Msg  string
	Data map[string]interface{}
}

func NewMessage(c int, m string, d map[string]interface{}) Message {
	return Message{
		Code: c,
		Msg:  m,
		Data: d,
	}
}

type MessageXml struct {
	Code int
	Msg  string
	Data string
}

func NewMessageXml(c int, m string, d string) MessageXml {
	return MessageXml{
		Code: c,
		Msg:  m,
		Data: d,
	}
}

type RpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request map[string]interface{} `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func NewRpcRequest() RpcRequest {
	reqMap := make(map[string]interface{})

	return RpcRequest{
		Request: reqMap,
	}
}

type RpcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response map[string]interface{} `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func NewRpcResponse() RpcResponse {
	respMap := make(map[string]interface{})

	return RpcResponse{
		Response: respMap,
	}
}

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

func (x *RpcResponse) Reset() {
	*x = RpcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcResponse) ProtoMessage() {}

func (x *RpcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest1.ProtoReflect.Descriptor instead.
func (*RpcResponse) Descriptor() ([]byte, []int) {
	return file_ProdService_rawDescGZIP(), []int{0}
}

func (x *RpcResponse) GetResponse() map[string]interface{} {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *RpcRequest) Reset() {
	*x = RpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ProdService_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RpcRequest) ProtoMessage() {}

func (x *RpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ProdService_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse1.ProtoReflect.Descriptor instead.
func (*RpcRequest) Descriptor() ([]byte, []int) {
	return file_ProdService_rawDescGZIP(), []int{1}
}

func (x *RpcRequest) GetRequest() map[string]interface{} {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_ProdService protoreflect.FileDescriptor

var file_ProdService_rawDesc = []byte{
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x06, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x1a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x4a,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x31, 0x12, 0x3a,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x31, 0x1a, 0x15, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ProdService_rawDescOnce sync.Once
	file_ProdService_rawDescData = file_ProdService_rawDesc
)

func file_ProdService_rawDescGZIP() []byte {
	file_ProdService_rawDescOnce.Do(func() {
		file_ProdService_rawDescData = protoimpl.X.CompressGZIP(file_ProdService_rawDescData)
	})
	return file_ProdService_rawDescData
}

var file_ProdService_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ProdService_goTypes = []interface{}{
	(*RpcRequest)(nil),  // 0: Models.ProdRequest1
	(*RpcResponse)(nil), // 1: Models.ProdResponse1
}
var file_ProdService_depIdxs = []int32{
	2, // 0: Models.ProdResponse1.data:type_name -> Models.ProdModel
	0, // 1: Models.ProdService1.GetProdList:input_type -> Models.ProdRequest1
	1, // 2: Models.ProdService1.GetProdList:output_type -> Models.ProdResponse1
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ProdService_init() }
func file_ProdService_init() {
	if File_ProdService != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ProdService_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_ProdService_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_ProdService_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ProdService_goTypes,
		DependencyIndexes: file_ProdService_depIdxs,
		MessageInfos:      file_ProdService_msgTypes,
	}.Build()
	File_ProdService = out.File
	file_ProdService_rawDesc = nil
	file_ProdService_goTypes = nil
	file_ProdService_depIdxs = nil
}
