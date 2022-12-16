package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) follow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow

	err := json.NewDecoder(r.Body).Decode(&follow)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	err = rt.db.PutFollow(follow.ToDatabaseFollow())
	if err != nil {

		ctx.Logger.WithError(err).Error("can't add follower")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	// bannedUser.FromDatabaseBanned(dbban)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(follow)
}
