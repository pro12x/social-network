package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ExtractIDFromRequest(r *http.Request) (uint, error) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 6 {
		return 0, fmt.Errorf("ID is missing in parameters")
	}

	idStr := pathSegments[len(pathSegments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}
	return uint(id), nil
}

func ExtractIdAndPrivacyFromRequest(r *http.Request) (uint, string, error) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 6 {
		return 0, "", fmt.Errorf("privacy is missing in parameters")
	}

	params := strings.Split(pathSegments[len(pathSegments)-1], "&")
	if len(params) < 2 || len(strings.TrimSpace(params[1])) == 0 {
		return 0, "", fmt.Errorf("privacy is missing in parameters")
	}

	if strings.TrimSpace(params[1]) != "public" && strings.TrimSpace(params[1]) != "private" && strings.TrimSpace(params[1]) != "almost_private" {
		return 0, "", fmt.Errorf("invalid privacy value")
	}

	id, err := strconv.Atoi(params[0])
	if err != nil {
		return 0, "", fmt.Errorf("invalid ID format")
	}

	return uint(id), params[1], nil
}

func GenerateToken() (string, error) {
	bytes := make([]byte, 16) // Adjust the length as needed
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
