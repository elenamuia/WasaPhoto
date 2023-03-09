package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var like Like

	userrec := ps.ByName("userid")

	photoid, err2 := strconv.Atoi(ps.ByName("photoid"))
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {
		var userput string
		err3 := json.NewDecoder(r.Body).Decode(&userput)
		if err3 != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err6 := rt.db.AddLike(like.ToDatabaseLike(userrec, userput, photoid))
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
