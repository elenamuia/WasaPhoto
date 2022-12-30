package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var login Login

	authToken := r.Header.Get("authToken")

	err1 := json.NewDecoder(r.Body).Decode(&login)

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, isNew, err2 := rt.db.LoginUser(login.ToDatabaseLogin())

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !isNew {
		bool, err := rt.db.CheckAuthToken(id, authToken)
		if err != nil {
			ctx.Logger.WithError(err).Error("Can't login")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Can't login")
			return
		} else if !bool {
			ctx.Logger.WithError(err).Error("Uncorrect token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Uncorrect token")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(id)

}
