package logic

import (
	"context"
	"errors"

	"github.com/nguyenhoang711/downloader/internal/configs"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Hash interface {
	Hash(ctx context.Context, data string) (string, error)
	IsHashEqual(ctx context.Context, data string, hashed string) (bool, error)
}

type hashService struct {
	authConfig configs.Auth
}

func NewHash(authConfig configs.Auth) Hash {
	return &hashService{
		authConfig: authConfig,
	}
}

// Hash implements Hash.
func (h *hashService) Hash(_ context.Context, data string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(data), h.authConfig.Hash.Cost)
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to hash data: %+v", err)
	}
	return string(hashed), nil
}

// isHashEqual implements Hash.
func (h *hashService) IsHashEqual(_ context.Context, data string, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, status.Errorf(codes.Internal, "failed to check if data equal hash: %+v", err)
	}
	return true, nil
}
