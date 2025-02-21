package retryable

import (
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"time"
)

func CreateRetryableClient(
	serviceName string,
	opts ...Option,
) micro.Service {
	// Default options
	retryOpts := &RetryOptions{
		MaxRetries: 3,
		RetryDelay: 1 * time.Second,
	}

	// Apply user-provided options
	for _, opt := range opts {
		opt(retryOpts)
	}

	customRetryWrapper := func(c client.Client) client.Client {
		return &RetryableClient{
			Client:     c,
			MaxRetries: retryOpts.MaxRetries,
			RetryDelay: retryOpts.RetryDelay,
		}
	}

	// Create a new service
	return micro.NewService(
		micro.Name(serviceName),
		micro.WrapClient(customRetryWrapper),
	)

}

// Options for creating a retryable client
type RetryOptions struct {
	MaxRetries int
	RetryDelay time.Duration
}

// Functional option pattern for customizable config
type Option func(*RetryOptions)

// Option to set the MaxRetries
func WithMaxRetries(maxRetries int) Option {
	return func(opts *RetryOptions) {
		opts.MaxRetries = maxRetries
	}
}

// Option to set the RetryDelay
func WithRetryDelay(retryDelay time.Duration) Option {
	return func(opts *RetryOptions) {
		opts.RetryDelay = retryDelay
	}
}
