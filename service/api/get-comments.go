package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("photo not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myId := GetIdFromBearer(r)

	hisId, err := rt.db.GetUserIdFromPhotoId(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during id getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isBan, err := rt.db.GetBoolBanned(myId, hisId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during ban getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBan {
		ctx.Logger.Error("Banned")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var commentList components.CommentList
	commentList, err = rt.db.GetPhotoComments(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during comment getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentStreamJSON, err := json.Marshal(commentList)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(commentStreamJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
