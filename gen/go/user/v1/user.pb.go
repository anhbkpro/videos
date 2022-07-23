// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: user/v1/user.proto

package userpb

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

type MaritalStatus int32

const (
	MaritalStatus_MARITAL_STATUS_UNSPECIFIED MaritalStatus = 0
	MaritalStatus_MARITAL_STATUS_SINGLE      MaritalStatus = 1
	MaritalStatus_MARITAL_STATUS_MARRIED     MaritalStatus = 2
)

// Enum value maps for MaritalStatus.
var (
	MaritalStatus_name = map[int32]string{
		0: "MARITAL_STATUS_UNSPECIFIED",
		1: "MARITAL_STATUS_SINGLE",
		2: "MARITAL_STATUS_MARRIED",
	}
	MaritalStatus_value = map[string]int32{
		"MARITAL_STATUS_UNSPECIFIED": 0,
		"MARITAL_STATUS_SINGLE":      1,
		"MARITAL_STATUS_MARRIED":     2,
	}
)

func (x MaritalStatus) Enum() *MaritalStatus {
	p := new(MaritalStatus)
	*p = x
	return p
}

func (x MaritalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MaritalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_v1_user_proto_enumTypes[0].Descriptor()
}

func (MaritalStatus) Type() protoreflect.EnumType {
	return &file_user_v1_user_proto_enumTypes[0]
}

func (x MaritalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MaritalStatus.Descriptor instead.
func (MaritalStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string        `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FullName      string        `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	BirthYear     int64         `protobuf:"varint,3,opt,name=birth_year,json=birthYear,proto3" json:"birth_year,omitempty"`
	Salary        uint32        `protobuf:"varint,4,opt,name=salary,proto3" json:"salary,omitempty"`
	Addresses     []*Address    `protobuf:"bytes,5,rep,name=addresses,proto3" json:"addresses,omitempty"`
	MaritalStatus MaritalStatus `protobuf:"varint,6,opt,name=marital_status,json=maritalStatus,proto3,enum=user.v1.MaritalStatus" json:"marital_status,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetBirthYear() int64 {
	if x != nil {
		return x.BirthYear
	}
	return 0
}

func (x *User) GetSalary() uint32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *User) GetAddresses() []*Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *User) GetMaritalStatus() MaritalStatus {
	if x != nil {
		return x.MaritalStatus
	}
	return MaritalStatus_MARITAL_STATUS_UNSPECIFIED
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street string `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City   string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xdd, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75,
	0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x72,
	0x74, 0x68, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x2e,
	0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x3d,
	0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0d,
	0x6d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x2a, 0x66, 0x0a, 0x0d, 0x4d, 0x61, 0x72, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x41, 0x52, 0x49, 0x54, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x52, 0x49, 0x54, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x41, 0x52, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4d, 0x41, 0x52, 0x52, 0x49, 0x45, 0x44, 0x10, 0x02, 0x42, 0x3d, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x68, 0x62, 0x6b,
	0x70, 0x72, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x2f, 0x32, 0x30, 0x32, 0x32, 0x2f,
	0x30, 0x37, 0x2f, 0x32, 0x33, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_v1_user_proto_goTypes = []interface{}{
	(MaritalStatus)(0), // 0: user.v1.MaritalStatus
	(*User)(nil),       // 1: user.v1.User
	(*Address)(nil),    // 2: user.v1.Address
}
var file_user_v1_user_proto_depIdxs = []int32{
	2, // 0: user.v1.User.addresses:type_name -> user.v1.Address
	0, // 1: user.v1.User.marital_status:type_name -> user.v1.MaritalStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		EnumInfos:         file_user_v1_user_proto_enumTypes,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}
