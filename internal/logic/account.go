package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
)

type CreateAccountParams struct {
	Username string
	Password string
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
	CreateAccount(context.Context, CreateAccountParams) (User, error)
	CreateSession(context.Context, CreateSessionParams) (Session, error)
}

type account struct {
	accDataAccessor database.AccountDataAccessor
}

func NewAccount() Account {
	return &account{}
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
func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (User, error) {
	isAccountNameTaken, err := a.isAccountnameTaken(ctx, params.Username)
	if err != nil {
		return User{}, nil
	}

	if isAccountNameTaken {
		return User{}, errors.New("accountname is already taken")
	}

	if err := a.accDataAccessor.CreateAccount(ctx, database.Account{
		AccountName: params.Username,
	}); err != nil {
		return User{}, err
	}

	return User{}, nil
}

// CreateSession implements Account.
func (a *account) CreateSession(context.Context, CreateSessionParams) (Session, error) {
	panic("unimplemented")
}
