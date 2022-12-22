package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) loginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var login Login
	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool == true {
		err := json.NewDecoder(r.Body).Decode(&login)
		defer r.Body.Close()
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		idlog, err := rt.db.LoginUser(login.ToDatabaseLogin())
		if err != nil {

			ctx.Logger.WithError(err).Error("Can't login")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
		// bannedUser.FromDatabaseBanned(dbban)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(idlog)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
