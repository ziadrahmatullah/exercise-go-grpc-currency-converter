// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/currency_converter.proto

package pb

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

type CurrencyConversionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount         float64 `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	SourceCurrency string  `protobuf:"bytes,2,opt,name=source_currency,json=sourceCurrency,proto3" json:"source_currency,omitempty"`
	TargetCurrency string  `protobuf:"bytes,3,opt,name=target_currency,json=targetCurrency,proto3" json:"target_currency,omitempty"`
}

func (x *CurrencyConversionRequest) Reset() {
	*x = CurrencyConversionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyConversionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyConversionRequest) ProtoMessage() {}

func (x *CurrencyConversionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_converter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyConversionRequest.ProtoReflect.Descriptor instead.
func (*CurrencyConversionRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_converter_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyConversionRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CurrencyConversionRequest) GetSourceCurrency() string {
	if x != nil {
		return x.SourceCurrency
	}
	return ""
}

func (x *CurrencyConversionRequest) GetTargetCurrency() string {
	if x != nil {
		return x.TargetCurrency
	}
	return ""
}

type CurrencyConversionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConvertedAmount float64 `protobuf:"fixed64,1,opt,name=converted_amount,json=convertedAmount,proto3" json:"converted_amount,omitempty"`
}

func (x *CurrencyConversionResponse) Reset() {
	*x = CurrencyConversionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currency_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyConversionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyConversionResponse) ProtoMessage() {}

func (x *CurrencyConversionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_converter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyConversionResponse.ProtoReflect.Descriptor instead.
func (*CurrencyConversionResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_converter_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyConversionResponse) GetConvertedAmount() float64 {
	if x != nil {
		return x.ConvertedAmount
	}
	return 0
}

var File_proto_currency_converter_proto protoreflect.FileDescriptor

var file_proto_currency_converter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x22, 0x85, 0x01, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x47, 0x0a, 0x1a,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x85, 0x01, 0x0a, 0x11, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x70, 0x0a, 0x0f, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x2d,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_converter_proto_rawDescOnce sync.Once
	file_proto_currency_converter_proto_rawDescData = file_proto_currency_converter_proto_rawDesc
)

func file_proto_currency_converter_proto_rawDescGZIP() []byte {
	file_proto_currency_converter_proto_rawDescOnce.Do(func() {
		file_proto_currency_converter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_converter_proto_rawDescData)
	})
	return file_proto_currency_converter_proto_rawDescData
}

var file_proto_currency_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_currency_converter_proto_goTypes = []interface{}{
	(*CurrencyConversionRequest)(nil),  // 0: currency_converter.CurrencyConversionRequest
	(*CurrencyConversionResponse)(nil), // 1: currency_converter.CurrencyConversionResponse
}
var file_proto_currency_converter_proto_depIdxs = []int32{
	0, // 0: currency_converter.CurrencyConverter.ConvertCurrency:input_type -> currency_converter.CurrencyConversionRequest
	1, // 1: currency_converter.CurrencyConverter.ConvertCurrency:output_type -> currency_converter.CurrencyConversionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_currency_converter_proto_init() }
func file_proto_currency_converter_proto_init() {
	if File_proto_currency_converter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currency_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyConversionRequest); i {
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
		file_proto_currency_converter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyConversionResponse); i {
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
			RawDescriptor: file_proto_currency_converter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_converter_proto_goTypes,
		DependencyIndexes: file_proto_currency_converter_proto_depIdxs,
		MessageInfos:      file_proto_currency_converter_proto_msgTypes,
	}.Build()
	File_proto_currency_converter_proto = out.File
	file_proto_currency_converter_proto_rawDesc = nil
	file_proto_currency_converter_proto_goTypes = nil
	file_proto_currency_converter_proto_depIdxs = nil
}