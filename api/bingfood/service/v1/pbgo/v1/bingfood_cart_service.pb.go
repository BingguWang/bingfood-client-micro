// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: v1/bingfood_cart_service.proto

package v1

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

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{0}
}

func (x *PageInfo) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageInfo) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type GetCartByCondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartCond *Cart     `protobuf:"bytes,1,opt,name=cartCond,proto3" json:"cartCond,omitempty"`
	PageInfo *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
}

func (x *GetCartByCondRequest) Reset() {
	*x = GetCartByCondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartByCondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartByCondRequest) ProtoMessage() {}

func (x *GetCartByCondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartByCondRequest.ProtoReflect.Descriptor instead.
func (*GetCartByCondRequest) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCartByCondRequest) GetCartCond() *Cart {
	if x != nil {
		return x.CartCond
	}
	return nil
}

func (x *GetCartByCondRequest) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type GetCartByCartIdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids      []uint64  `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	PageInfo *PageInfo `protobuf:"bytes,2,opt,name=pageInfo,proto3" json:"pageInfo,omitempty"`
}

func (x *GetCartByCartIdsRequest) Reset() {
	*x = GetCartByCartIdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartByCartIdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartByCartIdsRequest) ProtoMessage() {}

func (x *GetCartByCartIdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartByCartIdsRequest.ProtoReflect.Descriptor instead.
func (*GetCartByCartIdsRequest) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetCartByCartIdsRequest) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetCartByCartIdsRequest) GetPageInfo() *PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type GetCartByCartIdsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32          `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string          `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Data    *CartPagination `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCartByCartIdsReply) Reset() {
	*x = GetCartByCartIdsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartByCartIdsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartByCartIdsReply) ProtoMessage() {}

func (x *GetCartByCartIdsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartByCartIdsReply.ProtoReflect.Descriptor instead.
func (*GetCartByCartIdsReply) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetCartByCartIdsReply) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *GetCartByCartIdsReply) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *GetCartByCartIdsReply) GetData() *CartPagination {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCartByCondReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32          `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string          `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Data    *CartPagination `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetCartByCondReply) Reset() {
	*x = GetCartByCondReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartByCondReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartByCondReply) ProtoMessage() {}

func (x *GetCartByCondReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartByCondReply.ProtoReflect.Descriptor instead.
func (*GetCartByCondReply) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartByCondReply) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *GetCartByCondReply) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *GetCartByCondReply) GetData() *CartPagination {
	if x != nil {
		return x.Data
	}
	return nil
}

type CartPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List     []*Cart `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total    int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Page     int64   `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64   `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *CartPagination) Reset() {
	*x = CartPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartPagination) ProtoMessage() {}

func (x *CartPagination) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartPagination.ProtoReflect.Descriptor instead.
func (*CartPagination) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{5}
}

func (x *CartPagination) GetList() []*Cart {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *CartPagination) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CartPagination) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CartPagination) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type AddCartItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId   uint64 `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	ShopId   uint64 `protobuf:"varint,2,opt,name=shopId,proto3" json:"shopId,omitempty"`
	UserId   uint64 `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`
	SkuId    uint64 `protobuf:"varint,4,opt,name=skuId,proto3" json:"skuId,omitempty"`
	ProdNums int32  `protobuf:"varint,5,opt,name=prodNums,proto3" json:"prodNums,omitempty"`
}

func (x *AddCartItemRequest) Reset() {
	*x = AddCartItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartItemRequest) ProtoMessage() {}

func (x *AddCartItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartItemRequest.ProtoReflect.Descriptor instead.
func (*AddCartItemRequest) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{6}
}

func (x *AddCartItemRequest) GetCartId() uint64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *AddCartItemRequest) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *AddCartItemRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddCartItemRequest) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *AddCartItemRequest) GetProdNums() int32 {
	if x != nil {
		return x.ProdNums
	}
	return 0
}

type AddCartItemReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode uint32 `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
}

func (x *AddCartItemReply) Reset() {
	*x = AddCartItemReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCartItemReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCartItemReply) ProtoMessage() {}

func (x *AddCartItemReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCartItemReply.ProtoReflect.Descriptor instead.
func (*AddCartItemReply) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{7}
}

func (x *AddCartItemReply) GetRetCode() uint32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *AddCartItemReply) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartId   uint64    `protobuf:"varint,1,opt,name=cartId,proto3" json:"cartId,omitempty"`
	UserId   uint64    `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	ShopId   uint64    `protobuf:"varint,3,opt,name=shopId,proto3" json:"shopId,omitempty"`
	SkuId    uint64    `protobuf:"varint,4,opt,name=skuId,proto3" json:"skuId,omitempty"`
	ProdId   uint64    `protobuf:"varint,5,opt,name=prodId,proto3" json:"prodId,omitempty"`
	ProdNums int32     `protobuf:"varint,6,opt,name=prodNums,proto3" json:"prodNums,omitempty"`
	Sku      *Cart_Sku `protobuf:"bytes,9,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[8]
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
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{8}
}

