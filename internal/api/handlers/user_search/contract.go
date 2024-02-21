package user_search

import (
	"context"

	"monolith/internal/models"
)

type userGetter interface {
	Search(ctx context.Context, firstName string, secondName string) ([]models.User, error)
}
