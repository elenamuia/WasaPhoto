package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getFollowed(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// userid := ps.ByName("userid")
	userid := ps.ByName("userid")

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		listFollowed, err6 := rt.db.GetFollowed(userid)

		if err6 != nil {
			fmt.Println(err6)
			ctx.Logger.WithError(err6).Error("Can't get followeds")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(listFollowed)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
