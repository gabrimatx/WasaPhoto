package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("userId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var username string
	err = json.NewDecoder(r.Body).Decode(&username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rt.db.SetUsername(id, username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	responseString := fmt.Sprintf("Name successfully changed to %s", username)
	fmt.Fprintf(w, responseString)
}
