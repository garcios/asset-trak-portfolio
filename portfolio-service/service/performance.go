package service

import (
	"context"
	"fmt"
	"time"
)

type PerformanceService struct {
}

func NewPerformanceService() *PerformanceService {
	return &PerformanceService{}
}

// TransactionRecord represents a portfolio transactions
type TransactionRecord struct {
	AssetID                  string
	Quantity                 float64
	TradePrice               float64
	TradePriceCurrencyCode   string
	TransactionDate          time.Time
	TransactionType          string
	BrokerageFee             float64
	BrokerageFeeCurrencyCode string
}

// HistoricalRecord represents the value and cost for a specific day.
type HistoricalRecord struct {
	Date  time.Time
	Value float64 // Value in the target currency (based on current day price)
	Cost  float64 // Cost in the target currency (based on transaction day price)
}

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

type MarketData struct {
	AssetPrices   []*AssetPrice
	CurrencyRates []*CurrencyRate
}

// AssetPrice represents market data for an asset on a specific date, including the closing price.
type AssetPrice struct {
	Date         time.Time
	AssetID      string
	ClosingPrice float64
}

// CurrencyRate represents fx rate data for a currency pair on a specific date.
type CurrencyRate struct {
	Date         time.Time
	FromCurrency string
	ToCurrency   string
	Rate         float64
}

// CalculateDailyHistoricalValueAndCost calculates the portfolio's daily historical value and cost
// across a date range while supporting multi-currency and fetching historical asset prices.
func (s PerformanceService) CalculateDailyHistoricalValueAndCost(
	ctx context.Context,
	transactions []*TransactionRecord,
	marketData MarketData,
	targetCurrencyCode string,
	dateRange DateRange,
) ([]*HistoricalRecord, error) {
	// Step 1: Convert assetPrices to map
	priceDataMap := buildAssetPricesMap(marketData.AssetPrices, dateRange.StartDate, dateRange.EndDate)
	currencyRatesMap := buildCurrencyRatesMap(marketData.CurrencyRates)

	// Step 2: Calculate daily value and cost for each date
	var dailyRecords []*HistoricalRecord

	// Initialize running totals for cost
	var runningCost float64
	for currentDay := dateRange.StartDate; !currentDay.After(dateRange.EndDate); currentDay = currentDay.AddDate(0, 0, 1) {
		transactionsByDate := groupTransactionsByDate(transactions)

		// This may or may not contain transactions
		transactionsAtDate, _ := transactionsByDate[currentDay]

		// Calculate daily cost
		var dailyCost float64
		for _, txn := range transactionsAtDate {
			if txn.TransactionType != "BUY" && txn.TransactionType != "SELL" {
				continue
			}

			// Calculate cost for the txn
			tradePriceCurrencyRate, found := getLastAvailableCurrencyRate(
				currentDay,
				txn.TradePriceCurrencyCode,
				targetCurrencyCode,
				currencyRatesMap,
			)

			if !found {
				currencyPair := txn.TradePriceCurrencyCode + "-" + targetCurrencyCode
				return nil, fmt.Errorf("no currency rate found for currency pair %v on date %v", currencyPair, currentDay.Format("2006-01-02"))
			}

			brokerageFeeCurrencyRate, found := getLastAvailableCurrencyRate(
				currentDay,
				txn.BrokerageFeeCurrencyCode,
				targetCurrencyCode,
				currencyRatesMap,
			)

			if !found {
				return nil, fmt.Errorf("no currency rate found for currency pair %v on date %v", txn.TradePriceCurrencyCode, currentDay.Format("2006-01-02"))
			}

			dailyCost += txn.Quantity*txn.TradePrice*tradePriceCurrencyRate + (txn.BrokerageFee * brokerageFeeCurrencyRate)
		}

		// Calculate daily value
		var dailyValue float64
		assetIDs := extractUniqueAssetIDsByDateRange(transactionsByDate, dateRange.StartDate, currentDay)
		for assetID, currencyCode := range assetIDs {
			// Add the value of any assets already held
			quantity := getHoldingQuantity(assetID, dateRange.StartDate, currentDay, transactionsByDate)
			price, found := getLastAvailablePrice(
				currentDay,
				assetID,
				priceDataMap)

			if !found {
				return nil, fmt.Errorf("no price found for asset %v on date %v", assetID, currentDay.Format("2006-01-02"))
			}

			tradePriceCurrencyRate, found := getLastAvailableCurrencyRate(
				currentDay,
				currencyCode,
				targetCurrencyCode,
				currencyRatesMap,
			)

			if !found {
				currencyPair := currencyCode + "-" + targetCurrencyCode
				return nil, fmt.Errorf("no currency rate found for currency pair %v on date %v", currencyPair, currentDay.Format("2006-01-02"))
			}

			dailyValue += quantity * price * tradePriceCurrencyRate
		}

		// Accumulate the cost and value over days
		runningCost += dailyCost

		// Append daily historical record
		dailyRecords = append(dailyRecords, &HistoricalRecord{
			Date:  currentDay,
			Value: dailyValue,
			Cost:  runningCost,
		})

	}

	return dailyRecords, nil
}
