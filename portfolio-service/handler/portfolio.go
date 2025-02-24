package handler

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/db"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/model"
	"github.com/garcios/asset-trak-portfolio/portfolio-service/service"
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
	req *pb.SummaryTotalsRequest,
	res *pb.SummaryTotalsResponse,
) error {
	log.Println("GetSummaryTotals...")
	holdingReq := &pb.HoldingsRequest{AccountId: req.GetAccountId()}
	holdingsRes := &pb.HoldingsResponse{}
	err := h.GetHoldings(ctx, holdingReq, holdingsRes)
	if err != nil {
		return fmt.Errorf("failed to get holdings: %w", err)
	}

	portfolioValue := 0.0
	totalCost := 0.0
	totalCapitalReturnAmt := 0.0
	totalDividendReturnAmt := 0.0
	totalCurrencyReturnAmt := 0.0

	financeInvestments := make([]*finance.Investment, 0)

	for _, investment := range holdingsRes.GetInvestments() {
		portfolioValue += investment.GetTotalValue().GetAmount()
		totalCapitalReturnAmt += investment.GetCapitalReturn().GetAmount()
		totalDividendReturnAmt += investment.GetDividendReturn().GetAmount()
		totalCurrencyReturnAmt += investment.GetCurrencyReturn().GetAmount()
		totalCost += investment.GetTotalCost().GetAmount()
		financeInvestments = append(financeInvestments, &finance.Investment{
			AssetID:      investment.AssetSymbol,
			TotalValue:   investment.GetTotalValue().GetAmount(),
			CapitalGain:  investment.GetCapitalReturn().GetAmount(),
			CurrencyGain: investment.GetCurrencyReturn().GetAmount(),
			Dividend:     investment.GetDividendReturn().GetAmount(),
		})
	}

	res.PortfolioValue = &pb.Money{
		Amount:       portfolioValue,
		CurrencyCode: targetCurrency,
	}

	_, capitalReturnPct := finance.CalculateReturn(totalCost, portfolioValue)
	res.CapitalReturn = &pb.InvestmentReturn{
		Amount:           totalCapitalReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capitalReturnPct,
	}

	weightedDividendReturnPct := finance.CalculateTotalDividendGainPercentage(financeInvestments)
	res.DividendReturn = &pb.InvestmentReturn{
		Amount:           totalDividendReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: weightedDividendReturnPct,
	}

	weightedCurrencyReturnPct := finance.CalculateTotalCurrencyGainPercentage(financeInvestments)
	res.CurrencyReturn = &pb.InvestmentReturn{
		Amount:           totalCurrencyReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: weightedCurrencyReturnPct,
	}

	res.TotalReturn = &pb.InvestmentReturn{
		Amount:           totalCapitalReturnAmt + totalDividendReturnAmt + totalCurrencyReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capitalReturnPct + weightedDividendReturnPct + weightedCurrencyReturnPct,
	}

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
			AccountID:        req.GetAccountId(),
			AssetID:          s.AssetID,
			TransactionTypes: []string{service.TransactionTypeBuy, service.TransactionTypeSell},
		}

		txns, err := h.transactionManager.GetTransactions(ctx, dbFilter)

		if err != nil {
			return err
		}

		trades := toTrades(txns)

		totalCost := h.computeTotalCost(trades, targetCurrency)
		capitalReturn := h.computeCapitalReturn(totalCost.Amount, totalValue)
		currencyReturn := h.computeCurrencyReturn(trades, currencyRates.ExchangeRate)

		dbFilterDiv := db.TransactionFilter{
			AccountID:        req.GetAccountId(),
			AssetID:          s.AssetID,
			TransactionTypes: []string{service.TransactionTypeDividend},
		}

		dividends, err := h.transactionManager.GetTransactions(ctx, dbFilterDiv)
		tradesDiv := toTrades(dividends)
		dividendReturn := h.computeDividendReturn(tradesDiv, totalCost.Amount)

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

func (h *Transaction) computeDividendReturn(trades []*finance.Trade, totalCost float64) *pb.InvestmentReturn {
	amt, pct := finance.CalculateTotalDividendAndReturn(trades, totalCost)

	return &pb.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeCurrencyReturn(trades []*finance.Trade, currencyRate float64) *pb.InvestmentReturn {
	amt, pct := finance.CalculateCurrencyReturns(trades, currencyRate, targetCurrency)

	return &pb.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeTotalReturn(capital, dividend, currency *pb.InvestmentReturn) *pb.InvestmentReturn {
	return &pb.InvestmentReturn{
		Amount:           capital.Amount + dividend.Amount + currency.Amount,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capital.ReturnPercentage + dividend.ReturnPercentage + currency.ReturnPercentage,
	}
}

func toTrades(txns []*model.Transaction) []*finance.Trade {
	trades := make([]*finance.Trade, 0)

	for _, txn := range txns {
		trade := &finance.Trade{
			AssetID:      txn.AssetID,
			Quantity:     txn.Quantity,
			Price:        finance.Money{Amount: txn.TradePrice, CurrencyCode: txn.TradePriceCurrencyCode},
			Commission:   finance.Money{Amount: txn.BrokerageFee, CurrencyCode: txn.FeeCurrencyCode},
			CurrencyRate: 1 / txn.ExchangeRate,
			AmountCash:   finance.Money{Amount: txn.AmountCash, CurrencyCode: txn.AmountCurrencyCode},
		}
		trades = append(trades, trade)
	}

	return trades
}
