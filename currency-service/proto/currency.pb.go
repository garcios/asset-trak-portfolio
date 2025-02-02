// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: proto/currency.proto

package proto

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

// The request message to get the exchange rate between two currencies
type GetExchangeRateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"` // Origin currency code
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`       // Target currency code
	TradeDate     string                 `protobuf:"bytes,3,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`          // The date for which exchange rate is requested in the format 'yyyy-mm-dd'
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetExchangeRateRequest) Reset() {
	*x = GetExchangeRateRequest{}
	mi := &file_proto_currency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExchangeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateRequest) ProtoMessage() {}

func (x *GetExchangeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateRequest.ProtoReflect.Descriptor instead.
func (*GetExchangeRateRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{0}
}

func (x *GetExchangeRateRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *GetExchangeRateRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *GetExchangeRateRequest) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

// The response message containing the exchange rate
type GetExchangeRateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExchangeRate  float64                `protobuf:"fixed64,1,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"` // Exchange rate from origin to target currency
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetExchangeRateResponse) Reset() {
	*x = GetExchangeRateResponse{}
	mi := &file_proto_currency_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetExchangeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExchangeRateResponse) ProtoMessage() {}

func (x *GetExchangeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExchangeRateResponse.ProtoReflect.Descriptor instead.
func (*GetExchangeRateResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{1}
}

func (x *GetExchangeRateResponse) GetExchangeRate() float64 {
	if x != nil {
		return x.ExchangeRate
	}
	return 0
}

// The request message to get historical exchange rates between two currencies
type GetHistoricalExchangeRatesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FromCurrency  string                 `protobuf:"bytes,1,opt,name=from_currency,json=fromCurrency,proto3" json:"from_currency,omitempty"` // Origin currency code
	ToCurrency    string                 `protobuf:"bytes,2,opt,name=to_currency,json=toCurrency,proto3" json:"to_currency,omitempty"`       // Target currency code
	StartDate     string                 `protobuf:"bytes,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`          // The start date of the range for which historical exchange rates are requested in the format 'yyyy-mm-dd'
	EndDate       string                 `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`                // The end date of the range for which historical exchange rates are requested in the format 'yyyy-mm-dd'
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHistoricalExchangeRatesRequest) Reset() {
	*x = GetHistoricalExchangeRatesRequest{}
	mi := &file_proto_currency_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoricalExchangeRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoricalExchangeRatesRequest) ProtoMessage() {}

func (x *GetHistoricalExchangeRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoricalExchangeRatesRequest.ProtoReflect.Descriptor instead.
func (*GetHistoricalExchangeRatesRequest) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{2}
}

func (x *GetHistoricalExchangeRatesRequest) GetFromCurrency() string {
	if x != nil {
		return x.FromCurrency
	}
	return ""
}

func (x *GetHistoricalExchangeRatesRequest) GetToCurrency() string {
	if x != nil {
		return x.ToCurrency
	}
	return ""
}

func (x *GetHistoricalExchangeRatesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetHistoricalExchangeRatesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// Object to hold a date and the exchange rate for that date
type HistoricalRate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeDate     string                 `protobuf:"bytes,1,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"`            // Date in 'yyyy-mm-dd' format
	ExchangeRate  float64                `protobuf:"fixed64,2,opt,name=exchange_rate,json=exchangeRate,proto3" json:"exchange_rate,omitempty"` // Exchange rate for the given date
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HistoricalRate) Reset() {
	*x = HistoricalRate{}
	mi := &file_proto_currency_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HistoricalRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoricalRate) ProtoMessage() {}

func (x *HistoricalRate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoricalRate.ProtoReflect.Descriptor instead.
func (*HistoricalRate) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{3}
}

func (x *HistoricalRate) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

func (x *HistoricalRate) GetExchangeRate() float64 {
	if x != nil {
		return x.ExchangeRate
	}
	return 0
}

// The response message containing a list of historical exchange rates
type GetHistoricalExchangeRatesResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	HistoricalRates []*HistoricalRate      `protobuf:"bytes,1,rep,name=historical_rates,json=historicalRates,proto3" json:"historical_rates,omitempty"` // List of historical exchange rates for the requested date range
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetHistoricalExchangeRatesResponse) Reset() {
	*x = GetHistoricalExchangeRatesResponse{}
	mi := &file_proto_currency_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHistoricalExchangeRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoricalExchangeRatesResponse) ProtoMessage() {}

func (x *GetHistoricalExchangeRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currency_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoricalExchangeRatesResponse.ProtoReflect.Descriptor instead.
func (*GetHistoricalExchangeRatesResponse) Descriptor() ([]byte, []int) {
	return file_proto_currency_proto_rawDescGZIP(), []int{4}
}

func (x *GetHistoricalExchangeRatesResponse) GetHistoricalRates() []*HistoricalRate {
	if x != nil {
		return x.HistoricalRates
	}
	return nil
}

var File_proto_currency_proto protoreflect.FileDescriptor

var file_proto_currency_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x54, 0x0a, 0x0e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0c, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x22, 0x60, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63,
	0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x32, 0xbb, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x63, 0x61, 0x6c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currency_proto_rawDescOnce sync.Once
	file_proto_currency_proto_rawDescData = file_proto_currency_proto_rawDesc
)

func file_proto_currency_proto_rawDescGZIP() []byte {
	file_proto_currency_proto_rawDescOnce.Do(func() {
		file_proto_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currency_proto_rawDescData)
	})
	return file_proto_currency_proto_rawDescData
}

var file_proto_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_currency_proto_goTypes = []any{
	(*GetExchangeRateRequest)(nil),             // 0: GetExchangeRateRequest
	(*GetExchangeRateResponse)(nil),            // 1: GetExchangeRateResponse
	(*GetHistoricalExchangeRatesRequest)(nil),  // 2: GetHistoricalExchangeRatesRequest
	(*HistoricalRate)(nil),                     // 3: HistoricalRate
	(*GetHistoricalExchangeRatesResponse)(nil), // 4: GetHistoricalExchangeRatesResponse
}
var file_proto_currency_proto_depIdxs = []int32{
	3, // 0: GetHistoricalExchangeRatesResponse.historical_rates:type_name -> HistoricalRate
	0, // 1: Currency.GetExchangeRate:input_type -> GetExchangeRateRequest
	2, // 2: Currency.GetHistoricalExchangeRates:input_type -> GetHistoricalExchangeRatesRequest
	1, // 3: Currency.GetExchangeRate:output_type -> GetExchangeRateResponse
	4, // 4: Currency.GetHistoricalExchangeRates:output_type -> GetHistoricalExchangeRatesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_currency_proto_init() }
func file_proto_currency_proto_init() {
	if File_proto_currency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currency_proto_goTypes,
		DependencyIndexes: file_proto_currency_proto_depIdxs,
		MessageInfos:      file_proto_currency_proto_msgTypes,
	}.Build()
	File_proto_currency_proto = out.File
	file_proto_currency_proto_rawDesc = nil
	file_proto_currency_proto_goTypes = nil
	file_proto_currency_proto_depIdxs = nil
}
