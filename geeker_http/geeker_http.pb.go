// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: geeker_http/geeker_http.proto

package geeker_http

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geeker_http_geeker_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_geeker_http_geeker_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_geeker_http_geeker_http_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geeker_http_geeker_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_geeker_http_geeker_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_geeker_http_geeker_http_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SayHelloResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_geeker_http_geeker_http_proto protoreflect.FileDescriptor

var file_geeker_http_geeker_http_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x67, 0x65,
	0x65, 0x6b, 0x65, 0x72, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x65, 0x77, 0x62, 0x65, 0x65, 0x2e, 0x67, 0x65, 0x65,
	0x6b, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x40, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x32, 0x82, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x65, 0x6b, 0x65, 0x72, 0x48, 0x74,
	0x74, 0x70, 0x12, 0x74, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x26,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x65, 0x77, 0x62, 0x65, 0x65, 0x2e, 0x67, 0x65, 0x65,
	0x6b, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6e, 0x65,
	0x77, 0x62, 0x65, 0x65, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x6e, 0x65, 0x77, 0x62,
	0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x65, 0x65, 0x6b,
	0x65, 0x72, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geeker_http_geeker_http_proto_rawDescOnce sync.Once
	file_geeker_http_geeker_http_proto_rawDescData = file_geeker_http_geeker_http_proto_rawDesc
)

func file_geeker_http_geeker_http_proto_rawDescGZIP() []byte {
	file_geeker_http_geeker_http_proto_rawDescOnce.Do(func() {
		file_geeker_http_geeker_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_geeker_http_geeker_http_proto_rawDescData)
	})
	return file_geeker_http_geeker_http_proto_rawDescData
}

var file_geeker_http_geeker_http_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_geeker_http_geeker_http_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),  // 0: grpc.newbee.geek_http.SayHelloRequest
	(*SayHelloResponse)(nil), // 1: grpc.newbee.geek_http.SayHelloResponse
}
var file_geeker_http_geeker_http_proto_depIdxs = []int32{
	0, // 0: grpc.newbee.geek_http.GeekerHttp.SayHello:input_type -> grpc.newbee.geek_http.SayHelloRequest
	1, // 1: grpc.newbee.geek_http.GeekerHttp.SayHello:output_type -> grpc.newbee.geek_http.SayHelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geeker_http_geeker_http_proto_init() }
func file_geeker_http_geeker_http_proto_init() {
	if File_geeker_http_geeker_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geeker_http_geeker_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_geeker_http_geeker_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
			RawDescriptor: file_geeker_http_geeker_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_geeker_http_geeker_http_proto_goTypes,
		DependencyIndexes: file_geeker_http_geeker_http_proto_depIdxs,
		MessageInfos:      file_geeker_http_geeker_http_proto_msgTypes,
	}.Build()
	File_geeker_http_geeker_http_proto = out.File
	file_geeker_http_geeker_http_proto_rawDesc = nil
	file_geeker_http_geeker_http_proto_goTypes = nil
	file_geeker_http_geeker_http_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GeekerHttpClient is the client API for GeekerHttp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GeekerHttpClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type geekerHttpClient struct {
	cc grpc.ClientConnInterface
}

func NewGeekerHttpClient(cc grpc.ClientConnInterface) GeekerHttpClient {
	return &geekerHttpClient{cc}
}

func (c *geekerHttpClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/grpc.newbee.geek_http.GeekerHttp/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeekerHttpServer is the server API for GeekerHttp service.
type GeekerHttpServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

// UnimplementedGeekerHttpServer can be embedded to have forward compatible implementations.
type UnimplementedGeekerHttpServer struct {
}

func (*UnimplementedGeekerHttpServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGeekerHttpServer(s *grpc.Server, srv GeekerHttpServer) {
	s.RegisterService(&_GeekerHttp_serviceDesc, srv)
}

func _GeekerHttp_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeekerHttpServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.newbee.geek_http.GeekerHttp/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeekerHttpServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeekerHttp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.newbee.geek_http.GeekerHttp",
	HandlerType: (*GeekerHttpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GeekerHttp_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "geeker_http/geeker_http.proto",
}
