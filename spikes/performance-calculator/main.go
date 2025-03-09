package main

import (
	"fmt"
	"time"
)

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

type DailyPortfolioSummary struct {
	Date       time.Time
	DailyValue float64
	DailyCost  float64
}

type DateRange struct {
	StartDate time.Time
	EndDate   time.Time
}

func CalculatePortfolioValueAndCost(
	transactions []*TransactionRecord,
	marketData MarketData,
	targetCurrencyCode string,
	dateRange DateRange,
) ([]*DailyPortfolioSummary, error) {
	// Step 1: Convert assetPrices to map
	priceDataMap := buildAssetPricesMap(marketData.AssetPrices, dateRange.StartDate, dateRange.EndDate)
	currencyRatesMap := buildCurrencyRatesMap(marketData.CurrencyRates)
	fmt.Printf("currencyRatesMap: %v\n", currencyRatesMap)

	// Step 2: Calculate daily value and cost for each date
	var dailySummaries []*DailyPortfolioSummary

	// Initialize running totals for cost
	var runningCost float64
	for currentDay := dateRange.StartDate; !currentDay.After(dateRange.EndDate); currentDay = currentDay.AddDate(0, 0, 1) {
		fmt.Printf("Processing date: %v\n", currentDay.Format("2006-01-02"))

		transactionsByDate := groupTransactionsByDate(transactions)

		// This may or may not contain transactions
		transactionsAtDate, exists := transactionsByDate[currentDay]
		if !exists {
			fmt.Printf("No transactions on date: %v\n", currentDay.Format("2006-01-02"))
		}

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

			fmt.Printf("tradePriceCurrencyRate: %v\n", tradePriceCurrencyRate)
			fmt.Printf("brokerageFeeCurrencyRate: %v\n", brokerageFeeCurrencyRate)

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

		// Append daily portfolio summary
		dailySummaries = append(dailySummaries, &DailyPortfolioSummary{
			Date:       currentDay,
			DailyValue: dailyValue,
			DailyCost:  runningCost,
		})
	}

	return dailySummaries, nil
}

func buildAssetPricesMap(assetPrices []*AssetPrice, startDate time.Time, endDate time.Time) map[time.Time]map[string]float64 {
	priceDataMap := make(map[time.Time]map[string]float64) // Date -> AssetID -> ClosingPrice
	for _, assetPrice := range assetPrices {
		if assetPrice.Date.After(startDate) && assetPrice.Date.Before(endDate) || assetPrice.Date.Equal(startDate) || assetPrice.Date.Equal(endDate) {
			if _, exists := priceDataMap[assetPrice.Date]; !exists {
				priceDataMap[assetPrice.Date] = make(map[string]float64)
			}
			priceDataMap[assetPrice.Date][assetPrice.AssetID] = assetPrice.ClosingPrice
		}
	}
	return priceDataMap
}

func buildCurrencyRatesMap(currencyRates []*CurrencyRate) map[time.Time]map[string]float64 {
	// Initialize the result map
	ratesMap := make(map[time.Time]map[string]float64)

	// Iterate over the CurrencyRate slice
	for _, rate := range currencyRates {
		// Ensure the date key exists in the rates map
		if _, exists := ratesMap[rate.Date]; !exists {
			ratesMap[rate.Date] = make(map[string]float64)
		}

		// Create a currency pair key "from-to"
		currencyPair := rate.FromCurrency + "-" + rate.ToCurrency

		// Add the rate to the nested map
		ratesMap[rate.Date][currencyPair] = rate.Rate
	}

	return ratesMap
}

// Helper function to search for the last available price
func getLastAvailablePrice(date time.Time, assetID string, assetPricesMap map[time.Time]map[string]float64) (float64, bool) {
	for {
		if prices, exists := assetPricesMap[date]; exists {
			if price, ok := prices[assetID]; ok {
				return price, true
			}
		}
		date = date.AddDate(0, 0, -1)                                 // Move to the previous day
		if date.Before(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)) { // Stop if date is unreasonably old
			break
		}
	}
	return 0, false // No price found
}

// Helper function to search for the last available currency rate
func getLastAvailableCurrencyRate(
	date time.Time,
	fromCurrency string,
	toCurrency string,
	currencyRatesMap map[time.Time]map[string]float64,
) (float64, bool) {
	if fromCurrency == toCurrency {
		return 1, true
	}

	currencyPair := fromCurrency + "-" + toCurrency
	reverseCurrencyPair := toCurrency + "-" + fromCurrency
	for {
		if rates, exists := currencyRatesMap[date]; exists {
			if rate, ok := rates[currencyPair]; ok {
				return rate, true
			}

			// try reverse currency pair
			if rate, ok := rates[reverseCurrencyPair]; ok {
				return 1 / rate, true
			}
		}
		date = date.AddDate(0, 0, -1)                                 // Move to the previous day
		if date.Before(time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)) { // Stop if date is unreasonably old
			break
		}
	}
	return 0, false // No price found
}

