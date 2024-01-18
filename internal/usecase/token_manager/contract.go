package token_manager

import "time"

type tokenCreator interface {
	CreateToken(privateKey string, ttl time.Duration, payload interface{}) (string, error)
}
