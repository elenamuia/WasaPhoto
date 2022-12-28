package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// getContextReply is an example of HTTP endpoint that returns "Hello World!" as a plain text. The signature of this
// handler accepts a reqcontext.RequestContext (see httpRouterHandler).
func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	idget, err := strconv.Atoi("id")
	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
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
