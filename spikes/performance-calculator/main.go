package main

import (
	"fmt"
	"time"
)

// TransactionRecord represents a portfolio transactions
type TransactionRecord struct {
	AssetID         string
	Quantity        float64
	TradePrice      float64
	TransactionDate time.Time
	TransactionType string
	BrokerageFee    float64
}

// PriceData represents market data for an asset on a specific date, including the closing price.
type PriceData struct {
	Date         time.Time
	AssetID      string
	ClosingPrice float64
}

type DailyPortfolioSummary struct {
	Date       time.Time
	DailyValue float64
	DailyCost  float64
}

func CalculatePortfolioValueAndCost(
	transactions []TransactionRecord,
	priceData []PriceData,
	startDate time.Time,
	endDate time.Time,
) ([]DailyPortfolioSummary, error) {
	// Step 1: Filter relevant market data within the date range
	priceDataMap := make(map[time.Time]map[string]float64) // Date -> AssetID -> ClosingPrice
	for _, data := range priceData {
		if data.Date.After(startDate) && data.Date.Before(endDate) || data.Date.Equal(startDate) || data.Date.Equal(endDate) {
			if _, exists := priceDataMap[data.Date]; !exists {
				priceDataMap[data.Date] = make(map[string]float64)
			}
			priceDataMap[data.Date][data.AssetID] = data.ClosingPrice
		}
	}

	// Step 2: Calculate daily value and cost for each date
	var dailySummaries []DailyPortfolioSummary

	// Initialize running totals for cost
	var runningCost float64
	for currentDay := startDate; !currentDay.After(endDate); currentDay = currentDay.AddDate(0, 0, 1) {
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
			// Calculate cost for the txn
			if txn.TransactionType == "BUY" || txn.TransactionType == "SELL" {
				dailyCost += txn.Quantity*txn.TradePrice + txn.BrokerageFee
			}
		}

		// Calculate daily value
		var dailyValue float64
		assetIDs := extractUniqueAssetIDsByDateRange(transactionsByDate, startDate, currentDay)
		for _, assetID := range assetIDs {
			// Add the value of any assets already held
			quantity := getHoldingQuantity(assetID, startDate, currentDay, transactionsByDate)
			price, found := getLastAvailablePrice(
				currentDay,
				assetID,
				priceDataMap)

			if !found {
				return nil, fmt.Errorf("no price found for asset %v on date %v", assetID, currentDay.Format("2006-01-02"))
			}

			dailyValue += quantity * price
		}

		// Accumulate the cost and value over days
		runningCost += dailyCost

		// Append daily portfolio summary
		dailySummaries = append(dailySummaries, DailyPortfolioSummary{
			Date:       currentDay,
			DailyValue: dailyValue,
			DailyCost:  runningCost,
		})
	}

	return dailySummaries, nil
}

