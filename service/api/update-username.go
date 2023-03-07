package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("userid")

	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var updatedUser Users
		err2 := json.NewDecoder(r.Body).Decode(&updatedUser)
		if err2 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updatedUser.Name = id

		username, err2 := rt.db.Updateusername(updatedUser.ToDatabaseUser())
		fmt.Println(username)
		if errors.Is(err, database.ErrUserDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err2 != nil {

			ctx.Logger.WithError(err).WithField("id", id).Error("can't update the username")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		updatedUser.Name = username
		updatedUser.AuthToken = authToken
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(updatedUser)

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
