package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
)

type Account struct {
	AccountID   uint64 `sql:"account_id"`
	AccountName string `sql:"accountname"`
}

type AccountDataAccessor interface {
	CreateAccount(context.Context, Account) error
	GetAccountByID(context.Context, uint64) (Account, error)
	GetAccountByAccountName(ctx context.Context, accname string) (Account, error)
}

type accountDataAccessor struct {
	database *goqu.Database
}

func NewAccountDataAccessor(
	database *goqu.Database,
) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}

// CreateAccount implements AccountDataAccessor.
func (a *accountDataAccessor) CreateAccount(ctx context.Context, acc Account) error {
	if _, err := a.database.
		Insert("accounts").
		Rows(goqu.Record{
			"accountname": acc.AccountName,
		}).
		Executor().
		ExecContext(ctx); err != nil {
		log.Printf("failed to create account, err= %+v", err)
	}
	return nil
}

// GetAccountByAccountName implements AccountDataAccessor.
func (a *accountDataAccessor) GetAccountByAccountName(ctx context.Context, accname string) (Account, error) {
	panic("unimplemented")
}

// GetAccountByID implements AccountDataAccessor.
func (a *accountDataAccessor) GetAccountByID(context.Context, uint64) (Account, error) {
	panic("unimplemented")
}
