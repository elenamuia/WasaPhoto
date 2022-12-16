package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) PostPhoto(photo Photo) (PhotoID int, err error) {
	res, err := db.c.Exec(`INSERT INTO Photo (PhotoID, UserID, Photo, NumComment, NumLike, DataPost) VALUES (?, ?, ?, ?,?,?)`,
		photo.ID, photo.UserID, photo.PhotoStructure, photo.NumComm, photo.NumLikes, photo.Datapost)
	if err != nil {
		return photo.ID, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return photo.ID, err
	}

	photo.ID = int(lastInsertID)
	return photo.ID, nil
}
