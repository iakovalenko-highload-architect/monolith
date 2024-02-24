package dialog

import (
	"context"

	"github.com/go-faster/errors"

	dto "monolith/internal/generated/rpc/clients/dialog"
)

func (c *Client) SendMessage(ctx context.Context, msg Message) error {
	_, err := c.cli.Create(ctx, &dto.CreateRequest{
		FromUserID: msg.FromID,
		ToUserID:   msg.ToID,
		Text:       msg.Text,
	})
	if err != nil {
		return errors.Wrap(err, "dialog client err")
	}

	return nil
}

func (c *Client) Get(ctx context.Context, fromID string, toID string) ([]Message, error) {
	res, err := c.cli.Get(ctx, &dto.GetRequest{
		FromUserID: fromID,
		ToUserID:   toID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "dialog client err")
	}
	messages := make([]Message, 0, len(res.Messages))
	for _, msg := range res.Messages {
		if msg == nil {
			continue
		}

		messages = append(messages, Message{
			FromID: msg.FromUserID,
			ToID:   msg.ToUserID,
			Text:   msg.Text,
		})
	}

	return messages, nil
}
