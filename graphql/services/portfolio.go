package services

import (
	"context"
	"fmt"
	lib "github.com/garcios/asset-trak-portfolio/lib/retryable"
	pb "github.com/garcios/asset-trak-portfolio/portfolio-service/proto"
	"time"
)

type IPortfolioService interface {
	GetHoldingsSummary(ctx context.Context, accountID string) (*pb.HoldingsResponse, error)
}

type PortfolioService struct {
	grpcPortfolioService pb.PortfolioService
}

func NewPortfolioService() IPortfolioService {
	serviceClient := lib.CreateRetryableClient(
		"portfolio-service.client",
		lib.WithMaxRetries(3),
		lib.WithRetryDelay(1*time.Second),
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
