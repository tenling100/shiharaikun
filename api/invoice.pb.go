// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: invoice.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
	InvoiceStatus_UNSPECIFIED InvoiceStatus = 0
	InvoiceStatus_UNPAID      InvoiceStatus = 1
	InvoiceStatus_PAID        InvoiceStatus = 2
	InvoiceStatus_REJECTED    InvoiceStatus = 3
	InvoiceStatus_PENDING     InvoiceStatus = 4
	InvoiceStatus_CANCELLED   InvoiceStatus = 5
	InvoiceStatus_DELETED     InvoiceStatus = 6
)

// Enum value maps for InvoiceStatus.
var (
	InvoiceStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNPAID",
		2: "PAID",
		3: "REJECTED",
		4: "PENDING",
		5: "CANCELLED",
		6: "DELETED",
	}
	InvoiceStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNPAID":      1,
		"PAID":        2,
		"REJECTED":    3,
		"PENDING":     4,
		"CANCELLED":   5,
		"DELETED":     6,
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
	return file_invoice_proto_enumTypes[0].Descriptor()
}

func (InvoiceStatus) Type() protoreflect.EnumType {
	return &file_invoice_proto_enumTypes[0]
}

func (x InvoiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvoiceStatus.Descriptor instead.
func (InvoiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_invoice_proto_rawDescGZIP(), []int{0}
}

type InvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyId      uint64                 `protobuf:"varint,2,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	CreatorId      uint64                 `protobuf:"varint,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	ClientId       uint64                 `protobuf:"varint,4,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	IssueDate      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	InvoiceAmount  float32                `protobuf:"fixed32,6,opt,name=invoice_amount,json=invoiceAmount,proto3" json:"invoice_amount,omitempty"`
	PaymentDueDate *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=payment_due_date,json=paymentDueDate,proto3" json:"payment_due_date,omitempty"`
	Status         InvoiceStatus          `protobuf:"varint,8,opt,name=status,proto3,enum=invoice.InvoiceStatus" json:"status,omitempty"`
}

func (x *InvoiceRequest) Reset() {
	*x = InvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRequest) ProtoMessage() {}

func (x *InvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[0]
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
	return file_invoice_proto_rawDescGZIP(), []int{0}
}

func (x *InvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvoiceRequest) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

func (x *InvoiceRequest) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *InvoiceRequest) GetClientId() uint64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *InvoiceRequest) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *InvoiceRequest) GetInvoiceAmount() float32 {
	if x != nil {
		return x.InvoiceAmount
	}
	return 0
}

func (x *InvoiceRequest) GetPaymentDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDueDate
	}
	return nil
}

func (x *InvoiceRequest) GetStatus() InvoiceStatus {
	if x != nil {
		return x.Status
	}
	return InvoiceStatus_UNSPECIFIED
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invoice *Invoice `protobuf:"bytes,1,opt,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[1]
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
	return file_invoice_proto_rawDescGZIP(), []int{1}
}

func (x *InvoiceResponse) GetInvoice() *Invoice {
	if x != nil {
		return x.Invoice
	}
	return nil
}

type DateRangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *DateRangeRequest) Reset() {
	*x = DateRangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRangeRequest) ProtoMessage() {}

func (x *DateRangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[2]
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
	return file_invoice_proto_rawDescGZIP(), []int{2}
}

func (x *DateRangeRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *DateRangeRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
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
		mi := &file_invoice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoicesResponse) ProtoMessage() {}

func (x *InvoicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[3]
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
	return file_invoice_proto_rawDescGZIP(), []int{3}
}

func (x *InvoicesResponse) GetInvoices() []*Invoice {
	if x != nil {
		return x.Invoices
	}
	return nil
}

type Company struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name           string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Representative string                  `protobuf:"bytes,3,opt,name=representative,proto3" json:"representative,omitempty"`
	Phone          string                  `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	PostalCode     string                  `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Address        string                  `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Company) Reset() {
	*x = Company{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Company) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Company) ProtoMessage() {}

