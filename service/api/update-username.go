package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.

	id, err1 := strconv.Atoi(ps.ByName("id"))
	if err1 != nil {
		fmt.Print("Message3", err1)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool {
		if err != nil {
			fmt.Print("Message1")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Read the new content for the fountain from the request body.
		var updatedUsername Users
		err = json.NewDecoder(r.Body).Decode(&updatedUsername)
		if err != nil {
			// The body was not a parseable JSON, reject it
			fmt.Print("Message2")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
		// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
		updatedUsername.ID = id

		// Update the fountain in the database.
		username, err2 := rt.db.Updateusername(updatedUsername.ToDatabaseUser())
		fmt.Println(username)
		if errors.Is(err, database.ErrUserDoesNotExist) {
			// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err2 != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
			// the identifier of the fountain that triggered the error.
			ctx.Logger.WithError(err).WithField("id", id).Error("can't update the username")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		updatedUsername.Username = username
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(username)

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
