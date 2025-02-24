package model

import "time"

type Transaction struct {
	ID                      string
	AccountID               string
	AssetID                 string
	TransactionType         string
	TransactionDate         *time.Time
	Quantity                float64
	TradePrice              float64
	TradePriceCurrencyCode  string
	BrokerageFee            float64
	FeeCurrencyCode         string
	AmountCash              float64
	AmountCurrencyCode      string
	ExchangeRate            float64
	WithheldTaxAmount       float64
	WithheldTaxCurrencyCode string
}
