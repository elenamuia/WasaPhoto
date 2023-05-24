package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getLikeList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// userid := ps.ByName("userid")
	photoid, err1 := strconv.Atoi(ps.ByName("photoid"))
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		arraylike, err6 := rt.db.GetListLike(photoid)

		if err6 != nil {
			ctx.Logger.WithError(err6).Error("Can't get list of likes")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(arraylike)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
