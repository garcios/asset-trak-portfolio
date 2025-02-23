package service

import (
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"time"
)

type ICurrencyManager interface {
	GetExchangeRate(fromCurrency string, toCurrency string, tradeDate time.Time) (float64, error)
}

// verify interface compliance
var _ ICurrencyManager = (*db.CurrencyRepository)(nil)