// GroupTransactionsByDate groups a slice of TransactionRecord by their TransactionDate
func groupTransactionsByDate(records []*TransactionRecord) map[time.Time][]*TransactionRecord {
	transactionMap := make(map[time.Time][]*TransactionRecord)

	for _, record := range records {
		// Group by the date (only year, month, and day, ignoring time of day)
		date := time.Date(record.TransactionDate.Year(), record.TransactionDate.Month(), record.TransactionDate.Day(), 0, 0, 0, 0, record.TransactionDate.Location())

		// Append the transaction to the map
		transactionMap[date] = append(transactionMap[date], record)
	}

	return transactionMap
}

func getHoldingQuantity(
	assetID string,
	startDate time.Time,
	endDate time.Time,
	transactionsByDate map[time.Time][]*TransactionRecord,
) float64 {
	var totalQuantity float64
	for txnDate, transactions := range transactionsByDate {
		// Check if the date is within the specified range
		if txnDate.Before(startDate) || txnDate.After(endDate) {
			continue
		}

		for _, txn := range transactions {
			if txn.AssetID != assetID {
				continue
			}
			if txn.TransactionType == "BUY" || txn.TransactionType == "SELL" || txn.TransactionType == "SPLIT" {
				totalQuantity += txn.Quantity
			}
		}
	}

	return totalQuantity
}

// extractUniqueAssetIDsByDateRange takes transactions grouped by date and a date range,
// and returns a slice of unique asset IDs within that range.
func extractUniqueAssetIDsByDateRange(
	transactionsByDate map[time.Time][]*TransactionRecord,
	startDate time.Time,
	endDate time.Time,
) map[string]string {
	// Create a map to store unique asset IDs and corresponding currency
	uniqueAssetIDs := make(map[string]string)

	// Iterate through the transactions grouped by date
	for txnDate, transactions := range transactionsByDate {
		// Check if the date is within the specified range
		if txnDate.Before(startDate) || txnDate.After(endDate) {
			continue
		}

		// Collect asset IDs for transactions within the date range
		for _, txn := range transactions {
			if txn.TradePriceCurrencyCode == "" { // this might happen e.g. SPLIT transaction
				continue
			}

			if _, found := uniqueAssetIDs[txn.AssetID]; !found {
				uniqueAssetIDs[txn.AssetID] = txn.TradePriceCurrencyCode
			}
		}

	}

	return uniqueAssetIDs
}

