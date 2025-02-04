package finance

import (
	"math"
	"testing"
)

func TestCalculateAveragePrice(t *testing.T) {
	tests := []struct {
		name   string
		trades []*Trade
		want   float64
	}{
		{
			name: "Single Trade",
			trades: []*Trade{
				{Quantity: 100, Price: 1.0, Commission: 0.0},
			},
			want: 1.0,
		},
		{
			name: "Multiple Trades",
			trades: []*Trade{
				{Quantity: 100, Price: 1.0, Commission: 0.0},
				{Quantity: 200, Price: 1.5, Commission: 0.0},
			},
			want: 1.33333,
		},
		{
			name:   "No Trades",
			trades: []*Trade{},
			want:   0,
		},
		{
			name: "Zero Quantity",
			trades: []*Trade{
				{Quantity: 0, Price: 1.0, Commission: 0.0},
			},
			want: 0,
		},
		{
			name: "Single Trade With Commission",
			trades: []*Trade{
				{Quantity: 100, Price: 1.0, Commission: 50.0},
			},
			want: 1.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAveragePrice(tt.trades)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("CalculateAveragePrice() = %.5f, want %.5f", got, tt.want)
			}
		})
	}
}

func TestCalculateProfitLoss(t *testing.T) {
	tests := []struct {
		name         string
		totalShares  int
		totalCost    float64
		currentPrice float64
		want         float64
	}{
		{"Profit scenario", 10, 100.0, 15.0, 50.0},
		{"Loss scenario", 10, 150.0, 10.0, -50.0},
		{"Break-even scenario", 10, 100.0, 10.0, 0.0},
		{"Zero shares", 0, 150.0, 10.0, -150.0},
		{"Negative price", 10, 150.0, -10.0, -250.0},
		{"Zero cost", 10, 0.0, 10.0, 100.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateProfitLoss(tt.totalShares, tt.totalCost, tt.currentPrice)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("failed %s: expected %v but got %v", tt.name, tt.want, got)
			}
		})
	}
}

func TestCalculateTotalCost(t *testing.T) {
	tests := []struct {
		name   string
		trades []Trade
		want   float64
	}{
		{
			name:   "no trade",
			trades: []Trade{},
			want:   0,
		},
		{
			name: "single trade",
			trades: []Trade{
				{Quantity: 10, Price: 30.0, Commission: 3.0},
			},
			want: 303.0,
		},
		{
			name: "multiple trades",
			trades: []Trade{
				{Quantity: 10, Price: 25.0, Commission: 2.0},
				{Quantity: 5, Price: 35.0, Commission: 1.0},
			},
			want: 428.0,
		},
		{
			name: "commission not included in some trades",
			trades: []Trade{
				{Quantity: 10, Price: 20.0},
				{Quantity: 5, Price: 30.0, Commission: 2.0},
			},
			want: 352.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateTotalCost(tt.trades)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("CalculateTotalCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputeNewAveragePrice(t *testing.T) {
	tests := []struct {
		name          string
		oldAvgPrice   float64
		oldShares     int
		newShares     int
		newPrice      float64
		newCommission float64
		want          float64
	}{
		{
			name:          "Base case",
			oldAvgPrice:   207.67,
			oldShares:     30,
			newShares:     15,
			newPrice:      215.0,
			newCommission: 15.0,
			want:          210.44666666666663,
		},
		{
			name:          "Zero old shares",
			oldAvgPrice:   0,
			oldShares:     0,
			newShares:     50,
			newPrice:      20,
			newCommission: 10,
			want:          20.2,
		},
		{
			name:          "Zero new shares",
			oldAvgPrice:   10,
			oldShares:     100,
			newShares:     0,
			newPrice:      0,
			newCommission: 0,
			want:          10,
		},
		{
			name:          "Zero new and old shares",
			oldAvgPrice:   0,
			oldShares:     0,
			newShares:     0,
			newPrice:      0,
			newCommission: 0,
			want:          0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateNewAveragePrice(tt.oldAvgPrice, tt.oldShares, tt.newShares, tt.newPrice, tt.newCommission)
			if math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("CalculateNewAveragePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePercentageReturn(t *testing.T) {
	var tests = []struct {
		name         string
		totalShares  int
		totalCost    float64
		currentPrice float64
		want         float64
	}{
		{
			name:         "Basic Positive Test",
			totalShares:  10,
			totalCost:    100.0,
			currentPrice: 20.0,
			want:         100.0,
		},
		{
			name:         "Loss Test",
			totalShares:  5,
			totalCost:    100.0,
			currentPrice: 10.0,
			want:         -50.0,
		},
		{
			name:         "Zero Shares Test",
			totalShares:  0,
			totalCost:    200.0,
			currentPrice: 20.0,
			want:         0.0,
		},
		{
			name:         "Zero Total Cost Test",
			totalShares:  5,
			totalCost:    0.0,
			currentPrice: 100.0,
			want:         0.0,
		},
		{
			name:         "Floating Point Precision Test",
			totalShares:  3,
			totalCost:    17.25,
			currentPrice: 5.65,
			want:         -1.7391304347826,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePercentageReturn(tt.totalShares, tt.totalCost, tt.currentPrice)
			if math.Abs(got-tt.want) > 1e-8 {
				t.Errorf("CalculatePercentageReturn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatePortfolioReturn(t *testing.T) {
	testCases := []struct {
		name        string
		trades      []Trade
		stockPrices map[string]float64
		want        float64
	}{
		{
			name:        "empty trades",
			trades:      []Trade{},
			stockPrices: map[string]float64{},
			want:        0,
		},
		{
			name: "single trade",
			trades: []Trade{
				{
					Quantity: 1,
					Price:    100,
					Ticker:   "GOOGL",
				},
			},
			stockPrices: map[string]float64{
				"GOOGL": 150,
			},
			want: 50,
		},
		{
			name: "multiple trades",
			trades: []Trade{
				{
					Quantity: 1,
					Price:    100,
					Ticker:   "GOOGL",
				},
				{
					Quantity: 2,
					Price:    50,
					Ticker:   "AAPL",
				},
			},
			stockPrices: map[string]float64{
				"GOOGL": 150,
				"AAPL":  75,
			},
			want: 50,
		},
		{
			name: "multiple trades with different prices",
			trades: []Trade{
				{
					Quantity: 1,
					Price:    100,
					Ticker:   "GOOGL",
				},
				{
					Quantity: 2,
					Price:    50,
					Ticker:   "AAPL",
				},
			},
			stockPrices: map[string]float64{
				"GOOGL": 120,
				"AAPL":  60,
			},
			want: 20,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatePortfolioReturn(tt.trades, tt.stockPrices)
			if math.Abs(got-tt.want) > 1e-8 {
				t.Fatalf("CalculatePortfolioReturn() output = %.2f; want %.2f", got, tt.want)
			}
		})
	}
}
