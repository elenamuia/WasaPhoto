package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var bannedUser Banned
	id, err1 := strconv.Atoi("id")
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)

	if bool {

		err = json.NewDecoder(r.Body).Decode(&bannedUser)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		err = rt.db.BanUser(bannedUser.ToDatabase())
		if err != nil {

			ctx.Logger.WithError(err).Error("can't ban the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if bannedUser.BannedID != bannedUser.BanningID {
			ctx.Logger.WithError(err).WithField("BannedID", bannedUser).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(bannedUser)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
