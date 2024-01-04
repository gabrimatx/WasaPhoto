package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authHeader := r.Header.Get("Authorization")
	authParts := strings.Fields(authHeader)
	token := authParts[1]
	PublisherId, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		ctx.Logger.Error("Bad header parsing")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("Bad id")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var commentBody struct {
		Text string `json:"commentText"`
	}

	err = json.NewDecoder(r.Body).Decode(&commentBody)
	if err != nil {
		ctx.Logger.WithError(err).Error("Bad json content")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment := components.Comment{
		PhotoId:      id,
		UserId:       PublisherId,
		Text_Comment: commentBody.Text,
	}

	err = rt.db.AddComment(comment)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't post comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
