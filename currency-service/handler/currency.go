package handler

import (
	"context"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/currency-service/service"
	"log"
	"time"
)

func New(currencyManager service.ICurrencyManager) *Currency {
	return &Currency{
		currencyManager: currencyManager,
	}
}

type Currency struct {
	currencyManager service.ICurrencyManager
}

func (h *Currency) GetExchangeRate(
	_ context.Context,
	in *pb.GetExchangeRateRequest,
	out *pb.GetExchangeRateResponse,
) error {
	log.Println("handler.GetExchangeRate...")
	tradeDate, err := time.Parse("2006-01-02", in.GetTradeDate())
	if err != nil {
		return err
	}

	exchangeRate, err := h.currencyManager.GetExchangeRate(
		in.GetFromCurrency(),
		in.GetToCurrency(),
		tradeDate)
	if err != nil {
		return err
	}

	out.ExchangeRate = exchangeRate

	log.Printf("out: %#v\n", out)

	return nil
}

func (h *Currency) GetHistoricalExchangeRates(
	ctx context.Context,
	in *pb.GetHistoricalExchangeRatesRequest,
	out *pb.GetHistoricalExchangeRatesResponse,
) error {

	return nil
}
