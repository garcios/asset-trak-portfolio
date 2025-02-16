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
	client.Client
}

func (r *RetryableClient) Call(
	ctx context.Context,
	req client.Request,
	rsp interface{},
	opts ...client.CallOption,
) error {
	var retryCount int

	// Retry loop
	for {
		// Perform the actual request
		err := r.Client.Call(ctx, req, rsp, opts...)
		if err == nil {
			return nil // No error, return success
		}

		// Increment retry count
		retryCount++

		if retryCount > defaultRetryCount { //exceed retry count
			return err
		}

		// Parse and evaluate the error for retry conditions
		parsedErr := errors.Parse(err.Error())
		if !shouldRetry(parsedErr) {
			return err // Error not retryable or parsable
		}

		// Perform retry with delay
		time.Sleep(retryDelay)
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
