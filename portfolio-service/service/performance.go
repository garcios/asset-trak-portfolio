package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	apm "github.com/garcios/asset-trak-portfolio/asset-price-service/model"
	"github.com/garcios/asset-trak-portfolio/lib/finance"
	"github.com/go-redis/redis/v8"
)

const (
	cacheExpiration = 24 * time.Hour
)

type ICacheClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type PerformanceService struct {
	cacheClient ICacheClient
}

func NewPerformanceService(cacheClient ICacheClient) *PerformanceService {
	return &PerformanceService{
		cacheClient: cacheClient,
	}
}

// HistoricalRecord represents the value and cost for a specific day.
type HistoricalRecord struct {
	Date  time.Time
	Value float64 // Value in the target currency (based on current day price)
	Cost  float64 // Cost in the target currency (based on transaction day price)
}

// ExchangeRateFunc retrieves the exchange rate for a given currency and date.
type ExchangeRateFunc func(fromCurrency string, toCurrency string, date time.Time) (float64, error)

// AssetPriceFunc retrieves the price of an asset on a specific date.
type AssetPriceFunc func(assetID string, date time.Time) (*apm.AssetPrice, error)

// getCachedValue retrieves a cached value or computes and caches it if not present.
func getCachedValue[T any](
	ctx context.Context,
	cacheClient ICacheClient,
	key string,
	fetchFunc func() (T, error),
	expiration time.Duration,
) (T, error) {
	// Check Redis cache for the key
	val, err := cacheClient.Get(ctx, key).Result()
	if err == nil {
		// If found, deserialize and return cached value
		var result T
		err = deserialize(val, &result) // Implement your deserialization logic
		if err != nil {
			return result, fmt.Errorf("failed to deserialize cache value: %w", err)
		}
		return result, nil
	} else if !errors.Is(err, redis.Nil) {
		// Return error if Redis has other issues (not a cache miss)
		var zero T
		return zero, fmt.Errorf("failed to access Redis: %w", err)
	}

	// Cache miss: fetch the value and store it in Redis
	result, fetchErr := fetchFunc()
	if fetchErr != nil {
		// Return the error from the fetch function
		var zero T
		return zero, fetchErr
	}

	// Serialize value and store it in cache
	serializedValue, err := serialize(result) // Implement your serialization logic
	if err == nil {
		_ = cacheClient.Set(ctx, key, serializedValue, expiration)
	}

	return result, nil
}

// serialize converts a Go value to a JSON string, which can be stored in Redis.
func serialize[T any](value T) (string, error) {
	// Convert the Go value into a JSON string
	data, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("failed to serialize value: %w", err)
	}
	return string(data), nil
}

// deserialize converts a JSON string from Redis back into a Go value.
func deserialize[T any](data string, value *T) error {
	// Convert the JSON string back into the Go value
	if err := json.Unmarshal([]byte(data), value); err != nil {
		return fmt.Errorf("failed to deserialize value: %w", err)
	}
	return nil
}

