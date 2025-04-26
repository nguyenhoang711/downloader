package database

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	TabNameAccountPasswords = goqu.T("account_passwords")
)

const (
	ColNameAccountPasswordsOfAccountID = "of_account_id"
	ColNameAccountPasswordsHash        = "hash"
)

type AccountPassword struct {
	OfAccountID uint64 `db:"of_account_id" goqu:"skipinsert,skipupdate"`
	Hash        string `db:"hash"`
}

type AccountPasswordDataAccessor interface {
	CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	GetAccountPassword(ctx context.Context, ofAccountID uint64) (AccountPassword, error)
	UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error
	WithDatabase(database Database) AccountPasswordDataAccessor
}

type accountPasswordDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewAccountPasswordDataAccesor(
	database *goqu.Database,
	logger *zap.Logger,
) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		logger:   logger,
	}
}

// CreateAccountPassword implements AccountPasswordDataAccessor.
func (a accountPasswordDataAccessor) CreateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	log := utils.LoggerWithContext(ctx, a.logger)
	_, err := a.database.
		Insert(TabNameAccountPasswords).
		Rows(goqu.Record{
			ColNameAccountPasswordsOfAccountID: accountPassword.OfAccountID,
			ColNameAccountPasswordsHash:        accountPassword.Hash,
		}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		log.With(zap.Error(err)).Error("fail to save account password")
		return status.Errorf(codes.Internal, "failed to create account password: %+v", err)
	}

	return nil
}

// UpdateAccountPassword implements AccountPasswordDataAccessor.
func (a accountPasswordDataAccessor) UpdateAccountPassword(ctx context.Context, accountPassword AccountPassword) error {
	logger := utils.LoggerWithContext(ctx, a.logger)

	_, err := a.database.
		Update(TabNameAccountPasswords).
		Set(goqu.Record{ColNameAccountPasswordsHash: accountPassword.Hash}).
		Where(goqu.Ex{ColNameAccountPasswordsOfAccountID: accountPassword.OfAccountID}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("fail to update password")
		return status.Errorf(codes.Internal, "failed to update account password: %+v", err)
	}
	return nil
}

// WithDatabase implements AccountPasswordDataAccessor.
func (a accountPasswordDataAccessor) WithDatabase(database Database) AccountPasswordDataAccessor {
	return &accountPasswordDataAccessor{
		database: database,
		logger:   a.logger,
	}
}

// GetAccountPassword implements AccountPasswordDataAccessor.
func (a accountPasswordDataAccessor) GetAccountPassword(
	ctx context.Context,
	ofAccountID uint64,
) (AccountPassword, error) {
	logger := utils.LoggerWithContext(ctx, a.logger).With(zap.Uint64("of_account_id", ofAccountID))

	accountPass := AccountPassword{}

	found, err := a.database.
		From(TabNameAccountPasswords).
		Where(goqu.Ex{
			ColNameAccountPasswordsOfAccountID: ofAccountID,
		}).
		ScanStructContext(ctx, &accountPass)

	if err != nil {
		logger.With(zap.Error(err)).Error("fail to get account password by account id")
		return AccountPassword{}, status.Errorf(codes.Internal, "failed to get password by account id: %+v", err)
	}
	if !found {
		logger.Warn("cannot get account password by account id")
		return AccountPassword{}, status.Errorf(codes.Internal, "cannot find password via account id: %+v", err)
	}

	return accountPass, nil
}
