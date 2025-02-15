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

type SummaryTotals struct {
	PortfolioValue *Money               `json:"portfolioValue"`
	CapitalGain    *MoneyWithPercentage `json:"capitalGain"`
	Dividends      *MoneyWithPercentage `json:"dividends"`
	CurrencyGain   *MoneyWithPercentage `json:"currencyGain"`
	TotalReturn    *MoneyWithPercentage `json:"totalReturn"`
}
