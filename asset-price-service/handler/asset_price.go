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
	res.TradeDate = req.TradeDate

	return nil
}

func (a *AssetPrice) GetAssetPriceHistory(ctx context.Context, req *pba.GetAssetPriceHistoryRequest, res *pba.GetAssetPriceHistoryResponse) error {
	ap, err := a.repo.GetAssetPriceHistory(req.AssetId, req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	if ap == nil {
		return fmt.Errorf("asset prices not found for assetId: %s, tradeDates: %s-%s",
			req.AssetId,
			req.StartDate,
			req.EndDate,
		)
	}

	assetPrices := make([]*pba.AssetPriceEntry, 0, len(ap))

	for _, ap := range ap {
		apEntry := &pba.AssetPriceEntry{
			Date:     ap.TradeDate.Format("2006-01-02"),
			Price:    ap.Price,
			Currency: ap.CurrencyCode,
		}
		assetPrices = append(assetPrices, apEntry)
	}

	res.AssetId = req.AssetId
	res.Prices = assetPrices

	return nil
}
