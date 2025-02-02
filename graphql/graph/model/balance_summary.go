package model

type BalanceSummary struct {
	AccountID    string
	TotalValue   *Money
	BalanceItems []*BalanceItem
}

type BalanceItem struct {
	AssetSymbol string
	AssetName   string
	Price       *Money
	Quantity    float64
	Value       *Money
	TotalGain   float64
	MarketCode  string
}

type Money struct {
	Amount       float64
	CurrencyCode string
}
