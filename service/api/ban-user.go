package api

import (
	"encoding/json"

	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var bannedUser Banned

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {

		err2 := json.NewDecoder(r.Body).Decode(&bannedUser)
		if bannedUser.BannedID == bannedUser.BanningID {
			ctx.Logger.WithError(err).WithField("BannedID", bannedUser).Error("A user cannot ban himself/herself")
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err2 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		err = rt.db.BanUser(bannedUser.ToDatabase())

		if err != nil {

			ctx.Logger.WithError(err).Error("can't ban the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(bannedUser)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
