package database

import (
	"time"
)

func (db *appdbimpl) PostPhoto(userid int, photoStruct string) error {

	_, err1 := db.c.Exec(`INSERT INTO Photo (PhotoID, UserID, Photo, NumComment, NumLike, DataPost) VALUES ((SELECT MAX (PhotoID) from Photo) + 1, ?, ?, 0,0,?)`,
		userid, photoStruct, time.Now())
	if err1 != nil {
		return err1
	}
	/*
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}

		photo.ID = int(lastInsertID)
	*/
	return nil
}
