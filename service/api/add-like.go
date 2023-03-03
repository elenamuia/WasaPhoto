package api

import (
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like

	userrecid, err1 := strconv.Atoi(ps.ByName("userid"))
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	photoid, err2 := strconv.Atoi(ps.ByName("photoid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userputid, err3 := strconv.Atoi(ps.ByName("likeid"))
	if err3 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		err6 := rt.db.AddLike(like.ToDatabaseLike(userrecid, userputid, photoid))
		if err6 != nil {

			ctx.Logger.WithError(err).Error("can't comment the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
