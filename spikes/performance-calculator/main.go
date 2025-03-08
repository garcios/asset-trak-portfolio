package main

import (
	"fmt"
	"time"
)

// Define structures for portfolio transactions and market data
type TransactionRecord struct {
	AssetID         string
	Quantity        float64
	TradePrice      float64
	TransactionDate time.Time
	TransactionType string
}

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

func CalculatePortfolioValueAndCost(transactions []TransactionRecord, priceData []PriceData, startDate, endDate time.Time) []DailyPortfolioSummary {
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
	for currentDay := startDate; !currentDay.After(endDate); currentDay = currentDay.AddDate(0, 0, 1) {
		fmt.Printf("Processing date: %v\n", currentDay.Format("2006-01-02"))
		currentPrices, _ := priceDataMap[currentDay]

		transactionsByDate := groupTransactionsByDate(transactions)
		transactionsAtDate, exists := transactionsByDate[currentDay]
		if !exists {
			continue
		}

		var dailyValue, dailyCost float64
		for _, txn := range transactionsAtDate {
			// Calculate value for the txn
			var price float64
			var found bool
			if currentPrices != nil {
				price, found = currentPrices[txn.AssetID]
			}

			// If price not found, look for the last available price backward
			if !found {
				price, found = getLastAvailablePrice(
					currentDay.AddDate(0, 0, -1),
					txn.AssetID,
					priceDataMap)
			}

			// If price is found, calculate the value
			if found {
				dailyValue += txn.Quantity * price
			}

			// Calculate cost for the txn (static over time)
			if txn.TransactionType == "BUY" || txn.TransactionType == "SELL" {
				dailyCost += txn.Quantity * txn.TradePrice
			}
		}

		// Append daily portfolio summary
		dailySummaries = append(dailySummaries, DailyPortfolioSummary{
			Date:       currentDay,
			DailyValue: dailyValue,
			DailyCost:  dailyCost,
		})
	}

	return dailySummaries
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

func main() {
	// Example transactions data
	transactions := []TransactionRecord{
		{AssetID: "AAPL", Quantity: 50, TradePrice: 100.0, TransactionDate: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), TransactionType: "BUY"},
		{AssetID: "AAPL", Quantity: 10, TradePrice: 101.0, TransactionDate: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), TransactionType: "BUY"},
		{AssetID: "GOOGL", Quantity: 20, TradePrice: 200.0, TransactionDate: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), TransactionType: "BUY"},
		{AssetID: "GOOGL", Quantity: 2, TradePrice: 250.0, TransactionDate: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), TransactionType: "BUY"},
		{AssetID: "AAPL", Quantity: -5, TradePrice: 110.0, TransactionDate: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), TransactionType: "SELL"},
	}

	// Day 1
	//--------------------------------
	// Cost (APPL)=  50 * 100 = 5,000
	// Cost (GOOGL) = 20 * 200 =4,000
	// Total Cost = 5000 + 4000 = 9,000
	// -------------------------------
	// Value (APPL) = 50 * 120 = 6,000
	// Value (GOOGL) = 20 * 201 = 4,020
	// Total Vale =10,020

	// Day 2
	//--------------------------------
	// Cost (APPL)=  10 * 101 = 1,010
	// Cost (GOOGL) = 2 * 250 = 500
	// Cost (APPL) = -5 * 110 = -550
	// Total Cost = 1010 + 500 -550 = 960
	// -------------------------------
	// Value (APPL) = 10 * 125 = 1,250
	// Value (GOOGL) = 2 * 240 = 480
	// Value (APPL) = -5 * 125 = -625
	// Total Vale = 1,105

	// Example market price data
	marketPriceData := []PriceData{
		{Date: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), AssetID: "AAPL", ClosingPrice: 120.0},
		{Date: time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC), AssetID: "GOOGL", ClosingPrice: 201.0},
		{Date: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), AssetID: "AAPL", ClosingPrice: 125.0},
		{Date: time.Date(2023, 10, 2, 0, 0, 0, 0, time.UTC), AssetID: "GOOGL", ClosingPrice: 240.0},
		{Date: time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC), AssetID: "AAPL", ClosingPrice: 126.0},
		{Date: time.Date(2023, 10, 3, 0, 0, 0, 0, time.UTC), AssetID: "GOOGL", ClosingPrice: 260.0},
		{Date: time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC), AssetID: "AAPL", ClosingPrice: 127.0},
		{Date: time.Date(2023, 10, 4, 0, 0, 0, 0, time.UTC), AssetID: "GOOGL", ClosingPrice: 265.0},
		{Date: time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC), AssetID: "AAPL", ClosingPrice: 128.0},
		{Date: time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC), AssetID: "GOOGL", ClosingPrice: 265.0},
	}

	// Specify time range
	startDate := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)

	// Calculate daily portfolio values and costs
	dailySummaries := CalculatePortfolioValueAndCost(transactions, marketPriceData, startDate, endDate)

	// Output results
	for _, summary := range dailySummaries {
		fmt.Printf("Date: %v | Daily Value: %.2f | Daily Cost: %.2f\n", summary.Date.Format("2006-01-02"), summary.DailyValue, summary.DailyCost)
	}
}