func (x *Company) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Company.ProtoReflect.Descriptor instead.
func (*Company) Descriptor() ([]byte, []int) {
	return file_invoice_proto_rawDescGZIP(), []int{4}
}

func (x *Company) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Company) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Company) GetRepresentative() string {
	if x != nil {
		return x.Representative
	}
	return ""
}

func (x *Company) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Company) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Company) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email     string                  `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CompanyId uint64                  `protobuf:"varint,4,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[5]
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
	return file_invoice_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetCompanyId() uint64 {
	if x != nil {
		return x.CompanyId
	}
	return 0
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Company        *Company                `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
	Client         *Company                `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	IssueDate      *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	PaymentAmount  float32                 `protobuf:"fixed32,5,opt,name=payment_amount,json=paymentAmount,proto3" json:"payment_amount,omitempty"`
	Fee            float32                 `protobuf:"fixed32,6,opt,name=fee,proto3" json:"fee,omitempty"`
	FeeRate        float32                 `protobuf:"fixed32,7,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
	Tax            float32                 `protobuf:"fixed32,8,opt,name=tax,proto3" json:"tax,omitempty"`
	TaxRate        float32                 `protobuf:"fixed32,9,opt,name=tax_rate,json=taxRate,proto3" json:"tax_rate,omitempty"`
	InvoiceAmount  float32                 `protobuf:"fixed32,10,opt,name=invoice_amount,json=invoiceAmount,proto3" json:"invoice_amount,omitempty"`
	PaymentDueDate *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=payment_due_date,json=paymentDueDate,proto3" json:"payment_due_date,omitempty"`
	Status         InvoiceStatus           `protobuf:"varint,12,opt,name=status,proto3,enum=invoice.InvoiceStatus" json:"status,omitempty"`
	CreatedAt      *timestamppb.Timestamp  `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp  `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_proto_msgTypes[6]
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
	return file_invoice_proto_rawDescGZIP(), []int{6}
}

func (x *Invoice) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Invoice) GetCompany() *Company {
	if x != nil {
		return x.Company
	}
	return nil
}

func (x *Invoice) GetClient() *Company {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Invoice) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *Invoice) GetPaymentAmount() float32 {
	if x != nil {
		return x.PaymentAmount
	}
	return 0
}

func (x *Invoice) GetFee() float32 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Invoice) GetFeeRate() float32 {
	if x != nil {
		return x.FeeRate
	}
	return 0
}

func (x *Invoice) GetTax() float32 {
	if x != nil {
		return x.Tax
	}
	return 0
}

func (x *Invoice) GetTaxRate() float32 {
	if x != nil {
		return x.TaxRate
	}
	return 0
}

func (x *Invoice) GetInvoiceAmount() float32 {
	if x != nil {
		return x.InvoiceAmount
	}
	return 0
}

func (x *Invoice) GetPaymentDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PaymentDueDate
	}
	return nil
}

func (x *Invoice) GetStatus() InvoiceStatus {
	if x != nil {
		return x.Status
	}
	return InvoiceStatus_UNSPECIFIED
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

var File_invoice_proto protoreflect.FileDescriptor

var file_invoice_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3d, 0x0a,
	0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a,
	0x10, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x7d, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x22, 0xdc, 0x04, 0x0a, 0x07,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x28, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x66, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x74, 0x61, 0x78, 0x12, 0x19, 0x0a, 0x08,
	0x74, 0x61, 0x78, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x74, 0x61, 0x78, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44,
	0x0a, 0x10, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x75, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x6d, 0x0a, 0x0d, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x4e, 0x50, 0x41, 0x49, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x32, 0xe4, 0x02, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x4c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x12, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x1a, 0x10, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a,
	0x22, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x40,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x65, 0x6e, 0x6c, 0x69, 0x6e, 0x67, 0x31, 0x30, 0x30, 0x2f, 0x73, 0x68, 0x69, 0x68, 0x61, 0x72,
	0x61, 0x69, 0x6b, 0x75, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_invoice_proto_rawDescOnce sync.Once
	file_invoice_proto_rawDescData = file_invoice_proto_rawDesc
)

