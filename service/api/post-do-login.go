package api

import (
	"encoding/json"
	"net/http"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var requestData struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", requestData.Username).Error("Can't login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	username := requestData.Username

	userID, err := rt.db.GetUser(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't operate database")
		w.WriteHeader(http.StatusInternalServerError)
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
