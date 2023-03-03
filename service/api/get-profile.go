package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var ban Banned

	id, err2 := strconv.Atoi(ps.ByName("userid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		var idget int
		err3 := json.NewDecoder(r.Body).Decode(&idget)
		if err3 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		if id == ban.BanningID && idget == ban.BannedID {
			ctx.Logger.WithError(err).Error("Banned USer")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Banend User")
			return
		}

		profile, err := rt.db.GetProfile(idget)
		if err != nil {

			ctx.Logger.WithError(err).Error("Can't get profile")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
		// bannedUser.FromDatabaseBanned(dbban)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(profile)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
