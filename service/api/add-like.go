package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like

	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		err = json.NewDecoder(r.Body).Decode(&like)
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		// Create the fountain in the database. Note that this function will return a new instance of the fountain with the
		// same information, plus the ID.
		err = rt.db.AddLike(like.ToDatabase())
		if err != nil {

			ctx.Logger.WithError(err).Error("can't comment the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(like)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
