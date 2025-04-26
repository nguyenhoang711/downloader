package logic

import (
	"context"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/cache"
	"github.com/nguyenhoang711/downloader/internal/dataaccess/database"
	go_load "github.com/nguyenhoang711/downloader/internal/generated/go_load/v1"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	AccountName string
	Password    string
}

type User struct {
	ID       uint64
	Username string
}

type Session struct {
	ID       uint64
	Username string
}

type CreateSessionOutput struct {
	Account *go_load.Account
	Token   string
}

type Account interface {
	CreateAccount(context.Context, CreateAccountParams) (CreateAccountOutput, error)
	CreateSession(context.Context, CreateSessionParams) (CreateSessionOutput, error)
}

type account struct {
	takenAccountNameCache cache.TakenAccountName
	goquDatabase          *goqu.Database
	accDataAccessor       database.AccountDataAccessor
	accPassDataAccessor   database.AccountPasswordDataAccessor
	hashLogic             Hash
	tokenLogic            Token
	logger                *zap.Logger
}

func NewAccount(
	goquDatabase *goqu.Database,
	accDataAccessor database.AccountDataAccessor,
	accPassDataAccessor database.AccountPasswordDataAccessor,
	hashLogic Hash,
	tokenLogic Token,
	takenNameCache cache.TakenAccountName,
	logger *zap.Logger,
) Account {
	return &account{
		accDataAccessor:       accDataAccessor,
		accPassDataAccessor:   accPassDataAccessor,
		hashLogic:             hashLogic,
		goquDatabase:          goquDatabase,
		tokenLogic:            tokenLogic,
		takenAccountNameCache: takenNameCache,
		logger:                logger,
	}
}

func (a account) isAccountnameTaken(ctx context.Context, accountName string) (bool, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.String("account_name", accountName))

	// get account name from cache
	accountNameTaken, err := a.takenAccountNameCache.Has(ctx, accountName)
	if err != nil {
		logger.With(zap.Error(err)).Warn("failed to get account name from taken set in cache, will fall back to database")
	} else if accountNameTaken {
		return true, nil
	}

	// query to database
	_, err = a.accDataAccessor.GetAccountByAccountName(ctx, accountName)
	if err != nil {
		if errors.Is(err, database.ErrAccountNotFound) {
			return false, nil
		}
		return false, err
	}

	if err := a.takenAccountNameCache.Add(ctx, accountName); err != nil {
		logger.With(zap.Error(err)).Warn("failed to set account name into taken set in cache")
	}

	return true, nil
}

// CreateAccount implements Account.
func (a account) CreateAccount(ctx context.Context, params CreateAccountParams) (CreateAccountOutput, error) {
	var accountID uint64
	accountNameTaken, err := a.isAccountnameTaken(ctx, params.AccountName)
	if err != nil {
		return CreateAccountOutput{}, status.Error(codes.Internal, "failed to check if account name is taken")
	}

	if accountNameTaken {
		return CreateAccountOutput{}, status.Error(codes.AlreadyExists, "account name is already taken")
	}
	txErr := a.goquDatabase.WithTx(func(td *goqu.TxDatabase) error {
		accountID, err = a.accDataAccessor.CreateAccount(ctx, database.Account{
			AccountName: params.AccountName,
		})
		if err != nil {
			return err
		}

		// create hashed password
		hashPassword, hashErr := a.hashLogic.Hash(ctx, params.Password)
		if hashErr != nil {
			return hashErr
		}
		// create database interface for general use
		if err := a.accPassDataAccessor.WithDatabase(td).CreateAccountPassword(ctx, database.AccountPassword{
			OfAccountID: accountID,
			Hash:        hashPassword,
		}); err != nil {
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
func (a account) databaseAccountToProtoAccount(account database.Account) *go_load.Account {
	return &go_load.Account{
		Id:          account.ID,
		AccountName: account.AccountName,
	}
}

// CreateSession implements Account.
func (a account) CreateSession(ctx context.Context, params CreateSessionParams) (CreateSessionOutput, error) {
	// check tai khoan ton tai
	existingAccount, err := a.accDataAccessor.GetAccountByAccountName(ctx, params.AccountName)
	if err != nil {
		return CreateSessionOutput{}, err
	}
	// kiem tra password ton tai
	existingAccPass, err := a.accPassDataAccessor.GetAccountPassword(ctx, existingAccount.ID)
	if err != nil {
		return CreateSessionOutput{}, err
	}
	// kiem tra hash bang nhau ko
	isHashEqual, err := a.hashLogic.IsHashEqual(ctx, params.Password, existingAccPass.Hash)
	if err != nil {
		return CreateSessionOutput{}, err
	}

	if !isHashEqual {
		return CreateSessionOutput{}, status.Error(codes.Unauthenticated, "incorrect password")
	}

	token, _, err := a.tokenLogic.GetToken(ctx, existingAccount.ID)
	if err != nil {
		return CreateSessionOutput{}, err
	}
	return CreateSessionOutput{
		Account: a.databaseAccountToProtoAccount(existingAccount),
		Token:   token,
	}, nil
}
