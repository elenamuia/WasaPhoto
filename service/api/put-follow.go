package api

import (
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var follow Follow
	follower := ps.ByName("userid")
	followed := ps.ByName("followeduser")
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {

		if followed == follower {
			ctx.Logger.Error("A user can't follow himself/herself")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		err2 := rt.db.PutFollow(follow.ToDatabaseFollow(follower, followed))
		if err2 != nil {

			ctx.Logger.WithError(err2).Error("Can't add follower")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
