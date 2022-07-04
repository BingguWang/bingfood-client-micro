// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: v1/order_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SettleOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartIds []uint64 `protobuf:"varint,1,rep,packed,name=cartIds,proto3" json:"cartIds,omitempty"`
}

func (x *SettleOrderRequest) Reset() {
	*x = SettleOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleOrderRequest) ProtoMessage() {}

func (x *SettleOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleOrderRequest.ProtoReflect.Descriptor instead.
func (*SettleOrderRequest) Descriptor() ([]byte, []int) {
	return file_v1_order_service_proto_rawDescGZIP(), []int{0}
}

func (x *SettleOrderRequest) GetCartIds() []uint64 {
	if x != nil {
		return x.CartIds
	}
	return nil
}

type SettleOrderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32                 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string                 `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Data    *SettleOrderReply_Data `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SettleOrderReply) Reset() {
	*x = SettleOrderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleOrderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleOrderReply) ProtoMessage() {}

func (x *SettleOrderReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleOrderReply.ProtoReflect.Descriptor instead.
func (*SettleOrderReply) Descriptor() ([]byte, []int) {
	return file_v1_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *SettleOrderReply) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *SettleOrderReply) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *SettleOrderReply) GetData() *SettleOrderReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderItemId uint64 `protobuf:"varint,1,opt,name=OrderItemId,proto3" json:"OrderItemId,omitempty"`
	OrderNumber string `protobuf:"bytes,2,opt,name=OrderNumber,proto3" json:"OrderNumber,omitempty"`
	UserId      uint64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Score       int32  `protobuf:"varint,4,opt,name=Score,proto3" json:"Score,omitempty"`
	ShopId      uint64 `protobuf:"varint,5,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	ProdId      uint64 `protobuf:"varint,6,opt,name=ProdId,proto3" json:"ProdId,omitempty"`
	ProdName    string `protobuf:"bytes,7,opt,name=ProdName,proto3" json:"ProdName,omitempty"`
	ProdNums    uint32 `protobuf:"varint,8,opt,name=ProdNums,proto3" json:"ProdNums,omitempty"`
	Pic         string `protobuf:"bytes,9,opt,name=Pic,proto3" json:"Pic,omitempty"`
	ProdAmount  uint32 `protobuf:"varint,10,opt,name=ProdAmount,proto3" json:"ProdAmount,omitempty"`
	SkuId       uint32 `protobuf:"varint,11,opt,name=SkuId,proto3" json:"SkuId,omitempty"`
	SkuName     string `protobuf:"bytes,12,opt,name=SkuName,proto3" json:"SkuName,omitempty"`
	Price       int32  `protobuf:"varint,13,opt,name=Price,proto3" json:"Price,omitempty"`
	OriPrice    int32  `protobuf:"varint,14,opt,name=oriPrice,proto3" json:"oriPrice,omitempty"`
	PropId      uint64 `protobuf:"varint,15,opt,name=PropId,proto3" json:"PropId,omitempty"`
	PropName    string `protobuf:"bytes,16,opt,name=PropName,proto3" json:"PropName,omitempty"`
	IsCommented uint32 `protobuf:"varint,17,opt,name=IsCommented,proto3" json:"IsCommented,omitempty"`
	IsGood      uint32 `protobuf:"varint,18,opt,name=IsGood,proto3" json:"IsGood,omitempty"`
	Comment     string `protobuf:"bytes,19,opt,name=Comment,proto3" json:"Comment,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_v1_order_service_proto_rawDescGZIP(), []int{2}
}

func (x *OrderItem) GetOrderItemId() uint64 {
	if x != nil {
		return x.OrderItemId
	}
	return 0
}

func (x *OrderItem) GetOrderNumber() string {
	if x != nil {
		return x.OrderNumber
	}
	return ""
}

func (x *OrderItem) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderItem) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *OrderItem) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *OrderItem) GetProdId() uint64 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *OrderItem) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *OrderItem) GetProdNums() uint32 {
	if x != nil {
		return x.ProdNums
	}
	return 0
}

func (x *OrderItem) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *OrderItem) GetProdAmount() uint32 {
	if x != nil {
		return x.ProdAmount
	}
	return 0
}

func (x *OrderItem) GetSkuId() uint32 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *OrderItem) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *OrderItem) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderItem) GetOriPrice() int32 {
	if x != nil {
		return x.OriPrice
	}
	return 0
}

func (x *OrderItem) GetPropId() uint64 {
	if x != nil {
		return x.PropId
	}
	return 0
}

func (x *OrderItem) GetPropName() string {
	if x != nil {
		return x.PropName
	}
	return ""
}

func (x *OrderItem) GetIsCommented() uint32 {
	if x != nil {
		return x.IsCommented
	}
	return 0
}

func (x *OrderItem) GetIsGood() uint32 {
	if x != nil {
		return x.IsGood
	}
	return 0
}

