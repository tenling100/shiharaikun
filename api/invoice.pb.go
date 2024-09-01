// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: proto/invoice.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// invoice status enum
type InvoiceStatus int32

const (
	InvoiceStatus_unknow    InvoiceStatus = 0
	InvoiceStatus_unpaid    InvoiceStatus = 1
	InvoiceStatus_paid      InvoiceStatus = 2
	InvoiceStatus_rejected  InvoiceStatus = 3
	InvoiceStatus_pending   InvoiceStatus = 4
	InvoiceStatus_cancelled InvoiceStatus = 5
	InvoiceStatus_deleted   InvoiceStatus = 6
)

// Enum value maps for InvoiceStatus.
var (
	InvoiceStatus_name = map[int32]string{
		0: "unknow",
		1: "unpaid",
		2: "paid",
		3: "rejected",
		4: "pending",
		5: "cancelled",
		6: "deleted",
	}
	InvoiceStatus_value = map[string]int32{
		"unknow":    0,
		"unpaid":    1,
		"paid":      2,
		"rejected":  3,
		"pending":   4,
		"cancelled": 5,
		"deleted":   6,
	}
)

func (x InvoiceStatus) Enum() *InvoiceStatus {
	p := new(InvoiceStatus)
	*p = x
	return p
}

func (x InvoiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvoiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_invoice_proto_enumTypes[0].Descriptor()
}

func (InvoiceStatus) Type() protoreflect.EnumType {
	return &file_proto_invoice_proto_enumTypes[0]
}

func (x InvoiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceStatus.Descriptor instead.
func (InvoiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{0}
}

type InvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount  float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	DueDate string  `protobuf:"bytes,2,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
}

func (x *InvoiceRequest) Reset() {
	*x = InvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRequest) ProtoMessage() {}

func (x *InvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceRequest.ProtoReflect.Descriptor instead.
func (*InvoiceRequest) Descriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *InvoiceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceRequest) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceResponse.ProtoReflect.Descriptor instead.
func (*InvoiceResponse) Descriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DateRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *DateRangeRequest) Reset() {
	*x = DateRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRangeRequest) ProtoMessage() {}

func (x *DateRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRangeRequest.ProtoReflect.Descriptor instead.
func (*DateRangeRequest) Descriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *DateRangeRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *DateRangeRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type InvoicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoices []*Invoice `protobuf:"bytes,1,rep,name=invoices,proto3" json:"invoices,omitempty"`
}

func (x *InvoicesResponse) Reset() {
	*x = InvoicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoicesResponse) ProtoMessage() {}

func (x *InvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoicesResponse.ProtoReflect.Descriptor instead.
func (*InvoicesResponse) Descriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{3}
}

func (x *InvoicesResponse) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount        uint64                 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	RepaymentDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=repayment_date,json=repaymentDate,proto3" json:"repayment_date,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_invoice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_invoice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_proto_invoice_proto_rawDescGZIP(), []int{4}
}

func (x *Invoice) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Invoice) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Invoice) GetRepaymentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.RepaymentDate
	}
	return nil
}

func (x *Invoice) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Invoice) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_proto_invoice_proto protoreflect.FileDescriptor

var file_proto_invoice_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x43, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x75, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x4c, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x40, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x22, 0xea, 0x01, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x68,
	0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x0a, 0x0a, 0x06, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x75,
	0x6e, 0x70, 0x61, 0x69, 0x64, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x64, 0x10,
	0x02, 0x12, 0x0c, 0x0a, 0x08, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x06, 0x32, 0xa4, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x6e, 0x6c, 0x69, 0x6e, 0x67, 0x31, 0x30, 0x30, 0x2f, 0x73, 0x68, 0x69, 0x68, 0x61, 0x72, 0x61,
	0x69, 0x6b, 0x75, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_invoice_proto_rawDescOnce sync.Once
	file_proto_invoice_proto_rawDescData = file_proto_invoice_proto_rawDesc
)

func file_proto_invoice_proto_rawDescGZIP() []byte {
	file_proto_invoice_proto_rawDescOnce.Do(func() {
		file_proto_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_invoice_proto_rawDescData)
	})
	return file_proto_invoice_proto_rawDescData
}

var file_proto_invoice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_invoice_proto_goTypes = []any{
	(InvoiceStatus)(0),            // 0: invoice.InvoiceStatus
	(*InvoiceRequest)(nil),        // 1: invoice.InvoiceRequest
	(*InvoiceResponse)(nil),       // 2: invoice.InvoiceResponse
	(*DateRangeRequest)(nil),      // 3: invoice.DateRangeRequest
	(*InvoicesResponse)(nil),      // 4: invoice.InvoicesResponse
	(*Invoice)(nil),               // 5: invoice.Invoice
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_proto_invoice_proto_depIdxs = []int32{
	5, // 0: invoice.InvoicesResponse.invoices:type_name -> invoice.Invoice
	6, // 1: invoice.Invoice.repayment_date:type_name -> google.protobuf.Timestamp
	6, // 2: invoice.Invoice.created_at:type_name -> google.protobuf.Timestamp
	6, // 3: invoice.Invoice.updated_at:type_name -> google.protobuf.Timestamp
	1, // 4: invoice.InvoiceService.CreateInvoice:input_type -> invoice.InvoiceRequest
	3, // 5: invoice.InvoiceService.GetInvoicesByDateRange:input_type -> invoice.DateRangeRequest
	2, // 6: invoice.InvoiceService.CreateInvoice:output_type -> invoice.InvoiceResponse
	4, // 7: invoice.InvoiceService.GetInvoicesByDateRange:output_type -> invoice.InvoicesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_invoice_proto_init() }
func file_proto_invoice_proto_init() {
	if File_proto_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_invoice_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InvoiceRequest); i {
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
		file_proto_invoice_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InvoiceResponse); i {
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
		file_proto_invoice_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DateRangeRequest); i {
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
		file_proto_invoice_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*InvoicesResponse); i {
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
		file_proto_invoice_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Invoice); i {
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
			RawDescriptor: file_proto_invoice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_invoice_proto_goTypes,
		DependencyIndexes: file_proto_invoice_proto_depIdxs,
		EnumInfos:         file_proto_invoice_proto_enumTypes,
		MessageInfos:      file_proto_invoice_proto_msgTypes,
	}.Build()
	File_proto_invoice_proto = out.File
	file_proto_invoice_proto_rawDesc = nil
	file_proto_invoice_proto_goTypes = nil
	file_proto_invoice_proto_depIdxs = nil
}
