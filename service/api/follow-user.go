package api

import (
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract parameters
	userIDStr := ps.ByName("userId")
	followIDStr := ps.ByName("followId")

	//Extract userId
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid userID"))
		return
	}

	//Extract followId
	followID, err := strconv.Atoi(followIDStr)
	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid followID"))
		return
	}

	//Database logic
	err = rt.db.FollowUser(followID, userID)
	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Error following user"))
		return
	}

	//Write output
	w.Header().Set("content-type", "text/plain")
	_, _ = w.Write([]byte("Followed user successfully"))
}
