package auth_manager

import (
	"context"

	"github.com/go-faster/errors"

	uErr "monolith/internal/errors"
	"monolith/internal/models"
)

type AuthManager struct {
	userStorage     userStorage
	passwordManager passwordManager
	tokenManager    tokenManager
}

func New(
	userStorage userStorage,
	passwordManager passwordManager,
	tokenManager tokenManager,
) *AuthManager {
	return &AuthManager{
		userStorage:     userStorage,
		passwordManager: passwordManager,
		tokenManager:    tokenManager,
	}
}

func (l *AuthManager) Login(ctx context.Context, req LoginRequest) (string, error) {
	userData, err := l.userStorage.FindByUserID(ctx, req.UserID)
	if err != nil {
		return "", errors.Wrap(err, "find user by id error")
	}
	if userData == nil {
		return "", uErr.UserNotFoundErr
	}

	if !l.passwordManager.Compare(userData.Password, req.Password) {
		return "", uErr.WrongPasswordErr
	}

	token, err := l.tokenManager.CreateAuthToken(userData.ID)
	if err != nil {
		return "", errors.Wrap(err, "create auth token error")
	}

	return token, nil
}

func (u *AuthManager) Register(ctx context.Context, user models.User) (string, error) {
	hashedPassword, err := u.passwordManager.Generate(user.Password)
	if err != nil {
		return "", errors.Wrap(err, "generate password hash error")
	}
	user.Password = hashedPassword

	userID, err := u.userStorage.Insert(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "save user in db error")
	}

	return userID, nil
}
