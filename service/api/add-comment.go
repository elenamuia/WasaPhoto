package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var comment Comment
	photoid, err1 := strconv.Atoi(ps.ByName("photoid"))
	userrec := ps.ByName("userid")
	if err1 != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var elements []string
		err3 := json.NewDecoder(r.Body).Decode(&elements)
		if err3 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var comment_cont string
		var userput string

		comment_cont = elements[0]
		userput = elements[1]

		commentid, err5 := rt.db.AddComment(comment.ToDatabaseComment(userrec, userput, photoid, comment_cont))
		if err5 != nil {

			ctx.Logger.WithError(err5).Error("can't comment the photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(commentid)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
