package resolvers

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
)

func (r *queryResolver) GetSummaryTotals(ctx context.Context, accountID string) (*models.SummaryTotals, error) {
	// TODO: implement this with call to transaction service
	summaryTotals := &models.SummaryTotals{
		PortfolioValue: &models.Money{
			Amount:       250000,
			CurrencyCode: "AUD",
		},
		CapitalGain: &models.MoneyWithPercentage{
			Amount:       42000,
			CurrencyCode: "AUD",
			Percentage:   18,
		},
		Dividends: &models.MoneyWithPercentage{
			Amount:       1200,
			CurrencyCode: "AUD",
			Percentage:   2.5,
		},
		CurrencyGain: &models.MoneyWithPercentage{
			Amount:       259,
			CurrencyCode: "AUD",
			Percentage:   0.54,
		},
		TotalReturn: &models.MoneyWithPercentage{
			Amount:       52000,
			CurrencyCode: "AUD",
			Percentage:   24,
		},
	}

	return summaryTotals, nil
}
