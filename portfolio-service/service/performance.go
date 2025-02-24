package service

import (
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"time"
)

// HistoricalRecord represents the value and cost for a specific day.
type HistoricalRecord struct {
	Date  time.Time
	Value float64 // Value in the target currency (based on current day price)
	Cost  float64 // Cost in the target currency (based on transaction day price)
}

// ExchangeRateFunc retrieves the exchange rate for a given currency and date.
type ExchangeRateFunc func(fromCurrency string, toCurrency string, date time.Time) float64

// AssetPriceFunc retrieves the price of an asset on a specific date.
type AssetPriceFunc func(assetID string, date time.Time) float64

// CalculateDailyHistoricalValueAndCost calculates the portfolio's daily historical value and cost
// across a date range while supporting multi-currency and fetching historical asset prices.
func CalculateDailyHistoricalValueAndCost(
	trades []*finance.Trade,
	startDate, endDate time.Time,
	targetCurrency string,
	getExchangeRate ExchangeRateFunc,
	getAssetPrice AssetPriceFunc,
) []HistoricalRecord {
	// Initialize a map to store daily updates for portfolio value and cost
	dailyRecords := make(map[time.Time]*HistoricalRecord)

	// Iterate over each trade to update the portfolio data for the trade's date and subsequent days
	for _, trade := range trades {
		tradeDate := trade.TransactionDate // Assume Trade struct has a Date field of type time.Time
		if tradeDate.Before(startDate) || tradeDate.After(endDate) {
			continue // Skip trades outside the date range
		}

		currentDay := tradeDate
		for !currentDay.After(endDate) {
			if dailyRecords[currentDay] == nil {
				dailyRecords[currentDay] = &HistoricalRecord{
					Date:  currentDay,
					Value: 0,
					Cost:  0,
				}
			}

			// Get exchange rates for trade currency and commission currency for the trade date
			priceExchangeRateOnTradeDate := getExchangeRate(trade.Price.CurrencyCode, targetCurrency, tradeDate)
			commissionExchangeRateOnTradeDate := getExchangeRate(trade.Commission.CurrencyCode, targetCurrency, tradeDate)

			// Get the exchange rate for the target currency for the current day (for value calculation)
			priceExchangeRateOnCurrentDay := getExchangeRate(trade.Price.CurrencyCode, targetCurrency, currentDay)

			// Get the asset price on the transaction date (for cost calculation)
			assetPriceOnTradeDate := getAssetPrice(trade.AssetID, tradeDate)
			if assetPriceOnTradeDate == 0 {
				assetPriceOnTradeDate = trade.Price.Amount // Fallback if no historical price is available
			}

			// Get the asset price for the current day (for value calculation)
			assetPriceOnCurrentDay := getAssetPrice(trade.AssetID, currentDay)
			if assetPriceOnCurrentDay == 0 {
				assetPriceOnCurrentDay = trade.Price.Amount // Fallback if no price available
			}

			// Convert amounts to the target currency
			priceInTargetCurrencyForCost := assetPriceOnTradeDate * priceExchangeRateOnTradeDate
			priceInTargetCurrencyForValue := assetPriceOnCurrentDay * priceExchangeRateOnCurrentDay
			commissionInTargetCurrency := trade.Commission.Amount * commissionExchangeRateOnTradeDate

			// Update the value and cost based on trade details
			dailyRecord := dailyRecords[currentDay]

			// BUY transactions (Quantity > 0)
			if trade.Quantity > 0 {
				dailyRecord.Cost += trade.Quantity*priceInTargetCurrencyForCost + commissionInTargetCurrency
				dailyRecord.Value += trade.Quantity * priceInTargetCurrencyForValue
			}

			// SELL transactions (Quantity < 0)
			if trade.Quantity < 0 {
				sellQuantity := -trade.Quantity
				if dailyRecord.Value > 0 && sellQuantity > 0 {
					averageCost := dailyRecord.Cost / dailyRecord.Value // Average cost per unit
					dailyRecord.Cost -= sellQuantity * averageCost
					dailyRecord.Value -= sellQuantity * priceInTargetCurrencyForValue
				}
			}

			// Ensure final values are converted to two decimal place
			dailyRecord.Cost = finance.ToTwoDecimalPlaces(dailyRecord.Cost)
			dailyRecord.Value = finance.ToTwoDecimalPlaces(dailyRecord.Value)

			// Move to the next day
			currentDay = currentDay.AddDate(0, 0, 1)
		}
	}

	// Collect the results into an array sorted by date
	var result []HistoricalRecord
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if dailyRecords[d] != nil {
			result = append(result, *dailyRecords[d])
		} else {
			// Add a record for dates without updates
			result = append(result, HistoricalRecord{
				Date:  d,
				Value: 0,
				Cost:  0,
			})
		}
	}

	return result
}
