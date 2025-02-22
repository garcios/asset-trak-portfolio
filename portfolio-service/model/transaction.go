package model

import "time"

type Transaction struct {
	ID                     string
	AccountID              string
	AssetID                string
	TransactionType        string
	TransactionDate        *time.Time
	Quantity               float64
	TradePrice             float64
	AssetPriceCurrencyCode string
	TradeCommission        float64
	CommissionCurrencyCode string
}
