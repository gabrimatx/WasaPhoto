package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(ps.ByName("photoId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var commentBody struct {
		UserID uint64 `json:"commenterId"`
		Text   string `json:"text_comment"`
	}

	err = json.NewDecoder(r.Body).Decode(&commentBody)
	fmt.Fprint(w, commentBody.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment := components.Comment{
		PhotoId:      id,
		UserId:       commentBody.UserID,
		Text_Comment: commentBody.Text,
	}

	err = rt.db.AddComment(comment)
	if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't post comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
