package handler

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
	"log"
	"time"

	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

const (
	foreignCurrency = "USD"
	targetCurrency  = "AUD"
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
	log.Println("GetBalanceSummary...")
	now := time.Now()

	currencyRates, err := h.currencyService.GetExchangeRate(
		context.Background(),
		&pbc.GetExchangeRateRequest{
			FromCurrency: foreignCurrency,
			ToCurrency:   targetCurrency,
			TradeDate:    now.Format("2006-01-02"),
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

	var totalValue float64
	for _, s := range summaryItems {
		value := s.Quantity * s.Price
		if s.CurrencyCode == foreignCurrency {
			value = value * currencyRates.ExchangeRate
		}

		totalValue += value

		protoBalanceItem := &pb.BalanceItem{
			AssetSymbol: s.AssetSymbol,
			AssetName:   s.AssetName,
			Quantity:    s.Quantity,
			Price: &pb.Money{
				Amount:       s.Price,
				CurrencyCode: s.CurrencyCode,
			},
			Value: &pb.Money{
				Amount:       value,
				CurrencyCode: targetCurrency,
			},
			TotalGain:  computeTotalGain(s.AssetSymbol),
			MarketCode: s.MarketCode,
		}
		res.BalanceItems = append(res.BalanceItems, protoBalanceItem)
	}

	res.TotalValue = &pb.Money{
		Amount:       totalValue,
		CurrencyCode: targetCurrency,
	}

	return nil
}

// TODO: will compute total gain for specific asset.
func computeTotalGain(symbol string) float64 {
	return 0
}
