// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: proto/ticket.proto

package traintkt

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[0]
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
	return file_proto_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type PurchaseTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User  *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Price float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *PurchaseTicketRequest) Reset() {
	*x = PurchaseTicketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseTicketRequest) ProtoMessage() {}

func (x *PurchaseTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseTicketRequest.ProtoReflect.Descriptor instead.
func (*PurchaseTicketRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseTicketRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseTicketRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseTicketRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseTicketRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From  string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To    string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User  *User   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Price float64 `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Seat  string  `protobuf:"bytes,6,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *Receipt) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Receipt) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Receipt) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Receipt) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Receipt) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Receipt) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

type GetReceiptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetReceiptRequest) Reset() {
	*x = GetReceiptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptRequest) ProtoMessage() {}

func (x *GetReceiptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptRequest.ProtoReflect.Descriptor instead.
func (*GetReceiptRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *GetReceiptRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ViewSeatAllocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *ViewSeatAllocationRequest) Reset() {
	*x = ViewSeatAllocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewSeatAllocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewSeatAllocationRequest) ProtoMessage() {}

func (x *ViewSeatAllocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewSeatAllocationRequest.ProtoReflect.Descriptor instead.
func (*ViewSeatAllocationRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ViewSeatAllocationRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

type SeatAllocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Allocations []*SeatAllocation `protobuf:"bytes,1,rep,name=allocations,proto3" json:"allocations,omitempty"`
}

func (x *SeatAllocationResponse) Reset() {
	*x = SeatAllocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatAllocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatAllocationResponse) ProtoMessage() {}

func (x *SeatAllocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatAllocationResponse.ProtoReflect.Descriptor instead.
func (*SeatAllocationResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *SeatAllocationResponse) GetAllocations() []*SeatAllocation {
	if x != nil {
		return x.Allocations
	}
	return nil
}

type SeatAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seat string `protobuf:"bytes,1,opt,name=seat,proto3" json:"seat,omitempty"`
	User *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *SeatAllocation) Reset() {
	*x = SeatAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeatAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeatAllocation) ProtoMessage() {}

func (x *SeatAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeatAllocation.ProtoReflect.Descriptor instead.
func (*SeatAllocation) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *SeatAllocation) GetSeat() string {
	if x != nil {
		return x.Seat
	}
	return ""
}

func (x *SeatAllocation) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NewSeat string `protobuf:"bytes,2,opt,name=new_seat,json=newSeat,proto3" json:"new_seat,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *ModifySeatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

type ModifySeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	NewSeat string `protobuf:"bytes,2,opt,name=new_seat,json=newSeat,proto3" json:"new_seat,omitempty"`
}

func (x *ModifySeatResponse) Reset() {
	*x = ModifySeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ticket_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatResponse) ProtoMessage() {}

func (x *ModifySeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ticket_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatResponse.ProtoReflect.Descriptor instead.
func (*ModifySeatResponse) Descriptor() ([]byte, []int) {
	return file_proto_ticket_proto_rawDescGZIP(), []int{10}
}

func (x *ModifySeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ModifySeatResponse) GetNewSeat() string {
	if x != nil {
		return x.NewSeat
	}
	return ""
}

var File_proto_ticket_proto protoreflect.FileDescriptor

var file_proto_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0x58, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x78, 0x0a, 0x15, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x56,
	0x69, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x0e, 0x53,
	0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a,
	0x12, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x0a,
	0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x22, 0x49, 0x0a,
	0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x32, 0xaa, 0x03, 0x0a, 0x0d, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0e, 0x50, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x00, 0x12, 0x63,
	0x0a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x74, 0x41,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x61, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x72, 0x65, 0x61, 0x6b, 0x74, 0x68, 0x65, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x6b, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_ticket_proto_rawDescOnce sync.Once
	file_proto_ticket_proto_rawDescData = file_proto_ticket_proto_rawDesc
)

func file_proto_ticket_proto_rawDescGZIP() []byte {
	file_proto_ticket_proto_rawDescOnce.Do(func() {
		file_proto_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ticket_proto_rawDescData)
	})
	return file_proto_ticket_proto_rawDescData
}

var file_proto_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_ticket_proto_goTypes = []any{
	(*User)(nil),                      // 0: trainticket.User
	(*PurchaseTicketRequest)(nil),     // 1: trainticket.PurchaseTicketRequest
	(*Receipt)(nil),                   // 2: trainticket.Receipt
	(*GetReceiptRequest)(nil),         // 3: trainticket.GetReceiptRequest
	(*ViewSeatAllocationRequest)(nil), // 4: trainticket.ViewSeatAllocationRequest
	(*SeatAllocationResponse)(nil),    // 5: trainticket.SeatAllocationResponse
	(*SeatAllocation)(nil),            // 6: trainticket.SeatAllocation
	(*RemoveUserRequest)(nil),         // 7: trainticket.RemoveUserRequest
	(*RemoveUserResponse)(nil),        // 8: trainticket.RemoveUserResponse
	(*ModifySeatRequest)(nil),         // 9: trainticket.ModifySeatRequest
	(*ModifySeatResponse)(nil),        // 10: trainticket.ModifySeatResponse
}
var file_proto_ticket_proto_depIdxs = []int32{
	0,  // 0: trainticket.PurchaseTicketRequest.user:type_name -> trainticket.User
	0,  // 1: trainticket.Receipt.user:type_name -> trainticket.User
	6,  // 2: trainticket.SeatAllocationResponse.allocations:type_name -> trainticket.SeatAllocation
	0,  // 3: trainticket.SeatAllocation.user:type_name -> trainticket.User
	1,  // 4: trainticket.TicketService.PurchaseTicket:input_type -> trainticket.PurchaseTicketRequest
	3,  // 5: trainticket.TicketService.GetReceipt:input_type -> trainticket.GetReceiptRequest
	4,  // 6: trainticket.TicketService.ViewSeatAllocation:input_type -> trainticket.ViewSeatAllocationRequest
	7,  // 7: trainticket.TicketService.RemoveUser:input_type -> trainticket.RemoveUserRequest
	9,  // 8: trainticket.TicketService.ModifySeat:input_type -> trainticket.ModifySeatRequest
	2,  // 9: trainticket.TicketService.PurchaseTicket:output_type -> trainticket.Receipt
	2,  // 10: trainticket.TicketService.GetReceipt:output_type -> trainticket.Receipt
	5,  // 11: trainticket.TicketService.ViewSeatAllocation:output_type -> trainticket.SeatAllocationResponse
	8,  // 12: trainticket.TicketService.RemoveUser:output_type -> trainticket.RemoveUserResponse
	10, // 13: trainticket.TicketService.ModifySeat:output_type -> trainticket.ModifySeatResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ticket_proto_init() }
func file_proto_ticket_proto_init() {
	if File_proto_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ticket_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_proto_ticket_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseTicketRequest); i {
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
		file_proto_ticket_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Receipt); i {
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
		file_proto_ticket_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetReceiptRequest); i {
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
		file_proto_ticket_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ViewSeatAllocationRequest); i {
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
		file_proto_ticket_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SeatAllocationResponse); i {
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
		file_proto_ticket_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SeatAllocation); i {
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
		file_proto_ticket_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserRequest); i {
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
		file_proto_ticket_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserResponse); i {
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
		file_proto_ticket_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatRequest); i {
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
		file_proto_ticket_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatResponse); i {
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
			RawDescriptor: file_proto_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ticket_proto_goTypes,
		DependencyIndexes: file_proto_ticket_proto_depIdxs,
		MessageInfos:      file_proto_ticket_proto_msgTypes,
	}.Build()
	File_proto_ticket_proto = out.File
	file_proto_ticket_proto_rawDesc = nil
	file_proto_ticket_proto_goTypes = nil
	file_proto_ticket_proto_depIdxs = nil
}
