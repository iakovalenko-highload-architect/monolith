package cache

import (
	"encoding/json"

	"monolith/internal/models"
)

const CacheLen = 1000

type post struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

func exported(internal post) models.Post {
	return models.Post{
		ID:     internal.ID,
		UserID: internal.UserID,
		Text:   internal.Text,
	}
}

func imported(external models.Post) post {
	return post{
		ID:     external.ID,
		UserID: external.UserID,
		Text:   external.Text,
	}
}

func encode(p post) ([]byte, error) {
	return json.Marshal(p)
}

func decode(b string) (post, error) {
	var p post
	if err := json.Unmarshal([]byte(b), &p); err != nil {
		return post{}, err
	}
	return p, nil
}
