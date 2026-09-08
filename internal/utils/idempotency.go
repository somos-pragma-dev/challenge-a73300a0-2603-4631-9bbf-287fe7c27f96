package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateIdempotencyKey() string {
	hash := sha256.New()
	hash.Write([]byte(time.Now().String()))
	return hex.EncodeToString(hash.Sum(nil))
}