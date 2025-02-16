package resolvers

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/graphql/generated"
	"github.com/garcios/asset-trak-portfolio/graphql/middlewares"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// GetHoldingsSummary is the resolver for the getBalanceSummary field.
func (r *queryResolver) GetHoldingsSummary(ctx context.Context, accountID string) ([]*models.Investment, error) {
	svcs := middlewares.GetServices(ctx)
	resp, err := svcs.TransactionService.GetHoldingsSummary(ctx, accountID)
	if err != nil {
		return nil, err
	}

	investments := make([]*models.Investment, 0, len(resp.BalanceItems))

	for _, item := range resp.BalanceItems {
		investments = append(investments, &models.Investment{
			AssetSymbol: item.AssetSymbol,
			AssetName:   item.AssetName,
			MarketCode:  item.MarketCode,
			Price:       toMoney(item.Price),
			Quantity:    item.Quantity,
			Value:       toMoney(item.Value),
			CapitalGain: &models.MoneyWithPercentage{
				Amount:       0,
				CurrencyCode: "AUD",
				Percentage:   0,
			},
			Dividend: &models.MoneyWithPercentage{
				Amount:       0,
				CurrencyCode: "AUD",
				Percentage:   0,
			},
			CurrencyGain: &models.MoneyWithPercentage{
				Amount:       0,
				CurrencyCode: "AUD",
				Percentage:   0,
			},
			TotalReturn: &models.MoneyWithPercentage{
				Amount:       0,
				CurrencyCode: "AUD",
				Percentage:   0,
			},
		})
	}

	return investments, nil
}

func toMoney(m *pb.Money) *models.Money {
	return &models.Money{
		Amount:       m.Amount,
		CurrencyCode: m.CurrencyCode,
	}
}
