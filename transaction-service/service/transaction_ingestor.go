package service

import (
	"fmt"
	"github.com/garcios/asset-trak-portfolio/lib/excel"
	"github.com/garcios/asset-trak-portfolio/transaction-service/db"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	fieldLength = 8
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
	AssetManager       IAssetManager
}

func NewTransactionIngestor(
	tm ITransactionManager,
	am IAccountManager,
	astm IAssetManager,
) *TransactionIngestor {
	return &TransactionIngestor{
		TransactionManager: tm,
		AccountManager:     am,
		AssetManager:       astm,
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

	rows, err := excel.GetRows(filePath, tabName)
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

		transaction, err := ingestor.mapColumnsToTransaction(row)
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

func (ingestor *TransactionIngestor) mapColumnsToTransaction(row []string) (*model.Transaction, error) {
	if len(row) < fieldLength {
		log.Printf("row: %v", row)
		return nil, fmt.Errorf("row contains less than expected fields: %d", len(row))
	}

	assetSymbol := row[0]
	var assetID string
	a, err := ingestor.AssetManager.FindAssetBySymbol(assetSymbol)
	if err != nil {
		return nil, err
	}

	assetID = a.ID

	dateString := row[3]
	var transactionDate time.Time
	dateValue, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return nil, err
	}
	transactionDate = dateValue

	transactionType := normalizeString(row[4])

	quantityString := row[5]

	var quantity float64
	if quantityString != "" {
		quantity, err = strconv.ParseFloat(normalizeNumber(quantityString), 64)
		if err != nil {
			log.Printf("row: %v", row)
			return nil, err
		}
	}

	priceString := row[6]

	var price float64
	if priceString != "" {
		price, err = strconv.ParseFloat(normalizeNumber(priceString), 64)
		if err != nil {
			log.Printf("row: %v", row)
			return nil, err
		}
	}

	transaction := &model.Transaction{
		AssetID:         assetID,
		TransactionType: transactionType,
		TransactionDate: transactionDate,
		Quantity:        int(quantity),
		Price:           price,
		CurrencyCode:    row[7],
	}

	return transaction, nil
}

func (ingestor *TransactionIngestor) addTransaction(rec *model.Transaction) error {

	// TODO:
	// Insert asset if it does not exist yet.

	// The following steps should be atomic.
	// 1. insert into transaction table
	// 2. insert/update asset balance
	//     - if asset for account_id already exists, insert; otherwise update balance.

	err := ingestor.TransactionManager.AddTransaction(rec)
	if err != nil {
		return err
	}

	return nil
}

func normalizeString(s string) string {
	return strings.ToUpper(strings.TrimSpace(s))
}

func normalizeNumber(s string) string {
	return strings.Replace(s, ",", "", -1)
}
