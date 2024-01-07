package api

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		ctx.Logger.Error("Bad id")
		w.WriteHeader(http.StatusNotFound)
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
	// get photo from filesystem
	path := "/tmp/filesystem/" + strconv.FormatUint(id, 10) + ".jpg"
	photofile, err := os.Open(path)
	if err != nil {
		ctx.Logger.WithError(err).Error("Photo does not exists")
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		w.Header().Set("Content-Type", "image/png")
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, photofile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Error in buffer")
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf.Bytes())
			if err != nil {
				ctx.Logger.WithError(err).Error("Response sending error")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
}
