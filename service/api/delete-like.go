package api

import (
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.WithField("id", photoId).Error("Can't find photo")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithField("id", userId).Error("Can't find user")
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

	err = rt.db.DeleteLike(photoId, userId)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", photoId).Error("Can't unlike photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = rt.db.DecrementLikeCount(photoId)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", photoId).Error("Can't decrement like counter of photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
