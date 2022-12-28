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

func (rt *_router) deleteMyProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedUser Users
	id, err1 := strconv.Atoi("id")
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		err := json.NewDecoder(r.Body).Decode(&deletedUser)
		if err != nil {

			return
		}

		err = rt.db.DeleteProfile(deletedUser.ToDatabaseUser())
		if errors.Is(err, database.ErrUserDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {

			ctx.Logger.WithError(err).WithField("UserID", deletedUser).Error("can't delete profile")
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
