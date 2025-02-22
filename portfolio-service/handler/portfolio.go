package handler

import (
	"context"
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
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

func New(
	currencyService pbc.CurrencyService,
	portfolioSummaryManager PortfolioManager,
	transactionManager TransactionManager,
) *Transaction {
	return &Transaction{
		currencyService:    currencyService,
		portfolioManager:   portfolioSummaryManager,
		transactionManager: transactionManager,
	}
}

type Transaction struct {
	currencyService    pbc.CurrencyService
	portfolioManager   PortfolioManager
	transactionManager TransactionManager
}

type PortfolioManager interface {
	GetHoldings(ctx context.Context, accountID string) ([]*model.BalanceSummary, error)
}

type TransactionManager interface {
	GetTransactions(ctx context.Context, filter db.TransactionFilter) ([]*model.Transaction, error)
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

	summaryItems, err := h.portfolioManager.GetHoldings(ctx, req.GetAccountId())
	if err != nil {
		return err
	}

	res.Investments = make([]*pb.Investment, 0)

	for _, s := range summaryItems {
		totalValue := s.Quantity * s.Price
		if s.CurrencyCode == foreignCurrency {
			totalValue = finance.ConvertCurrency(totalValue, currencyRates.ExchangeRate)
		}

		dbFilter := db.TransactionFilter{
			AccountID: req.GetAccountId(),
			AssetID:   s.AssetID,
		}

		txns, err := h.transactionManager.GetTransactions(ctx, dbFilter)

		if err != nil {
			return err
		}

		trades := toTrades(txns, currencyRates.ExchangeRate)

		totalCost := h.computeTotalCost(trades, targetCurrency)
		capitalReturn := h.computeCapitalReturn(totalCost.Amount, totalValue)
		dividendReturn := h.computeDividendReturn(trades)
		currencyReturn := h.computeCurrencyReturn(trades)

		investment := &pb.Investment{
			AssetSymbol: s.AssetSymbol,
			AssetName:   s.AssetName,
			MarketCode:  s.MarketCode,
			Quantity:    s.Quantity,
			CurrentPrice: &pb.Money{
				Amount:       s.Price,
				CurrencyCode: s.CurrencyCode,
			},
			AveragePrice: h.computeAveragePrice(trades, targetCurrency),
			TotalValue: &pb.Money{
				Amount:       totalValue,
				CurrencyCode: targetCurrency,
			},
			TotalCost:      totalCost,
			CapitalReturn:  capitalReturn,
			DividendReturn: dividendReturn,
			CurrencyReturn: currencyReturn,
			TotalReturn:    h.computeTotalReturn(capitalReturn, dividendReturn, currencyReturn),
		}
		res.Investments = append(res.Investments, investment)
	}

	// Sort the investments by Value.Amount in descending order
	sort.Slice(res.Investments, func(i, j int) bool {
		return res.Investments[i].TotalValue.Amount > res.Investments[j].TotalValue.Amount
	})

	return nil
}

func (h *Transaction) computeTotalCost(trades []*finance.Trade, targetCurrency string) *pb.Money {
	return &pb.Money{
		Amount:       finance.CalculateTotalCost(trades, targetCurrency),
		CurrencyCode: targetCurrency,
	}
}

func (h *Transaction) computeAveragePrice(trades []*finance.Trade, targetCurrency string) *pb.Money {
	return &pb.Money{
		Amount:       finance.CalculateAveragePrice(trades, targetCurrency),
		CurrencyCode: targetCurrency,
	}
}

func (h *Transaction) computeCapitalReturn(totalCost, totalValue float64) *pb.InvestmentReturn {
	amt, pct := finance.CalculateReturn(totalCost, totalValue)

	return &pb.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeDividendReturn(trades []*finance.Trade) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           0,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: 0,
	}
}

func (h *Transaction) computeCurrencyReturn(trades []*finance.Trade) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           0,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: 0,
	}
}

func (h *Transaction) computeTotalReturn(capital, dividend, currency *pb.InvestmentReturn) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           capital.Amount + dividend.Amount + currency.Amount,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capital.ReturnPercentage + dividend.ReturnPercentage + currency.ReturnPercentage,
	}
}

func toTrades(txns []*model.Transaction, currencyRate float64) []*finance.Trade {
	trades := make([]*finance.Trade, 0)

	for _, txn := range txns {
		trade := &finance.Trade{
			AssetID:      txn.AssetID,
			Quantity:     int(txn.Quantity),
			Price:        finance.Money{Amount: txn.TradePrice, CurrencyCode: txn.AssetPriceCurrencyCode},
			Commission:   finance.Money{Amount: txn.TradeCommission, CurrencyCode: txn.CommissionCurrencyCode},
			CurrencyRate: currencyRate,
		}
		trades = append(trades, trade)
	}

	return trades
}
