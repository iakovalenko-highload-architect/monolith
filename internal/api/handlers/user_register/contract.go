package user_register

import (
	"context"

	"monolith/internal/models"
)

type authManager interface {
	Register(ctx context.Context, user models.User) (string, error)
}
