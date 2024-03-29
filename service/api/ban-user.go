package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	banned := ps.ByName("banneduser")
	banning := ps.ByName("userid")
	BearerToken := r.Header.Get("Authorization")
	authToken := strings.Split(BearerToken, " ")[1]

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		banname, err1 := rt.db.BanUser(banning, banned)

		if err1 != nil {

			ctx.Logger.WithError(err1).Error("can't ban the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(banname)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
