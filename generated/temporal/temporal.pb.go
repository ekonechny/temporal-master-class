// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: temporal.proto

package temporal

import (
	_ "github.com/cludden/protoc-gen-go-temporal/gen/temporal/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentType int32

const (
	PaymentType_CASH   PaymentType = 0
	PaymentType_ONLINE PaymentType = 1
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0: "CASH",
		1: "ONLINE",
	}
	PaymentType_value = map[string]int32{
		"CASH":   0,
		"ONLINE": 1,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_temporal_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{0}
}

type CustomerFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *CustomerFlowRequest) Reset() {
	*x = CustomerFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerFlowRequest) ProtoMessage() {}

func (x *CustomerFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerFlowRequest.ProtoReflect.Descriptor instead.
func (*CustomerFlowRequest) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerFlowRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomerFlowRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Lat   string `protobuf:"bytes,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long  string `protobuf:"bytes,3,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[1]
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
	return file_temporal_proto_rawDescGZIP(), []int{1}
}

func (x *Address) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Address) GetLat() string {
	if x != nil {
		return x.Lat
	}
	return ""
}

func (x *Address) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone   string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Address *Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Profile) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type UpdateProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateProfileRequest) Reset() {
	*x = UpdateProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileRequest) ProtoMessage() {}

func (x *UpdateProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SetAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *Address `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *SetAddressRequest) Reset() {
	*x = SetAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAddressRequest) ProtoMessage() {}

func (x *SetAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAddressRequest.ProtoReflect.Descriptor instead.
func (*SetAddressRequest) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{4}
}

func (x *SetAddressRequest) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type UpdateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *UpdateCartRequest) Reset() {
	*x = UpdateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartRequest) ProtoMessage() {}

func (x *UpdateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartRequest) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCartRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentType PaymentType `protobuf:"varint,1,opt,name=paymentType,proto3,enum=temporal.PaymentType" json:"paymentType,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{6}
}

func (x *CheckoutRequest) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_CASH
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Customer    *Profile    `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	Cart        *Cart       `protobuf:"bytes,3,opt,name=cart,proto3" json:"cart,omitempty"`
	PaymentType PaymentType `protobuf:"varint,4,opt,name=paymentType,proto3,enum=temporal.PaymentType" json:"paymentType,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{7}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetCustomer() *Profile {
	if x != nil {
		return x.Customer
	}
	return nil
}

