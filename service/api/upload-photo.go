package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

type FormFile struct {
	File   multipart.File
	Header *multipart.FileHeader
	Mime   string
}

const uploadDirectory = "service/filesystem/"

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// check auth
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // memory limit
	if err != nil {
		ctx.Logger.Error("Bad multiform data")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get photo from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		ctx.Logger.Error("Bad photo file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get publisher id and caption into struct
	additionalData := r.FormValue("additionalData")
	var photoData components.Photo
	err = json.Unmarshal([]byte(additionalData), &photoData)
	if err != nil {
		ctx.Logger.Error("Bad json data parsing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add photo info to database
	photoID, err := rt.db.UploadPhoto(photoData)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating database row for the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// save photo to filesystem
	err = saveUploadedFile(file, handler, photoID)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error creating database file in the filesystem")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// saveUploadedFile saves the uploaded file to the filesystem
func saveUploadedFile(file multipart.File, handler *multipart.FileHeader, photoID uint64) error {
	// filename such that id.ext
	fileName := filepath.Join(uploadDirectory, fmt.Sprintf("%d%s", photoID, ".jpg"))

	// create file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// copy file content to new file
	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	return nil
}
