package auth_manager

import (
	"context"

	"monolith/internal/models"
)

type userStorage interface {
	Insert(ctx context.Context, user models.User) (string, error)
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
}

type passwordManager interface {
	Compare(hash string, secret string) bool
	Generate(secret string) (string, error)
}

type tokenManager interface {
	CreateAuthToken(userID string) (string, error)
}
