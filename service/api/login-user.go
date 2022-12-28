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

	err := json.NewDecoder(r.Body).Decode(&login)
	defer r.Body.Close()
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return

	}

	id, err := rt.db.LoginUser(login.ToDatabaseLogin())
	bool, err := rt.db.CheckAuthToken(id, authToken)
	if err != nil {

		ctx.Logger.WithError(err).Error("Can't login")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if bool != true {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	login.FromDatabase(id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(login)

}
