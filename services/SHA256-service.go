package services

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

func SHA256Encoder(s string) (string, error) {
	if s == "" {
		return "", errors.New("cannot encode an empty string")
	}

	pass := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", pass), nil
}
