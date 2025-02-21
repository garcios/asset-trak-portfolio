package services

import (
	"context"
	"fmt"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"time"
)

type IPortfolioService interface {
	GetHoldingsSummary(ctx context.Context, accountID string) (*pb.HoldingsResponse, error)
}

type PortfolioService struct {
	grpcPortfolioService pb.PortfolioService
}

func NewPortfolioService() IPortfolioService {
	// Define a custom client wrapper to leverage retry on error
	customRetryWrapper := func(c client.Client) client.Client {
		return &lib.RetryableClient{
			Client:     c,
			MaxRetries: 3,
			RetryDelay: 1 * time.Second,
		}
	}

	serviceClient := micro.NewService(
		micro.Name("portfolio-service.client"),
		micro.WrapClient(customRetryWrapper),
	)
	serviceClient.Init()

	return &PortfolioService{
		grpcPortfolioService: pb.NewPortfolioService("portfolio-service", serviceClient.Client()),
	}
}

func (t PortfolioService) GetHoldingsSummary(
	ctx context.Context,
	accountID string,
) (*pb.HoldingsResponse, error) {

	req := &pb.HoldingsRequest{
		AccountId: accountID,
	}

	resp, err := t.grpcPortfolioService.GetHoldings(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get holdings summary error: %w", err)
	}

	return resp, nil
}
