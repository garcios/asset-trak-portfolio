package resolvers

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/middlewares"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
)

func (r *queryResolver) GetPerformanceHistory(
	ctx context.Context,
	accountID string,
	startDate string,
	endDate string,
) ([]*models.PerformanceData, error) {
	svcs := middlewares.GetServices(ctx)
	resp, err := svcs.PortfolioService.GetPerformanceHistory(ctx, accountID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	performanceData := make([]*models.PerformanceData, 0, len(resp.GetRecords()))

	for _, item := range resp.GetRecords() {
		performanceData = append(performanceData, &models.PerformanceData{
			TradeDate:    item.TradeDate,
			Cost:         item.Cost,
			Value:        item.Value,
			CurrencyCode: item.CurrencyCode,
		})
	}

	return performanceData, nil
}
