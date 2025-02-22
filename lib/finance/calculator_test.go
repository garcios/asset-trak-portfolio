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

func TestCalculateTotalCurrencyGainPercentage(t *testing.T) {
	tests := []struct {
		name        string
		investments []*Investment
		expected    float64
	}{
		{
			name: "Positive and negative gains cancel out",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 100, CurrencyGain: 10},  // 10% gain
				{AssetID: "B", TotalValue: 200, CurrencyGain: -10}, // -5% loss
			},
			expected: 0.0, // Gains cancel each other out
		},
		{
			name: "All positive gains",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 150, CurrencyGain: 15}, // 10% gain
				{AssetID: "B", TotalValue: 200, CurrencyGain: 20}, // 10% gain
			},
			expected: 10.0, // Overall gain is 10%
		},
		{
			name: "All negative gains",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 100, CurrencyGain: -10}, // -10% loss
				{AssetID: "B", TotalValue: 300, CurrencyGain: -30}, // -10% loss
			},
			expected: -10.0, // Overall loss is -10%
		},
		{
			name: "Mixed gains and weights",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 50, CurrencyGain: 5},    // 10% gain
				{AssetID: "B", TotalValue: 150, CurrencyGain: -15}, // -10% loss
			},
			expected: -5.0, // Weighted total loss is -5%
		},
		{
			name: "Zero investment value",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 0, CurrencyGain: 0},
			},
			expected: 0.0, // Zero total value results in 0% gain
		},
		{
			name:        "Empty input",
			investments: []*Investment{},
			expected:    0.0, // No investments, no gain
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateTotalCurrencyGainPercentage(test.investments)
			if result != test.expected {
				t.Errorf("got %.2f, want %.2f", result, test.expected)
			}
		})
	}
}

func TestCalculateTotalDividendGainPercentage(t *testing.T) {
	tests := []struct {
		name        string
		investments []*Investment
		expected    float64
	}{
		{
			name: "Multiple dividends",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 100, Dividend: 10}, // 10% gain
				{AssetID: "B", TotalValue: 200, Dividend: 20}, // 10% gain
			},
			expected: 10.0, // Weighted average gain
		},
		{
			name: "Zero total value",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 0, Dividend: 10}, // Value is 0
			},
			expected: 0.0, // Division by zero avoided
		},
		{
			name:        "Empty investments array",
			investments: []*Investment{},
			expected:    0.0, // No investments implies no dividend
		},
		{
			name: "All zero capital gains",
			investments: []*Investment{
				{AssetID: "A", TotalValue: 100, Dividend: 0}, // No dividend
				{AssetID: "B", TotalValue: 200, Dividend: 0}, // No dividend
			},
			expected: 0.0, // Zero gain percentage
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CalculateTotalDividendGainPercentage(test.investments)
			if result != test.expected {
				t.Errorf("Test %q failed: got %.2f, expected %.2f", test.name, result, test.expected)
			}
		})
	}
}
