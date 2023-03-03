package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedFollow Follow

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		err1 := json.NewDecoder(r.Body).Decode(&deletedFollow)
		if err1 != nil {

			return
		}

		err2 := rt.db.DeleteFollow(deletedFollow.ToDatabaseFollow())
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
