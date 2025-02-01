package handler

import (
	"context"
	"fmt"
	"go-micro.dev/v4/client"

	pb "github.com/garcios/asset-trak-portfolio/transaction-service/proto"
)

func New(c client.Client) *Transaction {
	return &Transaction{}
}

type Transaction struct {
}

func (h *Transaction) GetBalanceSummary(
	ctx context.Context,
	req *pb.BalanceSummaryRequest, res *pb.BalanceSummaryResponse) error {

	fmt.Println("GetBalanceSummary...")

	return nil
}
