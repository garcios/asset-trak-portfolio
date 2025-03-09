package service

import (
	"log"
	"time"
)

func displayRow(row []string) {
	log.Printf("row: %v", row)
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
