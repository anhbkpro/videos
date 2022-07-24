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
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x32, 0x70, 0x0a, 0x0f, 0x57, 0x65, 0x61, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x42, 0x65, 0x61,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x77, 0x65,
	0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65,
	0x61, 0x74, 0x73, 0x50, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x68, 0x62, 0x6b, 0x70, 0x72, 0x6f, 0x2f,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x30, 0x37, 0x2f, 0x32,
	0x33, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x65, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_wearable_v1_wearable_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_wearable_v1_wearable_service_proto_goTypes = []interface{}{
	(*BeatsPerMinuteRequest)(nil),  // 0: wearable.v1.BeatsPerMinuteRequest
	(*BeatsPerMinuteResponse)(nil), // 1: wearable.v1.BeatsPerMinuteResponse
}
var file_wearable_v1_wearable_service_proto_depIdxs = []int32{
	0, // 0: wearable.v1.WearableService.BeatsPerMinute:input_type -> wearable.v1.BeatsPerMinuteRequest
	1, // 1: wearable.v1.WearableService.BeatsPerMinute:output_type -> wearable.v1.BeatsPerMinuteResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wearable_v1_wearable_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
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
