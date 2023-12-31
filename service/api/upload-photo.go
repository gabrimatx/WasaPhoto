package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	filesystem "github.com/gabrimatx/WasaPhoto/service/filesystem"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var toPublish components.Photo
	err := json.NewDecoder(r.Body).Decode(&toPublish)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//add photo to database
	photoId, err := rt.db.UploadPhoto(toPublish)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating photo: %s", err.Error()), http.StatusInternalServerError)
		http.Error(w, "Error creating photo", http.StatusInternalServerError)
		return
	}
	//add photo to filesystem

	err = filesystem.SaveBase64Photo(toPublish.File, strconv.FormatUint(photoId, 10)+".jpeg")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving photo to filesystem: %s", err.Error()), http.StatusInternalServerError)
	}

	response := map[string]uint64{"photoId": photoId}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
