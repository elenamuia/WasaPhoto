package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow

	userid, err1 := strconv.Atoi(ps.ByName("userid"))
	if err1 != nil {
		fmt.Print("Message1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(userid, authToken)
	if bool {

		err := json.NewDecoder(r.Body).Decode(&follow)

		if err != nil {
			fmt.Print("Message2")
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		if follow.FollowedID == follow.FollowerID {
			ctx.Logger.WithError(err).Error("A user can't follow himself/herself")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		err = rt.db.PutFollow(follow.ToDatabaseFollow())
		if err != nil {

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
