package handler

import (
	"context"
	"fmt"
	"log"
	"time"

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

	transactions := h.toTransactions(txns)
	log.Printf("len(trades): %d\n", len(transactions))

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

	dateRange := service.DateRange{
		StartDate: startDate,
		EndDate:   endDate,
	}

	marketData := service.MarketData{
		AssetPrices:   h.getAssetPrices(),
		CurrencyRates: h.getCurrencyRates(),
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

func (h *Transaction) getCurrencyRates() []*service.CurrencyRate {
	return nil
}

func (h *Transaction) getAssetPrices() []*service.AssetPrice {
	//h.assetPriceService.GetAssetPricesByDateRange()

	return nil
}

func (h *Transaction) toHistoricalRecords(recs []*service.HistoricalRecord, res *pb.PerformanceHistoryResponse) {
	protoRecords := make([]*pb.HistoricalRecord, len(recs))
	for i, rec := range recs {
		protoRecords[i] = convertToProtoHistoricalRecord(rec)
	}

	res.Records = protoRecords
}

func (h *Transaction) toTransactions(txns []*model.Transaction) []*service.TransactionRecord {
	trades := make([]*service.TransactionRecord, len(txns))

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
