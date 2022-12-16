package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedLike Like

	err := json.NewDecoder(r.Body).Decode(&deletedLike)
	if err != nil {

		return
	}

	err = rt.db.DeleteLike(deletedLike.ToDatabaseLike())
	if errors.Is(err, database.ErrLikeDoesNotExist) {

		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {

		ctx.Logger.WithError(err).WithField("LikeID", deletedLike).Error("can't unlike")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
