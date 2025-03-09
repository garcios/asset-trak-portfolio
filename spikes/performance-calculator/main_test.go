package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculatePortfolioValueAndCost(t *testing.T) {
	tests := []struct {
		name               string
		transactions       []*TransactionRecord
		marketData         MarketData
		targetCurrencyCode string
		dateRange          DateRange
		expectedSummaries  []*DailyPortfolioSummary
		expectError        bool
	}{
		{
			name: "basic case with valid data",
			transactions: []*TransactionRecord{
				{
					TransactionDate:          time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
					TransactionType:          "BUY",
					AssetID:                  "ASSET123",
					TradePrice:               100,
					Quantity:                 10,
					TradePriceCurrencyCode:   "USD",
					BrokerageFee:             5,
					BrokerageFeeCurrencyCode: "USD",
				},
			},
			marketData: MarketData{
				AssetPrices: []*AssetPrice{
					{
						Date:         time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
						AssetID:      "ASSET123",
						ClosingPrice: 120,
					},
				},
				CurrencyRates: []*CurrencyRate{
					{
						Date:         time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
						FromCurrency: "USD",
						ToCurrency:   "USD",
						Rate:         1.0,
					},
				},
			},
			targetCurrencyCode: "USD",
			dateRange: DateRange{
				StartDate: time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			expectedSummaries: []*DailyPortfolioSummary{
				{
					Date:       time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
					DailyValue: 1200,
					DailyCost:  1005,
				},
			},
			expectError: false,
		},
		{
			name: "missing currency rate",
			transactions: []*TransactionRecord{
				{
					TransactionDate:          time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
					TransactionType:          "BUY",
					AssetID:                  "ASSET123",
					TradePrice:               100,
					Quantity:                 10,
					TradePriceCurrencyCode:   "EUR",
					BrokerageFee:             5,
					BrokerageFeeCurrencyCode: "EUR",
				},
			},
			marketData: MarketData{
				AssetPrices: []*AssetPrice{
					{
						Date:         time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
						AssetID:      "ASSET123",
						ClosingPrice: 100,
					},
				},
				CurrencyRates: nil,
			},
			targetCurrencyCode: "USD",
			dateRange: DateRange{
				StartDate: time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			expectedSummaries: nil,
			expectError:       true,
		},
		{
			name:         "no transactions",
			transactions: nil,
			marketData: MarketData{
				AssetPrices: []*AssetPrice{
					{
						Date:         time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
						AssetID:      "ASSET123",
						ClosingPrice: 100,
					},
				},
				CurrencyRates: []*CurrencyRate{
					{
						Date:         time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
						FromCurrency: "USD",
						ToCurrency:   "USD",
						Rate:         1.0,
					},
				},
			},
			targetCurrencyCode: "USD",
			dateRange: DateRange{
				StartDate: time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			expectedSummaries: []*DailyPortfolioSummary{
				{
					Date:       time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
					DailyValue: 0,
					DailyCost:  0,
				},
			},
			expectError: false,
		},
		{
			name: "asset price missing for some dates",
			transactions: []*TransactionRecord{
				{
					TransactionDate:          time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
					TransactionType:          "BUY",
					AssetID:                  "ASSET123",
					TradePrice:               50,
					Quantity:                 20,
					TradePriceCurrencyCode:   "USD",
					BrokerageFee:             1,
					BrokerageFeeCurrencyCode: "USD",
				},
			},
			marketData: MarketData{
				AssetPrices: []*AssetPrice{
					{
						Date:         time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
						AssetID:      "ASSET123",
						ClosingPrice: 55,
					},
				},
				CurrencyRates: []*CurrencyRate{
					{
						Date:         time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
						FromCurrency: "USD",
						ToCurrency:   "USD",
						Rate:         1.0,
					},
				},
			},
			targetCurrencyCode: "USD",
			dateRange: DateRange{
				StartDate: time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
				EndDate:   time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			expectedSummaries: []*DailyPortfolioSummary{
				{
					Date:       time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
					DailyValue: 1100,
					DailyCost:  1001,
				},
				{
					Date:       time.Date(2023, 10, 11, 0, 0, 0, 0, time.UTC),
					DailyValue: 1100,
					DailyCost:  1001,
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Running test case: %s\n", tt.name)
			fmt.Printf("Transactions: %+v\n", tt.transactions)
			result, err := CalculatePortfolioValueAndCost(tt.transactions, tt.marketData, tt.targetCurrencyCode, tt.dateRange)

			if (err != nil) != tt.expectError {
				t.Errorf("unexpected error outcome: got %v, expected %v", err != nil, tt.expectError)
			}

			for _, summary := range result {
				fmt.Printf("Summary: %+v\n", summary)
			}

			if !tt.expectError {
				if len(result) != len(tt.expectedSummaries) {
					t.Errorf("unexpected number of daily summaries: got %v, expected %v", len(result), len(tt.expectedSummaries))
					return
				}
				for i, summary := range result {
					expected := tt.expectedSummaries[i]
					if !summary.Date.Equal(expected.Date) || summary.DailyValue != expected.DailyValue || summary.DailyCost != expected.DailyCost {
						t.Errorf("unexpected summary at index %d: got %+v, expected %+v", i, summary, expected)
					}
				}
			}
		})
	}
}
