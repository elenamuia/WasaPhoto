package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	err = rt.db.AddComment(comment.ToDatabaseComment())
	if err != nil {

		ctx.Logger.WithError(err).Error("can't comment the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if comment.UserIDPut == comment.UserIDRec {
		ctx.Logger.WithError(err).Error("A user can't comment himself/herself")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	// bannedUser.FromDatabaseBanned(dbban)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(comment)
}
