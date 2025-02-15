package resolvers

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
)

func (r *queryResolver) GetHistoricalValues(ctx context.Context, accountID string) ([]*models.PerformanceData, error) {
	performanceData := []*models.PerformanceData{
		{
			TradeDate:    "2021-01-01",
			Amount:       100,
			CurrencyCode: "AUD",
		},
		{
			TradeDate: "2021-02-01",

			Amount:       101,
			CurrencyCode: "AUD",
		},
		{
			TradeDate:    "2021-03-01",
			Amount:       150,
			CurrencyCode: "AUD",
		},
		{
			TradeDate:    "2021-04-01",
			Amount:       120,
			CurrencyCode: "AUD",
		},
	}

	return performanceData, nil
}
