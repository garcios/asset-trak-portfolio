package handler

import (
	"context"
	"log"
	"sync"
	"time"

	ap "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	cr "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/service"

	con "github.com/garcios/asset-trak-portfolio/lib/concurrency"
)

const (
	maxConcurrency = 20
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

	transactions := h.toTransactions(txns)
	log.Printf("len(trades): %d\n", len(transactions))

	var (
		startDate time.Time
		endDate   time.Time
	)

	dateFormat := "2006-01-02"

	startDate, err = time.Parse(dateFormat, req.StartDate)
	if err != nil {
		return err
	}

	endDate, err = time.Parse(dateFormat, req.EndDate)
	if err != nil {
		return err
	}

	dateRange := service.DateRange{
		StartDate: startDate,
		EndDate:   endDate,
	}

	assetIds := collectUniqueAssetIds(transactions)
	assetPrices, err := h.getAssetPrices(ctx, assetIds, req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	currencies := collectUniqueCurrencies(transactions)
	currencyRates, err := h.getCurrencyRates(ctx, currencies, req.StartDate, req.EndDate)
	if err != nil {
		return err
	}

	marketData := service.MarketData{
		AssetPrices:   assetPrices,
		CurrencyRates: currencyRates,
	}

	recs, err := h.performanceService.CalculateDailyHistoricalValueAndCost(
		ctx,
		transactions,
		marketData,
		targetCurrency,
		dateRange,
	)

	if err != nil {
		return err
	}

	// populate response
	h.toHistoricalRecords(recs, res)

	return nil
}

func collectUniqueCurrencies(transactions []*service.TransactionRecord) []string {
	uniqueCurrencies := make(map[string]struct{}, 0)

	for _, txn := range transactions {
		if _, ok := uniqueCurrencies[txn.TradePriceCurrencyCode]; !ok {
			uniqueCurrencies[txn.TradePriceCurrencyCode] = struct{}{}
		}

		if _, ok := uniqueCurrencies[txn.BrokerageFeeCurrencyCode]; !ok {
			uniqueCurrencies[txn.BrokerageFeeCurrencyCode] = struct{}{}
		}
	}

	delete(uniqueCurrencies, targetCurrency)

	currencies := make([]string, 0, len(uniqueCurrencies))
	for currency := range uniqueCurrencies {
		currencies = append(currencies, currency)
	}

	return currencies
}

func collectUniqueAssetIds(transactions []*service.TransactionRecord) []string {
	uniqueAssetIds := make(map[string]struct{}, 0)

	for _, txn := range transactions {
		if _, ok := uniqueAssetIds[txn.AssetID]; !ok {
			uniqueAssetIds[txn.AssetID] = struct{}{}
		}
	}

	assetIds := make([]string, 0, len(uniqueAssetIds))
	for assetId := range uniqueAssetIds {
		assetIds = append(assetIds, assetId)
	}

	return assetIds
}

func (h *Transaction) getCurrencyRates(
	ctx context.Context,
	currencies []string,
	startDate string,
	endDate string,
) ([]*service.CurrencyRate, error) {
	currencyRates := make([]*service.CurrencyRate, 0)

	g, gctx := con.WithContext(ctx, maxConcurrency)
	var mutex sync.Mutex

	for _, currency := range currencies {
		g.Go(func() error {
			req := &cr.GetHistoricalExchangeRatesRequest{
				FromCurrency: currency,
				ToCurrency:   targetCurrency,
				StartDate:    startDate,
				EndDate:      endDate,
			}

			res, err := h.currencyService.GetHistoricalExchangeRates(gctx, req)
			if err != nil {
				return err
			}

			for _, r := range res.GetHistoricalRates() {
				parsedDate, err := time.Parse("2006-01-02", r.TradeDate)
				if err != nil {
					return err
				}
				rate := &service.CurrencyRate{
					Date:         parsedDate,
					FromCurrency: currency,
					ToCurrency:   targetCurrency,
					Rate:         r.ExchangeRate,
				}

				mutex.Lock()
				currencyRates = append(currencyRates, rate)
				mutex.Unlock()
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return currencyRates, nil
}

func (h *Transaction) getAssetPrices(
	ctx context.Context,
	assetIds []string,
	startDate string,
	endDate string,
) ([]*service.AssetPrice, error) {
	assetPrices := make([]*service.AssetPrice, 0)

	g, gctx := con.WithContext(ctx, maxConcurrency)
	var mutex sync.Mutex

	for _, assetId := range assetIds {
		g.Go(func() error {
			req := &ap.GetAssetPricesByDateRangeRequest{
				AssetId:   assetId,
				StartDate: startDate,
				EndDate:   endDate,
			}

			res, err := h.assetPriceService.GetAssetPricesByDateRange(gctx, req)
			if err != nil {
				return err
			}

			for _, p := range res.GetPrices() {
				parsedDate, err := time.Parse("2006-01-02", p.Date)
				if err != nil {
					return err
				}

				price := &service.AssetPrice{
					Date:         parsedDate,
					AssetID:      assetId,
					ClosingPrice: p.GetPrice(),
				}

				mutex.Lock()
				assetPrices = append(assetPrices, price)
				mutex.Unlock()
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return assetPrices, nil
}

func (h *Transaction) toHistoricalRecords(recs []*service.HistoricalRecord, res *pb.PerformanceHistoryResponse) {
	protoRecords := make([]*pb.HistoricalRecord, len(recs))
	for i, rec := range recs {
		protoRecords[i] = convertToProtoHistoricalRecord(rec)
	}

	res.Records = protoRecords
}

func (h *Transaction) toTransactions(txns []*model.Transaction) []*service.TransactionRecord {
	trades := make([]*service.TransactionRecord, 0, len(txns))

	for _, txn := range txns {
		trades = append(trades, &service.TransactionRecord{
			AssetID:                  txn.AssetID,
			Quantity:                 txn.Quantity,
			TradePrice:               txn.TradePrice,
			TradePriceCurrencyCode:   txn.TradePriceCurrencyCode,
			BrokerageFee:             txn.BrokerageFee,
			BrokerageFeeCurrencyCode: txn.FeeCurrencyCode,
			TransactionType:          txn.TransactionType,
			TransactionDate:          *txn.TransactionDate,
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
