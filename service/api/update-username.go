package api

import (
	"encoding/json"
	"errors"

	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("userid")

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]

	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {
		var body BodyReq

		err2 := json.NewDecoder(r.Body).Decode(&body)

		if err2 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		username, err3 := rt.db.Updateusername(id, body.Username)

		if errors.Is(err3, database.ErrUserDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err3 != nil {

			ctx.Logger.WithError(err).WithField("id", id).Error("can't update the username")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(username)

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
