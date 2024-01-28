package api

import (
	"encoding/json"
	"net/http"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	queryParams := r.URL.Query()
	userName := queryParams.Get("userName")

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var userList components.UserSearchList
	userList, err := rt.db.GetUserSearch(userName)
	if err != nil {
		ctx.Logger.WithError(err).Error("Can't search users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userListJSON, err := json.Marshal(userList)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(userListJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
