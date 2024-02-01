package api

import (
	"errors"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	filesystem "github.com/gabrimatx/WasaPhoto/service/filesystem"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("Bad id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := rt.db.GetUserIdFromPhotoId(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Bad database fetch")
		w.WriteHeader(http.StatusInternalServerError)
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

	// delete photo from filesystem
	err = filesystem.RemovePhoto(id)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete photo from filesystem")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// delete photo from database
	err = rt.db.DeletePhoto(id)
	if errors.Is(err, components.ErrObjNotExists) {
		ctx.Logger.Error("Photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// cascade deletion of comments and likes
	err = rt.db.PhotoCascadeDeletion(id)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete photo's comments and likes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
