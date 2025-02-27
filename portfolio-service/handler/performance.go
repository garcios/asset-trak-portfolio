package handler

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"time"
)

func (h *Transaction) GetPerformanceHistory(
	ctx context.Context,
	req *pb.PerformanceHistoryRequest,
	res *pb.PerformanceHistoryResponse,
) error {

	// TODO: Get trades, use getTransactions
	var (
		trades    []*finance.Trade
		startDate time.Time
		endDate   time.Time
	)

	// TODO: use currency service
	getExchangeRate := func(fromCurrency string, toCurrency string, date time.Time) (float64, error) {
		return 0, nil
	}

	// TODO: use asset service
	getAssetPrice := func(assetID string, date time.Time) (float64, error) {
		return 0, nil
	}

	recs, err := h.performanceService.CalculateDailyHistoricalValueAndCost(
		ctx,
		trades,
		startDate,
		endDate,
		targetCurrency,
		getExchangeRate,
		getAssetPrice,
	)

	if err != nil {
		return err
	}

	// TODO: populate response
	fmt.Printf("recs: %#v\n", recs)

	return nil
}
