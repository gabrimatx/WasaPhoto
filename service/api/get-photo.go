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
	//check if not banned from photo

	// get photo from filesystem
	path := "service/filesystem/" + strconv.FormatUint(id, 10) + ".jpg"
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
