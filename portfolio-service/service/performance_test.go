package service

import (
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"testing"
	"time"
)

func TestCalculateDailyHistoricalValueAndCost(t *testing.T) {
	// Mock implementations
	mockExchangeRate := func(fromCurrency string, toCurrency string, date time.Time) (float64, error) {
		if fromCurrency == "USD" && toCurrency == "EUR" {
			return 0.85, nil
		}
		return 1.0, nil
	}

	mockAssetPrice := func(assetID string, date time.Time) (float64, error) {
		if assetID == "A1" {
			return 100.0, nil
		}
		return 0.0, nil
	}

	// Test cases
	tests := []struct {
		name           string
		trades         []*finance.Trade
		startDate      time.Time
		endDate        time.Time
		targetCurrency string
		expected       []HistoricalRecord
	}{
		{
			name: "single trade within date range",
			trades: []*finance.Trade{
				{
					AssetID:         "A1",
					TransactionDate: time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Quantity:        10,
					Price:           finance.Money{Amount: 90, CurrencyCode: "USD"},
					Commission:      finance.Money{Amount: 2, CurrencyCode: "USD"},
				},
			},
			startDate:      time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
			targetCurrency: "EUR",
			expected: []HistoricalRecord{
				{
					Date:  time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Value: 850.0,
					Cost:  851.7,
				},
				{
					Date:  time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
					Value: 850.0,
					Cost:  851.7,
				},
			},
		},
		{
			name: "trade outside date range",
			trades: []*finance.Trade{
				{
					AssetID:         "A1",
					TransactionDate: time.Date(2023, 9, 10, 0, 0, 0, 0, time.UTC),
					Quantity:        10,
					Price:           finance.Money{Amount: 90, CurrencyCode: "USD"},
					Commission:      finance.Money{Amount: 2, CurrencyCode: "USD"},
				},
			},
			startDate:      time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC),
			targetCurrency: "EUR",
			expected: []HistoricalRecord{
				{Date: time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
				{Date: time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
				{Date: time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
				{Date: time.Date(2023, 10, 18, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
				{Date: time.Date(2023, 10, 19, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
				{Date: time.Date(2023, 10, 20, 0, 0, 0, 0, time.UTC), Value: 0, Cost: 0},
			},
		},
		{
			name: "multiple trades within date range",
			trades: []*finance.Trade{
				{
					AssetID:         "A1",
					TransactionDate: time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Quantity:        5,
					Price:           finance.Money{Amount: 90, CurrencyCode: "USD"},
					Commission:      finance.Money{Amount: 1, CurrencyCode: "USD"},
				},
				{
					AssetID:         "A1",
					TransactionDate: time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
					Quantity:        10,
					Price:           finance.Money{Amount: 95, CurrencyCode: "USD"},
					Commission:      finance.Money{Amount: 2, CurrencyCode: "USD"},
				},
			},
			startDate:      time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
			targetCurrency: "EUR",
			expected: []HistoricalRecord{
				{
					Date:  time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Value: 425.0,
					Cost:  425.85,
				},
				{
					Date:  time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
					Value: 1275.0,
					Cost:  1277.55,
				},
				{
					Date:  time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
					Value: 1275.0,
					Cost:  1277.55,
				},
			},
		},
	}

	// Execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := CalculateDailyHistoricalValueAndCost(tt.trades, tt.startDate, tt.endDate, tt.targetCurrency, mockExchangeRate, mockAssetPrice)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %d records, got %d", len(tt.expected), len(result))
			}
			for i, r := range result {
				if r.Date != tt.expected[i].Date || r.Value != tt.expected[i].Value || r.Cost != tt.expected[i].Cost {
					t.Errorf("record %d mismatch: got %+v, expected %+v", i, r, tt.expected[i])
				}
			}
		})
	}
}
