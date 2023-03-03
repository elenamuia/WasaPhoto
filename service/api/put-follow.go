package api

import (
	"encoding/json"

	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		err1 := json.NewDecoder(r.Body).Decode(&follow)

		if err1 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}
		if follow.FollowedID == follow.FollowerID {
			ctx.Logger.WithError(err).Error("A user can't follow himself/herself")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		err2 := rt.db.PutFollow(follow.ToDatabaseFollow())
		if err2 != nil {

			ctx.Logger.WithError(err).Error("Can't add follower")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(follow)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
