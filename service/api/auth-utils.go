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

func CheckValidAuth(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != BEAR || authParts[1] == "null" {
		return false
	}

	return true
}

func GetIdFromBearer(r *http.Request) uint64 {
	authHeader := r.Header.Get("Authorization")
	authParts := strings.Fields(authHeader)
	token := authParts[1]
	myId, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		return 100000
	}
	return myId
}
