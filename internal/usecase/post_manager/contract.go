package post_manager

import (
	"context"

	"monolith/internal/models"
)

type storage interface {
	InsertPost(ctx context.Context, post models.Post) (string, error)
	UpdatePost(ctx context.Context, post models.Post) error
	DeletePost(ctx context.Context, postID string) error
	FindByPostID(ctx context.Context, postID string) (*models.Post, error)
	FindPostsByUserID(ctx context.Context, userID string, limit int64) ([]models.Post, error)
}

type cache interface {
	Append(ctx context.Context, userID string, post models.Post) error
	Clear(ctx context.Context, userID string) error
	Update(ctx context.Context, userID string, post models.Post) error
	Delete(ctx context.Context, userID string, post models.Post) error
}

type friendGetter interface {
	GetFriends(ctx context.Context, userID string) ([]models.Friendship, error)
	GetAllFriends(ctx context.Context) ([]models.Friendship, error)
}
