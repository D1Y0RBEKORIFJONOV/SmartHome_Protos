// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: device/speaker/speaker.proto

package speaker1

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

type OnAndOffUserSpeakerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	On     bool   `protobuf:"varint,2,opt,name=on,proto3" json:"on,omitempty"`
	Off    bool   `protobuf:"varint,3,opt,name=off,proto3" json:"off,omitempty"`
}

func (x *OnAndOffUserSpeakerReq) Reset() {
	*x = OnAndOffUserSpeakerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnAndOffUserSpeakerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnAndOffUserSpeakerReq) ProtoMessage() {}

func (x *OnAndOffUserSpeakerReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnAndOffUserSpeakerReq.ProtoReflect.Descriptor instead.
func (*OnAndOffUserSpeakerReq) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{0}
}

func (x *OnAndOffUserSpeakerReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OnAndOffUserSpeakerReq) GetOn() bool {
	if x != nil {
		return x.On
	}
	return false
}

func (x *OnAndOffUserSpeakerReq) GetOff() bool {
	if x != nil {
		return x.Off
	}
	return false
}

type OnAndOffUserSpeakerRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OnAndOffUserSpeakerRes) Reset() {
	*x = OnAndOffUserSpeakerRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnAndOffUserSpeakerRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnAndOffUserSpeakerRes) ProtoMessage() {}

func (x *OnAndOffUserSpeakerRes) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnAndOffUserSpeakerRes.ProtoReflect.Descriptor instead.
func (*OnAndOffUserSpeakerRes) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{1}
}

func (x *OnAndOffUserSpeakerRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChannelS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelName   string `protobuf:"bytes,1,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	ChannelNumber string `protobuf:"bytes,2,opt,name=channel_number,json=channelNumber,proto3" json:"channel_number,omitempty"`
}

func (x *ChannelS) Reset() {
	*x = ChannelS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelS) ProtoMessage() {}

func (x *ChannelS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelS.ProtoReflect.Descriptor instead.
func (*ChannelS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{2}
}

func (x *ChannelS) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *ChannelS) GetChannelNumber() string {
	if x != nil {
		return x.ChannelNumber
	}
	return ""
}

type GetUserChannelResS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channels []*ChannelS `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	Count    int64       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetUserChannelResS) Reset() {
	*x = GetUserChannelResS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserChannelResS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserChannelResS) ProtoMessage() {}

func (x *GetUserChannelResS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserChannelResS.ProtoReflect.Descriptor instead.
func (*GetUserChannelResS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserChannelResS) GetChannels() []*ChannelS {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *GetUserChannelResS) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddSpeakerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ModelName string `protobuf:"bytes,2,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
}

func (x *AddSpeakerReq) Reset() {
	*x = AddSpeakerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSpeakerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSpeakerReq) ProtoMessage() {}

func (x *AddSpeakerReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSpeakerReq.ProtoReflect.Descriptor instead.
func (*AddSpeakerReq) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{4}
}

func (x *AddSpeakerReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddSpeakerReq) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

type SpeakerStatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successfully bool `protobuf:"varint,1,opt,name=successfully,proto3" json:"successfully,omitempty"`
}

func (x *SpeakerStatusMessage) Reset() {
	*x = SpeakerStatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeakerStatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeakerStatusMessage) ProtoMessage() {}

func (x *SpeakerStatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeakerStatusMessage.ProtoReflect.Descriptor instead.
func (*SpeakerStatusMessage) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{5}
}

func (x *SpeakerStatusMessage) GetSuccessfully() bool {
	if x != nil {
		return x.Successfully
	}
	return false
}

type AddChannelReqS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
}

func (x *AddChannelReqS) Reset() {
	*x = AddChannelReqS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddChannelReqS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddChannelReqS) ProtoMessage() {}

func (x *AddChannelReqS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddChannelReqS.ProtoReflect.Descriptor instead.
func (*AddChannelReqS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{6}
}

func (x *AddChannelReqS) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddChannelReqS) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

type GetUserChannelReqS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
}

func (x *GetUserChannelReqS) Reset() {
	*x = GetUserChannelReqS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserChannelReqS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserChannelReqS) ProtoMessage() {}

func (x *GetUserChannelReqS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserChannelReqS.ProtoReflect.Descriptor instead.
func (*GetUserChannelReqS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserChannelReqS) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserChannelReqS) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

type DeleteChannelReqS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChannelName string `protobuf:"bytes,2,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
}

func (x *DeleteChannelReqS) Reset() {
	*x = DeleteChannelReqS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChannelReqS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChannelReqS) ProtoMessage() {}

func (x *DeleteChannelReqS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChannelReqS.ProtoReflect.Descriptor instead.
func (*DeleteChannelReqS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteChannelReqS) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteChannelReqS) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

type DownOrUpVolumeReqS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Down   bool   `protobuf:"varint,2,opt,name=down,proto3" json:"down,omitempty"`
	Up     bool   `protobuf:"varint,3,opt,name=up,proto3" json:"up,omitempty"`
}

