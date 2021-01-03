// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: pkg/protobuf/phone.proto

package blatherpb

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

type CreateRoomReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRoomReq) Reset() {
	*x = CreateRoomReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomReq) ProtoMessage() {}

func (x *CreateRoomReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomReq.ProtoReflect.Descriptor instead.
func (*CreateRoomReq) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{0}
}

type CreateRoomResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateRoomResp) Reset() {
	*x = CreateRoomResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomResp) ProtoMessage() {}

func (x *CreateRoomResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomResp.ProtoReflect.Descriptor instead.
func (*CreateRoomResp) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CallData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioData *AudioData `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
	UserId    uint64     `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CallData) Reset() {
	*x = CallData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallData) ProtoMessage() {}

func (x *CallData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[2]
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
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{2}
}

func (x *CallData) GetAudioData() *AudioData {
	if x != nil {
		return x.AudioData
	}
	return nil
}

func (x *CallData) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AudioData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples    []*Sample `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
	NumSamples uint32    `protobuf:"varint,2,opt,name=num_samples,json=numSamples,proto3" json:"num_samples,omitempty"`
	SampleRate uint32    `protobuf:"varint,3,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
}

func (x *AudioData) Reset() {
	*x = AudioData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioData) ProtoMessage() {}

func (x *AudioData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioData.ProtoReflect.Descriptor instead.
func (*AudioData) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{3}
}

func (x *AudioData) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

func (x *AudioData) GetNumSamples() uint32 {
	if x != nil {
		return x.NumSamples
	}
	return 0
}

func (x *AudioData) GetSampleRate() uint32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeftChannel  float64 `protobuf:"fixed64,1,opt,name=left_channel,json=leftChannel,proto3" json:"left_channel,omitempty"`
	RightChannel float64 `protobuf:"fixed64,2,opt,name=right_channel,json=rightChannel,proto3" json:"right_channel,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{4}
}

func (x *Sample) GetLeftChannel() float64 {
	if x != nil {
		return x.LeftChannel
	}
	return 0
}

func (x *Sample) GetRightChannel() float64 {
	if x != nil {
		return x.RightChannel
	}
	return 0
}

type Coordinates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Coordinates) Reset() {
	*x = Coordinates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinates) ProtoMessage() {}

func (x *Coordinates) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinates.ProtoReflect.Descriptor instead.
func (*Coordinates) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{5}
}

func (x *Coordinates) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinates) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Coordinates) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type UserSettingsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mute     bool            `protobuf:"varint,1,opt,name=mute,proto3" json:"mute,omitempty"`
	MuteList map[uint64]bool `protobuf:"bytes,2,rep,name=mute_list,json=muteList,proto3" json:"mute_list,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *UserSettingsData) Reset() {
	*x = UserSettingsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSettingsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettingsData) ProtoMessage() {}

func (x *UserSettingsData) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettingsData.ProtoReflect.Descriptor instead.
func (*UserSettingsData) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{6}
}

func (x *UserSettingsData) GetMute() bool {
	if x != nil {
		return x.Mute
	}
	return false
}

func (x *UserSettingsData) GetMuteList() map[uint64]bool {
	if x != nil {
		return x.MuteList
	}
	return nil
}

type UserSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UserSettingsResponse) Reset() {
	*x = UserSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_phone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettingsResponse) ProtoMessage() {}

func (x *UserSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_phone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettingsResponse.ProtoReflect.Descriptor instead.
func (*UserSettingsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_phone_proto_rawDescGZIP(), []int{7}
}

func (x *UserSettingsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_pkg_protobuf_phone_proto protoreflect.FileDescriptor

var file_pkg_protobuf_phone_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x62, 0x6c, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x70, 0x62, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x20, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x52, 0x09,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x7a, 0x0a, 0x09, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2b, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0x50,
	0x0a, 0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x65, 0x66, 0x74,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x6c, 0x65, 0x66, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0c, 0x72, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x22, 0x37, 0x0a, 0x0b, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x7a, 0x22, 0xab, 0x01, 0x0a, 0x10, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6d, 0x75,
	0x74, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x6d, 0x75, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x6d, 0x75, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x75,
	0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32,
	0xd0, 0x01, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x04,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x13, 0x2e, 0x62, 0x6c, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x1a, 0x1f, 0x2e, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x62, 0x6c, 0x61, 0x74, 0x68, 0x65, 0x72, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protobuf_phone_proto_rawDescOnce sync.Once
	file_pkg_protobuf_phone_proto_rawDescData = file_pkg_protobuf_phone_proto_rawDesc
)

func file_pkg_protobuf_phone_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_phone_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_phone_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_phone_proto_rawDescData)
	})
	return file_pkg_protobuf_phone_proto_rawDescData
}

var file_pkg_protobuf_phone_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_protobuf_phone_proto_goTypes = []interface{}{
	(*CreateRoomReq)(nil),        // 0: blatherpb.CreateRoomReq
	(*CreateRoomResp)(nil),       // 1: blatherpb.CreateRoomResp
	(*CallData)(nil),             // 2: blatherpb.CallData
	(*AudioData)(nil),            // 3: blatherpb.AudioData
	(*Sample)(nil),               // 4: blatherpb.Sample
	(*Coordinates)(nil),          // 5: blatherpb.Coordinates
	(*UserSettingsData)(nil),     // 6: blatherpb.UserSettingsData
	(*UserSettingsResponse)(nil), // 7: blatherpb.UserSettingsResponse
	nil,                          // 8: blatherpb.UserSettingsData.MuteListEntry
}
var file_pkg_protobuf_phone_proto_depIdxs = []int32{
	3, // 0: blatherpb.CallData.audio_data:type_name -> blatherpb.AudioData
	4, // 1: blatherpb.AudioData.samples:type_name -> blatherpb.Sample
	8, // 2: blatherpb.UserSettingsData.mute_list:type_name -> blatherpb.UserSettingsData.MuteListEntry
	0, // 3: blatherpb.Phone.CreateRoom:input_type -> blatherpb.CreateRoomReq
	2, // 4: blatherpb.Phone.Call:input_type -> blatherpb.CallData
	6, // 5: blatherpb.Phone.UpdateSettings:input_type -> blatherpb.UserSettingsData
	1, // 6: blatherpb.Phone.CreateRoom:output_type -> blatherpb.CreateRoomResp
	2, // 7: blatherpb.Phone.Call:output_type -> blatherpb.CallData
	7, // 8: blatherpb.Phone.UpdateSettings:output_type -> blatherpb.UserSettingsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_phone_proto_init() }
func file_pkg_protobuf_phone_proto_init() {
	if File_pkg_protobuf_phone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_phone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomReq); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomResp); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_protobuf_phone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioData); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinates); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSettingsData); i {
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
		file_pkg_protobuf_phone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSettingsResponse); i {
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
			RawDescriptor: file_pkg_protobuf_phone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protobuf_phone_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_phone_proto_depIdxs,
		MessageInfos:      file_pkg_protobuf_phone_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_phone_proto = out.File
	file_pkg_protobuf_phone_proto_rawDesc = nil
	file_pkg_protobuf_phone_proto_goTypes = nil
	file_pkg_protobuf_phone_proto_depIdxs = nil
}
