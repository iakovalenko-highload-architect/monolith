package post_update

import (
	"context"

	"monolith/internal/models"
)

type postUpdater interface {
	Update(ctx context.Context, post models.Post) error
}
