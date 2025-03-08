package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/garcios/asset-trak-portfolio/currency-service/db"
	pb "github.com/garcios/asset-trak-portfolio/currency-service/proto"
	"github.com/garcios/asset-trak-portfolio/currency-service/service"
)

const DateFormat = "2006-01-02"

func New(currencyManager service.ICurrencyManager) *Currency {
	return &Currency{
		currencyManager: currencyManager,
	}
}

type Currency struct {
	currencyManager service.ICurrencyManager
}

// extractGetExchangeRate extracts the exchange rate between two currencies for a given trade date, considering
// the possibility of reversing the pair if the original lookup fails.
func (h *Currency) extractGetExchangeRate(
	from string,
	to string,
	tradeDate time.Time,
) (float64, error) {
	exchangeRate, err := h.currencyManager.GetExchangeRate(from, to, tradeDate)
	if err != nil {
		if errors.Is(err, db.NotFound) { // Attempt reverse exchange rate
			log.Printf("Exchange rate not found for %s -> %s, trying reverse...", from, to)
			reverseRate, reverseErr := h.currencyManager.GetExchangeRate(to, from, tradeDate)
			if reverseErr != nil {
				return 0, reverseErr
			}

			return 1 / reverseRate, nil
		}

		return 0, err
	}

	return exchangeRate, nil
}

func (h *Currency) GetExchangeRate(
	_ context.Context,
	in *pb.GetExchangeRateRequest,
	out *pb.GetExchangeRateResponse,
) error {
	log.Println("handler.GetExchangeRate...")
	// Parse trade date
	tradeDate, err := time.Parse(DateFormat, in.GetTradeDate())
	if err != nil {
		return fmt.Errorf("invalid trade date format: %w", err)
	}

	// Fetch the exchange rate
	exchangeRate, err := h.extractGetExchangeRate(
		in.GetFromCurrency(),
		in.GetToCurrency(),
		tradeDate,
	)
	if err != nil {
		return err
	}

	// Set the response
	out.ExchangeRate = exchangeRate

	return nil
}

func (h *Currency) GetHistoricalExchangeRates(
	ctx context.Context,
	in *pb.GetHistoricalExchangeRatesRequest,
	out *pb.GetHistoricalExchangeRatesResponse,
) error {

	return nil
}
