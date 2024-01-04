package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != BEAR {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := authParts[1]

	if token != strconv.FormatUint(id, 10) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var photoStream components.PhotoList
	photoStream, err = rt.db.GetUserStream(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	photoStreamJSON, err := json.Marshal(photoStream)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(photoStreamJSON)
	w.WriteHeader(http.StatusOK)
}
