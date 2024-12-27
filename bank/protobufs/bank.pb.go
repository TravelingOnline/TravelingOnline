// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: bank.proto

package protobufs

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

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	mi := &file_bank_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWalletRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type CreateWalletRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateWalletRequestResponse) Reset() {
	*x = CreateWalletRequestResponse{}
	mi := &file_bank_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWalletRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequestResponse) ProtoMessage() {}

func (x *CreateWalletRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequestResponse.ProtoReflect.Descriptor instead.
func (*CreateWalletRequestResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWalletRequestResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderOwnerID   string `protobuf:"bytes,1,opt,name=senderOwnerID,proto3" json:"senderOwnerID,omitempty"`
	ReceiverOwnerID string `protobuf:"bytes,2,opt,name=receiverOwnerID,proto3" json:"receiverOwnerID,omitempty"`
	IsPaidToSystem  bool   `protobuf:"varint,3,opt,name=isPaidToSystem,proto3" json:"isPaidToSystem,omitempty"`
	Amount          uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_bank_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *TransferRequest) GetSenderOwnerID() string {
	if x != nil {
		return x.SenderOwnerID
	}
	return ""
}

func (x *TransferRequest) GetReceiverOwnerID() string {
	if x != nil {
		return x.ReceiverOwnerID
	}
	return ""
}

func (x *TransferRequest) GetIsPaidToSystem() bool {
	if x != nil {
		return x.IsPaidToSystem
	}
	return false
}

func (x *TransferRequest) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderOwnerID   string `protobuf:"bytes,1,opt,name=senderOwnerID,proto3" json:"senderOwnerID,omitempty"`
	ReceiverOwnerID string `protobuf:"bytes,2,opt,name=receiverOwnerID,proto3" json:"receiverOwnerID,omitempty"`
	IsPaidToSystem  bool   `protobuf:"varint,3,opt,name=isPaidToSystem,proto3" json:"isPaidToSystem,omitempty"`
	Amount          uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Status          string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	mi := &file_bank_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{3}
}

func (x *TransferResponse) GetSenderOwnerID() string {
	if x != nil {
		return x.SenderOwnerID
	}
	return ""
}

func (x *TransferResponse) GetReceiverOwnerID() string {
	if x != nil {
		return x.ReceiverOwnerID
	}
	return ""
}

func (x *TransferResponse) GetIsPaidToSystem() bool {
	if x != nil {
		return x.IsPaidToSystem
	}
	return false
}

func (x *TransferResponse) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61,
	0x6e, 0x6b, 0x22, 0x2d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x37, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26,
	0x0a, 0x0e, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x54, 0x6f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x54, 0x6f,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xba,
	0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x54, 0x6f, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x50,
	0x61, 0x69, 0x64, 0x54, 0x6f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x0b,
	0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x19, 0x2e, 0x62, 0x61,
	0x6e, 0x6b, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_bank_proto_goTypes = []any{
	(*CreateWalletRequest)(nil),         // 0: bank.CreateWalletRequest
	(*CreateWalletRequestResponse)(nil), // 1: bank.CreateWalletRequestResponse
	(*TransferRequest)(nil),             // 2: bank.TransferRequest
	(*TransferResponse)(nil),            // 3: bank.TransferResponse
}
var file_bank_proto_depIdxs = []int32{
	0, // 0: bank.BankService.CreateWallet:input_type -> bank.CreateWalletRequest
	2, // 1: bank.BankService.Transfer:input_type -> bank.TransferRequest
	1, // 2: bank.BankService.CreateWallet:output_type -> bank.CreateWalletRequestResponse
	3, // 3: bank.BankService.Transfer:output_type -> bank.TransferResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}