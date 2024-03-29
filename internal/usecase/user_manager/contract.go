package user_manager

import (
	"context"

	"monolith/internal/models"
)

type userStorage interface {
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
	Search(ctx context.Context, firstName string, secondName string) ([]models.User, error)
}
