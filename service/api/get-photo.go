package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get photo from filesystem
	path := "service/filesystem/" + fmt.Sprint(id) + ".jpg"
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
