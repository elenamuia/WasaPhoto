package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if err9 := r.ParseForm(); err9 != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)

		return
	}

	var login Login

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
		authToken := r.Header.Get("Authorization")
		authToken = strings.Split(authToken, " ")[1]
		bool, err := rt.db.CheckAuthToken(authToken)
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
