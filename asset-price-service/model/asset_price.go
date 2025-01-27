package model

import "time"

type AssetPrice struct {
	AssetID      string
	Price        float64
	CurrencyCode string
	TradeDate    *time.Time
}
