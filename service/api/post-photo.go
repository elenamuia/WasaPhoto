package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		err := json.NewDecoder(r.Body).Decode(&photo)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		err = rt.db.PostPhoto(photo.ToDatabasePhoto())
		if err != nil {

			ctx.Logger.WithError(err).Error("Can't post photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photo)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