func (x *Order) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *Order) GetPaymentType() PaymentType {
	if x != nil {
		return x.PaymentType
	}
	return PaymentType_CASH
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Inn   string `protobuf:"bytes,4,opt,name=inn,proto3" json:"inn,omitempty"`
	Qty   int32  `protobuf:"varint,5,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{8}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *Product) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total    int32      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_temporal_proto_rawDescGZIP(), []int{9}
}

func (x *Cart) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *Cart) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_temporal_proto protoreflect.FileDescriptor

var file_temporal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x45, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x70, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x2b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x53, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x4a, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x67, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0x4b, 0x0a, 0x04, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x2a, 0x23, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x48, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x32, 0xa2, 0x06, 0x0a,
	0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0xf1, 0x01, 0x0a, 0x0c, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0xa9, 0x01, 0x8a, 0xc4, 0x03, 0xa4, 0x01, 0x0a, 0x0c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x0a, 0x09, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x0f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0c, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x1a,
	0x0f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x1a, 0x0c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x1a, 0x0a,
	0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2a, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x24, 0x7b, 0x21, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x2e, 0x6f, 0x72, 0x28, 0x69, 0x64, 0x2e, 0x6f, 0x72, 0x28, 0x75, 0x75,
	0x69, 0x64, 0x5f, 0x76, 0x34, 0x28, 0x29, 0x29, 0x29, 0x20, 0x7d, 0x30, 0x01, 0x12, 0x3d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x04, 0x9a, 0xc4, 0x03, 0x00, 0x12, 0x48, 0x0a, 0x0d,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x04, 0xaa, 0xc4, 0x03, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x12, 0x47, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x04, 0xa2, 0xc4, 0x03, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x22, 0x04, 0x9a, 0xc4, 0x03, 0x00, 0x12,
	0x3f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x1b, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x22, 0x04, 0xaa, 0xc4, 0x03, 0x00,
	0x12, 0x42, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x04,
	0xa2, 0xc4, 0x03, 0x00, 0x12, 0x3c, 0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x12, 0x19, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x04, 0xaa, 0xc4,
	0x03, 0x00, 0x1a, 0x0d, 0x8a, 0xc4, 0x03, 0x09, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x74, 0x2d, 0x76,
	0x31, 0x42, 0x2a, 0x5a, 0x28, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_proto_rawDescOnce sync.Once
	file_temporal_proto_rawDescData = file_temporal_proto_rawDesc
)

func file_temporal_proto_rawDescGZIP() []byte {
	file_temporal_proto_rawDescOnce.Do(func() {
		file_temporal_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_proto_rawDescData)
	})
	return file_temporal_proto_rawDescData
}

var file_temporal_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_temporal_proto_goTypes = []any{
	(PaymentType)(0),             // 0: temporal.PaymentType
	(*CustomerFlowRequest)(nil),  // 1: temporal.CustomerFlowRequest
	(*Address)(nil),              // 2: temporal.Address
	(*Profile)(nil),              // 3: temporal.Profile
	(*UpdateProfileRequest)(nil), // 4: temporal.UpdateProfileRequest
	(*SetAddressRequest)(nil),    // 5: temporal.SetAddressRequest
	(*UpdateCartRequest)(nil),    // 6: temporal.UpdateCartRequest
	(*CheckoutRequest)(nil),      // 7: temporal.CheckoutRequest
	(*Order)(nil),                // 8: temporal.Order
	(*Product)(nil),              // 9: temporal.Product
	(*Cart)(nil),                 // 10: temporal.Cart
	(*emptypb.Empty)(nil),        // 11: google.protobuf.Empty
}
var file_temporal_proto_depIdxs = []int32{
	2,  // 0: temporal.Profile.address:type_name -> temporal.Address
	2,  // 1: temporal.SetAddressRequest.address:type_name -> temporal.Address
	9,  // 2: temporal.UpdateCartRequest.products:type_name -> temporal.Product
	0,  // 3: temporal.CheckoutRequest.paymentType:type_name -> temporal.PaymentType
	3,  // 4: temporal.Order.customer:type_name -> temporal.Profile
	10, // 5: temporal.Order.cart:type_name -> temporal.Cart
	0,  // 6: temporal.Order.paymentType:type_name -> temporal.PaymentType
	9,  // 7: temporal.Cart.products:type_name -> temporal.Product
	1,  // 8: temporal.Customer.CustomerFlow:input_type -> temporal.CustomerFlowRequest
	11, // 9: temporal.Customer.GetProfile:input_type -> google.protobuf.Empty
	4,  // 10: temporal.Customer.UpdateProfile:input_type -> temporal.UpdateProfileRequest
	11, // 11: temporal.Customer.DeleteProfile:input_type -> google.protobuf.Empty
	5,  // 12: temporal.Customer.SetAddress:input_type -> temporal.SetAddressRequest
	11, // 13: temporal.Customer.GetCart:input_type -> google.protobuf.Empty
	6,  // 14: temporal.Customer.UpdateCart:input_type -> temporal.UpdateCartRequest
	11, // 15: temporal.Customer.DeleteCart:input_type -> google.protobuf.Empty
	7,  // 16: temporal.Customer.Checkout:input_type -> temporal.CheckoutRequest
	11, // 17: temporal.Customer.CustomerFlow:output_type -> google.protobuf.Empty
	3,  // 18: temporal.Customer.GetProfile:output_type -> temporal.Profile
	3,  // 19: temporal.Customer.UpdateProfile:output_type -> temporal.Profile
	11, // 20: temporal.Customer.DeleteProfile:output_type -> google.protobuf.Empty
	11, // 21: temporal.Customer.SetAddress:output_type -> google.protobuf.Empty
	10, // 22: temporal.Customer.GetCart:output_type -> temporal.Cart
	10, // 23: temporal.Customer.UpdateCart:output_type -> temporal.Cart
	11, // 24: temporal.Customer.DeleteCart:output_type -> google.protobuf.Empty
	8,  // 25: temporal.Customer.Checkout:output_type -> temporal.Order
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_temporal_proto_init() }
func file_temporal_proto_init() {
	if File_temporal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomerFlowRequest); i {
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
		file_temporal_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_temporal_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Profile); i {
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
		file_temporal_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateProfileRequest); i {
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
		file_temporal_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SetAddressRequest); i {
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
		file_temporal_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCartRequest); i {
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
		file_temporal_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CheckoutRequest); i {
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
		file_temporal_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Order); i {
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
		file_temporal_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
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
		file_temporal_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Cart); i {
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
			RawDescriptor: file_temporal_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_temporal_proto_goTypes,
		DependencyIndexes: file_temporal_proto_depIdxs,
		EnumInfos:         file_temporal_proto_enumTypes,
		MessageInfos:      file_temporal_proto_msgTypes,
	}.Build()
	File_temporal_proto = out.File
	file_temporal_proto_rawDesc = nil
	file_temporal_proto_goTypes = nil
	file_temporal_proto_depIdxs = nil
}
