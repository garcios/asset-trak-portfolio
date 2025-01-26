package service

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transactions-service/db"
	"github.com/garcios/asset-trak-portfolio/transactions-service/model"
	"github.com/google/uuid"
	"log"
)

type ITransactionManager interface {
	AddTransaction(rec *model.Transaction) error
}

type IAccountManager interface {
	FindAccountByID(id string) (*model.Account, error)
}

// verify interface compliance
var _ ITransactionManager = &db.TransactionRepository{}
var _ IAccountManager = &db.AccountRepository{}

type TransactionIngestor struct {
	TransactionManager ITransactionManager
	AccountManager     IAccountManager
}

func NewTransactionIngestor(tm ITransactionManager, am IAccountManager) *TransactionIngestor {
	return &TransactionIngestor{
		TransactionManager: tm,
		AccountManager:     am,
	}
}

func (ingestor *TransactionIngestor) ProcessTransactions(
	filePath,
	tabName string,
	skipRows int,
	accountID string,
) error {
	log.Println("Processing transactions...")

	// verify account
	account, err := ingestor.AccountManager.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	if account == nil {
		return fmt.Errorf("account with ID %s not found", accountID)
	}

	rows, err := getRows(filePath, tabName)
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

		// populate IDs
		transaction.ID = uuid.New().String()
		transaction.AccountID = accountID

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
