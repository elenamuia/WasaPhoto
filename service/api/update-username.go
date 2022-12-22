package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) updateUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// The Fountain ID in the path is a 64-bit unsigned integer. Let's parse it.

	id, err := strconv.Atoi("id")
	authToken := r.Header.Get("authToken")

	bool, err := rt.db.CheckAuthToken(id, authToken)
	if bool == true {
		if err != nil {
			// The value was not uint64, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Read the new content for the fountain from the request body.
		var updatedUsername Users
		err = json.NewDecoder(r.Body).Decode(&updatedUsername)
		if err != nil {
			// The body was not a parseable JSON, reject it
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// The client is not supposed to send us the ID in the body, as the fountain ID is already specified in the path,
		// and it's immutable. So, here we overwrite the ID in the JSON with the `id` variable (that comes from the URL).
		updatedUsername.ID = int(id)

		// Update the fountain in the database.
		err = rt.db.Updateusername(updatedUsername.ToDatabaseUser())
		if errors.Is(err, database.ErrUserDoesNotExist) {
			// The fountain (indicated by `id`) does not exist, reject the action indicating an error on the client side.
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
			// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
			// Note (2): we are adding the error and an additional field (`id`) to the log entry, so that we will receive
			// the identifier of the fountain that triggered the error.
			ctx.Logger.WithError(err).WithField("id", id).Error("can't update the username")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
