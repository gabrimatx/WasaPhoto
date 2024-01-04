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

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("user not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var photoStream components.PhotoList
	photoStream, err = rt.db.GetProfilePhotos(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during photo getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authHeader := r.Header.Get("Authorization")
	authParts := strings.Fields(authHeader)
	token := authParts[1]
	myId, err := strconv.ParseUint(token, 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during id  auth getting")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var followCount int
	followCount, err = rt.db.GetFollowingUsers(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during follow count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var followingCount int
	followingCount, err = rt.db.GetFollowedUsers(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during following count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isBanned bool
	isBanned, err = rt.db.GetBoolBanned(id, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during ban bool getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isFollowed bool
	isFollowed, err = rt.db.GetBoolFollow(id, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during follow bool getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := components.Response{
		PhotoList:     photoStream,
		FollowCount:   followCount,
		FollowedCount: followingCount,
		IsBanned:      isBanned,
		IsFollowed:    isFollowed,
	}

	profileJSON, err := json.Marshal(response)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(profileJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
