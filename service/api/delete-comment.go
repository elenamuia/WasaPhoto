package api

import (
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedComment Comment
	commentid, err1 := strconv.Atoi(ps.ByName("commentid"))

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {

		err3 := rt.db.DeleteComment(commentid)
		if errors.Is(err, database.ErrCommentDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err3 != nil {

			ctx.Logger.WithError(err).WithField("commentID", deletedComment).Error("can't delete the comment")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if deletedComment.UserPut != deletedComment.UserRec {
			ctx.Logger.WithError(err).WithField("commentID", deletedComment).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
