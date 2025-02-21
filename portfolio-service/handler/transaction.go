package handler

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"log"
	"sort"
	"time"

	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
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
	GetHoldings(ctx context.Context, accountID string) ([]*model.BalanceSummary, error)
}

func (h *Transaction) GetSummaryTotals(
	ctx context.Context,
	request *pb.SummaryTotalsRequest,
	response *pb.SummaryTotalsResponse) error {

	return nil
}

func (h *Transaction) GetHoldings(
	ctx context.Context,
	req *pb.HoldingsRequest,
	res *pb.HoldingsResponse,
) error {
	log.Println("GetHoldings...")
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

	summaryItems, err := h.balanceSummaryManager.GetHoldings(ctx, req.GetAccountId())
	if err != nil {
		return err
	}

	res.Investments = make([]*pb.Investment, 0)

	for _, s := range summaryItems {
		value := s.Quantity * s.Price
		if s.CurrencyCode == foreignCurrency {
			value = value * currencyRates.ExchangeRate
		}

		capitalReturn := computeCapitalReturn(s.AssetSymbol)
		dividendReturn := computeDividendReturn(s.AssetSymbol)
		currencyReturn := computeCurrencyReturn(s.AssetSymbol)

		investment := &pb.Investment{
			AssetSymbol: s.AssetSymbol,
			AssetName:   s.AssetName,
			MarketCode:  s.MarketCode,
			Quantity:    s.Quantity,
			Price: &pb.Money{
				Amount:       s.Price,
				CurrencyCode: s.CurrencyCode,
			},
			Value: &pb.Money{
				Amount:       value,
				CurrencyCode: targetCurrency,
			},
			CapitalReturn:  capitalReturn,
			DividendReturn: dividendReturn,
			CurrencyReturn: currencyReturn,
			TotalReturn:    computeTotalReturn(capitalReturn, dividendReturn, currencyReturn),
		}
		res.Investments = append(res.Investments, investment)
	}

	// Sort the investments by Value.Amount in descending order
	sort.Slice(res.Investments, func(i, j int) bool {
		return res.Investments[i].Value.Amount > res.Investments[j].Value.Amount
	})

	return nil
}

// TODO: will compute total gain for specific asset.
func computeCapitalReturn(symbol string) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           0,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: 0,
	}
}

// TODO: will compute total gain for specific asset.
func computeDividendReturn(symbol string) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           0,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: 0,
	}
}

// TODO: will compute total gain for specific asset.
func computeCurrencyReturn(symbol string) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           0,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: 0,
	}
}

func computeTotalReturn(capital, dividend, currency *pb.InvestmentReturn) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           capital.Amount + dividend.Amount + currency.Amount,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capital.ReturnPercentage + dividend.ReturnPercentage + currency.ReturnPercentage,
	}
}
