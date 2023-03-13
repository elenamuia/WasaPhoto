package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("userid")

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {
		var idget string
		err3 := json.NewDecoder(r.Body).Decode(&idget)
		if err3 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		bool, err6 := rt.db.CheckIfBanned(id, idget)

		if bool {
			ctx.Logger.WithError(err6).Error("Banned User")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode("Banned User")
			return
		} else {

			profile, err4 := rt.db.GetProfile(idget)
			if err4 != nil {
				fmt.Println(err4)
				ctx.Logger.WithError(err4).Error("Can't get profile")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(profile)
		}

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
