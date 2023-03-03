package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedLike Like

	photoid, err2 := strconv.Atoi(ps.ByName("photoid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userputid, err1 := strconv.Atoi(ps.ByName("likeid"))
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var userrecid int
		err3 := json.NewDecoder(r.Body).Decode(&userrecid)
		if err3 != nil {

			return
		}
		err4 := rt.db.DeleteLike(deletedLike.ToDatabaseLike(userrecid, userputid, photoid))
		if errors.Is(err, database.ErrLikeDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err4 != nil {

			ctx.Logger.WithError(err).WithField("LikeID", deletedLike).Error("can't unlike")
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
