package dialog_manager

import (
	"context"

	"github.com/go-faster/errors"

	"monolith/internal/clients/dialog"
	uErr "monolith/internal/errors"
	"monolith/internal/models"
)

type DialogManager struct {
	userStorage  userStorage
	dialogClient dialogClient
}

func New(userStorage userStorage, dialogClient dialogClient) *DialogManager {
	return &DialogManager{
		userStorage:  userStorage,
		dialogClient: dialogClient,
	}
}

func (d *DialogManager) SendMessage(ctx context.Context, message models.Message) error {
	recipient, err := d.userStorage.FindByUserID(ctx, message.ToID)
	if err != nil {
		return errors.Wrap(err, "recipient does not exist")
	}
	if recipient == nil {
		return errors.Wrap(uErr.UserNotFoundErr, "recipient not found")
	}

	if err := d.dialogClient.SendMessage(ctx, dialog.Message{
		FromID: message.FromID,
		ToID:   message.ToID,
		Text:   message.Text,
	}); err != nil {
		return errors.Wrap(err, "send message error")
	}

	return nil
}

func (d *DialogManager) GetDialog(ctx context.Context, fromID string, toID string) ([]models.Message, error) {
	recipient, err := d.userStorage.FindByUserID(ctx, toID)
	if err != nil {
		return nil, errors.Wrap(err, "recipient does not exist")
	}
	if recipient == nil {
		return nil, errors.Wrap(uErr.UserNotFoundErr, "recipient not found")
	}

	res, err := d.dialogClient.Get(ctx, fromID, toID)
	if err != nil {
		return nil, errors.Wrap(err, "get dialog error")
	}

	messages := make([]models.Message, 0, len(res))
	for _, msg := range res {
		messages = append(messages, models.Message{
			FromID: msg.FromID,
			ToID:   msg.ToID,
			Text:   msg.Text,
		})
	}

	return messages, nil
}
