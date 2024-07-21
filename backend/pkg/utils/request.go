package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func ExtractIDFromRequest(r *http.Request) (int, error) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 6 {
		return 0, fmt.Errorf("ID is missing in parameters")
	}

	idStr := pathSegments[len(pathSegments)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format")
	}
	return id, nil
}

func GenerateToken() (string, error) {
	bytes := make([]byte, 16) // Adjust the length as needed
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
