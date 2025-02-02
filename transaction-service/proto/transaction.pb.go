// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: proto/transaction.proto

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

type BalanceSummaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalanceSummaryRequest) Reset() {
	*x = BalanceSummaryRequest{}
	mi := &file_proto_transaction_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceSummaryRequest) ProtoMessage() {}

func (x *BalanceSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceSummaryRequest.ProtoReflect.Descriptor instead.
func (*BalanceSummaryRequest) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceSummaryRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type BalanceSummaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountId     string                 `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	BalanceItems  []*BalanceItem         `protobuf:"bytes,2,rep,name=balance_items,json=balanceItems,proto3" json:"balance_items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalanceSummaryResponse) Reset() {
	*x = BalanceSummaryResponse{}
	mi := &file_proto_transaction_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceSummaryResponse) ProtoMessage() {}

func (x *BalanceSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceSummaryResponse.ProtoReflect.Descriptor instead.
func (*BalanceSummaryResponse) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceSummaryResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *BalanceSummaryResponse) GetBalanceItems() []*BalanceItem {
	if x != nil {
		return x.BalanceItems
	}
	return nil
}

type BalanceItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AssetSymbol   string                 `protobuf:"bytes,1,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	AssetName     string                 `protobuf:"bytes,2,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	Price         *Money                 `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity      float64                `protobuf:"fixed64,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Value         *Money                 `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	TotalGain     float64                `protobuf:"fixed64,6,opt,name=total_gain,json=totalGain,proto3" json:"total_gain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BalanceItem) Reset() {
	*x = BalanceItem{}
	mi := &file_proto_transaction_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BalanceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceItem) ProtoMessage() {}

func (x *BalanceItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceItem.ProtoReflect.Descriptor instead.
func (*BalanceItem) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceItem) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *BalanceItem) GetAssetName() string {
	if x != nil {
		return x.AssetName
	}
	return ""
}

func (x *BalanceItem) GetPrice() *Money {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *BalanceItem) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *BalanceItem) GetValue() *Money {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *BalanceItem) GetTotalGain() float64 {
	if x != nil {
		return x.TotalGain
	}
	return 0
}

type Money struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Amount        float64                `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	CurrencyCode  string                 `protobuf:"bytes,2,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Money) Reset() {
	*x = Money{}
	mi := &file_proto_transaction_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *Money) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

var File_proto_transaction_proto protoreflect.FileDescriptor

var file_proto_transaction_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x15, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x6a, 0x0a, 0x16, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc6, 0x01,
	0x0a, 0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06,
	0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x47, 0x61, 0x69, 0x6e, 0x22, 0x44, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x55, 0x0a, 0x0b,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x16, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_transaction_proto_rawDescOnce sync.Once
	file_proto_transaction_proto_rawDescData = file_proto_transaction_proto_rawDesc
)

func file_proto_transaction_proto_rawDescGZIP() []byte {
	file_proto_transaction_proto_rawDescOnce.Do(func() {
		file_proto_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transaction_proto_rawDescData)
	})
	return file_proto_transaction_proto_rawDescData
}

var file_proto_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_transaction_proto_goTypes = []any{
	(*BalanceSummaryRequest)(nil),  // 0: BalanceSummaryRequest
	(*BalanceSummaryResponse)(nil), // 1: BalanceSummaryResponse
	(*BalanceItem)(nil),            // 2: BalanceItem
	(*Money)(nil),                  // 3: Money
}
var file_proto_transaction_proto_depIdxs = []int32{
	2, // 0: BalanceSummaryResponse.balance_items:type_name -> BalanceItem
	3, // 1: BalanceItem.price:type_name -> Money
	3, // 2: BalanceItem.value:type_name -> Money
	0, // 3: Transaction.GetBalanceSummary:input_type -> BalanceSummaryRequest
	1, // 4: Transaction.GetBalanceSummary:output_type -> BalanceSummaryResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_transaction_proto_init() }
func file_proto_transaction_proto_init() {
	if File_proto_transaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_transaction_proto_goTypes,
		DependencyIndexes: file_proto_transaction_proto_depIdxs,
		MessageInfos:      file_proto_transaction_proto_msgTypes,
	}.Build()
	File_proto_transaction_proto = out.File
	file_proto_transaction_proto_rawDesc = nil
	file_proto_transaction_proto_goTypes = nil
	file_proto_transaction_proto_depIdxs = nil
}
