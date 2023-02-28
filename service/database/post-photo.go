package database

import "time"

func (db *appdbimpl) PostPhoto(photo Photo) error {

	res, err1 := db.c.Exec(`INSERT INTO Photo (PhotoID, UserID, Photo, NumComment, NumLike, DataPost) VALUES (?, ?, ?, ?,?,?)`,
		photo.ID, photo.UserID, photo.PhotoStructure, 0, 0, time.Now())
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