func (x *DownOrUpVolumeReqS) Reset() {
	*x = DownOrUpVolumeReqS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownOrUpVolumeReqS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownOrUpVolumeReqS) ProtoMessage() {}

func (x *DownOrUpVolumeReqS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownOrUpVolumeReqS.ProtoReflect.Descriptor instead.
func (*DownOrUpVolumeReqS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{9}
}

func (x *DownOrUpVolumeReqS) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DownOrUpVolumeReqS) GetDown() bool {
	if x != nil {
		return x.Down
	}
	return false
}

func (x *DownOrUpVolumeReqS) GetUp() bool {
	if x != nil {
		return x.Up
	}
	return false
}

type DownOrUpVolumeResS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sound int64 `protobuf:"varint,1,opt,name=sound,proto3" json:"sound,omitempty"`
}

func (x *DownOrUpVolumeResS) Reset() {
	*x = DownOrUpVolumeResS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownOrUpVolumeResS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownOrUpVolumeResS) ProtoMessage() {}

func (x *DownOrUpVolumeResS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownOrUpVolumeResS.ProtoReflect.Descriptor instead.
func (*DownOrUpVolumeResS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{10}
}

func (x *DownOrUpVolumeResS) GetSound() int64 {
	if x != nil {
		return x.Sound
	}
	return 0
}

type PreviousAndNextReqS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Next   bool   `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
	Back   bool   `protobuf:"varint,3,opt,name=back,proto3" json:"back,omitempty"`
}

func (x *PreviousAndNextReqS) Reset() {
	*x = PreviousAndNextReqS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousAndNextReqS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousAndNextReqS) ProtoMessage() {}

func (x *PreviousAndNextReqS) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousAndNextReqS.ProtoReflect.Descriptor instead.
func (*PreviousAndNextReqS) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{11}
}

func (x *PreviousAndNextReqS) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PreviousAndNextReqS) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *PreviousAndNextReqS) GetBack() bool {
	if x != nil {
		return x.Back
	}
	return false
}

type PreviousAndNextRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel *ChannelS `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *PreviousAndNextRes) Reset() {
	*x = PreviousAndNextRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_speaker_speaker_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreviousAndNextRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreviousAndNextRes) ProtoMessage() {}

func (x *PreviousAndNextRes) ProtoReflect() protoreflect.Message {
	mi := &file_device_speaker_speaker_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreviousAndNextRes.ProtoReflect.Descriptor instead.
func (*PreviousAndNextRes) Descriptor() ([]byte, []int) {
	return file_device_speaker_speaker_proto_rawDescGZIP(), []int{12}
}

func (x *PreviousAndNextRes) GetChannel() *ChannelS {
	if x != nil {
		return x.Channel
	}
	return nil
}

var File_device_speaker_speaker_proto protoreflect.FileDescriptor

var file_device_speaker_speaker_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53,
	0x0a, 0x16, 0x4f, 0x6e, 0x41, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x66, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x6f, 0x66, 0x66, 0x22, 0x32, 0x0a, 0x16, 0x4f, 0x6e, 0x41, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x53, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x51, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x53, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x47, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x53, 0x70, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x4c, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x12, 0x44, 0x6f, 0x77, 0x6e, 0x4f, 0x72,
	0x55, 0x70, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x53, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x75, 0x70, 0x22, 0x2a, 0x0a, 0x12, 0x44, 0x6f, 0x77,
	0x6e, 0x4f, 0x72, 0x55, 0x70, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x53, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x56, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75,
	0x73, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x53, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x63,
	0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x39, 0x0a,
	0x12, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x32, 0xb6, 0x03, 0x0a, 0x0e, 0x53, 0x70, 0x65,
	0x61, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x41, 0x64, 0x64, 0x53,
	0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x53, 0x70, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0f,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x1a,
	0x15, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x1a, 0x13, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65,
	0x73, 0x53, 0x12, 0x3a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x53, 0x1a, 0x15, 0x2e, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3a,
	0x0a, 0x0e, 0x44, 0x6f, 0x77, 0x6e, 0x4f, 0x72, 0x55, 0x70, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x13, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x4f, 0x72, 0x55, 0x70, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x53, 0x1a, 0x13, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x4f, 0x72, 0x55, 0x70,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x53, 0x12, 0x3c, 0x0a, 0x0f, 0x50, 0x72,
	0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x14, 0x2e,
	0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x6e, 0x64, 0x4e, 0x65, 0x78, 0x74, 0x52,
	0x65, 0x71, 0x53, 0x1a, 0x13, 0x2e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x41, 0x6e,
	0x64, 0x4e, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x13, 0x4f, 0x6e, 0x41, 0x6e,
	0x64, 0x4f, 0x66, 0x66, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x4f, 0x6e, 0x41, 0x6e, 0x64, 0x4f, 0x66, 0x66, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x4f, 0x6e, 0x41, 0x6e, 0x64,
	0x4f, 0x66, 0x66, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x64, 0x69, 0x79, 0x6f, 0x72, 0x62, 0x65, 0x6b, 0x2e, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_speaker_speaker_proto_rawDescOnce sync.Once
	file_device_speaker_speaker_proto_rawDescData = file_device_speaker_speaker_proto_rawDesc
)

