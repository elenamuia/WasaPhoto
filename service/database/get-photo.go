package database

import "fmt"

func (db *appdbimpl) GetPhoto(photoid int) (photo Photo, err error) {
	fmt.Println(photoid)
	err1 := db.c.QueryRow("SELECT * FROM Photo WHERE PhotoID=  ? ", photoid).Scan(&photo.ID, &photo.User, &photo.PhotoStructure, &photo.Datapost)
	fmt.Println(photo.ID)
	if err1 != nil {
		return photo, err
	}

	return photo, nil
}
