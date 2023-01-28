// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/datakeeper/data_keeper_service.proto

package proto

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

type DataType int32

const (
	DataType_UNSPECIFIED DataType = 0
	DataType_LOGG_PASS   DataType = 1
	DataType_TEXT        DataType = 2
	DataType_BANK_CARD   DataType = 3
	DataType_OTHER       DataType = 4
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "LOGG_PASS",
		2: "TEXT",
		3: "BANK_CARD",
		4: "OTHER",
	}
	DataType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"LOGG_PASS":   1,
		"TEXT":        2,
		"BANK_CARD":   3,
		"OTHER":       4,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_datakeeper_data_keeper_service_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_proto_datakeeper_data_keeper_service_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{0}
}

type NewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string            `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DataType DataType          `protobuf:"varint,2,opt,name=dataType,proto3,enum=data_keeper_service.DataType" json:"dataType,omitempty"`
	Data     []byte            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta     map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NewRequest) Reset() {
	*x = NewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRequest) ProtoMessage() {}

func (x *NewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRequest.ProtoReflect.Descriptor instead.
func (*NewRequest) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *NewRequest) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_UNSPECIFIED
}

func (x *NewRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NewRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type NewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewResponse) Reset() {
	*x = NewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewResponse) ProtoMessage() {}

func (x *NewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewResponse.ProtoReflect.Descriptor instead.
func (*NewResponse) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DataType DataType          `protobuf:"varint,2,opt,name=dataType,proto3,enum=data_keeper_service.DataType" json:"dataType,omitempty"`
	Data     []byte            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta     map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetResponse) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_UNSPECIFIED
}

func (x *GetResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string            `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Id     string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Data   []byte            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta   map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{4}
}

func (x *SetRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DataType DataType          `protobuf:"varint,2,opt,name=dataType,proto3,enum=data_keeper_service.DataType" json:"dataType,omitempty"`
	Data     []byte            `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Meta     map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{5}
}

func (x *SetResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetResponse) GetDataType() DataType {
	if x != nil {
		return x.DataType
	}
	return DataType_UNSPECIFIED
}

func (x *SetResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datakeeper_data_keeper_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP(), []int{7}
}

var File_proto_datakeeper_data_keeper_service_proto protoreflect.FileDescriptor

var file_proto_datakeeper_data_keeper_service_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6b, 0x65, 0x65, 0x70,
	0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0xeb, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x1d, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xe5, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65,
	0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x3e, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc0, 0x01, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3d, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65,
	0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xe5, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x39, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3e,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x4e, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x4c, 0x4f, 0x47, 0x47, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x41, 0x4e, 0x4b,
	0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x04, 0x32, 0xbd, 0x02, 0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x12, 0x48, 0x0a, 0x03, 0x4e, 0x65, 0x77, 0x12, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e,
	0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x51, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_datakeeper_data_keeper_service_proto_rawDescOnce sync.Once
	file_proto_datakeeper_data_keeper_service_proto_rawDescData = file_proto_datakeeper_data_keeper_service_proto_rawDesc
)

func file_proto_datakeeper_data_keeper_service_proto_rawDescGZIP() []byte {
	file_proto_datakeeper_data_keeper_service_proto_rawDescOnce.Do(func() {
		file_proto_datakeeper_data_keeper_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_datakeeper_data_keeper_service_proto_rawDescData)
	})
	return file_proto_datakeeper_data_keeper_service_proto_rawDescData
}

var file_proto_datakeeper_data_keeper_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_datakeeper_data_keeper_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_datakeeper_data_keeper_service_proto_goTypes = []interface{}{
	(DataType)(0),          // 0: data_keeper_service.DataType
	(*NewRequest)(nil),     // 1: data_keeper_service.NewRequest
	(*NewResponse)(nil),    // 2: data_keeper_service.NewResponse
	(*GetRequest)(nil),     // 3: data_keeper_service.GetRequest
	(*GetResponse)(nil),    // 4: data_keeper_service.GetResponse
	(*SetRequest)(nil),     // 5: data_keeper_service.SetRequest
	(*SetResponse)(nil),    // 6: data_keeper_service.SetResponse
	(*DeleteRequest)(nil),  // 7: data_keeper_service.DeleteRequest
	(*DeleteResponse)(nil), // 8: data_keeper_service.DeleteResponse
	nil,                    // 9: data_keeper_service.NewRequest.MetaEntry
	nil,                    // 10: data_keeper_service.GetResponse.MetaEntry
	nil,                    // 11: data_keeper_service.SetRequest.MetaEntry
	nil,                    // 12: data_keeper_service.SetResponse.MetaEntry
}
var file_proto_datakeeper_data_keeper_service_proto_depIdxs = []int32{
	0,  // 0: data_keeper_service.NewRequest.dataType:type_name -> data_keeper_service.DataType
	9,  // 1: data_keeper_service.NewRequest.meta:type_name -> data_keeper_service.NewRequest.MetaEntry
	0,  // 2: data_keeper_service.GetResponse.dataType:type_name -> data_keeper_service.DataType
	10, // 3: data_keeper_service.GetResponse.meta:type_name -> data_keeper_service.GetResponse.MetaEntry
	11, // 4: data_keeper_service.SetRequest.meta:type_name -> data_keeper_service.SetRequest.MetaEntry
	0,  // 5: data_keeper_service.SetResponse.dataType:type_name -> data_keeper_service.DataType
	12, // 6: data_keeper_service.SetResponse.meta:type_name -> data_keeper_service.SetResponse.MetaEntry
	1,  // 7: data_keeper_service.DataKeeper.New:input_type -> data_keeper_service.NewRequest
	3,  // 8: data_keeper_service.DataKeeper.Get:input_type -> data_keeper_service.GetRequest
	5,  // 9: data_keeper_service.DataKeeper.Set:input_type -> data_keeper_service.SetRequest
	7,  // 10: data_keeper_service.DataKeeper.Delete:input_type -> data_keeper_service.DeleteRequest
	2,  // 11: data_keeper_service.DataKeeper.New:output_type -> data_keeper_service.NewResponse
	4,  // 12: data_keeper_service.DataKeeper.Get:output_type -> data_keeper_service.GetResponse
	6,  // 13: data_keeper_service.DataKeeper.Set:output_type -> data_keeper_service.SetResponse
	8,  // 14: data_keeper_service.DataKeeper.Delete:output_type -> data_keeper_service.DeleteResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_datakeeper_data_keeper_service_proto_init() }
func file_proto_datakeeper_data_keeper_service_proto_init() {
	if File_proto_datakeeper_data_keeper_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRequest); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewResponse); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_proto_datakeeper_data_keeper_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_proto_datakeeper_data_keeper_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_datakeeper_data_keeper_service_proto_goTypes,
		DependencyIndexes: file_proto_datakeeper_data_keeper_service_proto_depIdxs,
		EnumInfos:         file_proto_datakeeper_data_keeper_service_proto_enumTypes,
		MessageInfos:      file_proto_datakeeper_data_keeper_service_proto_msgTypes,
	}.Build()
	File_proto_datakeeper_data_keeper_service_proto = out.File
	file_proto_datakeeper_data_keeper_service_proto_rawDesc = nil
	file_proto_datakeeper_data_keeper_service_proto_goTypes = nil
	file_proto_datakeeper_data_keeper_service_proto_depIdxs = nil
}