func main() {
	// Example transactions data
	tradeDate := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	tradeDatePlusOneDay := tradeDate.AddDate(0, 0, 1)
	tradeDatePlusTwoDays := tradeDate.AddDate(0, 0, 2)
	tradeDatePlusThreeDays := tradeDate.AddDate(0, 0, 3)
	tradeDatePlusFourDays := tradeDate.AddDate(0, 0, 4)

	transactions := []*TransactionRecord{
		{AssetID: "AAPL", Quantity: 50, TradePrice: 100.0, TradePriceCurrencyCode: "USD", TransactionDate: tradeDate, TransactionType: "BUY", BrokerageFee: 20, BrokerageFeeCurrencyCode: "AUD"},
		{AssetID: "GOOGL", Quantity: 20, TradePrice: 200.0, TradePriceCurrencyCode: "USD", TransactionDate: tradeDate, TransactionType: "BUY", BrokerageFee: 10, BrokerageFeeCurrencyCode: "AUD"},
		{AssetID: "AAPL", Quantity: 10, TradePrice: 101.0, TradePriceCurrencyCode: "USD", TransactionDate: tradeDatePlusOneDay, TransactionType: "BUY", BrokerageFee: 10, BrokerageFeeCurrencyCode: "AUD"},
		{AssetID: "GOOGL", Quantity: 2, TradePrice: 250.0, TradePriceCurrencyCode: "USD", TransactionDate: tradeDatePlusOneDay, TransactionType: "BUY", BrokerageFee: 10, BrokerageFeeCurrencyCode: "AUD"},
		{AssetID: "AAPL", Quantity: -5, TradePrice: 110.0, TradePriceCurrencyCode: "USD", TransactionDate: tradeDatePlusOneDay, TransactionType: "SELL", BrokerageFee: 10, BrokerageFeeCurrencyCode: "AUD"},
		{AssetID: "AAPL", Quantity: 495, TradePrice: 0, TransactionDate: tradeDatePlusThreeDays, TransactionType: "SPLIT"}, // represents adjustment after a stock split
	}

	// Day 1
	//--------------------------------
	// Cost (APPL)=  50 * 100 * 1.59 + 20 = 7,970
	// Cost (GOOGL) = 20 * 200 * 1.59 + 10 = 6,370
	// Total Cost (Day 1) = 7,970 + 6,370 = 14,340
	// Running cost = 14,340
	// -------------------------------
	// Value (APPL) = 50 * 120 * 1.59 = 9,540
	// Value (GOOGL) = 20 * 201 * 1.59 = 6,391.8
	// Total Value (Day 1) = 9,540 + 6,391.8 =  15,931.80
	// Running Value = 15,931.80

	// Day 2
	//--------------------------------
	// Cost (APPL)=  10 * 101 * 1.59 + 10= 1,615.9
	// Cost (GOOGL) = 2 * 250 * 1.59 + 10 = 805
	// Cost (APPL) = -5 * 110 * 1.59 + 10= -864.5
	// Total Cost (Day 2) = 1,615.9 + 805 -864.5 = 1556.4
	// Running Cost = 14,340 + 1556.4= 15,896.4
	// -------------------------------
	// Value (APPL) = 55 * 125 * 1.59 = 10,931.25
	// Value (GOOGL) = 22 * 240 * 1.59 = 8,395.2
	// Total Value = 10,931.25 +  8,395.2 =  19,326.45

	// Day 3 (no transactions on this day, so it will  the current market prices to compute total value of current holdings)
	// Running cost will remain unchanged.
	//--------------------------------
	// Cost (APPL)=  0
	// Cost (GOOGL) = 0
	// Running Total Cost = 0 + 15,896.4 = 15,896.4
	// -------------------------------
	// Value (APPL) = 55 * 126 * 1.59 = 11,018.7
	// Value (GOOGL) = 22 * 260 * 1.59 = 9,094.8
	// Total Value = 11,018.7 +  9,094.8 = 20,113.5

	// Day 4 (no transactions on this day, so it will  the current market prices to compute total value of current holdings)
	// Running cost will remain unchanged.
	//--------------------------------
	// Cost (APPL)=  0
	// Cost (GOOGL) = 0
	// Running Total Cost = 0 + 15,896.4 = 15,896.4
	// -------------------------------
	// Value (APPL) = 55 * 127 * 1.59 = 11,106.15
	// Value (GOOGL) = 22 * 265 * 1.59 = 9,269.7
	// Total Value = 11,106.15 +  9,269.7 = 20,375.85

	// Day 5 (only SPLIT transaction)
	// Running cost will remain unchanged.
	//--------------------------------
	// Cost (APPL)=  0
	// Cost (GOOGL) = 0
	// Running Total Cost = 0 + 15,896.4 = 15,896.4
	// -------------------------------
	// Value (APPL) = (55 + 495) * 12.80 * 1.59 = 11,193.6
	// Value (GOOGL) = 22 * 265 * 1.59 = 9,269.7
	// Total Value = 11,106.15 +  9,269.7 = 20,463.3

	// Example market asset price data
	assetPrices := []*AssetPrice{
		{Date: tradeDate, AssetID: "AAPL", ClosingPrice: 120.0},
		{Date: tradeDate, AssetID: "GOOGL", ClosingPrice: 201.0},
		{Date: tradeDatePlusOneDay, AssetID: "AAPL", ClosingPrice: 125.0},
		{Date: tradeDatePlusOneDay, AssetID: "GOOGL", ClosingPrice: 240.0},
		{Date: tradeDatePlusTwoDays, AssetID: "AAPL", ClosingPrice: 126.0},
		{Date: tradeDatePlusTwoDays, AssetID: "GOOGL", ClosingPrice: 260.0},
		{Date: tradeDatePlusThreeDays, AssetID: "AAPL", ClosingPrice: 12.70},
		{Date: tradeDatePlusThreeDays, AssetID: "GOOGL", ClosingPrice: 265.0},
		{Date: tradeDatePlusFourDays, AssetID: "AAPL", ClosingPrice: 12.80},
		{Date: tradeDatePlusFourDays, AssetID: "GOOGL", ClosingPrice: 265.0},
	}

	// Example currency rates data
	currencyRates := []*CurrencyRate{
		{Date: tradeDate, FromCurrency: "USD", ToCurrency: "AUD", Rate: 1.59},
	}

	// Specify time range
	dateRange := DateRange{
		StartDate: tradeDate,
		EndDate:   tradeDatePlusFourDays,
	}

	marketData := MarketData{
		AssetPrices:   assetPrices,
		CurrencyRates: currencyRates,
	}

	// Calculate daily portfolio values and costs
	targetCurrencyCode := "AUD"
	dailySummaries, err := CalculatePortfolioValueAndCost(transactions, marketData, targetCurrencyCode, dateRange)
	if err != nil {
		fmt.Printf("Error calculating portfolio value and cost: %v\n", err)
		return
	}

	// Output results
	fmt.Println("--------------------------")
	fmt.Println("Daily Portfolio Summary:")
	for _, summary := range dailySummaries {
		fmt.Printf("Date: %v | Daily Value: %.2f | Daily Cost: %.2f\n", summary.Date.Format("2006-01-02"), summary.DailyValue, summary.DailyCost)
	}

	//Daily Portfolio Summary:
	//Date: 2023-10-01 | Daily Value: 15931.80 | Daily Cost: 14340.00
	//Date: 2023-10-02 | Daily Value: 19326.45 | Daily Cost: 15896.40
	//Date: 2023-10-03 | Daily Value: 20113.50 | Daily Cost: 15896.40
	//Date: 2023-10-04 | Daily Value: 20375.85 | Daily Cost: 15896.40
	//Date: 2023-10-05 | Daily Value: 20463.30 | Daily Cost: 15896.40
}
