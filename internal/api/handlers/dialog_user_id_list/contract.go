package dialog_user_id_list

import (
	"context"

	"monolith/internal/models"
)

type dialogManager interface {
	GetDialog(ctx context.Context, fromID string, toID string) ([]models.Message, error)
}
