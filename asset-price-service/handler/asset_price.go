package handler

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/asset-price-service/db"
	pba "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	"time"
)

type AssetPrice struct {
	repo *db.AssetPriceRepository
}

func New(repo *db.AssetPriceRepository) *AssetPrice {
	return &AssetPrice{
		repo: repo,
	}
}

func (a *AssetPrice) GetAssetPrice(ctx context.Context, req *pba.GetAssetPriceRequest, res *pba.GetAssetPriceResponse) error {
	tradeDate, err := time.Parse("2006-01-02", req.TradeDate)
	if err != nil {
		return err
	}

	ap, err := a.repo.GetAssetPrice(req.AssetId, tradeDate)
	if err != nil {
		return err
	}

	if ap == nil {
		return fmt.Errorf("asset price not found for assetId: %s, tradeDate: %s",
			req.AssetId,
			req.TradeDate)
	}

	res.AssetId = req.AssetId
	res.Price = ap.Price
	res.Currency = ap.CurrencyCode

	return nil
}

func (a *AssetPrice) GetAssetPricesByDateRange(ctx context.Context, req *pba.GetAssetPricesByDateRangeRequest, res *pba.GetAssetPricesByDateRangeResponse) error {
	// TODO: implement this
	return nil
}
