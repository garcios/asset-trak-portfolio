package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type MockRedisClient struct {
	data map[string]string
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{
		data: make(map[string]string),
	}
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	// Attempt to retrieve the value from the mock storage
	val, found := m.data[key]
	cmd := redis.NewStringCmd(ctx, key)

	if !found {
		// Set the Redis command result to nil to mimic Redis behavior for a missing key
		cmd.Val()
		cmd.SetErr(redis.Nil) // `redis.Nil` is used to indicate "key not found"
		return cmd
	}

	// Set the actual value for the command and return
	cmd.SetVal(val)
	return cmd
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	// Safely convert value to string, mimicking Redis behavior
	strValue, ok := value.(string)
	if !ok {
		return redis.NewStatusCmd(ctx, "ERR value is not a string")
	}

	m.data[key] = strValue
	return redis.NewStatusCmd(ctx, "OK")
}
