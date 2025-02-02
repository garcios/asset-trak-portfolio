package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
	"go-micro.dev/v4"

	"github.com/garcios/asset-trak-portfolio/graphql/graph/model"
)

type Resolver struct{}

// GetBalanceSummary is the resolver for the getBalanceSummary field.
func (r *queryResolver) GetBalanceSummary(ctx context.Context, accountID string) (*model.BalanceSummary, error) {
	resolverClient := micro.NewService(micro.Name("transaction-resolver.client"))
	resolverClient.Init()

	transactionSrv := pb.NewTransactionService("transaction-service", resolverClient.Client())

	req := &pb.BalanceSummaryRequest{
		AccountId: "eb08df3c-958d-4ae8-b3ae-41ec04418786",
	}

	resp, err := transactionSrv.GetBalanceSummary(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("get balance summary error: %w", err)
	}

	balanceItems := make([]*model.BalanceItem, 0, len(resp.BalanceItems))

	for _, item := range resp.BalanceItems {
		balanceItems = append(balanceItems, &model.BalanceItem{
			AssetSymbol: item.AssetSymbol,
			AssetName:   item.AssetName,
			Price:       toMoney(item.Price),
			Quantity:    item.Quantity,
			Value:       toMoney(item.Value),
			TotalGain:   item.TotalGain,
			MarketCode:  item.MarketCode,
		})
	}

	balanceSummary := &model.BalanceSummary{
		AccountID:    accountID,
		TotalValue:   toMoney(resp.TotalValue),
		BalanceItems: balanceItems,
	}

	return balanceSummary, nil
}

func toMoney(m *pb.Money) *model.Money {
	return &model.Money{
		Amount:       m.Amount,
		CurrencyCode: m.CurrencyCode,
	}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