func file_device_speaker_speaker_proto_rawDescGZIP() []byte {
	file_device_speaker_speaker_proto_rawDescOnce.Do(func() {
		file_device_speaker_speaker_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_speaker_speaker_proto_rawDescData)
	})
	return file_device_speaker_speaker_proto_rawDescData
}

var file_device_speaker_speaker_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_device_speaker_speaker_proto_goTypes = []any{
	(*OnAndOffUserSpeakerReq)(nil), // 0: OnAndOffUserSpeakerReq
	(*OnAndOffUserSpeakerRes)(nil), // 1: OnAndOffUserSpeakerRes
	(*ChannelS)(nil),               // 2: ChannelS
	(*GetUserChannelResS)(nil),     // 3: GetUserChannelResS
	(*AddSpeakerReq)(nil),          // 4: AddSpeakerReq
	(*SpeakerStatusMessage)(nil),   // 5: SpeakerStatusMessage
	(*AddChannelReqS)(nil),         // 6: AddChannelReqS
	(*GetUserChannelReqS)(nil),     // 7: GetUserChannelReqS
	(*DeleteChannelReqS)(nil),      // 8: DeleteChannelReqS
	(*DownOrUpVolumeReqS)(nil),     // 9: DownOrUpVolumeReqS
	(*DownOrUpVolumeResS)(nil),     // 10: DownOrUpVolumeResS
	(*PreviousAndNextReqS)(nil),    // 11: PreviousAndNextReqS
	(*PreviousAndNextRes)(nil),     // 12: PreviousAndNextRes
}
var file_device_speaker_speaker_proto_depIdxs = []int32{
	2,  // 0: GetUserChannelResS.channels:type_name -> ChannelS
	2,  // 1: PreviousAndNextRes.channel:type_name -> ChannelS
	4,  // 2: SpeakerService.AddSpeaker:input_type -> AddSpeakerReq
	6,  // 3: SpeakerService.AddChannel:input_type -> AddChannelReqS
	7,  // 4: SpeakerService.GetUserChannel:input_type -> GetUserChannelReqS
	8,  // 5: SpeakerService.DeleteChannel:input_type -> DeleteChannelReqS
	9,  // 6: SpeakerService.DownOrUpVolume:input_type -> DownOrUpVolumeReqS
	11, // 7: SpeakerService.PreviousAndNext:input_type -> PreviousAndNextReqS
	0,  // 8: SpeakerService.OnAndOffUserSpeaker:input_type -> OnAndOffUserSpeakerReq
	5,  // 9: SpeakerService.AddSpeaker:output_type -> SpeakerStatusMessage
	5,  // 10: SpeakerService.AddChannel:output_type -> SpeakerStatusMessage
	3,  // 11: SpeakerService.GetUserChannel:output_type -> GetUserChannelResS
	5,  // 12: SpeakerService.DeleteChannel:output_type -> SpeakerStatusMessage
	10, // 13: SpeakerService.DownOrUpVolume:output_type -> DownOrUpVolumeResS
	12, // 14: SpeakerService.PreviousAndNext:output_type -> PreviousAndNextRes
	1,  // 15: SpeakerService.OnAndOffUserSpeaker:output_type -> OnAndOffUserSpeakerRes
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_device_speaker_speaker_proto_init() }
func file_device_speaker_speaker_proto_init() {
	if File_device_speaker_speaker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_speaker_speaker_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*OnAndOffUserSpeakerReq); i {
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
		file_device_speaker_speaker_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*OnAndOffUserSpeakerRes); i {
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
		file_device_speaker_speaker_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ChannelS); i {
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
		file_device_speaker_speaker_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserChannelResS); i {
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
		file_device_speaker_speaker_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddSpeakerReq); i {
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
		file_device_speaker_speaker_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SpeakerStatusMessage); i {
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
		file_device_speaker_speaker_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*AddChannelReqS); i {
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
		file_device_speaker_speaker_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserChannelReqS); i {
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
		file_device_speaker_speaker_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteChannelReqS); i {
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
		file_device_speaker_speaker_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DownOrUpVolumeReqS); i {
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
		file_device_speaker_speaker_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*DownOrUpVolumeResS); i {
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
		file_device_speaker_speaker_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*PreviousAndNextReqS); i {
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
		file_device_speaker_speaker_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*PreviousAndNextRes); i {
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
			RawDescriptor: file_device_speaker_speaker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_speaker_speaker_proto_goTypes,
		DependencyIndexes: file_device_speaker_speaker_proto_depIdxs,
		MessageInfos:      file_device_speaker_speaker_proto_msgTypes,
	}.Build()
	File_device_speaker_speaker_proto = out.File
	file_device_speaker_speaker_proto_rawDesc = nil
	file_device_speaker_speaker_proto_goTypes = nil
	file_device_speaker_speaker_proto_depIdxs = nil
}
