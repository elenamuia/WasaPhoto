package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like

	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	// Create the fountain in the database. Note that this function will return a new instance of the fountain with the
	// same information, plus the ID.
	err = rt.db.AddLike(like.ToDatabaseLike())
	if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't comment the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
	//bannedUser.FromDatabaseBanned(dbban)

	// Send the output to the user.
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(like)
}
