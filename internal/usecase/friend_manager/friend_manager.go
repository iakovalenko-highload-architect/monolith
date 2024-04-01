package friend_manager

import (
	"context"

	"github.com/go-faster/errors"

	"monolith/internal/models"
)

type FriendManager struct {
	storage storage
}

func New(storage storage) *FriendManager {
	return &FriendManager{
		storage: storage,
	}
}

func (f *FriendManager) SetFriend(ctx context.Context, userID string, friendID string) error {
	if err := f.storage.InsertFriendship(ctx, userID, friendID); err != nil {
		return errors.Wrap(err, "insert friend error")
	}
	return nil
}

func (f *FriendManager) DeleteFriendship(ctx context.Context, userID string, friendID string) error {
	if err := f.storage.DeleteFriendship(ctx, userID, friendID); err != nil {
		return errors.Wrap(err, "delete friend error")
	}
	return nil
}

func (f *FriendManager) GetFriends(ctx context.Context, userID string) ([]models.Friendship, error) {
	friendIDs, err := f.storage.FindFriendshipByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "find friend ids by user id error")
	}
	return friendIDs, nil
}

func (f *FriendManager) GetAllFriends(ctx context.Context) ([]models.Friendship, error) {
	friends, err := f.storage.GetAllFriends(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "find friends by user id error")
	}
	return friends, nil
}
