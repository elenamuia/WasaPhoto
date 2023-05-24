package api

import (
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like

	userrec := ps.ByName("userid")
	likeid := ps.ByName("likeid")
	photoid, err2 := strconv.Atoi(ps.ByName("photoid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		err6 := rt.db.AddLike(like.ToDatabaseLike(userrec, likeid, photoid))
		if err6 != nil {
			ctx.Logger.WithError(err).Error("can't like the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
