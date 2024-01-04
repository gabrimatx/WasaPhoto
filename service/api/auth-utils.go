package api

import (
	"net/http"
	"strconv"
	"strings"
)

func CheckIdAuthorized(r *http.Request, id uint64) int {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 1
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != BEAR {
		return 1
	}

	token := authParts[1]

	if token != strconv.FormatUint(id, 10) {
		return 2
	}

	return 0
}
