package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	id := ps.ByName("userid")

	authToken := r.Header.Get("Authorization")
	authToken = strings.Replace(authToken, "Bearer ", "", 1)
	fmt.Println(authToken)
	bool, err := rt.db.CheckAuthToken(authToken)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if bool {
		var newName string
		err2 := json.NewDecoder(r.Body).Decode(&newName)
		fmt.Println(newName)
		if err2 != nil {

			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println(id)
		username, err2 := rt.db.Updateusername(id, newName)
		fmt.Println(username)
		if errors.Is(err, database.ErrUserDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err2 != nil {

			ctx.Logger.WithError(err).WithField("id", id).Error("can't update the username")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(newName)

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
