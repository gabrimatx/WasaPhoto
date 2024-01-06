package api

import (
	"encoding/json"
	"net/http"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var username string
	err := json.NewDecoder(r.Body).Decode(&username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := rt.db.GetUser(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't operate database")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if userID != 0 {
		response := map[string]uint64{"userId": userID}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	newUserID, err := rt.db.InsertUser(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]uint64{"userId": newUserID}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}