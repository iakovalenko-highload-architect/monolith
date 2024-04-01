package post_delete

import (
	"context"

	"monolith/internal/models"
)

type postDeleter interface {
	Delete(ctx context.Context, post models.Post) error
}
