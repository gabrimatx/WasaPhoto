package api

import (
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract parameters
	userID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

	if err != nil {
		w.Header().Set("content-type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid userID"))
		return
	}

	followID, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)

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
