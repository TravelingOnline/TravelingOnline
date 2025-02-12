// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: agency.proto

package pb

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

type AgencyCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,json=agencyName,proto3" json:"name,omitempty"`
}

func (x *AgencyCreateRequest) Reset() {
	*x = AgencyCreateRequest{}
	mi := &file_agency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgencyCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyCreateRequest) ProtoMessage() {}

func (x *AgencyCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyCreateRequest.ProtoReflect.Descriptor instead.
func (*AgencyCreateRequest) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{0}
}

func (x *AgencyCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AgencyUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,json=agencyName,proto3" json:"name,omitempty"`
}

func (x *AgencyUpdateRequest) Reset() {
	*x = AgencyUpdateRequest{}
	mi := &file_agency_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgencyUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyUpdateRequest) ProtoMessage() {}

func (x *AgencyUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyUpdateRequest.ProtoReflect.Descriptor instead.
func (*AgencyUpdateRequest) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{1}
}

func (x *AgencyUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AgencyCreateRespone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,json=agencyId,proto3" json:"id,omitempty"`
}

func (x *AgencyCreateRespone) Reset() {
	*x = AgencyCreateRespone{}
	mi := &file_agency_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgencyCreateRespone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyCreateRespone) ProtoMessage() {}

func (x *AgencyCreateRespone) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyCreateRespone.ProtoReflect.Descriptor instead.
func (*AgencyCreateRespone) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{2}
}

func (x *AgencyCreateRespone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AgencyUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,json=agencyID,proto3" json:"id,omitempty"`
}

func (x *AgencyUpdateResponse) Reset() {
	*x = AgencyUpdateResponse{}
	mi := &file_agency_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgencyUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyUpdateResponse) ProtoMessage() {}

func (x *AgencyUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyUpdateResponse.ProtoReflect.Descriptor instead.
func (*AgencyUpdateResponse) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{3}
}

func (x *AgencyUpdateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Agency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,json=agencyName,proto3" json:"name,omitempty"`
	OwnerID string `protobuf:"bytes,3,opt,name=ownerID,proto3" json:"ownerID,omitempty"`
}

func (x *Agency) Reset() {
	*x = Agency{}
	mi := &file_agency_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Agency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agency) ProtoMessage() {}

func (x *Agency) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agency.ProtoReflect.Descriptor instead.
func (*Agency) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{4}
}

func (x *Agency) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Agency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Agency) GetOwnerID() string {
	if x != nil {
		return x.OwnerID
	}
	return ""
}

var File_agency_proto protoreflect.FileDescriptor

var file_agency_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x2f, 0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2b, 0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x22, 0x2c, 0x0a,
	0x14, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x06, 0x41,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_agency_proto_rawDescOnce sync.Once
	file_agency_proto_rawDescData = file_agency_proto_rawDesc
)

func file_agency_proto_rawDescGZIP() []byte {
	file_agency_proto_rawDescOnce.Do(func() {
		file_agency_proto_rawDescData = protoimpl.X.CompressGZIP(file_agency_proto_rawDescData)
	})
	return file_agency_proto_rawDescData
}

var file_agency_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agency_proto_goTypes = []any{
	(*AgencyCreateRequest)(nil),  // 0: AgencyCreateRequest
	(*AgencyUpdateRequest)(nil),  // 1: AgencyUpdateRequest
	(*AgencyCreateRespone)(nil),  // 2: AgencyCreateRespone
	(*AgencyUpdateResponse)(nil), // 3: AgencyUpdateResponse
	(*Agency)(nil),               // 4: Agency
}
var file_agency_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_agency_proto_init() }
func file_agency_proto_init() {
	if File_agency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_agency_proto_goTypes,
		DependencyIndexes: file_agency_proto_depIdxs,
		MessageInfos:      file_agency_proto_msgTypes,
	}.Build()
	File_agency_proto = out.File
	file_agency_proto_rawDesc = nil
	file_agency_proto_goTypes = nil
	file_agency_proto_depIdxs = nil
}
