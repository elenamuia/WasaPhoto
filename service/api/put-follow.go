package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool == true {
		err := json.NewDecoder(r.Body).Decode(&follow)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return

		}

		err = rt.db.PutFollow(follow.ToDatabaseFollow())
		if err != nil {

			ctx.Logger.WithError(err).Error("Can't add follower")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if follow.FollowedID == follow.FollowerID {
			ctx.Logger.WithError(err).Error("A user can't follow himself/herself")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// Here we can re-use `fountain` as FromDatabase is overwriting every variabile in the structure.
		// bannedUser.FromDatabaseBanned(dbban)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(follow)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
