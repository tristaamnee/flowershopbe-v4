package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func ComputeHmac256(data string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
