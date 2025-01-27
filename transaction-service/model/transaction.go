package model

import "time"

type Transaction struct {
	ID              string
	AccountID       string
	AssetID         string
	TransactionType string
	TransactionDate *time.Time
	Quantity        int
	Price           float64
	CurrencyCode    string
}
