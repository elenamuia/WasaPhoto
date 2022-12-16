package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedComment Comment

	err := json.NewDecoder(r.Body).Decode(&deletedComment)
	if err != nil {

		return
	}

	err = rt.db.DeleteComment(deletedComment.ToDatabaseComment())
	if errors.Is(err, database.ErrCommentDoesNotExist) {

		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {

		ctx.Logger.WithError(err).WithField("commentID", deletedComment).Error("can't delete the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else if deletedComment.UserIDPut != deletedComment.UserIDRec {
		ctx.Logger.WithError(err).WithField("commentID", deletedComment).Error("Not Authorized")
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)

}
