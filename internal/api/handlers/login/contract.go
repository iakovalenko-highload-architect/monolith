package login

import (
	"context"

	loginUsecase "monolith/internal/usecase/auth_manager"
)

type login interface {
	Login(ctx context.Context, req loginUsecase.LoginRequest) (string, error)
}
