package api

import (
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	followId, err := strconv.ParseUint(ps.ByName("followId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ans := CheckIdAuthorized(r, userId)
	if ans != 0 {
		if ans == 2 {
			ctx.Logger.WithField("id", userId).Error("Can't authorize user")
			w.WriteHeader(http.StatusForbidden)
		} else {
			ctx.Logger.WithField("id", userId).Error("Auth header invalid")
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	err = rt.db.FollowUser(followId, userId)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", followId).Error("Can't follow user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
