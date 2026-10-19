package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return strings.TrimSpace(parts[1]), nil
}

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("authorization header must use Bearer")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if token == "" {
		return "", errors.New("bearer token is missing")
	}

	return token, nil
}

func MakeRefreshToken() string {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(randomBytes)
}
