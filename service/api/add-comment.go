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
	photoid, err1 := strconv.Atoi(ps.ByName("photoid"))

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err1)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var elements []string
		err3 := json.NewDecoder(r.Body).Decode(&elements)
		if err3 != nil {
			fmt.Println(err3)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var comment_cont string
		var useridput int
		var userid int
		comment_cont = elements[0]
		useridput, _ = strconv.Atoi(elements[1])
		userid, _ = strconv.Atoi(elements[2])

		err5 := rt.db.AddComment(comment.ToDatabaseComment(userid, useridput, photoid, comment_cont))
		if err5 != nil {

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
