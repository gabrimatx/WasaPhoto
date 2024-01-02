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
	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // memory limit
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing multipart form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// get photo from form
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving the file: %s", err.Error()), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// get publisher id and caption into struct
	additionalData := r.FormValue("additionalData")
	var photoData components.Photo
	err = json.Unmarshal([]byte(additionalData), &photoData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding additional data: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// add photo info to database
	photoID, err := rt.db.UploadPhoto(photoData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating photo: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	// save photo to filesystem
	err = saveUploadedFile(file, handler, photoID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving photo to filesystem: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]uint64{"photoID": photoID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusCreated)
}

// saveUploadedFile saves the uploaded file to the filesystem
func saveUploadedFile(file multipart.File, handler *multipart.FileHeader, photoID uint64) error {
	// filename such that id.ext
	fileName := filepath.Join(uploadDirectory, fmt.Sprintf("%d%s", photoID, ".jpg"))

	// create file
	out, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating the file: %s", err.Error())
	}
	defer out.Close()

	// copy file content to new file
	_, err = io.Copy(out, file)
	if err != nil {
		return fmt.Errorf("error copying the file: %s", err.Error())
	}

	return nil
}
