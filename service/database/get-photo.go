package database

func (db *appdbimpl) GetPhoto(photoid int) (photo Photo, err error) {

	err1 := db.c.QueryRow("SELECT * FROM Photo WHERE PhotoID=  ? ", photoid).Scan(&photo.ID, &photo.UserID, &photo.PhotoStructure, &photo.NumComm, &photo.NumLikes, &photo.Datapost)
	if err1 != nil {
		return photo, err
	}

	return photo, nil
}
