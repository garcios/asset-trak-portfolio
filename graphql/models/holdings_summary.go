package models

type HoldingsSummary struct {
	Investments []*Investment
}

type Investment struct {
	AssetSymbol  string
	AssetName    string
	MarketCode   string
	Price        *Money
	Quantity     float64
	Value        *Money
	CapitalGain  *MoneyWithPercentage
	Dividend     *MoneyWithPercentage
	CurrencyGain *MoneyWithPercentage
	TotalReturn  *MoneyWithPercentage
}

type Money struct {
	Amount       float64
	CurrencyCode string
}

type MoneyWithPercentage struct {
	Amount       float64
	CurrencyCode string
	Percentage   float64
}
