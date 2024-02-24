package dialog_manager

import (
	"context"

	"monolith/internal/clients/dialog"
	"monolith/internal/models"
)

type userStorage interface {
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
}

type dialogClient interface {
	SendMessage(ctx context.Context, msg dialog.Message) error
	Get(ctx context.Context, fromID string, toID string) ([]dialog.Message, error)
}
