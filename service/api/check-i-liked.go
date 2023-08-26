package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) checkILiked(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	userrec := ps.ByName("likeid")
	photoid, err2 := strconv.Atoi(ps.ByName("photoid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		//arrayban, err6 := rt.db.getBanningList(userid)
		hasbanned, err6 := rt.db.CheckILiked(userrec, photoid)

		if err6 != nil {
			ctx.Logger.WithError(err6).Error("Can't check if banned")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(hasbanned)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
