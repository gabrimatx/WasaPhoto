package api

import (
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract parameters from the request path
	userIDStr := ps.ByName("userId")
	followIDStr := ps.ByName("followId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		// Handle the error, for example, by returning a bad request response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid userID"))
		return
	}

	followID, err := strconv.Atoi(followIDStr)
	if err != nil {
		// Handle the error, for example, by returning a bad request response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid followID"))
		return
	}
	err = rt.db.FollowUser(followID, userID)
	if err != nil {
		// Handle the error, for example, by returning an error response
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error following user"))
		return
	}
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Followed user successfully"))
}
