package dialog_user_id_send

import (
	"context"

	"monolith/internal/models"
)

type dialogManager interface {
	SendMessage(ctx context.Context, message models.Message) error
}
