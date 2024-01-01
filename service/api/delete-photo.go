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
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
