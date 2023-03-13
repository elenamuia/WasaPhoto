package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedComment Comment
	commentid, err1 := strconv.Atoi(ps.ByName("commentid"))
	userid := ps.ByName("userid")
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var userput string
		err4 := json.NewDecoder(r.Body).Decode(&userput)
		if err4 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err3 := rt.db.DeleteComment(commentid)

		if errors.Is(err, database.ErrCommentDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err3 != nil {

			ctx.Logger.WithError(err3).WithField("commentID", deletedComment).Error("can't delete the comment")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if userid != userput {
			ctx.Logger.WithError(err3).WithField("commentID", deletedComment).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
