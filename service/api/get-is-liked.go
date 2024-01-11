package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLiked(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoId, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("photo not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.WithField("id", userId).Error("Can't find user")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isLiked, err := rt.db.GetIfLiked(photoId, userId)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", userId).Error("Can't get like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := struct {
		IsLiked bool `json:"isLiked"`
	}{
		IsLiked: isLiked,
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
