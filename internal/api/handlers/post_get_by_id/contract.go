package post_get_by_id

import (
	"context"

	"monolith/internal/models"
)

type postGetter interface {
	GetByID(ctx context.Context, postID string) (*models.Post, error)
}
