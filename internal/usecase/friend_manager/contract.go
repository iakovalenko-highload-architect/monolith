package friend_manager

import (
	"context"

	"monolith/internal/models"
)

type storage interface {
	InsertFriendship(ctx context.Context, userID string, friendID string) error
	DeleteFriendship(ctx context.Context, userID string, friendID string) error
	FindFriendshipByUserID(ctx context.Context, userID string) ([]models.Friendship, error)
	GetAllFriends(ctx context.Context) ([]models.Friendship, error)
}