func (x *Cart) GetCartId() uint64 {
	if x != nil {
		return x.CartId
	}
	return 0
}

func (x *Cart) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetShopId() uint64 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Cart) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *Cart) GetProdId() uint64 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *Cart) GetProdNums() int32 {
	if x != nil {
		return x.ProdNums
	}
	return 0
}

func (x *Cart) GetSku() *Cart_Sku {
	if x != nil {
		return x.Sku
	}
	return nil
}

type Cart_Sku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkuId      uint64 `protobuf:"varint,1,opt,name=skuId,proto3" json:"skuId,omitempty"`
	SkuName    string `protobuf:"bytes,2,opt,name=skuName,proto3" json:"skuName,omitempty"`
	ProdName   string `protobuf:"bytes,3,opt,name=prodName,proto3" json:"prodName,omitempty"`
	ProdId     uint64 `protobuf:"varint,4,opt,name=prodId,proto3" json:"prodId,omitempty"`
	Price      int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	OriPrice   int32  `protobuf:"varint,6,opt,name=oriPrice,proto3" json:"oriPrice,omitempty"`
	PackingFee int32  `protobuf:"varint,7,opt,name=packingFee,proto3" json:"packingFee,omitempty"`
	ShopId     uint32 `protobuf:"varint,8,opt,name=shopId,proto3" json:"shopId,omitempty"`
	Pic        string `protobuf:"bytes,9,opt,name=pic,proto3" json:"pic,omitempty"`
	Imags      string `protobuf:"bytes,10,opt,name=imags,proto3" json:"imags,omitempty"`
	Weight     uint32 `protobuf:"varint,11,opt,name=weight,proto3" json:"weight,omitempty"`
	SellStatus uint32 `protobuf:"varint,12,opt,name=sellStatus,proto3" json:"sellStatus,omitempty"`
	Stock      uint32 `protobuf:"varint,13,opt,name=stock,proto3" json:"stock,omitempty"`
}

func (x *Cart_Sku) Reset() {
	*x = Cart_Sku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_bingfood_cart_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart_Sku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart_Sku) ProtoMessage() {}

func (x *Cart_Sku) ProtoReflect() protoreflect.Message {
	mi := &file_v1_bingfood_cart_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart_Sku.ProtoReflect.Descriptor instead.
func (*Cart_Sku) Descriptor() ([]byte, []int) {
	return file_v1_bingfood_cart_service_proto_rawDescGZIP(), []int{8, 0}
}

func (x *Cart_Sku) GetSkuId() uint64 {
	if x != nil {
		return x.SkuId
	}
	return 0
}

func (x *Cart_Sku) GetSkuName() string {
	if x != nil {
		return x.SkuName
	}
	return ""
}

func (x *Cart_Sku) GetProdName() string {
	if x != nil {
		return x.ProdName
	}
	return ""
}

func (x *Cart_Sku) GetProdId() uint64 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

func (x *Cart_Sku) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Cart_Sku) GetOriPrice() int32 {
	if x != nil {
		return x.OriPrice
	}
	return 0
}

func (x *Cart_Sku) GetPackingFee() int32 {
	if x != nil {
		return x.PackingFee
	}
	return 0
}

func (x *Cart_Sku) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *Cart_Sku) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *Cart_Sku) GetImags() string {
	if x != nil {
		return x.Imags
	}
	return ""
}

func (x *Cart_Sku) GetWeight() uint32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Cart_Sku) GetSellStatus() uint32 {
	if x != nil {
		return x.SellStatus
	}
	return 0
}

func (x *Cart_Sku) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

var File_v1_bingfood_cart_service_proto protoreflect.FileDescriptor

var file_v1_bingfood_cart_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x61,
	0x72, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62,
	0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x08, 0x63, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x66, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69,
	0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73,
	0x67, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x85, 0x01, 0x0a, 0x0e, 0x43,
	0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69,
	0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e,
	0x75, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e,
	0x75, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x22, 0x95, 0x04, 0x0a, 0x04, 0x43, 0x61,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x6b, 0x75, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x4e, 0x75, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x4e, 0x75, 0x6d, 0x73, 0x12, 0x2f, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x2e, 0x53, 0x6b, 0x75,
	0x52, 0x03, 0x73, 0x6b, 0x75, 0x1a, 0xc9, 0x02, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x6b, 0x75, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x6b,
	0x75, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6b, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x46, 0x65,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x46, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x65, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x73, 0x65, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x32, 0xc5, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x27, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x69, 0x6e, 0x67,
	0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x65, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x6e, 0x64, 0x12, 0x29, 0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72,
	0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x6f,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x2e,
	0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x69,
	0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x70, 0x69,
	0x2f, 0x62, 0x69, 0x6e, 0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_bingfood_cart_service_proto_rawDescOnce sync.Once
	file_v1_bingfood_cart_service_proto_rawDescData = file_v1_bingfood_cart_service_proto_rawDesc
)

