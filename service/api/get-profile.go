package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var ban Banned

	id := ps.ByName("user")

	authToken := r.Header.Get("Authorization")

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {
		var idget string
		err3 := json.NewDecoder(r.Body).Decode(&idget)
		if err3 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		if id == ban.Banning && idget == ban.Banned {
			ctx.Logger.WithError(err).Error("Banned USer")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Banned User")
			return
		}

		profile, err4 := rt.db.GetProfile(idget)
		if err4 != nil {

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
