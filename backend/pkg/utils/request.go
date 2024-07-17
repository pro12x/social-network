package utils

import (
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