func file_v1_bingfood_cart_service_proto_rawDescGZIP() []byte {
	file_v1_bingfood_cart_service_proto_rawDescOnce.Do(func() {
		file_v1_bingfood_cart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_bingfood_cart_service_proto_rawDescData)
	})
	return file_v1_bingfood_cart_service_proto_rawDescData
}

var file_v1_bingfood_cart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_bingfood_cart_service_proto_goTypes = []interface{}{
	(*PageInfo)(nil),                // 0: bingfood.service.v1.PageInfo
	(*GetCartByCondRequest)(nil),    // 1: bingfood.service.v1.GetCartByCondRequest
	(*GetCartByCartIdsRequest)(nil), // 2: bingfood.service.v1.GetCartByCartIdsRequest
	(*GetCartByCartIdsReply)(nil),   // 3: bingfood.service.v1.GetCartByCartIdsReply
	(*GetCartByCondReply)(nil),      // 4: bingfood.service.v1.GetCartByCondReply
	(*CartPagination)(nil),          // 5: bingfood.service.v1.CartPagination
	(*AddCartItemRequest)(nil),      // 6: bingfood.service.v1.AddCartItemRequest
	(*AddCartItemReply)(nil),        // 7: bingfood.service.v1.AddCartItemReply
	(*Cart)(nil),                    // 8: bingfood.service.v1.Cart
	(*Cart_Sku)(nil),                // 9: bingfood.service.v1.Cart.Sku
}
var file_v1_bingfood_cart_service_proto_depIdxs = []int32{
	8,  // 0: bingfood.service.v1.GetCartByCondRequest.cartCond:type_name -> bingfood.service.v1.Cart
	0,  // 1: bingfood.service.v1.GetCartByCondRequest.pageInfo:type_name -> bingfood.service.v1.PageInfo
	0,  // 2: bingfood.service.v1.GetCartByCartIdsRequest.pageInfo:type_name -> bingfood.service.v1.PageInfo
	5,  // 3: bingfood.service.v1.GetCartByCartIdsReply.data:type_name -> bingfood.service.v1.CartPagination
	5,  // 4: bingfood.service.v1.GetCartByCondReply.data:type_name -> bingfood.service.v1.CartPagination
	8,  // 5: bingfood.service.v1.CartPagination.list:type_name -> bingfood.service.v1.Cart
	9,  // 6: bingfood.service.v1.Cart.sku:type_name -> bingfood.service.v1.Cart.Sku
	6,  // 7: bingfood.service.v1.CartService.AddCartItem:input_type -> bingfood.service.v1.AddCartItemRequest
	1,  // 8: bingfood.service.v1.CartService.GetCartByCond:input_type -> bingfood.service.v1.GetCartByCondRequest
	2,  // 9: bingfood.service.v1.CartService.GetCartByCartIds:input_type -> bingfood.service.v1.GetCartByCartIdsRequest
	7,  // 10: bingfood.service.v1.CartService.AddCartItem:output_type -> bingfood.service.v1.AddCartItemReply
	4,  // 11: bingfood.service.v1.CartService.GetCartByCond:output_type -> bingfood.service.v1.GetCartByCondReply
	3,  // 12: bingfood.service.v1.CartService.GetCartByCartIds:output_type -> bingfood.service.v1.GetCartByCartIdsReply
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_bingfood_cart_service_proto_init() }
func file_v1_bingfood_cart_service_proto_init() {
	if File_v1_bingfood_cart_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_bingfood_cart_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartByCondRequest); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartByCartIdsRequest); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartByCartIdsReply); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartByCondReply); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartPagination); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartItemRequest); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCartItemReply); i {
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
		file_v1_bingfood_cart_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_bingfood_cart_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart_Sku); i {
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
			RawDescriptor: file_v1_bingfood_cart_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_bingfood_cart_service_proto_goTypes,
		DependencyIndexes: file_v1_bingfood_cart_service_proto_depIdxs,
		MessageInfos:      file_v1_bingfood_cart_service_proto_msgTypes,
	}.Build()
	File_v1_bingfood_cart_service_proto = out.File
	file_v1_bingfood_cart_service_proto_rawDesc = nil
	file_v1_bingfood_cart_service_proto_goTypes = nil
	file_v1_bingfood_cart_service_proto_depIdxs = nil
}