package finance

import (
	"math"
	"testing"
)

func TestCalculateAveragePrice(t *testing.T) {
	tests := []struct {
		name           string
		trades         []*Trade
		targetCurrency string
		expected       float64
	}{
		{
			name: "single_trade_no_conversion",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "USD"}, Commission: Money{Amount: 10.0, CurrencyCode: "USD"}, TradeType: "BUY"},
			},
			targetCurrency: "USD",
			expected:       101,
		},
		{
			name: "multiple_trades_no_conversion",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "USD"}, Commission: Money{Amount: 10.0, CurrencyCode: "USD"}, TradeType: "BUY"},
				{AssetID: "ABC", Quantity: 20, Price: Money{Amount: 90.0, CurrencyCode: "USD"}, Commission: Money{Amount: 5.0, CurrencyCode: "USD"}, TradeType: "BUY"},
			},
			targetCurrency: "USD",
			expected:       93.83333333333333,
		},
		{
			name: "sell_trades_exceeds_shares",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "USD"}, Commission: Money{Amount: 10.0, CurrencyCode: "USD"}, TradeType: "BUY"},
				{AssetID: "ABC", Quantity: -20, Price: Money{Amount: 90.0, CurrencyCode: "USD"}, Commission: Money{Amount: 5.0, CurrencyCode: "USD"}, TradeType: "SELL"},
			},
			targetCurrency: "USD",
			expected:       0,
		},
		{
			name: "trade_with_conversion",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "EUR"}, Commission: Money{Amount: 10.0, CurrencyCode: "EUR"}, TradeType: "BUY", CurrencyRate: 1.1},
			},
			targetCurrency: "USD",
			expected:       111.10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CalculateAveragePrice(tc.trades, tc.targetCurrency)
			const epsilon = 0.01 // Define a small precision delta
			if math.Abs(actual-tc.expected) > epsilon {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestCalculateTotalCost(t *testing.T) {
	tests := []struct {
		name           string
		trades         []*Trade
		targetCurrency string
		expected       float64
	}{
		{
			name: "single_trade_no_conversion",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "USD"}, Commission: Money{Amount: 10.0, CurrencyCode: "USD"}, TradeType: "BUY"},
			},
			targetCurrency: "USD",
			expected:       1010.0,
		},
		{
			name: "multiple_trades_with_conversion",
			trades: []*Trade{
				{AssetID: "ABC", Quantity: 10, Price: Money{Amount: 100.0, CurrencyCode: "EUR"}, Commission: Money{Amount: 10.0, CurrencyCode: "EUR"}, TradeType: "BUY", CurrencyRate: 1.1},
				{AssetID: "ABC", Quantity: 5, Price: Money{Amount: 90.0, CurrencyCode: "EUR"}, Commission: Money{Amount: 5.0, CurrencyCode: "EUR"}, TradeType: "BUY", CurrencyRate: 1.1},
			},
			targetCurrency: "USD",
			expected:       1667.5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := CalculateTotalCost(tc.trades, tc.targetCurrency)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestCalculateReturn(t *testing.T) {
	tests := []struct {
		name       string
		totalCost  float64
		totalValue float64
		expectedVC float64
		expectedPC float64
	}{
		{
			name:       "no_cost",
			totalCost:  0,
			totalValue: 1000,
			expectedVC: 0,
			expectedPC: 0,
		},
		{
			name:       "positive_return",
			totalCost:  1000,
			totalValue: 1500,
			expectedVC: 500,
			expectedPC: 50,
		},
		{
			name:       "negative_return",
			totalCost:  1000,
			totalValue: 800,
			expectedVC: -200,
			expectedPC: -20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualVC, actualPC := CalculateReturn(tc.totalCost, tc.totalValue)
			if actualVC != tc.expectedVC || actualPC != tc.expectedPC {
				t.Errorf("expected (VC: %v, PC: %v), got (VC: %v, PC: %v)", tc.expectedVC, tc.expectedPC, actualVC, actualPC)
			}
		})
	}
}

func TestConvertCurrency(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		rate     float64
		expected float64
	}{
		{
			name:     "positive_conversion",
			amount:   100,
			rate:     1.2,
			expected: 120,
		},
		{
			name:     "zero_rate",
			amount:   100,
			rate:     0,
			expected: 0,
		},
		{
			name:     "negative_amount",
			amount:   -100,
			rate:     1.2,
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := ConvertCurrency(tc.amount, tc.rate)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
