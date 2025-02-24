package finance

import (
	"log"
	"math"
)

// Trade represents a stock purchase transaction.
type Trade struct {
	AssetID      string  // Asset ID
	Quantity     int     // Number of shares bought
	Price        Money   // Price per share
	Commission   Money   // Trade commission
	TradeType    string  // Trade type e.g. BUY or SELL
	CurrencyRate float64 // Exchange rate at transaction date
	AmountCash   Money   // Dividend
}

// Investment represents an investment holding.
type Investment struct {
	AssetID      string
	TotalValue   float64
	CapitalGain  float64
	Dividend     float64
	CurrencyGain float64
}

// Money represents a monetary value.
type Money struct {
	Amount       float64
	CurrencyCode string
}

// CalculateAveragePrice computes the weighted average cost per share
func CalculateAveragePrice(
	trades []*Trade,
	targetCurrency string,
) float64 {
	var totalCost float64
	var totalShares int

	for _, trade := range trades {
		assetPriceRate := 1.0
		if trade.Price.CurrencyCode != "" && trade.Price.CurrencyCode != targetCurrency {
			assetPriceRate = trade.CurrencyRate
		}

		commissionRate := 1.0
		if trade.Commission.CurrencyCode != "" && trade.Commission.CurrencyCode != targetCurrency {
			commissionRate = trade.CurrencyRate
		}

		// Handle BUY transactions
		if trade.Quantity > 0 {
			totalCost += float64(trade.Quantity)*(trade.Price.Amount*assetPriceRate) + (trade.Commission.Amount * commissionRate)
			totalShares += trade.Quantity
		}

		// Handle SELL transactions
		if trade.Quantity < 0 {
			// Ensure we only offset the cost for the shares being sold (reducing total cost proportionally)
			sellQuantity := -trade.Quantity // Convert to positive quantity for logic
			if sellQuantity > totalShares {
				// Avoid selling more shares than owned
				log.Println("SELL transaction exceeds currently owned shares")
				return 0
			}

			// Reduce the total cost proportionally to the shares being sold
			averagePricePerShare := totalCost / float64(totalShares)
			totalCost -= float64(sellQuantity) * averagePricePerShare

			// Decrease the total shares owned
			totalShares -= sellQuantity
		}

	}

	if totalShares == 0 {
		return 0 // Avoid division by zero
	}

	return totalCost / float64(totalShares)
}

// CalculateTotalCost handles multi-currencies by converting all trades to the target currency
// before computing the total cost.
func CalculateTotalCost(
	trades []*Trade,
	targetCurrency string,
) float64 {
	var totalCost float64

	for _, trade := range trades {
		// Ignore SELL transactions (Quantity < 0)
		if trade.Quantity > 0 {
			// Get the exchange rate for the trade's currency

			assetPriceRate := 1.0
			if trade.Price.CurrencyCode != "" && trade.Price.CurrencyCode != targetCurrency {
				assetPriceRate = trade.CurrencyRate
			}

			commissionRate := 1.0
			if trade.Commission.CurrencyCode != "" && trade.Commission.CurrencyCode != targetCurrency {
				commissionRate = trade.CurrencyRate
			}

			// Convert price and commission to the target currency
			priceInBase := trade.Price.Amount * assetPriceRate
			commissionInBase := trade.Commission.Amount * commissionRate

			// Calculate total cost in base currency
			totalCost += float64(trade.Quantity)*priceInBase + commissionInBase
		}
	}

	return totalCost
}

// CalculateReturn computes the total return percentage of the portfolio
func CalculateReturn(totalCost, totalValue float64) (float64, float64) {
	if totalCost == 0 {
		return 0, 0 // Avoid division by zero
	}

	valueChange := totalValue - totalCost
	pctChange := ((valueChange) / totalCost) * 100

	return valueChange, pctChange
}

// ConvertCurrency converts an amount from one currency to another given an exchange rate.
func ConvertCurrency(amount float64, rate float64) float64 {
	if rate <= 0 || amount < 0 {
		return 0
	}

	return amount * rate
}

// CalculateTotalCurrencyGainPercentage calculates the total percentage currency gain for a slice of investments.
func CalculateTotalCurrencyGainPercentage(investments []*Investment) float64 {
	var totalValue float64 = 0
	var totalCurrencyGain float64 = 0

	// Iterate through each investment to calculate total value and total currency gain
	for _, investment := range investments {
		totalValue += investment.TotalValue
		totalCurrencyGain += (investment.CurrencyGain / investment.TotalValue) * investment.TotalValue
	}

	// Avoid division by zero
	if totalValue == 0 {
		return 0
	}

	// Calculate total percentage of currency gain
	totalPercentage := (totalCurrencyGain / totalValue) * 100

	// Round to two decimal places
	return math.Round(totalPercentage*100) / 100
}

// CalculateTotalDividendGainPercentage calculates the total return percentage for dividend gains
// from an array of investments. The result is rounded to two decimal places.
func CalculateTotalDividendGainPercentage(investments []*Investment) float64 {
	var totalValue float64 = 0
	var totalDividendGain float64 = 0

	// Iterate through each investment to calculate total value and capital gains
	for _, investment := range investments {
		totalValue += investment.TotalValue
		totalDividendGain += investment.Dividend
	}

	// Avoid division by zero
	if totalValue == 0 {
		return 0
	}

	// Calculate total return percentage for capital gains
	totalPercentage := (totalDividendGain / totalValue) * 100

	// Round to two decimal places
	return math.Round(totalPercentage*100) / 100
}

// CalculateTotalDividendAndReturn computes the total dividend amount and return percentage.
func CalculateTotalDividendAndReturn(trades []*Trade, totalCost float64) (float64, float64) {
	var totalDividends float64

	// Sum up the dividends from AmountCash
	for _, trade := range trades {
		if trade.AmountCash.Amount > 0 {
			totalDividends += trade.AmountCash.Amount
		}
	}

	// Calculate return percentage (Total Dividends / Total Cost * 100)
	var returnPct float64
	if totalCost > 0 {
		returnPct = (totalDividends / totalCost) * 100
		returnPct = math.Round(returnPct*100) / 100
	}

	return totalDividends, returnPct
}
