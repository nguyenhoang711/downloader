package logic

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
)

type CreateAccountParams struct {
	AccountName string
	Password    string
}

type CreateAccountOutput struct {
	AccountName string
	ID          uint64
}

type CreateSessionParams struct {
	Username string
	Password string
}

type User struct {
	ID       uint64
	Username string
}

type Session struct {
	ID       uint64
	Username string
}

type Account interface {
	CreateAccount(context.Context, CreateAccountParams) (CreateAccountOutput, error)
	CreateSession(context.Context, CreateSessionParams) (Session, error)
}

type account struct {
	goquDatabase        *goqu.Database
	accDataAccessor     database.AccountDataAccessor
	accPassDataAccessor database.AccountPasswordDataAccessor
	hashLogic           Hash
}

func NewAccount(
	goquDatabase *goqu.Database,
	accDataAccessor database.AccountDataAccessor,
	accPassDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash,
) Account {
	return &account{
		accDataAccessor:     accDataAccessor,
		accPassDataAccessor: accPassDataAccessor,
		hashLogic:           hashLogic,
		goquDatabase:        goquDatabase,
	}
}

func (a account) isAccountnameTaken(ctx context.Context, accountName string) (bool, error) {
	if _, err := a.accDataAccessor.GetAccountByAccountName(ctx, accountName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateAccount implements Account.
func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	var accountID uint64
	txErr := a.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		isAccountNameTaken, err := a.isAccountnameTaken(ctx, params.AccountName)
		if err != nil {
			return nil
		}

		if isAccountNameTaken {
			return errors.New("accountname is already taken")
		}

		accountID, err = a.accDataAccessor.CreateAccount(ctx, database.Account{
			AccountName: params.AccountName,
		})
		if err != nil {
			return err
		}

		// create hashed password
		encryptedPass, err := a.hashLogic.Hash(ctx, params.Password)
		if err != nil {
			log.Printf("error when encrypt password, err=%+v\n", err)
			return err
		}
		// create database interface for general use
		if err := a.accPassDataAccessor.WithDatabase(td).CreateAccountPassword(ctx, database.AccountPassword{
			OfAccountID: accountID,
			Hash:        encryptedPass,
		}); err != nil {
			log.Printf("create password is wrong with err=%+v\n", err)
			return err
		}
		return nil
	})

	if txErr != nil {
		return CreateAccountOutput{}, txErr
	}

	return CreateAccountOutput{
		ID:          accountID,
		AccountName: params.AccountName,
	}, nil
}

// CreateSession implements Account.
func (a *account) CreateSession(context.Context, CreateSessionParams) (Session, error) {
	panic("unimplemented")
}
