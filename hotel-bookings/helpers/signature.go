package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSignature(data, secretKey string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifySignature(data, signature, secretKey string) bool {
	internalSignature := GenerateSignature(data, secretKey)
	return hmac.Equal([]byte(internalSignature), []byte(signature))
}
