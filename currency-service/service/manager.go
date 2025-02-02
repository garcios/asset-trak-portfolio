package service

import (
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"github.com/garcios/asset-trak-portfolio/currency-service/model"
	"time"
)

type ICurrencyManager interface {
	AddCurrencyRate(rec *model.CurrencyRate) error
	Truncate() error
	GetExchangeRate(fromCurrency string, toCurrency string, tradeDate time.Time) (float64, error)
}

// verify interface compliance
var _ ICurrencyManager = (*db.CurrencyRepository)(nil)
