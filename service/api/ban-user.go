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
	id, err := strconv.Atoi("id")

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)

	if bool == true {

		err = json.NewDecoder(r.Body).Decode(&bannedUser)
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return

		}

		err = rt.db.BanUser(bannedUser.ToDatabase())
		if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			ctx.Logger.WithError(err).Error("can't ban the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if bannedUser.BannedID != bannedUser.BanningID {
			ctx.Logger.WithError(err).WithField("BannedID", bannedUser).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
		//bannedUser.FromDatabaseBanned(dbban)

		// Send the output to the user.
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(bannedUser)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
