// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/asset_price.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for AssetPrice service

func NewAssetPriceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AssetPrice service

type AssetPriceService interface {
	// Get the price of a specific asset by ID
	GetAssetPrice(ctx context.Context, in *GetAssetPriceRequest, opts ...client.CallOption) (*GetAssetPriceResponse, error)
	// Get the prices of an asset by ID and a date range
	GetAssetPriceHistory(ctx context.Context, in *GetAssetPriceHistoryRequest, opts ...client.CallOption) (*GetAssetPriceHistoryResponse, error)
}

type assetPriceService struct {
	c    client.Client
	name string
}

func NewAssetPriceService(name string, c client.Client) AssetPriceService {
	return &assetPriceService{
		c:    c,
		name: name,
	}
}

func (c *assetPriceService) GetAssetPrice(ctx context.Context, in *GetAssetPriceRequest, opts ...client.CallOption) (*GetAssetPriceResponse, error) {
	req := c.c.NewRequest(c.name, "AssetPrice.GetAssetPrice", in)
	out := new(GetAssetPriceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetPriceService) GetAssetPriceHistory(ctx context.Context, in *GetAssetPriceHistoryRequest, opts ...client.CallOption) (*GetAssetPriceHistoryResponse, error) {
	req := c.c.NewRequest(c.name, "AssetPrice.GetAssetPriceHistory", in)
	out := new(GetAssetPriceHistoryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AssetPrice service

type AssetPriceHandler interface {
	// Get the price of a specific asset by ID
	GetAssetPrice(context.Context, *GetAssetPriceRequest, *GetAssetPriceResponse) error
	// Get the prices of an asset by ID and a date range
	GetAssetPriceHistory(context.Context, *GetAssetPriceHistoryRequest, *GetAssetPriceHistoryResponse) error
}

func RegisterAssetPriceHandler(s server.Server, hdlr AssetPriceHandler, opts ...server.HandlerOption) error {
	type assetPrice interface {
		GetAssetPrice(ctx context.Context, in *GetAssetPriceRequest, out *GetAssetPriceResponse) error
		GetAssetPriceHistory(ctx context.Context, in *GetAssetPriceHistoryRequest, out *GetAssetPriceHistoryResponse) error
	}
	type AssetPrice struct {
		assetPrice
	}
	h := &assetPriceHandler{hdlr}
	return s.Handle(s.NewHandler(&AssetPrice{h}, opts...))
}

type assetPriceHandler struct {
	AssetPriceHandler
}

func (h *assetPriceHandler) GetAssetPrice(ctx context.Context, in *GetAssetPriceRequest, out *GetAssetPriceResponse) error {
	return h.AssetPriceHandler.GetAssetPrice(ctx, in, out)
}

func (h *assetPriceHandler) GetAssetPriceHistory(ctx context.Context, in *GetAssetPriceHistoryRequest, out *GetAssetPriceHistoryResponse) error {
	return h.AssetPriceHandler.GetAssetPriceHistory(ctx, in, out)
}
