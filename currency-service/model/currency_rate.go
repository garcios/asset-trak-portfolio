package model

import "time"

type CurrencyRate struct {
	ID             int        `json:"id"`
	BaseCurrency   string     `json:"base_currency"`
	TargetCurrency string     `json:"target_currency"`
	ExchangeRate   float64    `json:"exchange_rate"`
	TradeDate      *time.Time `json:"trade_date"`
}
