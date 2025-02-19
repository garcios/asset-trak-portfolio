package services

import (
	"context"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/retryable"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

type ITransactionService interface {
	GetHoldingsSummary(ctx context.Context, accountID string) (*pb.BalanceSummaryResponse, error)
}

type TransactionService struct {
	grpcTransactionService pb.TransactionService
}

func NewTransactionService() ITransactionService {
	// Define a custom client wrapper to leverage retry on error
	customRetryWrapper := func(c client.Client) client.Client {
		return &retryable.RetryableClient{Client: c}
	}

	serviceClient := micro.NewService(
		micro.Name("transaction-service.client"),
		micro.WrapClient(customRetryWrapper),
	)
	serviceClient.Init()

	return &TransactionService{
		grpcTransactionService: pb.NewTransactionService("transaction-service", serviceClient.Client()),
	}
}

func (t TransactionService) GetHoldingsSummary(
	ctx context.Context,
	accountID string,
) (*pb.BalanceSummaryResponse, error) {

	req := &pb.BalanceSummaryRequest{
		AccountId: accountID,
	}

	resp, err := t.grpcTransactionService.GetHoldings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get balance summary error: %w", err)
	}

	return resp, nil
}
