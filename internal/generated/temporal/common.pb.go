// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: common.proto

package temporal

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
	return file_common_proto_enumTypes[0].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type PaymentStatus int32

const (
	PaymentStatus_PaymentStatusNew     PaymentStatus = 0
	PaymentStatus_PaymentStatusHold    PaymentStatus = 1
	PaymentStatus_PaymentStatusCharged PaymentStatus = 2
	PaymentStatus_PaymentStatusFailed  PaymentStatus = 3
)

// Enum value maps for PaymentStatus.
var (
	PaymentStatus_name = map[int32]string{
		0: "PaymentStatusNew",
		1: "PaymentStatusHold",
		2: "PaymentStatusCharged",
		3: "PaymentStatusFailed",
	}
	PaymentStatus_value = map[string]int32{
		"PaymentStatusNew":     0,
		"PaymentStatusHold":    1,
		"PaymentStatusCharged": 2,
		"PaymentStatusFailed":  3,
	}
)

func (x PaymentStatus) Enum() *PaymentStatus {
	p := new(PaymentStatus)
	*p = x
	return p
}

func (x PaymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (PaymentStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x PaymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentStatus.Descriptor instead.
func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type VendorOrderStatus int32

const (
	VendorOrderStatus_VendorOrderStatusNew        VendorOrderStatus = 0
	VendorOrderStatus_VendorOrderStatusConfirmed  VendorOrderStatus = 1
	VendorOrderStatus_VendorOrderStatusPicking    VendorOrderStatus = 2
	VendorOrderStatus_VendorOrderStatusReady      VendorOrderStatus = 3
	VendorOrderStatus_VendorOrderInStatusDelivery VendorOrderStatus = 4
	VendorOrderStatus_VendorOrderStatusCancelled  VendorOrderStatus = 99
)

// Enum value maps for VendorOrderStatus.
var (
	VendorOrderStatus_name = map[int32]string{
		0:  "VendorOrderStatusNew",
		1:  "VendorOrderStatusConfirmed",
		2:  "VendorOrderStatusPicking",
		3:  "VendorOrderStatusReady",
		4:  "VendorOrderInStatusDelivery",
		99: "VendorOrderStatusCancelled",
	}
	VendorOrderStatus_value = map[string]int32{
		"VendorOrderStatusNew":        0,
		"VendorOrderStatusConfirmed":  1,
		"VendorOrderStatusPicking":    2,
		"VendorOrderStatusReady":      3,
		"VendorOrderInStatusDelivery": 4,
		"VendorOrderStatusCancelled":  99,
	}
)

func (x VendorOrderStatus) Enum() *VendorOrderStatus {
	p := new(VendorOrderStatus)
	*p = x
	return p
}

func (x VendorOrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VendorOrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (VendorOrderStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x VendorOrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VendorOrderStatus.Descriptor instead.
func (VendorOrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

type OrderStatus int32

const (
	OrderStatus_OrderStatusNew             OrderStatus = 0
	OrderStatus_OrderStatusConfirmed       OrderStatus = 1
	OrderStatus_OrderStatusVendorConfirmed OrderStatus = 2
	OrderStatus_OrderStatusPicking         OrderStatus = 3
	OrderStatus_OrderStatusReady           OrderStatus = 4
	OrderStatus_OrderStatusInDelivery      OrderStatus = 5
	OrderStatus_OrderStatusDelivered       OrderStatus = 6
	OrderStatus_OrderStatusDone            OrderStatus = 7
	OrderStatus_OrderStatusCancelled       OrderStatus = 99
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0:  "OrderStatusNew",
		1:  "OrderStatusConfirmed",
		2:  "OrderStatusVendorConfirmed",
		3:  "OrderStatusPicking",
		4:  "OrderStatusReady",
		5:  "OrderStatusInDelivery",
		6:  "OrderStatusDelivered",
		7:  "OrderStatusDone",
		99: "OrderStatusCancelled",
	}
	OrderStatus_value = map[string]int32{
		"OrderStatusNew":             0,
		"OrderStatusConfirmed":       1,
		"OrderStatusVendorConfirmed": 2,
		"OrderStatusPicking":         3,
		"OrderStatusReady":           4,
		"OrderStatusInDelivery":      5,
		"OrderStatusDelivered":       6,
		"OrderStatusDone":            7,
		"OrderStatusCancelled":       99,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[3].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[3]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
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
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
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
	return file_common_proto_rawDescGZIP(), []int{0}
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
	Id       string     `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
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
	return file_common_proto_rawDescGZIP(), []int{1}
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

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Customer    *Profile               `protobuf:"bytes,2,opt,name=customer,proto3" json:"customer,omitempty"`
	Cart        *Cart                  `protobuf:"bytes,3,opt,name=cart,proto3" json:"cart,omitempty"`
	PaymentType PaymentType            `protobuf:"varint,4,opt,name=paymentType,proto3,enum=temporal.PaymentType" json:"paymentType,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Status      OrderStatus            `protobuf:"varint,7,opt,name=status,proto3,enum=temporal.OrderStatus" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
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
	return file_common_proto_rawDescGZIP(), []int{2}
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

func (x *Order) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Order) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Order) GetStatus() OrderStatus {
	if x != nil {
		return x.Status
	}
	return OrderStatus_OrderStatusNew
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
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
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
	return file_common_proto_rawDescGZIP(), []int{3}
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
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
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
	return file_common_proto_rawDescGZIP(), []int{4}
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

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6e, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x6e, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71,
	0x74, 0x79, 0x22, 0x5b, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xc6, 0x02, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x38, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x70, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x2b, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x45, 0x0a, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x2a, 0x23, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x4e,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x2a, 0x6f, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x65, 0x77, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x6f,
	0x6c, 0x64, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x03, 0x2a, 0xc8, 0x01, 0x0a, 0x11, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x14, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4e, 0x65, 0x77, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x69, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10,
	0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x10, 0x63, 0x2a, 0xed, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4e, 0x65, 0x77, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x01,
	0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50,
	0x69, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x04, 0x12, 0x19,
	0x0a, 0x15, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x49, 0x6e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65,
	0x64, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x07, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x10, 0x63, 0x42, 0x33, 0x5a, 0x31, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2d, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_common_proto_goTypes = []any{
	(PaymentType)(0),              // 0: temporal.PaymentType
	(PaymentStatus)(0),            // 1: temporal.PaymentStatus
	(VendorOrderStatus)(0),        // 2: temporal.VendorOrderStatus
	(OrderStatus)(0),              // 3: temporal.OrderStatus
	(*Product)(nil),               // 4: temporal.Product
	(*Cart)(nil),                  // 5: temporal.Cart
	(*Order)(nil),                 // 6: temporal.Order
	(*Profile)(nil),               // 7: temporal.Profile
	(*Address)(nil),               // 8: temporal.Address
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_common_proto_depIdxs = []int32{
	4, // 0: temporal.Cart.products:type_name -> temporal.Product
	7, // 1: temporal.Order.customer:type_name -> temporal.Profile
	5, // 2: temporal.Order.cart:type_name -> temporal.Cart
	0, // 3: temporal.Order.paymentType:type_name -> temporal.PaymentType
	9, // 4: temporal.Order.createdAt:type_name -> google.protobuf.Timestamp
	9, // 5: temporal.Order.updatedAt:type_name -> google.protobuf.Timestamp
	3, // 6: temporal.Order.status:type_name -> temporal.OrderStatus
	8, // 7: temporal.Profile.address:type_name -> temporal.Address
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
