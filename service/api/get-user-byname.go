package api

import (
	"encoding/json"
	"net/http"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserId(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	queryParams := r.URL.Query()
	userName := queryParams.Get("userName")

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId, err := rt.db.GetUser(userName)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", userId).Error("Can't get userId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userId == 0 {
		ctx.Logger.WithField("Name", userName).Error("User does not exists")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseData := struct {
		UserId uint64 `json:"userId"`
	}{
		UserId: userId,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
