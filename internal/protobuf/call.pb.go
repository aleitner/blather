// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: internal/protobuf/call.proto

package call

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_call_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_call_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_call_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_call_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_call_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_call_proto_rawDescGZIP(), []int{1}
}

type CallData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AudioEncoding string `protobuf:"bytes,2,opt,name=audio_encoding,json=audioEncoding,proto3" json:"audio_encoding,omitempty"`
	AudioData     []byte `protobuf:"bytes,3,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
}

func (x *CallData) Reset() {
	*x = CallData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_call_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallData) ProtoMessage() {}

func (x *CallData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_call_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallData.ProtoReflect.Descriptor instead.
func (*CallData) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_call_proto_rawDescGZIP(), []int{2}
}

func (x *CallData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CallData) GetAudioEncoding() string {
	if x != nil {
		return x.AudioEncoding
	}
	return ""
}

func (x *CallData) GetAudioData() []byte {
	if x != nil {
		return x.AudioData
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_call_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_call_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_internal_protobuf_call_proto_rawDescGZIP(), []int{3}
}

var File_internal_protobuf_call_proto protoreflect.FileDescriptor

var file_internal_protobuf_call_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x63, 0x61, 0x6c, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x08, 0x43, 0x61, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x22, 0x0e, 0x0a, 0x0c, 0x43,
	0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6c, 0x0a, 0x04, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x36, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x14,
	0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x43,
	0x61, 0x6c, 0x6c, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_call_proto_rawDescOnce sync.Once
	file_internal_protobuf_call_proto_rawDescData = file_internal_protobuf_call_proto_rawDesc
)

func file_internal_protobuf_call_proto_rawDescGZIP() []byte {
	file_internal_protobuf_call_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_call_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_call_proto_rawDescData)
	})
	return file_internal_protobuf_call_proto_rawDescData
}

var file_internal_protobuf_call_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_protobuf_call_proto_goTypes = []interface{}{
	(*ConnectRequest)(nil),  // 0: call.ConnectRequest
	(*ConnectResponse)(nil), // 1: call.ConnectResponse
	(*CallData)(nil),        // 2: call.CallData
	(*CallResponse)(nil),    // 3: call.CallResponse
}
var file_internal_protobuf_call_proto_depIdxs = []int32{
	0, // 0: call.Call.Connect:input_type -> call.ConnectRequest
	2, // 1: call.Call.Call:input_type -> call.CallData
	1, // 2: call.Call.Connect:output_type -> call.ConnectResponse
	3, // 3: call.Call.Call:output_type -> call.CallResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_protobuf_call_proto_init() }
func file_internal_protobuf_call_proto_init() {
	if File_internal_protobuf_call_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_call_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_internal_protobuf_call_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_internal_protobuf_call_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallData); i {
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
		file_internal_protobuf_call_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
			RawDescriptor: file_internal_protobuf_call_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_call_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_call_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_call_proto_msgTypes,
	}.Build()
	File_internal_protobuf_call_proto = out.File
	file_internal_protobuf_call_proto_rawDesc = nil
	file_internal_protobuf_call_proto_goTypes = nil
	file_internal_protobuf_call_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CallClient is the client API for Call service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CallClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
	Call(ctx context.Context, opts ...grpc.CallOption) (Call_CallClient, error)
}

type callClient struct {
	cc grpc.ClientConnInterface
}

func NewCallClient(cc grpc.ClientConnInterface) CallClient {
	return &callClient{cc}
}

func (c *callClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/call.Call/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callClient) Call(ctx context.Context, opts ...grpc.CallOption) (Call_CallClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Call_serviceDesc.Streams[0], "/call.Call/Call", opts...)
	if err != nil {
		return nil, err
	}
	x := &callCallClient{stream}
	return x, nil
}

type Call_CallClient interface {
	Send(*CallData) error
	CloseAndRecv() (*CallResponse, error)
	grpc.ClientStream
}

type callCallClient struct {
	grpc.ClientStream
}

func (x *callCallClient) Send(m *CallData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *callCallClient) CloseAndRecv() (*CallResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CallServer is the server API for Call service.
type CallServer interface {
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	Call(Call_CallServer) error
}

// UnimplementedCallServer can be embedded to have forward compatible implementations.
type UnimplementedCallServer struct {
}

func (*UnimplementedCallServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedCallServer) Call(Call_CallServer) error {
	return status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterCallServer(s *grpc.Server, srv CallServer) {
	s.RegisterService(&_Call_serviceDesc, srv)
}

func _Call_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/call.Call/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Call_Call_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CallServer).Call(&callCallServer{stream})
}

type Call_CallServer interface {
	SendAndClose(*CallResponse) error
	Recv() (*CallData, error)
	grpc.ServerStream
}

type callCallServer struct {
	grpc.ServerStream
}

func (x *callCallServer) SendAndClose(m *CallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *callCallServer) Recv() (*CallData, error) {
	m := new(CallData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Call_serviceDesc = grpc.ServiceDesc{
	ServiceName: "call.Call",
	HandlerType: (*CallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _Call_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Call",
			Handler:       _Call_Call_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/protobuf/call.proto",
}
