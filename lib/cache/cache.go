package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type ICacheClient interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

// GetCachedValue retrieves a cached value or computes and caches it if not present.
func GetCachedValue[T any](
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
