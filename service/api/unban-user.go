package api

import (
	"errors"

	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var unban Banned
	banning := ps.ByName("userid")
	banned := ps.ByName("banneduser")
	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		unban.Banned = banned
		unban.Banning = banning

		err3 := rt.db.UnbanUser(unban.ToDatabase(banned, banning))
		if errors.Is(err3, database.ErrPhotoDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err3 != nil {

			ctx.Logger.WithError(err3).WithField("banneduser", unban).Error("can't unban user")
			w.WriteHeader(http.StatusInternalServerError)
			return

		} else if banned == banning {
			ctx.Logger.WithField("banneduser", unban).Error("Not Authorized")
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
