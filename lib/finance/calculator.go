package finance

// Trade represents a stock purchase transaction
type Trade struct {
	Ticker     string  // Stock symbol (e.g., AMZN, GOOGL)
	Quantity   int     // Number of shares bought
	Price      float64 // Price per share
	Commission float64 // Trade commission
}

// CalculateAveragePrice computes the weighted average cost per share
func CalculateAveragePrice(trades []*Trade) float64 {
	var totalCost float64
	var totalShares int

	for _, trade := range trades {
		totalCost += float64(trade.Quantity)*trade.Price + trade.Commission
		totalShares += trade.Quantity
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
func CalculateTotalCost(trades []Trade) float64 {
	var totalCost float64

	for _, trade := range trades {
		totalCost += float64(trade.Quantity)*trade.Price + trade.Commission
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
func CalculateTotalValue(trades []Trade, stockPrices map[string]float64) float64 {
	var totalValue float64
	for _, trade := range trades {
		if currentPrice, exists := stockPrices[trade.Ticker]; exists {
			totalValue += float64(trade.Quantity) * currentPrice
		}
	}
	return totalValue
}

// CalculatePortfolioReturn computes the total return percentage of the portfolio
func CalculatePortfolioReturn(trades []Trade, stockPrices map[string]float64) float64 {
	totalCost := CalculateTotalCost(trades)
	totalValue := CalculateTotalValue(trades, stockPrices)

	if totalCost == 0 {
		return 0 // Avoid division by zero
	}

	return ((totalValue - totalCost) / totalCost) * 100
}
