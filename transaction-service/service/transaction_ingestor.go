package service

import (
	"github.com/garcios/asset-trak-portfolio/transactions-service/db"
	"github.com/garcios/asset-trak-portfolio/transactions-service/model"
	"log"
)

type ITransactionManager interface {
	AddTransaction(rec *model.Transaction) error
}

// verify interface compliance
var _ ITransactionManager = &db.TransactionRepository{}

type TransactionIngestor struct {
	TransactionManager ITransactionManager
}

func NewTransactionIngestor(tm ITransactionManager) *TransactionIngestor {
	return &TransactionIngestor{TransactionManager: tm}
}

func (ingestor *TransactionIngestor) ProcessTransactions(tabName string, skipRows int) error {
	log.Println("Processing transactions...")

	rows, err := getRows(tabName)
	if err != nil {
		return err
	}

	var rowCount int
	for _, row := range rows {
		if rowCount < skipRows {
			rowCount++
			continue
		}

		if len(row) == 0 || row[0] == "Total" {
			break
		}

		transaction, err := mapColumnsToTransaction(row)
		if err != nil {
			return err
		}

		err = ingestor.addTransaction(transaction)
		if err != nil {
			return err
		}

	}

	return nil
}

func mapColumnsToTransaction(row []string) (*model.Transaction, error) {

	return nil, nil
}

func (ingestor *TransactionIngestor) addTransaction(rec *model.Transaction) error {

	// TODO:
	// Insert asset if it does not exist yet.

	// The following steps should be atomic.
	// 1. insert into transaction table
	// 2. insert/update asset balance
	//     - if asset for account_id already exists, insert; otherwise update balance.

	ingestor.TransactionManager.AddTransaction(rec)

	return nil
}