func file_invoice_proto_rawDescGZIP() []byte {
	file_invoice_proto_rawDescOnce.Do(func() {
		file_invoice_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoice_proto_rawDescData)
	})
	return file_invoice_proto_rawDescData
}

var file_invoice_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_invoice_proto_goTypes = []any{
	(InvoiceStatus)(0),             // 0: invoice.InvoiceStatus
	(*InvoiceRequest)(nil),         // 1: invoice.InvoiceRequest
	(*InvoiceResponse)(nil),        // 2: invoice.InvoiceResponse
	(*DateRangeRequest)(nil),       // 3: invoice.DateRangeRequest
	(*InvoicesResponse)(nil),       // 4: invoice.InvoicesResponse
	(*Company)(nil),                // 5: invoice.company
	(*User)(nil),                   // 6: invoice.user
	(*Invoice)(nil),                // 7: invoice.Invoice
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil), // 9: google.protobuf.UInt64Value
}
var file_invoice_proto_depIdxs = []int32{
	8,  // 0: invoice.InvoiceRequest.issue_date:type_name -> google.protobuf.Timestamp
	8,  // 1: invoice.InvoiceRequest.payment_due_date:type_name -> google.protobuf.Timestamp
	0,  // 2: invoice.InvoiceRequest.status:type_name -> invoice.InvoiceStatus
	7,  // 3: invoice.InvoiceResponse.invoice:type_name -> invoice.Invoice
	8,  // 4: invoice.DateRangeRequest.start_date:type_name -> google.protobuf.Timestamp
	8,  // 5: invoice.DateRangeRequest.end_date:type_name -> google.protobuf.Timestamp
	7,  // 6: invoice.InvoicesResponse.invoices:type_name -> invoice.Invoice
	9,  // 7: invoice.company.id:type_name -> google.protobuf.UInt64Value
	9,  // 8: invoice.user.id:type_name -> google.protobuf.UInt64Value
	9,  // 9: invoice.Invoice.id:type_name -> google.protobuf.UInt64Value
	5,  // 10: invoice.Invoice.company:type_name -> invoice.company
	5,  // 11: invoice.Invoice.client:type_name -> invoice.company
	8,  // 12: invoice.Invoice.issue_date:type_name -> google.protobuf.Timestamp
	8,  // 13: invoice.Invoice.payment_due_date:type_name -> google.protobuf.Timestamp
	0,  // 14: invoice.Invoice.status:type_name -> invoice.InvoiceStatus
	8,  // 15: invoice.Invoice.created_at:type_name -> google.protobuf.Timestamp
	8,  // 16: invoice.Invoice.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 17: invoice.InvoiceService.CreateInvoice:input_type -> invoice.InvoiceRequest
	3,  // 18: invoice.InvoiceService.GetInvoicesByDateRange:input_type -> invoice.DateRangeRequest
	5,  // 19: invoice.InvoiceService.CreateCompany:input_type -> invoice.company
	6,  // 20: invoice.InvoiceService.CreateUser:input_type -> invoice.user
	2,  // 21: invoice.InvoiceService.CreateInvoice:output_type -> invoice.InvoiceResponse
	4,  // 22: invoice.InvoiceService.GetInvoicesByDateRange:output_type -> invoice.InvoicesResponse
	5,  // 23: invoice.InvoiceService.CreateCompany:output_type -> invoice.company
	6,  // 24: invoice.InvoiceService.CreateUser:output_type -> invoice.user
	21, // [21:25] is the sub-list for method output_type
	17, // [17:21] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_invoice_proto_init() }
func file_invoice_proto_init() {
	if File_invoice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoice_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_invoice_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_invoice_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_invoice_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_invoice_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Company); i {
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
		file_invoice_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_invoice_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_invoice_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoice_proto_goTypes,
		DependencyIndexes: file_invoice_proto_depIdxs,
		EnumInfos:         file_invoice_proto_enumTypes,
		MessageInfos:      file_invoice_proto_msgTypes,
	}.Build()
	File_invoice_proto = out.File
	file_invoice_proto_rawDesc = nil
	file_invoice_proto_goTypes = nil
	file_invoice_proto_depIdxs = nil
}
