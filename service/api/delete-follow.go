package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedFollow Follow

	err := json.NewDecoder(r.Body).Decode(&deletedFollow)
	if err != nil {

		return
	}

	err = rt.db.DeleteFollow(deletedFollow.ToDatabaseFollow())
	if errors.Is(err, database.ErrUserDoesNotExist) {

		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {

		ctx.Logger.WithError(err).WithField("FollowerID", deletedFollow).Error("can't unfollow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
