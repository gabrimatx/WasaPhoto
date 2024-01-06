package api

import (
	"errors"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		ctx.Logger.WithField("id", id).Error("Can't find comment")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := rt.db.GetUserIdFromCommentId(id)
	if err != nil {
		ctx.Logger.WithField("id", id).Error("Comment not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ans := CheckIdAuthorized(r, userId)
	if ans != 0 {
		if ans == 2 {
			ctx.Logger.WithField("id", id).Error("Can't authorize user")
			w.WriteHeader(http.StatusForbidden)
		} else {
			ctx.Logger.WithField("id", id).Error("Auth header invalid")
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	err = rt.db.DeleteComment(id)
	if errors.Is(err, components.ErrObjNotExists) {
		ctx.Logger.WithField("id", id).Error("Comment not found in database")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
