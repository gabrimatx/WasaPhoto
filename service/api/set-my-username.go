package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't change username")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ans := CheckIdAuthorized(r, id)
	if ans != 0 {
		if ans == 2 {
			ctx.Logger.WithField("id", id).Error("Can't authorize user")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			ctx.Logger.WithField("id", id).Error("Auth header invalid")
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	var username string
	err = json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		ctx.Logger.WithField("new username", username).Error("invalid json encoding")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.SetUsername(id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Database error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
