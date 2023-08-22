package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userID := r.URL.Query().Get("profile")
	page, err2 := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, err3 := strconv.Atoi(r.URL.Query().Get("perpage"))

	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		photos, err2 := rt.db.GetPagePhoto(userID, page, perPage)
		if err2 != nil {

			ctx.Logger.WithError(err2).Error("Can't post photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photos)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