func (x *OrderItem) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type SettleOrderReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId         uint64     `protobuf:"varint,1,opt,name=ShopId,proto3" json:"ShopId,omitempty"`
	UserId         uint64     `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	UserMobile     string     `protobuf:"bytes,3,opt,name=UserMobile,proto3" json:"UserMobile,omitempty"`
	ReceiverMobile string     `protobuf:"bytes,4,opt,name=ReceiverMobile,proto3" json:"ReceiverMobile,omitempty"`
	ProdName       string     `protobuf:"bytes,5,opt,name=ProdName,proto3" json:"ProdName,omitempty"`
	ProdNums       int32      `protobuf:"varint,6,opt,name=ProdNums,proto3" json:"ProdNums,omitempty"`
	PackingAmount  int32      `protobuf:"varint,7,opt,name=PackingAmount,proto3" json:"PackingAmount,omitempty"`
	DeliverAmount  int32      `protobuf:"varint,8,opt,name=DeliverAmount,proto3" json:"DeliverAmount,omitempty"`
	ProdAmount     int32      `protobuf:"varint,9,opt,name=ProdAmount,proto3" json:"ProdAmount,omitempty"`
	DiscountAmount int32      `protobuf:"varint,10,opt,name=DiscountAmount,proto3" json:"DiscountAmount,omitempty"`
	FinalAmount    int32      `protobuf:"varint,11,opt,name=FinalAmount,proto3" json:"FinalAmount,omitempty"`
	DeliverType    uint32     `protobuf:"varint,12,opt,name=DeliverType,proto3" json:"DeliverType,omitempty"`
	OrderItems     *OrderItem `protobuf:"bytes,13,opt,name=OrderItems,proto3" json:"OrderItems,omitempty"`
}

func (x *SettleOrderReply_Data) Reset() {
	*x = SettleOrderReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_order_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettleOrderReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettleOrderReply_Data) ProtoMessage() {}

func (x *SettleOrderReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_v1_order_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettleOrderReply_Data.ProtoReflect.Descriptor instead.
func (*SettleOrderReply_Data) Descriptor() ([]byte, []int) {
	return file_v1_order_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SettleOrderReply_Data) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *SettleOrderReply_Data) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SettleOrderReply_Data) GetUserMobile() string {
	if x != nil {
		return x.UserMobile
	}
	return ""
}

func (x *SettleOrderReply_Data) GetReceiverMobile() string {
	if x != nil {
		return x.ReceiverMobile
	}
	return ""
}

func (x *SettleOrderReply_Data) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *SettleOrderReply_Data) GetProdNums() int32 {
	if x != nil {
		return x.ProdNums
	}
	return 0
}

func (x *SettleOrderReply_Data) GetPackingAmount() int32 {
	if x != nil {
		return x.PackingAmount
	}
	return 0
}

func (x *SettleOrderReply_Data) GetDeliverAmount() int32 {
	if x != nil {
		return x.DeliverAmount
	}
	return 0
}

func (x *SettleOrderReply_Data) GetProdAmount() int32 {
	if x != nil {
		return x.ProdAmount
	}
	return 0
}

func (x *SettleOrderReply_Data) GetDiscountAmount() int32 {
	if x != nil {
		return x.DiscountAmount
	}
	return 0
}

func (x *SettleOrderReply_Data) GetFinalAmount() int32 {
	if x != nil {
		return x.FinalAmount
	}
	return 0
}

func (x *SettleOrderReply_Data) GetDeliverType() uint32 {
	if x != nil {
		return x.DeliverType
	}
	return 0
}

func (x *SettleOrderReply_Data) GetOrderItems() *OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

var File_v1_order_service_proto protoreflect.FileDescriptor

var file_v1_order_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x07, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x73, 0x22, 0xcf, 0x04, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12,
	0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xcb, 0x03, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x4e, 0x75, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x72, 0x6f,
	0x64, 0x4e, 0x75, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x50, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x46, 0x69, 0x6e, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a,
	0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x81, 0x04, 0x0a, 0x09, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x68,
	0x6f, 0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x53, 0x68, 0x6f, 0x70,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x75,
	0x6d, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x4e, 0x75,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x50, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x6b,
	0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x6b, 0x75,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72,
	0x69, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x72,
	0x69, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x70, 0x49, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x50, 0x72, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x49, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x47, 0x6f, 0x6f, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x49, 0x73,
	0x47, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x83,
	0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x73, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x0f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x74, 0x6f, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x3a, 0x01, 0x2a, 0x42, 0x19, 0x5a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_order_service_proto_rawDescOnce sync.Once
	file_v1_order_service_proto_rawDescData = file_v1_order_service_proto_rawDesc
)

func file_v1_order_service_proto_rawDescGZIP() []byte {
	file_v1_order_service_proto_rawDescOnce.Do(func() {
		file_v1_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_order_service_proto_rawDescData)
	})
	return file_v1_order_service_proto_rawDescData
}

var file_v1_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_v1_order_service_proto_goTypes = []interface{}{
	(*SettleOrderRequest)(nil),    // 0: order.service.v1.SettleOrderRequest
	(*SettleOrderReply)(nil),      // 1: order.service.v1.SettleOrderReply
	(*OrderItem)(nil),             // 2: order.service.v1.OrderItem
	(*SettleOrderReply_Data)(nil), // 3: order.service.v1.SettleOrderReply.Data
}
var file_v1_order_service_proto_depIdxs = []int32{
	3, // 0: order.service.v1.SettleOrderReply.data:type_name -> order.service.v1.SettleOrderReply.Data
	2, // 1: order.service.v1.SettleOrderReply.Data.OrderItems:type_name -> order.service.v1.OrderItem
	0, // 2: order.service.v1.OrderService.SettleOrder:input_type -> order.service.v1.SettleOrderRequest
	1, // 3: order.service.v1.OrderService.SettleOrder:output_type -> order.service.v1.SettleOrderReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_order_service_proto_init() }
func file_v1_order_service_proto_init() {
	if File_v1_order_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_order_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleOrderRequest); i {
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
		file_v1_order_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleOrderReply); i {
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
		file_v1_order_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_v1_order_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettleOrderReply_Data); i {
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
			RawDescriptor: file_v1_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_order_service_proto_goTypes,
		DependencyIndexes: file_v1_order_service_proto_depIdxs,
		MessageInfos:      file_v1_order_service_proto_msgTypes,
	}.Build()
	File_v1_order_service_proto = out.File
	file_v1_order_service_proto_rawDesc = nil
	file_v1_order_service_proto_goTypes = nil
	file_v1_order_service_proto_depIdxs = nil
}