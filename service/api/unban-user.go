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

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var unban Banned
	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		err := json.NewDecoder(r.Body).Decode(&unban)
		if err != nil {

			return
		}

		err = rt.db.UnbanUser(unban.ToDatabase())
		if errors.Is(err, database.ErrPhotoDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {

			ctx.Logger.WithError(err).WithField("BannedID", unban).Error("can't unban user")
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else if unban.BannedID == unban.BanningID {
			ctx.Logger.WithError(err).WithField("BannedID", unban).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
