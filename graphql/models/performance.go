package models

type PerformanceData struct {
	TradeDate    string  `json:"tradeDate"`
	Cost         float64 `json:"cost"`
	Value        float64 `json:"value"`
	CurrencyCode string  `json:"currencyCode"`
}
