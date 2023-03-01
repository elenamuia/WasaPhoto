package database

import (
	"time"
)

func (db *appdbimpl) PostPhoto(photo Photo) (Photo, error) {

	_, err1 := db.c.Exec(`INSERT INTO Photo (UserID, Photo, NumComment, NumLike, DataPost) VALUES ( ?, ?, ?,?,?)`,
		photo.UserID, photo.PhotoStructure, 0, 0, time.Now())
	if err1 != nil {
		return photo, err1
	}

	return photo, nil
}
