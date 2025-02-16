package services

import (
	"context"
	"fmt"
	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
	"go-micro.dev/v4"
)

type ITransactionService interface {
	GetHoldingsSummary(ctx context.Context, accountID string) (*pb.BalanceSummaryResponse, error)
}

type TransactionService struct {
}

func NewTransactionService() ITransactionService {
	return &TransactionService{}
}

func (t TransactionService) GetHoldingsSummary(ctx context.Context,
	accountID string,
) (*pb.BalanceSummaryResponse, error) {
	resolverClient := micro.NewService(micro.Name("transaction-resolver.client"))
	resolverClient.Init()

	transactionSrv := pb.NewTransactionService("transaction-service", resolverClient.Client())

	req := &pb.BalanceSummaryRequest{
		AccountId: accountID,
	}

	resp, err := transactionSrv.GetBalanceSummary(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("get balance summary error: %w", err)
	}

	return resp, nil
}
