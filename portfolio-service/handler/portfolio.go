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

	pba "github.com/garcios/asset-trak-portfolio/asset-price-service/proto"
	pbc "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	pbp "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
)

const (
	foreignCurrency = "USD"
	targetCurrency  = "AUD"
)

func New(
	currencyService pbc.CurrencyService,
	portfolioRepository PortfolioRepository,
	transactionRepository TransactionRepository,
	performanceService *service.PerformanceService,
	assetPriceService pba.AssetPriceService,
) *Transaction {
	return &Transaction{
		currencyService:       currencyService,
		portfolioRepository:   portfolioRepository,
		transactionRepository: transactionRepository,
		performanceService:    performanceService,
		assetPriceService:     assetPriceService,
	}
}

type Transaction struct {
	currencyService       pbc.CurrencyService
	portfolioRepository   PortfolioRepository
	transactionRepository TransactionRepository
	performanceService    *service.PerformanceService
	assetPriceService     pba.AssetPriceService
}

type PortfolioRepository interface {
	GetHoldings(ctx context.Context, accountID string) ([]*model.BalanceSummary, error)
	GetHoldingAtDateRange(
		ctx context.Context,
		accountID string,
		startDate,
		endDate string,
	) ([]db.Holding, error)
}

type TransactionRepository interface {
	GetTransactions(ctx context.Context, filter db.TransactionFilter) ([]*model.Transaction, error)
}

func (h *Transaction) GetSummaryTotals(
	ctx context.Context,
	req *pbp.SummaryTotalsRequest,
	res *pbp.SummaryTotalsResponse,
) error {
	log.Println("GetSummaryTotals...")
	holdingReq := &pbp.HoldingsRequest{AccountId: req.GetAccountId()}
	holdingsRes := &pbp.HoldingsResponse{}
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

	res.PortfolioValue = &pbp.Money{
		Amount:       portfolioValue,
		CurrencyCode: targetCurrency,
	}

	_, capitalReturnPct := finance.CalculateReturn(totalCost, portfolioValue)
	res.CapitalReturn = &pbp.InvestmentReturn{
		Amount:           totalCapitalReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capitalReturnPct,
	}

	weightedDividendReturnPct := finance.CalculateTotalDividendGainPercentage(financeInvestments)
	res.DividendReturn = &pbp.InvestmentReturn{
		Amount:           totalDividendReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: weightedDividendReturnPct,
	}

	weightedCurrencyReturnPct := finance.CalculateTotalCurrencyGainPercentage(financeInvestments)
	res.CurrencyReturn = &pbp.InvestmentReturn{
		Amount:           totalCurrencyReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: weightedCurrencyReturnPct,
	}

	res.TotalReturn = &pbp.InvestmentReturn{
		Amount:           totalCapitalReturnAmt + totalDividendReturnAmt + totalCurrencyReturnAmt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: capitalReturnPct + weightedDividendReturnPct + weightedCurrencyReturnPct,
	}

	return nil
}

func (h *Transaction) GetHoldings(
	ctx context.Context,
	req *pbp.HoldingsRequest,
	res *pbp.HoldingsResponse,
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

	summaryItems, err := h.portfolioRepository.GetHoldings(ctx, req.GetAccountId())
	if err != nil {
		return err
	}

	res.Investments = make([]*pbp.Investment, 0)

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

		txns, err := h.transactionRepository.GetTransactions(ctx, dbFilter)

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

		dividends, err := h.transactionRepository.GetTransactions(ctx, dbFilterDiv)
		if err != nil {
			return err
		}

		tradesDiv := toTrades(dividends)
		dividendReturn := h.computeDividendReturn(tradesDiv, totalCost.Amount)

		investment := &pbp.Investment{
			AssetSymbol: s.AssetSymbol,
			AssetName:   s.AssetName,
			MarketCode:  s.MarketCode,
			Quantity:    s.Quantity,
			CurrentPrice: &pbp.Money{
				Amount:       s.Price,
				CurrencyCode: s.CurrencyCode,
			},
			AveragePrice: h.computeAveragePrice(trades, targetCurrency),
			TotalValue: &pbp.Money{
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

func (h *Transaction) computeTotalCost(trades []*finance.Trade, targetCurrency string) *pbp.Money {
	return &pbp.Money{
		Amount:       finance.CalculateTotalCost(trades, targetCurrency),
		CurrencyCode: targetCurrency,
	}
}

func (h *Transaction) computeAveragePrice(trades []*finance.Trade, targetCurrency string) *pbp.Money {
	return &pbp.Money{
		Amount:       finance.CalculateAveragePrice(trades, targetCurrency),
		CurrencyCode: targetCurrency,
	}
}

func (h *Transaction) computeCapitalReturn(totalCost, totalValue float64) *pbp.InvestmentReturn {
	amt, pct := finance.CalculateReturn(totalCost, totalValue)

	return &pbp.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeDividendReturn(trades []*finance.Trade, totalCost float64) *pbp.InvestmentReturn {
	amt, pct := finance.CalculateTotalDividendAndReturn(trades, totalCost)

	return &pbp.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeCurrencyReturn(trades []*finance.Trade, currencyRate float64) *pbp.InvestmentReturn {
	amt, pct := finance.CalculateCurrencyReturns(trades, currencyRate, targetCurrency)

	return &pbp.InvestmentReturn{
		Amount:           amt,
		CurrencyCode:     targetCurrency,
		ReturnPercentage: pct,
	}
}

func (h *Transaction) computeTotalReturn(capital, dividend, currency *pbp.InvestmentReturn) *pbp.InvestmentReturn {
	return &pbp.InvestmentReturn{
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
