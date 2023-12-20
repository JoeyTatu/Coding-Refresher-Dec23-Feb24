package helpers

import (
	"crypto/rand"
)

func GenerateSecretKey(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	key := make([]byte, length)

	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}

	for i := range key {
		key[i] = charset[int(key[i])%len(charset)]
	}

	return string(key), nil
}
