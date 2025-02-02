package handler

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
	"log"

	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

func New(currencyService pbc.CurrencyService,
	balanceSummaryManager BalanceSummaryManager,
) *Transaction {
	return &Transaction{
		currencyService:       currencyService,
		balanceSummaryManager: balanceSummaryManager,
	}
}

type Transaction struct {
	currencyService       pbc.CurrencyService
	balanceSummaryManager BalanceSummaryManager
}

type BalanceSummaryManager interface {
	GetBalanceSummary(ctx context.Context, accountID string) ([]*model.BalanceSummary, error)
}

func (h *Transaction) GetBalanceSummary(
	ctx context.Context,
	req *pb.BalanceSummaryRequest,
	res *pb.BalanceSummaryResponse,
) error {
	fmt.Println("GetBalanceSummary...")

	currencyRates, err := h.currencyService.GetExchangeRate(
		context.Background(),
		&pbc.GetExchangeRateRequest{
			FromCurrency: "USD",
			ToCurrency:   "AUD",
			TradeDate:    "2025-01-30",
		},
	)
	if err != nil {
		log.Printf("Error calling GetExchangeRate: %v", err)
		return err
	}

	summaryItems, err := h.balanceSummaryManager.GetBalanceSummary(ctx, req.GetAccountId())
	if err != nil {
		return err
	}

	res.AccountId = req.GetAccountId()
	res.BalanceItems = make([]*pb.BalanceItem, 0)

	for _, s := range summaryItems {
		amount := s.Quantity * s.Price
		if s.CurrencyCode == "USD" {
			amount = amount * currencyRates.ExchangeRate
		}

		protoBalanceItem := &pb.BalanceItem{
			AssetSymbol: s.AssetSymbol,
			Amount: &pb.Money{
				Amount:       amount,
				CurrencyCode: "AUD",
			},
		}
		res.BalanceItems = append(res.BalanceItems, protoBalanceItem)
	}

	return nil
}
