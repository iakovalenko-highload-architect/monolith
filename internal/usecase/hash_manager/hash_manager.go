package hash_manager

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

type HashManager struct {
	config Config
}

func New(config Config) *HashManager {
	return &HashManager{
		config: config,
	}
}

func (h *HashManager) Generate(secret string) (string, error) {
	salt := make([]byte, h.config.SaltLen)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hashedSecret := argon2.IDKey(
		[]byte(secret),
		salt,
		h.config.Times,
		h.config.Memory,
		h.config.Threads,
		h.config.KeyLen,
	)
	return string(append(salt, hashedSecret...)), nil
}

func (h *HashManager) Compare(hash string, secret string) bool {
	hashBytes := []byte(hash)
	salt := hashBytes[0:h.config.SaltLen]
	hashedSecret := argon2.IDKey(
		[]byte(secret),
		salt,
		h.config.Times,
		h.config.Memory,
		h.config.Threads,
		h.config.KeyLen,
	)
	return bytes.Equal(hashedSecret, hashBytes[h.config.SaltLen:])
}
