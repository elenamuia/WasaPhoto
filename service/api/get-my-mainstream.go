package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userid := ps.ByName("userid")
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		photoRes, err1 := rt.db.GetMyMainstream(userid)

		if err1 != nil {

			ctx.Logger.WithError(err1).Error("can't load mainstream")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photoRes)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
