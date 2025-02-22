package finance

import "fmt"

// Trade represents a stock purchase transaction
type Trade struct {
	AssetID      string  // Asset ID
	Quantity     int     // Number of shares bought
	Price        Money   // Price per share
	Commission   Money   // Trade commission
	TradeType    string  // Trade type e.g. BUY or SELL
	CurrencyRate float64 // Exchange rate at transaction date
}

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
				fmt.Println("SELL transaction exceeds currently owned shares")
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
