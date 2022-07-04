// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: v1/cart_error_reason.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type CartErrorReason int32

const (
	// 为某个枚举单独设置错误码
	CartErrorReason_INVALID_ARGUMENT CartErrorReason = 0
	//  USER_NOT_FOUND = 1 [(errors.code) = 404];
	CartErrorReason_UNAUTHENTICATED CartErrorReason = 2 //  PASSWORD_FALSE = 3 [(errors.code) = 405];
	CartErrorReason_INTERNAL        CartErrorReason = 4 // 服务器内部错误
)

// Enum value maps for CartErrorReason.
var (
	CartErrorReason_name = map[int32]string{
		0: "INVALID_ARGUMENT",
		2: "UNAUTHENTICATED",
		4: "INTERNAL",
	}
	CartErrorReason_value = map[string]int32{
		"INVALID_ARGUMENT": 0,
		"UNAUTHENTICATED":  2,
		"INTERNAL":         4,
	}
)

func (x CartErrorReason) Enum() *CartErrorReason {
	p := new(CartErrorReason)
	*p = x
	return p
}

func (x CartErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CartErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_cart_error_reason_proto_enumTypes[0].Descriptor()
}

func (CartErrorReason) Type() protoreflect.EnumType {
	return &file_v1_cart_error_reason_proto_enumTypes[0]
}

func (x CartErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CartErrorReason.Descriptor instead.
func (CartErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_v1_cart_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_v1_cart_error_reason_proto protoreflect.FileDescriptor

var file_v1_cart_error_reason_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x61,
	0x72, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x62, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x90,
	0x03, 0x12, 0x19, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x12, 0x12, 0x0a, 0x08,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x14, 0x5a, 0x12, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cart_error_reason_proto_rawDescOnce sync.Once
	file_v1_cart_error_reason_proto_rawDescData = file_v1_cart_error_reason_proto_rawDesc
)

func file_v1_cart_error_reason_proto_rawDescGZIP() []byte {
	file_v1_cart_error_reason_proto_rawDescOnce.Do(func() {
		file_v1_cart_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cart_error_reason_proto_rawDescData)
	})
	return file_v1_cart_error_reason_proto_rawDescData
}

var file_v1_cart_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_cart_error_reason_proto_goTypes = []interface{}{
	(CartErrorReason)(0), // 0: cart.service.v1.CartErrorReason
}
var file_v1_cart_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_cart_error_reason_proto_init() }
func file_v1_cart_error_reason_proto_init() {
	if File_v1_cart_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_cart_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_cart_error_reason_proto_goTypes,
		DependencyIndexes: file_v1_cart_error_reason_proto_depIdxs,
		EnumInfos:         file_v1_cart_error_reason_proto_enumTypes,
	}.Build()
	File_v1_cart_error_reason_proto = out.File
	file_v1_cart_error_reason_proto_rawDesc = nil
	file_v1_cart_error_reason_proto_goTypes = nil
	file_v1_cart_error_reason_proto_depIdxs = nil
}
