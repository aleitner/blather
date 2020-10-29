// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: internal/protobuf/phone.proto

package phone

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

type CallData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioEncoding string `protobuf:"bytes,1,opt,name=audio_encoding,json=audioEncoding,proto3" json:"audio_encoding,omitempty"`
	AudioData     []byte `protobuf:"bytes,2,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
	Length        uint64 `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *CallData) Reset() {
	*x = CallData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallData) ProtoMessage() {}

func (x *CallData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_phone_proto_msgTypes[0]
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
	return file_internal_protobuf_phone_proto_rawDescGZIP(), []int{0}
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

func (x *CallData) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_protobuf_phone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protobuf_phone_proto_msgTypes[1]
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
	return file_internal_protobuf_phone_proto_rawDescGZIP(), []int{1}
}

var File_internal_protobuf_phone_proto protoreflect.FileDescriptor

var file_internal_protobuf_phone_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x68, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x65, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x64,
	0x69, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x61,
	0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x35, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x43, 0x61, 0x6c,
	0x6c, 0x12, 0x0f, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x44,
	0x61, 0x74, 0x61, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_protobuf_phone_proto_rawDescOnce sync.Once
	file_internal_protobuf_phone_proto_rawDescData = file_internal_protobuf_phone_proto_rawDesc
)

func file_internal_protobuf_phone_proto_rawDescGZIP() []byte {
	file_internal_protobuf_phone_proto_rawDescOnce.Do(func() {
		file_internal_protobuf_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_protobuf_phone_proto_rawDescData)
	})
	return file_internal_protobuf_phone_proto_rawDescData
}

var file_internal_protobuf_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_protobuf_phone_proto_goTypes = []interface{}{
	(*CallData)(nil),     // 0: phone.CallData
	(*CallResponse)(nil), // 1: phone.CallResponse
}
var file_internal_protobuf_phone_proto_depIdxs = []int32{
	0, // 0: phone.Phone.Call:input_type -> phone.CallData
	0, // 1: phone.Phone.Call:output_type -> phone.CallData
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_protobuf_phone_proto_init() }
func file_internal_protobuf_phone_proto_init() {
	if File_internal_protobuf_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_protobuf_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_protobuf_phone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_internal_protobuf_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_protobuf_phone_proto_goTypes,
		DependencyIndexes: file_internal_protobuf_phone_proto_depIdxs,
		MessageInfos:      file_internal_protobuf_phone_proto_msgTypes,
	}.Build()
	File_internal_protobuf_phone_proto = out.File
	file_internal_protobuf_phone_proto_rawDesc = nil
	file_internal_protobuf_phone_proto_goTypes = nil
	file_internal_protobuf_phone_proto_depIdxs = nil
}