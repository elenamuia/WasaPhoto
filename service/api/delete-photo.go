package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedPhoto Photo
	id, err1 := strconv.Atoi("id")
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)

	if bool {
		err := json.NewDecoder(r.Body).Decode(&deletedPhoto)
		if err != nil {

			return
		}

		err = rt.db.DeletePhoto(deletedPhoto.ToDatabasePhoto())
		if errors.Is(err, database.ErrPhotoDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {

			ctx.Logger.WithError(err).WithField("PhotoID", deletedPhoto).Error("can't delete photo")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		w.WriteHeader(http.StatusNoContent)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
