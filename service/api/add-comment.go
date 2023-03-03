package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	userput, err1 := strconv.Atoi(ps.ByName("useridputting"))
	photoid, err1 := strconv.Atoi(ps.ByName("photoid"))
	userid, err1 := strconv.Atoi(ps.ByName("userid"))

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err1)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var comment_cont string
		err3 := json.NewDecoder(r.Body).Decode(&comment_cont)
		if err3 != nil {
			fmt.Println(err3)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = rt.db.AddComment(comment.ToDatabaseComment(userid, userput, photoid, comment_cont))
		if err != nil {

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
