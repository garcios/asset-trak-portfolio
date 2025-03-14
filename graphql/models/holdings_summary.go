package models

type HoldingsSummary struct {
	Investments []*Investment
}

type Investment struct {
	AssetSymbol  string
	AssetName    string
	MarketCode   string
	Price        *Money
	Weight       float64
	Quantity     float64
	Value        *Money
	Cost         *Money
	CapitalGain  *MoneyWithPercentage
	Dividend     *MoneyWithPercentage
	CurrencyGain *MoneyWithPercentage
	TotalReturn  *MoneyWithPercentage
}

type Money struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}

type MoneyWithPercentage struct {
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
	Percentage   float64 `json:"percentage"`
}

type SummaryTotals struct {
	PortfolioValue *Money               `json:"portfolioValue"`
	CapitalGain    *MoneyWithPercentage `json:"capitalGain"`
	Dividends      *MoneyWithPercentage `json:"dividends"`
	CurrencyGain   *MoneyWithPercentage `json:"currencyGain"`
	TotalReturn    *MoneyWithPercentage `json:"totalReturn"`
}
