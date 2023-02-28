package database

import "time"

func (db *appdbimpl) PostPhoto(photo Photo) error {

	res, err1 := db.c.Exec(`INSERT INTO Photo (PhotoID, UserID, Photo, NumComment, NumLike, DataPost) VALUES ((SELECT MAX (PhotoID) from Photo, ?, ?, 0,0,?)`,
		photo.ID, photo.UserID, photo.PhotoStructure, time.Now())
	if err1 != nil {
		return err1
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	photo.ID = int(lastInsertID)
	return nil
}
