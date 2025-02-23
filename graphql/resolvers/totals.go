package resolvers

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/middlewares"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
	"log"
)

func (r *queryResolver) GetSummaryTotals(ctx context.Context, accountID string) (*models.SummaryTotals, error) {
	svcs := middlewares.GetServices(ctx)
	resp, err := svcs.PortfolioService.GetSummaryTotals(ctx, accountID)
	if err != nil {
		return nil, err
	}

	log.Printf("GetSummaryTotals:resp: %+v\n", resp)

	var portfolioValue *models.Money
	if resp.PortfolioValue != nil {
		portfolioValue = &models.Money{
			Amount:       resp.PortfolioValue.Amount,
			CurrencyCode: resp.PortfolioValue.CurrencyCode,
		}
	}

	var capitalGain *models.MoneyWithPercentage
	if resp.CapitalReturn != nil {
		capitalGain = &models.MoneyWithPercentage{
			Amount:       resp.CapitalReturn.Amount,
			CurrencyCode: resp.CapitalReturn.CurrencyCode,
			Percentage:   resp.CapitalReturn.ReturnPercentage,
		}
	}

	var currencyGain *models.MoneyWithPercentage
	if resp.CurrencyReturn != nil {
		currencyGain = &models.MoneyWithPercentage{
			Amount:       resp.CurrencyReturn.Amount,
			CurrencyCode: resp.CurrencyReturn.CurrencyCode,
			Percentage:   resp.CurrencyReturn.ReturnPercentage,
		}
	}

	var dividends *models.MoneyWithPercentage
	if resp.DividendReturn != nil {
		dividends = &models.MoneyWithPercentage{
			Amount:       resp.DividendReturn.Amount,
			CurrencyCode: resp.DividendReturn.CurrencyCode,
			Percentage:   resp.DividendReturn.ReturnPercentage,
		}
	}

	var totalReturn *models.MoneyWithPercentage
	if resp.TotalReturn != nil {
		totalReturn = &models.MoneyWithPercentage{
			Amount:       resp.TotalReturn.Amount,
			CurrencyCode: resp.TotalReturn.CurrencyCode,
			Percentage:   resp.TotalReturn.ReturnPercentage,
		}
	}

	summaryTotals := &models.SummaryTotals{
		PortfolioValue: portfolioValue,
		CapitalGain:    capitalGain,
		Dividends:      dividends,
		CurrencyGain:   currencyGain,
		TotalReturn:    totalReturn,
	}

	return summaryTotals, nil
}
