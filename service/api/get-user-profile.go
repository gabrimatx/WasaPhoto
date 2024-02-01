package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("user not found")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var photoStream components.PhotoList
	photoStream, err = rt.db.GetProfilePhotos(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during photo getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myId := GetIdFromBearer(r)

	hisId := id

	userName, err := rt.db.GetUserName(hisId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during name getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userName == "" {
		ctx.Logger.WithError(err).Error("Error user not found")
		w.WriteHeader(http.StatusNotFound)
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
		w.WriteHeader(http.StatusForbidden)
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

	var photoCount int
	photoCount, err = rt.db.GetNumberOfPhotos(id)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during photo count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := components.Response{
		PhotoList:     photoStream,
		UserName:      userName,
		FollowCount:   followCount,
		FollowedCount: followingCount,
		PhotoCount:    photoCount,
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
	_, err = w.Write(profileJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
