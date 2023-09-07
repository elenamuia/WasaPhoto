package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userid := ps.ByName("userid")
	page, err2 := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, err3 := strconv.Atoi(r.URL.Query().Get("perpage"))
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		photoRes, err1 := rt.db.GetMyMainstream(userid, page, perPage)

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
