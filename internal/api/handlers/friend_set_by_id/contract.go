package friend_set_by_id

import (
	"context"
)

type friendSetter interface {
	SetFriend(ctx context.Context, userID string, friendID string) error
}
