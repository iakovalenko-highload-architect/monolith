package user_get_by_id

import (
	"context"

	"monolith/internal/models"
)

type userGetter interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
}
