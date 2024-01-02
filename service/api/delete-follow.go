package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	followId, err := strconv.ParseUint(ps.ByName("followId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	token := authParts[1]
	if token == fmt.Sprint(userId) {
		fmt.Fprint(w, "Access granted!")
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = rt.db.DeleteFollow(followId, userId)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", followId).Error("Can't unfollow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
