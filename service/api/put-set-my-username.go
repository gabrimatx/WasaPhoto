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
		ctx.Logger.WithError(err).WithField("id", id).Error("User not found")
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

	var requestData struct {
		Username string `json:"username"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
	    ctx.Logger.WithError(err).WithField("username", requestData.Username).Error("Can't get new name for the user")
	    w.WriteHeader(http.StatusBadRequest)
	    return
	}

	username := requestData.Username

	err = rt.db.SetUsername(id, username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Database error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
