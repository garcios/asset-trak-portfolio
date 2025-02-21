package model

type BalanceSummary struct {
	AssetSymbol  string
	AssetName    string
	Quantity     float64
	Price        float64
	CurrencyCode string
	MarketCode   string
}
