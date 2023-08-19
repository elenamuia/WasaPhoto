package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getBanList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userid := ps.ByName("userid")

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		//arrayban, err6 := rt.db.getBanningList(userid)
		arrayban, err6 := rt.db.GetBanningList(userid)

		if err6 != nil {
			ctx.Logger.WithError(err6).Error("Can't get list of likes")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arrayban)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}