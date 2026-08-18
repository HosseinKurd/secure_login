package security

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateSessionID() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
