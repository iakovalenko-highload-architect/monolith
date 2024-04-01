package friend_delete_by_id

import (
	"context"
)

type friendDeleter interface {
	DeleteFriendship(ctx context.Context, userID string, friendID string) error
}
