package post_feed

import (
	"context"

	"monolith/internal/models"
)

type feedGetter interface {
	Get(ctx context.Context, userID string, limit int64, offset int64) ([]models.Post, error)
}
