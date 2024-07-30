// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: device/smart_alarm/smart_alarm.proto

package alarm1

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

type Alarm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlarmTime     string `protobuf:"bytes,1,opt,name=alarm_time,json=alarmTime,proto3" json:"alarm_time,omitempty"`
	RemainingTime string `protobuf:"bytes,2,opt,name=remaining_time,json=remainingTime,proto3" json:"remaining_time,omitempty"`
}

func (x *Alarm) Reset() {
	*x = Alarm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alarm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alarm) ProtoMessage() {}

func (x *Alarm) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alarm.ProtoReflect.Descriptor instead.
func (*Alarm) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{0}
}

func (x *Alarm) GetAlarmTime() string {
	if x != nil {
		return x.AlarmTime
	}
	return ""
}

func (x *Alarm) GetRemainingTime() string {
	if x != nil {
		return x.RemainingTime
	}
	return ""
}

type RemainingTimRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alarms []*Alarm `protobuf:"bytes,1,rep,name=alarms,proto3" json:"alarms,omitempty"`
	Count  int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *RemainingTimRes) Reset() {
	*x = RemainingTimRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemainingTimRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemainingTimRes) ProtoMessage() {}

func (x *RemainingTimRes) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemainingTimRes.ProtoReflect.Descriptor instead.
func (*RemainingTimRes) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{1}
}

func (x *RemainingTimRes) GetAlarms() []*Alarm {
	if x != nil {
		return x.Alarms
	}
	return nil
}

func (x *RemainingTimRes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RemainingTimeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
}

func (x *RemainingTimeReq) Reset() {
	*x = RemainingTimeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemainingTimeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemainingTimeReq) ProtoMessage() {}

func (x *RemainingTimeReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemainingTimeReq.ProtoReflect.Descriptor instead.
func (*RemainingTimeReq) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{2}
}

func (x *RemainingTimeReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemainingTimeReq) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

type AlarmStatusMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Successfully bool `protobuf:"varint,1,opt,name=successfully,proto3" json:"successfully,omitempty"`
}

func (x *AlarmStatusMessage) Reset() {
	*x = AlarmStatusMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlarmStatusMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlarmStatusMessage) ProtoMessage() {}

func (x *AlarmStatusMessage) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlarmStatusMessage.ProtoReflect.Descriptor instead.
func (*AlarmStatusMessage) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{3}
}

func (x *AlarmStatusMessage) GetSuccessfully() bool {
	if x != nil {
		return x.Successfully
	}
	return false
}

type AddSmartAlarmReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	ModelName  string `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
}

func (x *AddSmartAlarmReq) Reset() {
	*x = AddSmartAlarmReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSmartAlarmReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSmartAlarmReq) ProtoMessage() {}

func (x *AddSmartAlarmReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSmartAlarmReq.ProtoReflect.Descriptor instead.
func (*AddSmartAlarmReq) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{4}
}

func (x *AddSmartAlarmReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddSmartAlarmReq) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *AddSmartAlarmReq) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

type OpenAndCloseCurtainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	Open       bool   `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
}

func (x *OpenAndCloseCurtainReq) Reset() {
	*x = OpenAndCloseCurtainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAndCloseCurtainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAndCloseCurtainReq) ProtoMessage() {}

func (x *OpenAndCloseCurtainReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAndCloseCurtainReq.ProtoReflect.Descriptor instead.
func (*OpenAndCloseCurtainReq) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{5}
}

func (x *OpenAndCloseCurtainReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OpenAndCloseCurtainReq) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *OpenAndCloseCurtainReq) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

type OpenAndCloseCurtainRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OpenAndCloseCurtainRes) Reset() {
	*x = OpenAndCloseCurtainRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAndCloseCurtainRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAndCloseCurtainRes) ProtoMessage() {}

func (x *OpenAndCloseCurtainRes) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAndCloseCurtainRes.ProtoReflect.Descriptor instead.
func (*OpenAndCloseCurtainRes) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{6}
}

func (x *OpenAndCloseCurtainRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAlarmClockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	ClockTime  string `protobuf:"bytes,3,opt,name=ClockTime,proto3" json:"ClockTime,omitempty"`
}

func (x *CreateAlarmClockReq) Reset() {
	*x = CreateAlarmClockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlarmClockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlarmClockReq) ProtoMessage() {}

func (x *CreateAlarmClockReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlarmClockReq.ProtoReflect.Descriptor instead.
func (*CreateAlarmClockReq) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{7}
}

func (x *CreateAlarmClockReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateAlarmClockReq) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *CreateAlarmClockReq) GetClockTime() string {
	if x != nil {
		return x.ClockTime
	}
	return ""
}

type OpenAndCloseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceName string `protobuf:"bytes,2,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	Open       bool   `protobuf:"varint,3,opt,name=open,proto3" json:"open,omitempty"`
}

func (x *OpenAndCloseReq) Reset() {
	*x = OpenAndCloseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAndCloseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAndCloseReq) ProtoMessage() {}

func (x *OpenAndCloseReq) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAndCloseReq.ProtoReflect.Descriptor instead.
func (*OpenAndCloseReq) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{8}
}

func (x *OpenAndCloseReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OpenAndCloseReq) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *OpenAndCloseReq) GetOpen() bool {
	if x != nil {
		return x.Open
	}
	return false
}

type OpenAndCloseRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OpenAndCloseRes) Reset() {
	*x = OpenAndCloseRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenAndCloseRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenAndCloseRes) ProtoMessage() {}

func (x *OpenAndCloseRes) ProtoReflect() protoreflect.Message {
	mi := &file_device_smart_alarm_smart_alarm_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenAndCloseRes.ProtoReflect.Descriptor instead.
func (*OpenAndCloseRes) Descriptor() ([]byte, []int) {
	return file_device_smart_alarm_smart_alarm_proto_rawDescGZIP(), []int{9}
}

func (x *OpenAndCloseRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_device_smart_alarm_smart_alarm_proto protoreflect.FileDescriptor

var file_device_smart_alarm_smart_alarm_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x05, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x54, 0x69, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x06, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d,
	0x52, 0x06, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c,
	0x0a, 0x10, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x38, 0x0a, 0x12,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x22, 0x6b, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x66, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x16, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x74, 0x61,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x6d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x43, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5f,
	0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x22,
	0x2b, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc1, 0x02, 0x0a,
	0x11, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x12, 0x11, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x13, 0x4f,
	0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x74, 0x61,
	0x69, 0x6e, 0x12, 0x17, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x43, 0x75, 0x72, 0x74, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x43, 0x75, 0x72, 0x74, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x41, 0x6e, 0x64, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x52, 0x65,
	0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x52, 0x65, 0x73,
	0x42, 0x1a, 0x5a, 0x18, 0x64, 0x69, 0x79, 0x6f, 0x72, 0x62, 0x65, 0x6b, 0x2e, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_device_smart_alarm_smart_alarm_proto_rawDescOnce sync.Once
	file_device_smart_alarm_smart_alarm_proto_rawDescData = file_device_smart_alarm_smart_alarm_proto_rawDesc
)

func file_device_smart_alarm_smart_alarm_proto_rawDescGZIP() []byte {
	file_device_smart_alarm_smart_alarm_proto_rawDescOnce.Do(func() {
		file_device_smart_alarm_smart_alarm_proto_rawDescData = protoimpl.X.CompressGZIP(file_device_smart_alarm_smart_alarm_proto_rawDescData)
	})
	return file_device_smart_alarm_smart_alarm_proto_rawDescData
}

var file_device_smart_alarm_smart_alarm_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_device_smart_alarm_smart_alarm_proto_goTypes = []any{
	(*Alarm)(nil),                  // 0: Alarm
	(*RemainingTimRes)(nil),        // 1: RemainingTimRes
	(*RemainingTimeReq)(nil),       // 2: RemainingTimeReq
	(*AlarmStatusMessage)(nil),     // 3: AlarmStatusMessage
	(*AddSmartAlarmReq)(nil),       // 4: AddSmartAlarmReq
	(*OpenAndCloseCurtainReq)(nil), // 5: OpenAndCloseCurtainReq
	(*OpenAndCloseCurtainRes)(nil), // 6: OpenAndCloseCurtainRes
	(*CreateAlarmClockReq)(nil),    // 7: CreateAlarmClockReq
	(*OpenAndCloseReq)(nil),        // 8: OpenAndCloseReq
	(*OpenAndCloseRes)(nil),        // 9: OpenAndCloseRes
}
var file_device_smart_alarm_smart_alarm_proto_depIdxs = []int32{
	0, // 0: RemainingTimRes.alarms:type_name -> Alarm
	4, // 1: SmartAlarmService.AddSmartAlarm:input_type -> AddSmartAlarmReq
	5, // 2: SmartAlarmService.OpenAndCloseCurtain:input_type -> OpenAndCloseCurtainReq
	7, // 3: SmartAlarmService.CreateAlarmClock:input_type -> CreateAlarmClockReq
	8, // 4: SmartAlarmService.OpenAndClose:input_type -> OpenAndCloseReq
	2, // 5: SmartAlarmService.GetRemainingTime:input_type -> RemainingTimeReq
	3, // 6: SmartAlarmService.AddSmartAlarm:output_type -> AlarmStatusMessage
	6, // 7: SmartAlarmService.OpenAndCloseCurtain:output_type -> OpenAndCloseCurtainRes
	3, // 8: SmartAlarmService.CreateAlarmClock:output_type -> AlarmStatusMessage
	9, // 9: SmartAlarmService.OpenAndClose:output_type -> OpenAndCloseRes
	1, // 10: SmartAlarmService.GetRemainingTime:output_type -> RemainingTimRes
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_device_smart_alarm_smart_alarm_proto_init() }
func file_device_smart_alarm_smart_alarm_proto_init() {
	if File_device_smart_alarm_smart_alarm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_device_smart_alarm_smart_alarm_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Alarm); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RemainingTimRes); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RemainingTimeReq); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AlarmStatusMessage); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AddSmartAlarmReq); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*OpenAndCloseCurtainReq); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*OpenAndCloseCurtainRes); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAlarmClockReq); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*OpenAndCloseReq); i {
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
		file_device_smart_alarm_smart_alarm_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*OpenAndCloseRes); i {
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
			RawDescriptor: file_device_smart_alarm_smart_alarm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_device_smart_alarm_smart_alarm_proto_goTypes,
		DependencyIndexes: file_device_smart_alarm_smart_alarm_proto_depIdxs,
		MessageInfos:      file_device_smart_alarm_smart_alarm_proto_msgTypes,
	}.Build()
	File_device_smart_alarm_smart_alarm_proto = out.File
	file_device_smart_alarm_smart_alarm_proto_rawDesc = nil
	file_device_smart_alarm_smart_alarm_proto_goTypes = nil
	file_device_smart_alarm_smart_alarm_proto_depIdxs = nil
}
