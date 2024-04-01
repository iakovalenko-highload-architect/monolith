package post

import (
	"monolith/internal/models"
)

type Post struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
	Text   string `db:"text_"`
}

func Exported(internal Post) models.Post {
	return models.Post{
		ID:     internal.ID,
		UserID: internal.UserID,
		Text:   internal.Text,
	}
}

func Imported(external models.Post) Post {
	return Post{
		ID:     external.ID,
		UserID: external.UserID,
		Text:   external.Text,
	}
}
