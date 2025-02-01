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

	res.AccountId = req.GetAccountId()
	res.BalanceItems = make([]*pb.BalanceItem, 0)

	res.BalanceItems = append(res.BalanceItems,
		&pb.BalanceItem{
			AssetSymbol: "AMZN",
			Amount: &pb.Money{
				Amount:       1000,
				CurrencyCode: "USD",
			},
		},
		&pb.BalanceItem{
			AssetSymbol: "MSFT",
			Amount: &pb.Money{
				Amount:       3000,
				CurrencyCode: "USD",
			},
		},
	)

	return nil
}
