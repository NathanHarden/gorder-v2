// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: orderpb/order.proto

package orderpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CustomerID    string                 `protobuf:"bytes,1,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Items         []*ItemWithQuantity    `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_orderpb_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateOrderRequest) GetItems() []*ItemWithQuantity {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderID       string                 `protobuf:"bytes,1,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	mi := &file_orderpb_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetOrderRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

type ItemWithQuantity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Quantity      int32                  `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ItemWithQuantity) Reset() {
	*x = ItemWithQuantity{}
	mi := &file_orderpb_order_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ItemWithQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemWithQuantity) ProtoMessage() {}

func (x *ItemWithQuantity) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemWithQuantity.ProtoReflect.Descriptor instead.
func (*ItemWithQuantity) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{2}
}

func (x *ItemWithQuantity) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ItemWithQuantity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	PriceID       string                 `protobuf:"bytes,4,opt,name=PriceID,proto3" json:"PriceID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Item) Reset() {
	*x = Item{}
	mi := &file_orderpb_order_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_orderpb_order_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Item) GetPriceID() string {
	if x != nil {
		return x.PriceID
	}
	return ""
}

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	CustomerID    string                 `protobuf:"bytes,2,opt,name=CustomerID,proto3" json:"CustomerID,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=Status,proto3" json:"Status,omitempty"`
	Items         []*Item                `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
	PaymentLink   string                 `protobuf:"bytes,5,opt,name=PaymentLink,proto3" json:"PaymentLink,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_orderpb_order_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_orderpb_order_proto_msgTypes[4]
	if x != nil {
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
	return file_orderpb_order_proto_rawDescGZIP(), []int{4}
}

func (x *Order) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Order) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Order) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetPaymentLink() string {
	if x != nil {
		return x.PaymentLink
	}
	return ""
}

var File_orderpb_order_proto protoreflect.FileDescriptor

const file_orderpb_order_proto_rawDesc = "" +
	"\n" +
	"\x13orderpb/order.proto\x12\aorderpb\x1a\x1bgoogle/protobuf/empty.proto\"e\n" +
	"\x12CreateOrderRequest\x12\x1e\n" +
	"\n" +
	"CustomerID\x18\x01 \x01(\tR\n" +
	"CustomerID\x12/\n" +
	"\x05Items\x18\x02 \x03(\v2\x19.orderpb.ItemWithQuantityR\x05Items\"K\n" +
	"\x0fGetOrderRequest\x12\x18\n" +
	"\aOrderID\x18\x01 \x01(\tR\aOrderID\x12\x1e\n" +
	"\n" +
	"CustomerID\x18\x02 \x01(\tR\n" +
	"CustomerID\">\n" +
	"\x10ItemWithQuantity\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x1a\n" +
	"\bQuantity\x18\x02 \x01(\x05R\bQuantity\"`\n" +
	"\x04Item\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x12\n" +
	"\x04Name\x18\x02 \x01(\tR\x04Name\x12\x1a\n" +
	"\bQuantity\x18\x03 \x01(\x05R\bQuantity\x12\x18\n" +
	"\aPriceID\x18\x04 \x01(\tR\aPriceID\"\x96\x01\n" +
	"\x05Order\x12\x0e\n" +
	"\x02ID\x18\x01 \x01(\tR\x02ID\x12\x1e\n" +
	"\n" +
	"CustomerID\x18\x02 \x01(\tR\n" +
	"CustomerID\x12\x16\n" +
	"\x06Status\x18\x03 \x01(\tR\x06Status\x12#\n" +
	"\x05Items\x18\x04 \x03(\v2\r.orderpb.ItemR\x05Items\x12 \n" +
	"\vPaymentLink\x18\x05 \x01(\tR\vPaymentLink2\xbf\x01\n" +
	"\fOrderService\x12B\n" +
	"\vCreateOrder\x12\x1b.orderpb.CreateOrderRequest\x1a\x16.google.protobuf.Empty\x124\n" +
	"\bGetOrder\x12\x18.orderpb.GetOrderRequest\x1a\x0e.orderpb.Order\x125\n" +
	"\vUpdateOrder\x12\x0e.orderpb.Order\x1a\x16.google.protobuf.EmptyB;Z9github.com/NathanHarden/gorder-v2/common/genproto/orderpbb\x06proto3"

var (
	file_orderpb_order_proto_rawDescOnce sync.Once
	file_orderpb_order_proto_rawDescData []byte
)

func file_orderpb_order_proto_rawDescGZIP() []byte {
	file_orderpb_order_proto_rawDescOnce.Do(func() {
		file_orderpb_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)))
	})
	return file_orderpb_order_proto_rawDescData
}

var file_orderpb_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orderpb_order_proto_goTypes = []any{
	(*CreateOrderRequest)(nil), // 0: orderpb.CreateOrderRequest
	(*GetOrderRequest)(nil),    // 1: orderpb.GetOrderRequest
	(*ItemWithQuantity)(nil),   // 2: orderpb.ItemWithQuantity
	(*Item)(nil),               // 3: orderpb.Item
	(*Order)(nil),              // 4: orderpb.Order
	(*emptypb.Empty)(nil),      // 5: google.protobuf.Empty
}
var file_orderpb_order_proto_depIdxs = []int32{
	2, // 0: orderpb.CreateOrderRequest.Items:type_name -> orderpb.ItemWithQuantity
	3, // 1: orderpb.Order.Items:type_name -> orderpb.Item
	0, // 2: orderpb.OrderService.CreateOrder:input_type -> orderpb.CreateOrderRequest
	1, // 3: orderpb.OrderService.GetOrder:input_type -> orderpb.GetOrderRequest
	4, // 4: orderpb.OrderService.UpdateOrder:input_type -> orderpb.Order
	5, // 5: orderpb.OrderService.CreateOrder:output_type -> google.protobuf.Empty
	4, // 6: orderpb.OrderService.GetOrder:output_type -> orderpb.Order
	5, // 7: orderpb.OrderService.UpdateOrder:output_type -> google.protobuf.Empty
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_orderpb_order_proto_init() }
func file_orderpb_order_proto_init() {
	if File_orderpb_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_orderpb_order_proto_rawDesc), len(file_orderpb_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderpb_order_proto_goTypes,
		DependencyIndexes: file_orderpb_order_proto_depIdxs,
		MessageInfos:      file_orderpb_order_proto_msgTypes,
	}.Build()
	File_orderpb_order_proto = out.File
	file_orderpb_order_proto_goTypes = nil
	file_orderpb_order_proto_depIdxs = nil
}
