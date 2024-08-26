package session

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

func encodeBase64URL(data string) string {
	base64Str := base64.StdEncoding.EncodeToString([]byte(data))
	urlSafeStr := strings.ReplaceAll(base64Str, "+", "-")
	urlSafeStr = strings.ReplaceAll(urlSafeStr, "/", "_")
	urlSafeStr = strings.TrimRight(urlSafeStr, "=")
	return urlSafeStr
}

func hmacSHA256(data, secret []byte) []byte {
	hash := hmac.New(sha256.New, secret)
	hash.Write(data)
	return hash.Sum(nil)
}

func generateJWT(secret string, data interface{}) string {
	header := `{"alg":"HS256","typ":"JWT"}`

	payload := fmt.Sprintf(`%v`, data)

	headerEncoded := encodeBase64URL(header)
	payloadEncoded := encodeBase64URL(payload)

	signature := hmacSHA256([]byte(headerEncoded+"."+payloadEncoded), []byte(secret))
	signatureEncoded := encodeBase64URL(string(signature))

	return headerEncoded + "." + payloadEncoded + "." + signatureEncoded
}
