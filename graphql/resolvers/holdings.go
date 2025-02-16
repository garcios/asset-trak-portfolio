package resolvers

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/graphql/generated"
	"github.com/garcios/asset-trak-portfolio/graphql/models"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
	"go-micro.dev/v4"
)

type queryResolver struct{ *Resolver }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// GetHoldingsSummary is the resolver for the getBalanceSummary field.
func (r *queryResolver) GetHoldingsSummary(ctx context.Context, accountID string) ([]*models.Investment, error) {
	resolverClient := micro.NewService(micro.Name("transaction-resolver.client"))
	resolverClient.Init()

	transactionSrv := pb.NewTransactionService("transaction-service", resolverClient.Client())

	req := &pb.BalanceSummaryRequest{
		AccountId: accountID,
	}

	resp, err := transactionSrv.GetBalanceSummary(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("get balance summary error: %w", err)
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