// Helper function to search for the last available price
func getLastAvailablePrice(date time.Time, assetID string, marketDataMap map[time.Time]map[string]float64) (float64, bool) {
	for {
		if prices, exists := marketDataMap[date]; exists {
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

// GroupTransactionsByDate groups a slice of TransactionRecord by their TransactionDate
func groupTransactionsByDate(records []TransactionRecord) map[time.Time][]TransactionRecord {
	transactionMap := make(map[time.Time][]TransactionRecord)

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
	transactionsByDate map[time.Time][]TransactionRecord,
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
	transactionsByDate map[time.Time][]TransactionRecord,
	startDate time.Time,
	endDate time.Time,
) []string {
	// Create a map to store unique asset IDs
	uniqueAssetIDs := make(map[string]struct{})

	// Iterate through the transactions grouped by date
	for txnDate, transactions := range transactionsByDate {
		// Check if the date is within the specified range
		if txnDate.Before(startDate) || txnDate.After(endDate) {
			continue
		}

		// Collect asset IDs for transactions within the date range
		for _, txn := range transactions {
			uniqueAssetIDs[txn.AssetID] = struct{}{}
		}
	}

	// Convert the map keys to a slice
	assetIDSlice := make([]string, 0, len(uniqueAssetIDs))
	for assetID := range uniqueAssetIDs {
		assetIDSlice = append(assetIDSlice, assetID)
	}

	return assetIDSlice
}

func main() {
	// Example transactions data
	tradeDate := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	tradeDatePlusOneDay := tradeDate.AddDate(0, 0, 1)
	tradeDatePlusTwoDays := tradeDate.AddDate(0, 0, 2)
	tradeDatePlusThreeDays := tradeDate.AddDate(0, 0, 3)
	tradeDatePlusFourDays := tradeDate.AddDate(0, 0, 4)

	transactions := []TransactionRecord{
		{AssetID: "AAPL", Quantity: 50, TradePrice: 100.0, TransactionDate: tradeDate, TransactionType: "BUY"},
		{AssetID: "AAPL", Quantity: 10, TradePrice: 101.0, TransactionDate: tradeDatePlusOneDay, TransactionType: "BUY"},
		{AssetID: "GOOGL", Quantity: 20, TradePrice: 200.0, TransactionDate: tradeDate, TransactionType: "BUY"},
		{AssetID: "GOOGL", Quantity: 2, TradePrice: 250.0, TransactionDate: tradeDatePlusOneDay, TransactionType: "BUY"},
		{AssetID: "AAPL", Quantity: -5, TradePrice: 110.0, TransactionDate: tradeDatePlusOneDay, TransactionType: "SELL"},
		{AssetID: "AAPL", Quantity: 495, TradePrice: 0, TransactionDate: tradeDatePlusThreeDays, TransactionType: "SPLIT"}, // represents adjustment after a stock split
	}

	// Day 1
	//--------------------------------
	// Cost (APPL)=  50 * 100 = 5,000
	// Cost (GOOGL) = 20 * 200 =4,000
	// Total Cost (Day 1) = 5000 + 4000 = 9,000
	// Running cost = 9,000
	// -------------------------------
	// Value (APPL) = 50 * 120 = 6,000
	// Value (GOOGL) = 20 * 201 = 4,020
	// Total Value (Day 1) =10,020
	// Running Value = 10,020

	// Day 2
	//--------------------------------
	// Cost (APPL)=  10 * 101 = 1,010
	// Cost (GOOGL) = 2 * 250 = 500
	// Cost (APPL) = -5 * 110 = -550
	// Total Cost (Day 2) = 1010 + 500 -550 = 960
	// Running Cost = 9,000 + 960 = 9,960
	// -------------------------------
	// Value (APPL) = 55 * 125 = 6,875
	// Value (GOOGL) = 22 * 240 = 5,280
	// Total Value = 6,875 + 5,280 =  12,155

	// Day 3 (no transactions on this day, so it will  the current market prices to compute total value of current holdings)
	// Running cost will remain unchanged.
	//--------------------------------
	// Cost (APPL)=  0
	// Cost (GOOGL) = 0
	// Running Total Cost = 0 + 9,960 = 9,960
	// -------------------------------
	// Value (APPL) = 55 * 126 = 6,930
	// Value (GOOGL) = 22 * 260 = 5,720
	// Total Value = 6,930 +  5,720 = 12,650

	// Day 4 (no transactions on this day, so it will  the current market prices to compute total value of current holdings)
	// Running cost will remain unchanged.
	//--------------------------------
	// Cost (APPL)=  0
	// Cost (GOOGL) = 0
	// Running Total Cost = 0 + 9,960 = 9,960
	// -------------------------------
	// Value (APPL) = 55 * 127 = 6,985
	// Value (GOOGL) = 22 * 265 = 5,830
	// Total Value = 6,985 +  5,830 = 12,815

	// Example market price data
	marketPriceData := []PriceData{
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

	// Specify time range
	startDate := tradeDate
	endDate := tradeDatePlusFourDays

	// Calculate daily portfolio values and costs
	dailySummaries, err := CalculatePortfolioValueAndCost(transactions, marketPriceData, startDate, endDate)
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
	//Date: 2023-10-01 | Daily Value: 10020.00 | Daily Cost: 9000.00
	//Date: 2023-10-02 | Daily Value: 12155.00 | Daily Cost: 9960.00
	//Date: 2023-10-03 | Daily Value: 12650.00 | Daily Cost: 9960.00
	//Date: 2023-10-04 | Daily Value: 12815.00 | Daily Cost: 9960.00
	//Date: 2023-10-05 | Daily Value: 12870.00 | Daily Cost: 9960.00
}
