package models

type PerformanceData struct {
	TradeDate    string  `json:"tradeDate"`
	Amount       float64 `json:"amount"`
	CurrencyCode string  `json:"currencyCode"`
}
