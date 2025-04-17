package database

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nguyenhoang711/downloader/internal/utils"
	"go.uber.org/zap"
)

const (
	TabNameTokenPublicKeys          = "token_public_keys"
	ColNameTokenPublicKeysID        = "id"
	ColNameTokenPublicKeysPublicKey = "public_key"
)

type TokenPublicKey struct {
	ID        uint64 `sql:"id"`
	PublicKey []byte `sql:"public_key"`
}

type TokenPublicKeyDataAccessor interface {
	CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error)
	GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error)
	WithDatabase(database Database) TokenPublicKeyDataAccessor
}

type tokenPublicKeyDataAccessor struct {
	database Database
	logger   *zap.Logger
}

func NewTokenPublicKeyDataAccessor(
	database *goqu.Database,
	logger *zap.Logger,
) TokenPublicKeyDataAccessor {
	return &tokenPublicKeyDataAccessor{
		database: database,
		logger:   logger,
	}
}

// CreatePublicKey implements TokenPublicKeyDataAccessor.
func (t tokenPublicKeyDataAccessor) CreatePublicKey(ctx context.Context, tokenPublicKey TokenPublicKey) (uint64, error) {
	logger := utils.LoggerWithContext(ctx, t.logger)
	result, err := t.database.
		Insert(TabNameTokenPublicKeys).
		Rows(goqu.Record{ColNameTokenPublicKeysPublicKey: tokenPublicKey.PublicKey}).
		Executor().
		ExecContext(ctx)
	if err != nil {
		logger.With(zap.Error(err)).Error("fail to create public key")
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get inserted id")
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

// GetPublicKey implements TokenPublicKeyDataAccessor.
func (t tokenPublicKeyDataAccessor) GetPublicKey(ctx context.Context, id uint64) (TokenPublicKey, error) {
	logger := utils.LoggerWithContext(ctx, t.logger).With(zap.Uint64("id", id))

	tokenPublicKey := TokenPublicKey{}
	found, err := t.database.Select().From(TabNameTokenPublicKeys).Where(goqu.Ex{
		ColNameTokenPublicKeysID: id,
	}).
		Executor().
		ScanStructContext(ctx, &tokenPublicKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("failed to get public key")
		return TokenPublicKey{}, err
	}

	if !found {
		logger.Warn("public key not found")
		return TokenPublicKey{}, sql.ErrNoRows
	}

	return tokenPublicKey, nil
}

// WithDatabase implements TokenPublicKeyDataAccessor.
func (t tokenPublicKeyDataAccessor) WithDatabase(database Database) TokenPublicKeyDataAccessor {
	t.database = database
	return t
}
