package logic

import (
	"context"
	"errors"

	"github.com/nguyenhoang711/downloader/internal/configs"
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data string, hashed string) (bool, error)
}

type hashService struct {
	accountConfig configs.Account
}

func NewHash(accConfig configs.Account) Hash {
	return &hashService{
		accountConfig: accConfig,
	}
}

// Hash implements Hash.
func (h *hashService) Hash(ctx context.Context, data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), h.accountConfig.HashCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// isHashEqual implements Hash.
func (h *hashService) IsHashEqual(ctx context.Context, data string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}
	return true, nil
}
