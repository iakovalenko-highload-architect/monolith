package friendship

import "monolith/internal/models"

type Friendship struct {
	UserID   string `db:"user_id"`
	FriendID string `db:"friend_id"`
}

func Exported(internal Friendship) models.Friendship {
	return models.Friendship{
		UserID:   internal.UserID,
		FriendID: internal.FriendID,
	}
}

func Imported(external models.Friendship) Friendship {
	return Friendship{
		UserID:   external.UserID,
		FriendID: external.FriendID,
	}
}
