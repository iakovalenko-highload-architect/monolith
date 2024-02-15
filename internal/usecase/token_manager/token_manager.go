package token_manager

import (
	"fmt"
	"time"
)

const (
	TtlAccessTokenDefault = time.Hour * 24
)

type TokenManager struct {
	config       Config
	tokenCreator tokenCreator
}

func New(tokenCreator tokenCreator, config Config) *TokenManager {
	return &TokenManager{
		tokenCreator: tokenCreator,
		config:       config,
	}
}

func (m *TokenManager) CreateAuthToken(userID string) (string, error) {
	accessToken, err := m.tokenCreator.CreateToken(
		m.config.PrivateKey,
		m.config.TtlAccessToken,
		Data{
			UserID: userID,
		})
	if err != nil {
		return "", fmt.Errorf("failed create access token: %w", err)
	}

	return accessToken, nil
}
