package api

import (
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedFollow Follow
	follower := ps.ByName("userid")
	followed := ps.ByName("followeduser")
	authToken := r.Header.Get("Authorization")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {

		deletedFollow.Followed = followed
		deletedFollow.Follower = follower

		err2 := rt.db.DeleteFollow(deletedFollow.ToDatabaseFollow(follower, followed))
		if errors.Is(err, database.ErrUserDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err2 != nil {

			ctx.Logger.WithError(err).WithField("FollowerID", deletedFollow).Error("can't unfollow")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
