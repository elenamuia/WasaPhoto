package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	file, _, err2 := r.FormFile("photo")

	userid := ps.ByName("userid")

	if err2 != nil {
		fmt.Println("err2:", err2)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	defer file.Close()

	bytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println("err3:", err3)
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}
	resultEncoding := base64.StdEncoding.EncodeToString(bytes)

	// Append the base64 encoded output
	base64Encoding += resultEncoding

	// fileBytes is now a []byte containing the contents of the file

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {
		// err := json.NewDecoder(r.Body).Decode(&photo)
		// if err != nil {

		// 	w.WriteHeader(http.StatusBadRequest)
		//	fmt.Printf("err")
		// 	return

		photoRet, err4 := rt.db.PostPhoto(photo.ToDatabasePhoto(userid, bytes))
		if err4 != nil {

			ctx.Logger.WithError(err4).Error("Can't post photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photoRet)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
