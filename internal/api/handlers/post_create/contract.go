package post_create

import (
	"context"

	"monolith/internal/models"
)

type postCreator interface {
	Create(ctx context.Context, post models.Post) (string, error)
}
