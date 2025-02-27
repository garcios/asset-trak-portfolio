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
	GetSummaryTotals(ctx context.Context, accountID string) (*pb.SummaryTotalsResponse, error)
	GetPerformanceHistory(
		ctx context.Context,
		accountID string,
		startDate string,
		endDate string,
	) (*pb.PerformanceHistoryResponse, error)
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

func (t PortfolioService) GetSummaryTotals(ctx context.Context, accountID string) (*pb.SummaryTotalsResponse, error) {

	req := &pb.SummaryTotalsRequest{
		AccountId: accountID,
	}

	resp, err := t.grpcPortfolioService.GetSummaryTotals(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get summary totals error: %w", err)
	}

	return resp, nil
}

func (t PortfolioService) GetPerformanceHistory(
	ctx context.Context,
	accountID string,
	startDate string,
	endDate string,
) (*pb.PerformanceHistoryResponse, error) {

	req := &pb.PerformanceHistoryRequest{
		AccountId: accountID,
		StartDate: startDate,
		EndDate:   endDate,
	}

	resp, err := t.grpcPortfolioService.GetPerformanceHistory(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get summary totals error: %w", err)
	}

	return resp, nil
}
