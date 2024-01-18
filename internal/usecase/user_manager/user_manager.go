package user_manager

import (
	"context"

	uErr "monolith/internal/errors"
	"monolith/internal/models"
)

type User struct {
	userStorage userStorage
}

func New(userStorage userStorage) *User {
	return &User{
		userStorage: userStorage,
	}
}

func (u *User) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := u.userStorage.FindByUserID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, uErr.UserNotFoundErr
	}

	return user, nil
}
