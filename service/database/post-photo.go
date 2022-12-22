package database

import "time"

// GetName is an example that shows you how to query data
func (db *appdbimpl) PostPhoto(photo Photo) (err error) {

	res, err := db.c.Exec(`INSERT INTO Photo (PhotoID, UserID, Photo, NumComment, NumLike, DataPost) VALUES (?, ?, ?, ?,?,?)`,
		photo.ID, photo.UserID, photo.PhotoStructure, photo.NumComm, photo.NumLikes, time.Now())
	if err != nil {
		return err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	photo.ID = int(lastInsertID)
	return nil
}
