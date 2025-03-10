package service

import (
	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	"github.com/garcios/asset-trak-portfolio/currency-service/model"
	"time"
)

type ICurrencyRepository interface {
	GetExchangeRate(fromCurrency string, toCurrency string, tradeDate time.Time) (float64, error)
	GeExchangeRates(
		fromCurrency string,
		toCurrency string,
		startDate string,
		endDate string,
	) ([]*model.CurrencyRate, error)
}

// verify interface compliance
var _ ICurrencyRepository = (*db.CurrencyRepository)(nil)
