package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/garcios/asset-trak-portfolio/lib/finance"
)

func TestCalculateDailyHistoricalValueAndCost(t *testing.T) {
	// Mock implementations

	fxRateMap := make(map[string]float64)
	fxRateMap[createFXRateKey("USDAUD", time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC))] = 1.50
	fxRateMap[createFXRateKey("USDAUD", time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC))] = 1.60
	fxRateMap[createFXRateKey("USDAUD", time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC))] = 1.61

	mockExchangeRate := func(fromCurrency string, toCurrency string, date time.Time) (float64, error) {
		key := createFXRateKey(fromCurrency+toCurrency, date)
		if rate, ok := fxRateMap[key]; ok {
			return rate, nil
		}

		return 1.0, nil
	}

	assetPricesMap := make(map[string]float64)
	// Add example asset prices (key: "assetID:date", value: price)
	assetPricesMap[createAssetPriceKey("A1", time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC))] = 90
	assetPricesMap[createAssetPriceKey("A1", time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC))] = 100
	assetPricesMap[createAssetPriceKey("A1", time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC))] = 101

	mockAssetPrice := func(assetID string, date time.Time) (float64, error) {
		key := createAssetPriceKey(assetID, date)
		if price, ok := assetPricesMap[key]; ok {
			return price, nil
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
			targetCurrency: "AUD",
			expected: []HistoricalRecord{
				{
					Date:  time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Value: 1350,
					Cost:  1353,
				},
				{
					Date:  time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
					Value: 1600,
					Cost:  1353,
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
			targetCurrency: "AUD",
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
					Price:           finance.Money{Amount: 100, CurrencyCode: "USD"},
					Commission:      finance.Money{Amount: 2, CurrencyCode: "USD"},
				},
			},
			startDate:      time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
			endDate:        time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
			targetCurrency: "AUD",
			expected: []HistoricalRecord{
				{
					Date:  time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					Value: 675,
					Cost:  676.5,
				},
				{
					Date:  time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC),
					Value: 2400,
					Cost:  2279.7,
				},
				{
					Date:  time.Date(2023, 10, 17, 0, 0, 0, 0, time.UTC),
					Value: 2439.15,
					Cost:  2279.7,
				},
			},
		},
	}

	// Execute test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRedisClient := NewMockRedisClient()
			s := NewPerformanceService(mockRedisClient)
			result, _ := s.CalculateDailyHistoricalValueAndCost(
				context.Background(),
				tt.trades,
				tt.startDate,
				tt.endDate,
				tt.targetCurrency,
				mockExchangeRate,
				mockAssetPrice,
			)
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

// createAssetPriceKey generates the key in the format "assetID:YYYY-MM-DD"
func createAssetPriceKey(assetID string, date time.Time) string {
	return fmt.Sprintf("%s:%s", assetID, date.Format("2006-01-02"))
}

// createFXRateKey generates the key in the format "currencyPair:YYYY-MM-DD"
func createFXRateKey(currencyPair string, date time.Time) string {
	return fmt.Sprintf("%s:%s", currencyPair, date.Format("2006-01-02"))
}
