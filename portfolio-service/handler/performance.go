package handler

import (
	"context"
	"fmt"
	"log"
	"time"

	apm "github.com/garcios/asset-trak-portfolio/asset-price-service/model"
	pba "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	"github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/service"
)

func (h *Transaction) GetPerformanceHistory(
	ctx context.Context,
	req *pb.PerformanceHistoryRequest,
	res *pb.PerformanceHistoryResponse,
) error {

	txns, err := h.transactionRepository.GetTransactions(ctx, db.TransactionFilter{
		AccountID: req.AccountId,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	})

	if err != nil {
		return err
	}

	trades := h.toTrades(txns)
	log.Printf("len(trades): %d\n", len(trades))

	holdings, err := h.portfolioRepository.GetHoldingAtDateRange(ctx, req.AccountId, req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	log.Printf("len(holdings): %d\n", len(holdings))

	var (
		startDate time.Time
		endDate   time.Time
	)

	dateFormat := "2006-01-02"

	fmt.Printf("req.StartDate: %s\n", req.StartDate)
	startDate, err = time.Parse(dateFormat, req.StartDate)
	if err != nil {
		return err
	}

	endDate, err = time.Parse(dateFormat, req.EndDate)
	if err != nil {
		return err
	}

	recs, err := h.performanceService.CalculateDailyHistoricalValueAndCost(
		ctx,
		trades,
		startDate,
		endDate,
		targetCurrency,
		h.exchangeRateFn(ctx, dateFormat),
		h.assetPriceFn(ctx, dateFormat),
	)

	if err != nil {
		return err
	}

	// populate response
	h.toHistoricalRecords(recs, res)

	return nil
}

func (h *Transaction) assetPriceFn(ctx context.Context, dateFormat string) func(assetID string, date time.Time) (*apm.AssetPrice, error) {
	getAssetPrice := func(assetID string, date time.Time) (*apm.AssetPrice, error) {
		log.Printf("get asset price for assetID: %s, date: %s\n", assetID, date.Format(dateFormat))
		res, err := h.assetPriceService.GetAssetPrice(ctx, &pba.GetAssetPriceRequest{
			AssetId:   assetID,
			TradeDate: date.Format(dateFormat),
		})

		if err != nil {
			return nil, err
		}

		var parsedDate time.Time
		if res.TradeDate != "" {
			parsedDate, err = time.Parse(dateFormat, res.TradeDate)
			if err != nil {
				return nil, err
			}
		}

		return &apm.AssetPrice{
			AssetID:      res.AssetId,
			Price:        res.Price,
			CurrencyCode: res.Currency,
			TradeDate:    &parsedDate,
		}, nil
	}

	return getAssetPrice
}

func (h *Transaction) exchangeRateFn(ctx context.Context, dateFormat string) func(fromCurrency string, toCurrency string, date time.Time) (float64, error) {
	getExchangeRate := func(fromCurrency string, toCurrency string, date time.Time) (float64, error) {
		if fromCurrency == toCurrency {
			return 1.0, nil
		}

		log.Printf("get exchange rate for fromCurrency: %s, toCurrency: %s, trade date: %s\n",
			fromCurrency,
			toCurrency,
			date.Format(dateFormat))

		res, err := h.currencyService.GetExchangeRate(ctx, &proto.GetExchangeRateRequest{
			FromCurrency: fromCurrency,
			ToCurrency:   toCurrency,
			TradeDate:    date.Format(dateFormat),
		})

		if err != nil {
			return 0, err
		}

		return res.ExchangeRate, nil
	}
	return getExchangeRate
}

func (h *Transaction) toHistoricalRecords(recs []*service.HistoricalRecord, res *pb.PerformanceHistoryResponse) {
	protoRecords := make([]*pb.HistoricalRecord, len(recs))
	for i, rec := range recs {
		protoRecords[i] = convertToProtoHistoricalRecord(rec)
	}

	res.Records = protoRecords
}

func (h *Transaction) toTrades(txns []*model.Transaction) []*finance.Trade {
	trades := make([]*finance.Trade, len(txns))

	for _, txn := range txns {
		trades = append(trades, &finance.Trade{
			AssetID:  txn.AssetID,
			Quantity: txn.Quantity,
			Price: finance.Money{
				Amount:       txn.TradePrice,
				CurrencyCode: txn.AmountCurrencyCode,
			},
			Commission: finance.Money{
				Amount:       txn.BrokerageFee,
				CurrencyCode: txn.FeeCurrencyCode,
			},
			TradeType:    txn.TransactionType,
			CurrencyRate: txn.ExchangeRate,
			AmountCash: finance.Money{
				Amount:       txn.AmountCash,
				CurrencyCode: txn.AmountCurrencyCode,
			},
			TransactionDate: *txn.TransactionDate,
		})
	}

	return trades
}

func convertToProtoHistoricalRecord(record *service.HistoricalRecord) *pb.HistoricalRecord {
	return &pb.HistoricalRecord{
		TradeDate:    record.Date.Format("2006-01-02"),
		Value:        record.Value,
		Cost:         record.Cost,
		CurrencyCode: targetCurrency,
	}
}
