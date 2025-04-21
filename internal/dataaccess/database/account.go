package database

import (
	"context"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	TabNameAccounts = goqu.T("accounts")
)

const (
	ColNameAccountsID          = "id"
	ColNameAccountsAccountName = "account_name"
)

type Account struct {
	ID          uint64 `sql:"id"`
	AccountName string `sql:"account_name"`
}

type AccountDataAccessor interface {
	CreateAccount(context.Context, Account) (uint64, error)
	GetAccountByID(context.Context, uint64) (Account, error)
	GetAccountByAccountName(ctx context.Context, accname string) (Account, error)
	WithDatabase(database Database) AccountDataAccessor
}

type accountDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewAccountDataAccessor(
	database *goqu.Database,
	logger *zap.Logger,
) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
		logger:   logger,
	}
}

// CreateAccount implements AccountDataAccessor.
func (a accountDataAccessor) CreateAccount(ctx context.Context, acc Account) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, a.logger)

	result, err := a.database.
		Insert(TabNameAccounts).
		Rows(goqu.Record{
			ColNameAccountsAccountName: acc.AccountName,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		log.Printf("failed to create account, err=%+v\n", err)
		logger.With(zap.Error(err)).Error("failed to create account")
		return 0, status.Errorf(codes.Internal, "failed to create account: %+v", err)
	}
	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get last inserted id")
		return 0, status.Errorf(codes.Internal, "failed to get account id: %+v", err)
	}
	return uint64(lastInsertedID), err
}

// GetAccountByAccountName implements AccountDataAccessor.
func (a accountDataAccessor) GetAccountByAccountName(ctx context.Context, accname string) (Account, error) {
	logger := utils.LoggerWithContext(ctx, a.logger)

	account := Account{}
	found, err := a.database.
		From(TabNameAccounts).
		Where(goqu.Ex{ColNameAccountsAccountName: accname}).
		ScanStructContext(ctx, &account)
	if err != nil {
		logger.With(zap.Error(err)).Error("fail to get account by acc_name")
	}

	if !found {
		logger.Warn("cannot find account with account name")
		return Account{}, status.Errorf(codes.Internal, "failed to get account by name: %+v", err)
	}

	return account, nil
}

// GetAccountByID implements AccountDataAccessor.
func (a accountDataAccessor) GetAccountByID(ctx context.Context, acc_id uint64) (Account, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Uint64("of_account_id", acc_id))

	account := Account{}
	found, err := a.database.
		From(TabNameAccounts).
		Where(goqu.Ex{ColNameAccountsID: acc_id}).
		ScanStructContext(ctx, &account)
	if err != nil {
		logger.With(zap.Error(err)).Error("fail to get account by account id")
	}

	if !found {
		logger.Warn("cannot find account with account id")
		return Account{}, status.Errorf(codes.Internal, "failed to get account by id: %+v", err)
	}

	return account, nil
}

// WithDatabase implements AccountDataAccessor.
func (a *accountDataAccessor) WithDatabase(database Database) AccountDataAccessor {
	return &accountDataAccessor{
		database: database,
	}
}
