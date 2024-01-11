package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("user not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ans := CheckIdAuthorized(r, id)
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

	var photoStream components.PhotoStreamList
	photoStream, err = rt.db.GetUserStream(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during stream getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoStreamJSON, err := json.Marshal(photoStream)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(photoStreamJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