// CalculateDailyHistoricalValueAndCost calculates the portfolio's daily historical value and cost
// across a date range while supporting multi-currency and fetching historical asset prices.
func (s PerformanceService) CalculateDailyHistoricalValueAndCost(
	ctx context.Context,
	trades []*finance.Trade,
	startDate, endDate time.Time,
	targetCurrency string,
	getExchangeRate ExchangeRateFunc,
	getAssetPrice AssetPriceFunc,
) ([]*HistoricalRecord, error) {
	// Initialize a map to store daily updates for portfolio value and cost
	dailyRecords := make(map[time.Time]*HistoricalRecord)

	// Iterate over each trade to update the portfolio data for the trade's date and subsequent days
	for _, trade := range trades {
		if trade == nil {
			continue
		}

		log.Printf("trade: %#v\n", trade)
		tradeDate := trade.TransactionDate // Assume Trade struct has a Date field of type time.Time
		if tradeDate.Before(startDate) || tradeDate.After(endDate) {
			continue // Skip trades outside the date range
		}

		currentDay := tradeDate
		for !currentDay.After(endDate) {
			if dailyRecords[currentDay] == nil {
				dailyRecords[currentDay] = &HistoricalRecord{
					Date:  currentDay,
					Value: 0,
					Cost:  0,
				}
			}

			// Define cache keys
			priceExchangeRateKey := fmt.Sprintf("exchangeRate:%s:%s:%s",
				trade.Price.CurrencyCode,
				targetCurrency,
				tradeDate.Format("2006-01-02"),
			)

			commissionExchangeRateKey := fmt.Sprintf("exchangeRate:%s:%s:%s",
				trade.Commission.CurrencyCode,
				targetCurrency,
				tradeDate.Format("2006-01-02"),
			)

			currentDayExchangeRateKey := fmt.Sprintf("exchangeRate:%s:%s:%s",
				trade.Price.CurrencyCode,
				targetCurrency,
				currentDay.Format("2006-01-02"),
			)

			assetPriceOnTradeDateKey := fmt.Sprintf("assetPrice:%s:%s",
				trade.AssetID,
				tradeDate.Format("2006-01-02"),
			)

			assetPriceOnCurrentDayKey := fmt.Sprintf("assetPrice:%s:%s",
				trade.AssetID,
				currentDay.Format("2006-01-02"),
			)

			// Get exchange rates and asset prices from cache (or compute and cache them)
			priceExchangeRateOnTradeDate, err := getCachedValue(
				ctx,
				s.cacheClient,
				priceExchangeRateKey,
				func() (float64, error) {
					return getExchangeRate(trade.Price.CurrencyCode, targetCurrency, tradeDate)
				},
				cacheExpiration,
			)
			if err != nil {
				return nil, err
			}

			commissionExchangeRateOnTradeDate, err := getCachedValue(
				ctx,
				s.cacheClient,
				commissionExchangeRateKey,
				func() (float64, error) {
					return getExchangeRate(trade.Commission.CurrencyCode, targetCurrency, tradeDate)
				},
				cacheExpiration,
			)
			if err != nil {
				return nil, err
			}

			// Get the exchange rate for the target currency for the current day (for value calculation)
			priceExchangeRateOnCurrentDay, err := getCachedValue(
				ctx,
				s.cacheClient,
				currentDayExchangeRateKey,
				func() (float64, error) {
					return getExchangeRate(trade.Price.CurrencyCode, targetCurrency, currentDay)
				},
				cacheExpiration,
			)
			if err != nil {
				return nil, err
			}

			// Get the asset price on the transaction date (for cost calculation)
			assetPriceOnTradeDate, err := getCachedValue(
				ctx,
				s.cacheClient,
				assetPriceOnTradeDateKey,
				func() (*apm.AssetPrice, error) {
					return getAssetPrice(trade.AssetID, tradeDate)
				}, cacheExpiration)
			if err != nil {
				return nil, err
			}

			if assetPriceOnTradeDate.Price == 0 {
				assetPriceOnTradeDate.Price = trade.Price.Amount // Fallback if no historical price is available
			}

			// Get the asset price for the current day (for value calculation)
			assetPriceOnCurrentDay, err := getCachedValue(
				ctx,
				s.cacheClient,
				assetPriceOnCurrentDayKey,
				func() (*apm.AssetPrice, error) {
					return getAssetPrice(trade.AssetID, currentDay)
				}, cacheExpiration)
			if err != nil {
				return nil, err
			}

			if assetPriceOnCurrentDay.Price == 0 {
				assetPriceOnCurrentDay.Price = trade.Price.Amount // Fallback if no price available
			}

			// Convert amounts to the target currency
			priceInTargetCurrencyForCost := assetPriceOnTradeDate.Price * priceExchangeRateOnTradeDate
			priceInTargetCurrencyForValue := assetPriceOnCurrentDay.Price * priceExchangeRateOnCurrentDay
			commissionInTargetCurrency := trade.Commission.Amount * commissionExchangeRateOnTradeDate

			// Update the value and cost based on trade details
			dailyRecord := dailyRecords[currentDay]

			// BUY transactions (Quantity > 0)
			if trade.Quantity > 0 {
				dailyRecord.Cost += trade.Quantity*priceInTargetCurrencyForCost + commissionInTargetCurrency
				dailyRecord.Value += trade.Quantity * priceInTargetCurrencyForValue
			}

			// SELL transactions (Quantity < 0)
			if trade.Quantity < 0 {
				sellQuantity := -trade.Quantity
				if dailyRecord.Value > 0 && sellQuantity > 0 {
					averageCost := dailyRecord.Cost / dailyRecord.Value // Average cost per unit
					dailyRecord.Cost -= sellQuantity * averageCost
					dailyRecord.Value -= sellQuantity * priceInTargetCurrencyForValue
				}
			}

			// Ensure final values are converted to two decimal place
			dailyRecord.Cost = finance.ToTwoDecimalPlaces(dailyRecord.Cost)
			dailyRecord.Value = finance.ToTwoDecimalPlaces(dailyRecord.Value)

			// Move to the next day
			currentDay = currentDay.AddDate(0, 0, 1)
		}
	}

	// Collect the results into an array sorted by date
	var result []*HistoricalRecord
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		if dailyRecords[d] != nil {
			result = append(result, dailyRecords[d])
		} else {
			// Add a record for dates without updates
			result = append(result, &HistoricalRecord{
				Date:  d,
				Value: 0,
				Cost:  0,
			})
		}
	}

	return result, nil
}
