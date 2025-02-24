package service

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

func (m *MockRedisClient) Get(ctx context.Context, key string) (string, error) {
	val, found := m.data[key]
	if !found {
		return "", redis.Nil // Mimic Redis "key not found" behavior
	}

	return val, nil
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	m.data[key] = value
	return nil
}
