package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/garcios/asset-trak-portfolio/transaction-service/model"
)

type AccountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{
		DB: db,
	}
}

func (r *AccountRepository) FindAccountByID(id string) (*model.Account, error) {
	stmt, err := r.DB.Prepare("SELECT id, name, email FROM account WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, fmt.Errorf("FindAccountByID: %v", err)
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	if row == nil {
		return nil, nil
	}

	account := new(model.Account)

	err = row.Scan(&account.ID, &account.Name, &account.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("FindAccountByID: %v", err)
	}

	return account, nil
}
