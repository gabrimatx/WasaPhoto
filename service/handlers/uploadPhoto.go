package handlers

import (
	"encoding/json"
	"net/http"
)

func uploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newPhoto Photo
	err := json.NewDecoder(r.Body).Decode(&newPhoto)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Implement your logic to store the new photo in the database or handle as needed

}
