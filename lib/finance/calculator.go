package finance

import "fmt"

// Trade represents a stock purchase transaction
type Trade struct {
	AssetID    string  // Asset ID
	Quantity   int     // Number of shares bought
	Price      float64 // Price per share
	Commission float64 // Trade commission
	TradeType  string  // Trade type e.g. BUY or SELL
}

// CalculateAveragePrice computes the weighted average cost per share
func CalculateAveragePrice(trades []*Trade) float64 {
	var totalCost float64
	var totalShares int

	for _, trade := range trades {
		// Handle BUY transactions
		if trade.Quantity > 0 {
			totalCost += float64(trade.Quantity)*trade.Price + trade.Commission
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

// CalculateProfitLoss calculates profit or loss based on current price
func CalculateProfitLoss(totalShares int, totalCost, currentPrice float64) float64 {
	currentValue := float64(totalShares) * currentPrice
	return currentValue - totalCost
}

// CalculateTotalCost computes the total cost of all trades including commissions
func CalculateTotalCost(trades []*Trade) float64 {
	var totalCost float64

	for _, trade := range trades {
		// Ignore SELL transactions (Quantity < 0)
		if trade.Quantity > 0 {
			totalCost += float64(trade.Quantity)*trade.Price + trade.Commission
		}
	}

	return totalCost
}

// CalculateNewAveragePrice computes the new average price after adding a new trade
func CalculateNewAveragePrice(
	oldAvgPrice float64,
	oldShares int,
	newShares int,
	newPrice float64,
	newCommission float64,
) float64 {
	// Calculate total cost before new trade
	oldTotalCost := float64(oldShares) * oldAvgPrice

	// Calculate new total cost after adding new trade
	newTotalCost := oldTotalCost + (float64(newShares) * newPrice) + newCommission

	// Calculate new total shares
	newTotalShares := oldShares + newShares

	// Avoid division by zero
	if newTotalShares == 0 {
		return 0
	}

	// Compute new average price
	return newTotalCost / float64(newTotalShares)
}

// CalculatePercentageReturn computes the percentage return based on the current price
func CalculatePercentageReturn(totalShares int, totalCost, currentPrice float64) float64 {
	if totalShares == 0 || totalCost == 0 {
		return 0 // Avoid division by zero
	}

	// Compute current value of holdings
	currentValue := float64(totalShares) * currentPrice

	// Compute percentage return
	return ((currentValue - totalCost) / totalCost) * 100
}

// CalculateTotalValue computes the current total portfolio value based on market prices
func CalculateTotalValue(trades []*Trade, stockPrices map[string]float64) float64 {
	var totalValue float64
	for _, trade := range trades {
		if currentPrice, exists := stockPrices[trade.AssetID]; exists {
			totalValue += float64(trade.Quantity) * currentPrice
		}
	}
	return totalValue
}

// CalculatePortfolioReturn computes the total return percentage of the portfolio
func CalculatePortfolioReturn(totalCost, totalValue float64) (float64, float64) {
	if totalCost == 0 {
		return 0, 0 // Avoid division by zero
	}

	valueChange := totalValue - totalCost
	pctChange := ((valueChange) / totalCost) * 100

	return valueChange, pctChange
}
