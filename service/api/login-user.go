package api

import (
	"encoding/json"
	"net/http"

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

	user, err2 := rt.db.LoginUser(login.ToDatabaseLogin())

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)

}
