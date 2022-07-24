// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: wearable/v1/wearable_service.proto

package wearablepb

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

type BeatsPerMinuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *BeatsPerMinuteRequest) Reset() {
	*x = BeatsPerMinuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatsPerMinuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatsPerMinuteRequest) ProtoMessage() {}

func (x *BeatsPerMinuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatsPerMinuteRequest.ProtoReflect.Descriptor instead.
func (*BeatsPerMinuteRequest) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{0}
}

func (x *BeatsPerMinuteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type BeatsPerMinuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Minute uint32 `protobuf:"varint,2,opt,name=minute,proto3" json:"minute,omitempty"`
}

func (x *BeatsPerMinuteResponse) Reset() {
	*x = BeatsPerMinuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeatsPerMinuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeatsPerMinuteResponse) ProtoMessage() {}

func (x *BeatsPerMinuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeatsPerMinuteResponse.ProtoReflect.Descriptor instead.
func (*BeatsPerMinuteResponse) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{1}
}

func (x *BeatsPerMinuteResponse) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *BeatsPerMinuteResponse) GetMinute() uint32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type ConsumeBeatsPerMinuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Value  uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Minute uint32 `protobuf:"varint,3,opt,name=minute,proto3" json:"minute,omitempty"`
}

func (x *ConsumeBeatsPerMinuteRequest) Reset() {
	*x = ConsumeBeatsPerMinuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeBeatsPerMinuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeBeatsPerMinuteRequest) ProtoMessage() {}

func (x *ConsumeBeatsPerMinuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeBeatsPerMinuteRequest.ProtoReflect.Descriptor instead.
func (*ConsumeBeatsPerMinuteRequest) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumeBeatsPerMinuteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ConsumeBeatsPerMinuteRequest) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ConsumeBeatsPerMinuteRequest) GetMinute() uint32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type ConsumeBeatsPerMinuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ConsumeBeatsPerMinuteResponse) Reset() {
	*x = ConsumeBeatsPerMinuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeBeatsPerMinuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeBeatsPerMinuteResponse) ProtoMessage() {}

func (x *ConsumeBeatsPerMinuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeBeatsPerMinuteResponse.ProtoReflect.Descriptor instead.
func (*ConsumeBeatsPerMinuteResponse) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{3}
}

func (x *ConsumeBeatsPerMinuteResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CalculateBeatsPerMinuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Value  uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Minute uint32 `protobuf:"varint,3,opt,name=minute,proto3" json:"minute,omitempty"`
}

func (x *CalculateBeatsPerMinuteRequest) Reset() {
	*x = CalculateBeatsPerMinuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateBeatsPerMinuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateBeatsPerMinuteRequest) ProtoMessage() {}

func (x *CalculateBeatsPerMinuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateBeatsPerMinuteRequest.ProtoReflect.Descriptor instead.
func (*CalculateBeatsPerMinuteRequest) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{4}
}

func (x *CalculateBeatsPerMinuteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CalculateBeatsPerMinuteRequest) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *CalculateBeatsPerMinuteRequest) GetMinute() uint32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type CalculateBeatsPerMinuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *CalculateBeatsPerMinuteResponse) Reset() {
	*x = CalculateBeatsPerMinuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wearable_v1_wearable_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateBeatsPerMinuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateBeatsPerMinuteResponse) ProtoMessage() {}

func (x *CalculateBeatsPerMinuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wearable_v1_wearable_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateBeatsPerMinuteResponse.ProtoReflect.Descriptor instead.
func (*CalculateBeatsPerMinuteResponse) Descriptor() ([]byte, []int) {
	return file_wearable_v1_wearable_service_proto_rawDescGZIP(), []int{5}
}

func (x *CalculateBeatsPerMinuteResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

var File_wearable_v1_wearable_service_proto protoreflect.FileDescriptor

var file_wearable_v1_wearable_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65,
	0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x22, 0x2b, 0x0a, 0x15, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x46,
	0x0a, 0x16, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x60, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x35, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x62, 0x0a, 0x1e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x74,
	0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x22, 0x3b, 0x0a, 0x1f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x32, 0xe0, 0x02, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72,
	0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x65, 0x61,
	0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x12, 0x72, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x65,
	0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x77,
	0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x42, 0x65, 0x61,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x7a, 0x0a, 0x17, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x12, 0x2b, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x6e, 0x68, 0x62, 0x6b, 0x70, 0x72, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x73, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x30, 0x37, 0x2f, 0x32, 0x33, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_wearable_v1_wearable_service_proto_rawDescOnce sync.Once
	file_wearable_v1_wearable_service_proto_rawDescData = file_wearable_v1_wearable_service_proto_rawDesc
)

func file_wearable_v1_wearable_service_proto_rawDescGZIP() []byte {
	file_wearable_v1_wearable_service_proto_rawDescOnce.Do(func() {
		file_wearable_v1_wearable_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_wearable_v1_wearable_service_proto_rawDescData)
	})
	return file_wearable_v1_wearable_service_proto_rawDescData
}

var file_wearable_v1_wearable_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_wearable_v1_wearable_service_proto_goTypes = []interface{}{
	(*BeatsPerMinuteRequest)(nil),           // 0: wearable.v1.BeatsPerMinuteRequest
	(*BeatsPerMinuteResponse)(nil),          // 1: wearable.v1.BeatsPerMinuteResponse
	(*ConsumeBeatsPerMinuteRequest)(nil),    // 2: wearable.v1.ConsumeBeatsPerMinuteRequest
	(*ConsumeBeatsPerMinuteResponse)(nil),   // 3: wearable.v1.ConsumeBeatsPerMinuteResponse
	(*CalculateBeatsPerMinuteRequest)(nil),  // 4: wearable.v1.CalculateBeatsPerMinuteRequest
	(*CalculateBeatsPerMinuteResponse)(nil), // 5: wearable.v1.CalculateBeatsPerMinuteResponse
}
var file_wearable_v1_wearable_service_proto_depIdxs = []int32{
	0, // 0: wearable.v1.WearableService.BeatsPerMinute:input_type -> wearable.v1.BeatsPerMinuteRequest
	2, // 1: wearable.v1.WearableService.ConsumeBeatsPerMinute:input_type -> wearable.v1.ConsumeBeatsPerMinuteRequest
	4, // 2: wearable.v1.WearableService.CalculateBeatsPerMinute:input_type -> wearable.v1.CalculateBeatsPerMinuteRequest
	1, // 3: wearable.v1.WearableService.BeatsPerMinute:output_type -> wearable.v1.BeatsPerMinuteResponse
	3, // 4: wearable.v1.WearableService.ConsumeBeatsPerMinute:output_type -> wearable.v1.ConsumeBeatsPerMinuteResponse
	5, // 5: wearable.v1.WearableService.CalculateBeatsPerMinute:output_type -> wearable.v1.CalculateBeatsPerMinuteResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wearable_v1_wearable_service_proto_init() }
func file_wearable_v1_wearable_service_proto_init() {
	if File_wearable_v1_wearable_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wearable_v1_wearable_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatsPerMinuteRequest); i {
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
		file_wearable_v1_wearable_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeatsPerMinuteResponse); i {
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
		file_wearable_v1_wearable_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeBeatsPerMinuteRequest); i {
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
		file_wearable_v1_wearable_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeBeatsPerMinuteResponse); i {
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
		file_wearable_v1_wearable_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateBeatsPerMinuteRequest); i {
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
		file_wearable_v1_wearable_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateBeatsPerMinuteResponse); i {
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
			RawDescriptor: file_wearable_v1_wearable_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wearable_v1_wearable_service_proto_goTypes,
		DependencyIndexes: file_wearable_v1_wearable_service_proto_depIdxs,
		MessageInfos:      file_wearable_v1_wearable_service_proto_msgTypes,
	}.Build()
	File_wearable_v1_wearable_service_proto = out.File
	file_wearable_v1_wearable_service_proto_rawDesc = nil
	file_wearable_v1_wearable_service_proto_goTypes = nil
	file_wearable_v1_wearable_service_proto_depIdxs = nil
}
