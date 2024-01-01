package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	components "github.com/gabrimatx/WasaPhoto/service"
	"github.com/gabrimatx/WasaPhoto/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	authParts := strings.Fields(authHeader)
	if len(authParts) != 2 || authParts[0] != "Bearer" {
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	token := authParts[1]

	id, err := strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := rt.db.GetUserIdFromCommentId(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if token == fmt.Sprint(userId) {
		fmt.Fprint(w, "Access granted!")
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	err = rt.db.DeleteComment(id)
	if errors.Is(err, components.ErrObjNotExists) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("id", id).Error("Can't delete comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}