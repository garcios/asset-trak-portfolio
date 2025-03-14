// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: proto/asset_price.proto

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

// Request message for getting the price of an asset
type GetAssetPriceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`       // Identifier of the asset
	TradeDate     string                 `protobuf:"bytes,2,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"` // Trade date in ISO 8601 format (e.g., "2023-01-01")
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetPriceRequest) Reset() {
	*x = GetAssetPriceRequest{}
	mi := &file_proto_asset_price_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetPriceRequest) ProtoMessage() {}

func (x *GetAssetPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asset_price_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetPriceRequest.ProtoReflect.Descriptor instead.
func (*GetAssetPriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_asset_price_proto_rawDescGZIP(), []int{0}
}

func (x *GetAssetPriceRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *GetAssetPriceRequest) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

// Response message containing the price of the asset
type GetAssetPriceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`       // Identifier of the asset
	Price         float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`                        // Current price of the asset
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`                    // Currency of the price (e.g., USD, EUR)
	TradeDate     string                 `protobuf:"bytes,4,opt,name=trade_date,json=tradeDate,proto3" json:"trade_date,omitempty"` // Trade date in ISO 8601 format (e.g., "2023-01-01")
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetPriceResponse) Reset() {
	*x = GetAssetPriceResponse{}
	mi := &file_proto_asset_price_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetPriceResponse) ProtoMessage() {}

func (x *GetAssetPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asset_price_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetPriceResponse.ProtoReflect.Descriptor instead.
func (*GetAssetPriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_asset_price_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetPriceResponse) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *GetAssetPriceResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetAssetPriceResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetAssetPriceResponse) GetTradeDate() string {
	if x != nil {
		return x.TradeDate
	}
	return ""
}

// Request message for getting asset prices by date range
type GetAssetPriceHistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`       // Identifier of the asset
	StartDate     string                 `protobuf:"bytes,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"` // Start date in ISO 8601 format (e.g., "2023-01-01")
	EndDate       string                 `protobuf:"bytes,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`       // End date in ISO 8601 format (e.g., "2023-12-31")
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetPriceHistoryRequest) Reset() {
	*x = GetAssetPriceHistoryRequest{}
	mi := &file_proto_asset_price_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetPriceHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetPriceHistoryRequest) ProtoMessage() {}

func (x *GetAssetPriceHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asset_price_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetPriceHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetAssetPriceHistoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_asset_price_proto_rawDescGZIP(), []int{2}
}

func (x *GetAssetPriceHistoryRequest) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *GetAssetPriceHistoryRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *GetAssetPriceHistoryRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

// Response message containing a list of asset prices within the date range
type GetAssetPriceHistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetId       string                 `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"` // Identifier of the asset
	Prices        []*AssetPriceEntry     `protobuf:"bytes,2,rep,name=prices,proto3" json:"prices,omitempty"`                  // List of asset prices in the date range
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssetPriceHistoryResponse) Reset() {
	*x = GetAssetPriceHistoryResponse{}
	mi := &file_proto_asset_price_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssetPriceHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetPriceHistoryResponse) ProtoMessage() {}

func (x *GetAssetPriceHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asset_price_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetPriceHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetAssetPriceHistoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_asset_price_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetPriceHistoryResponse) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *GetAssetPriceHistoryResponse) GetPrices() []*AssetPriceEntry {
	if x != nil {
		return x.Prices
	}
	return nil
}

// Represents a price entry with the timestamp
type AssetPriceEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Date          string                 `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`         // Date of the price (in ISO 8601 format)
	Price         float64                `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`     // Price of the asset on the specified date
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"` // Currency of the price (e.g., USD, EUR)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetPriceEntry) Reset() {
	*x = AssetPriceEntry{}
	mi := &file_proto_asset_price_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetPriceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetPriceEntry) ProtoMessage() {}

func (x *AssetPriceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_proto_asset_price_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetPriceEntry.ProtoReflect.Descriptor instead.
func (*AssetPriceEntry) Descriptor() ([]byte, []int) {
	return file_proto_asset_price_proto_rawDescGZIP(), []int{4}
}

func (x *AssetPriceEntry) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *AssetPriceEntry) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AssetPriceEntry) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_proto_asset_price_proto protoreflect.FileDescriptor

var file_proto_asset_price_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x64, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x72, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x63, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x70, 0x72, 0x69, 0x63, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x0f, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x32, 0xa1, 0x01, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_asset_price_proto_rawDescOnce sync.Once
	file_proto_asset_price_proto_rawDescData = file_proto_asset_price_proto_rawDesc
)

func file_proto_asset_price_proto_rawDescGZIP() []byte {
	file_proto_asset_price_proto_rawDescOnce.Do(func() {
		file_proto_asset_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_asset_price_proto_rawDescData)
	})
	return file_proto_asset_price_proto_rawDescData
}

var file_proto_asset_price_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_asset_price_proto_goTypes = []any{
	(*GetAssetPriceRequest)(nil),         // 0: GetAssetPriceRequest
	(*GetAssetPriceResponse)(nil),        // 1: GetAssetPriceResponse
	(*GetAssetPriceHistoryRequest)(nil),  // 2: GetAssetPriceHistoryRequest
	(*GetAssetPriceHistoryResponse)(nil), // 3: GetAssetPriceHistoryResponse
	(*AssetPriceEntry)(nil),              // 4: AssetPriceEntry
}
var file_proto_asset_price_proto_depIdxs = []int32{
	4, // 0: GetAssetPriceHistoryResponse.prices:type_name -> AssetPriceEntry
	0, // 1: AssetPrice.GetAssetPrice:input_type -> GetAssetPriceRequest
	2, // 2: AssetPrice.GetAssetPriceHistory:input_type -> GetAssetPriceHistoryRequest
	1, // 3: AssetPrice.GetAssetPrice:output_type -> GetAssetPriceResponse
	3, // 4: AssetPrice.GetAssetPriceHistory:output_type -> GetAssetPriceHistoryResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_asset_price_proto_init() }
func file_proto_asset_price_proto_init() {
	if File_proto_asset_price_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_asset_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_asset_price_proto_goTypes,
		DependencyIndexes: file_proto_asset_price_proto_depIdxs,
		MessageInfos:      file_proto_asset_price_proto_msgTypes,
	}.Build()
	File_proto_asset_price_proto = out.File
	file_proto_asset_price_proto_rawDesc = nil
	file_proto_asset_price_proto_goTypes = nil
	file_proto_asset_price_proto_depIdxs = nil
}
