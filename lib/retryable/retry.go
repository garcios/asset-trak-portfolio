package retryable

import (
	"context"
	"fmt"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/errors"
	"net/http"
	"strings"
	"time"
)

const (
	defaultRetryCount = 3
	retryDelay        = 1 * time.Second // Retry delay between attempts
)

// RetryableClient implements a client with custom retry logic
type RetryableClient struct {
	Client     client.Client
	MaxRetries int
	RetryDelay time.Duration
}

func (r *RetryableClient) Init(option ...client.Option) error {
	return r.Client.Init(option...)
}

func (r *RetryableClient) Options() client.Options {
	return r.Client.Options()
}

func (r *RetryableClient) NewMessage(topic string, msg interface{}, opts ...client.MessageOption) client.Message {
	return r.Client.NewMessage(topic, msg, opts...)
}

func (r *RetryableClient) NewRequest(service, endpoint string, req interface{}, reqOpts ...client.RequestOption) client.Request {
	return r.Client.NewRequest(service, endpoint, req, reqOpts...)
}

func (r *RetryableClient) Stream(ctx context.Context, req client.Request, opts ...client.CallOption) (client.Stream, error) {
	return r.Client.Stream(ctx, req, opts...)
}

func (r *RetryableClient) Publish(ctx context.Context, msg client.Message, opts ...client.PublishOption) error {
	return r.Client.Publish(ctx, msg, opts...)
}

func (r *RetryableClient) String() string {
	return r.Client.String()
}

func (r *RetryableClient) Call(
	ctx context.Context,
	req client.Request,
	rsp interface{},
	opts ...client.CallOption,
) error {
	var retryCount int
	if r.MaxRetries == 0 {
		r.MaxRetries = defaultRetryCount
	}

	if r.RetryDelay == 0 {
		r.RetryDelay = retryDelay
	}

	// Retry loop
	for {
		// Perform the actual request
		err := r.Client.Call(ctx, req, rsp, opts...)
		if err == nil {
			return nil // No error, return success
		}

		// Increment retry count
		retryCount++

		if retryCount > r.MaxRetries { //exceed retry count
			return err
		}

		// Parse and evaluate the error for retry conditions
		parsedErr := errors.Parse(err.Error())
		if !shouldRetry(parsedErr) {
			return err // Error not retryable or parsable
		}

		// Perform retry with delay
		time.Sleep(r.RetryDelay)
		fmt.Printf("Retrying request: %+v (attempt: %d)\n", req, retryCount)
	}
}

// Determines whether the request should be retried based on the parsed error
func shouldRetry(err *errors.Error) bool {
	if err == nil {
		return false // Cannot parse error, no retry
	}
	// Retry for timeout or specific service not found error which might be due to service discovery.
	return err.Code == http.StatusRequestTimeout ||
		(err.Code == http.StatusInternalServerError && strings.Contains(err.Detail, "service: not found"))
}
