package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeletePhoto(photo Photo) error {
	res, err1 := db.c.Exec(`DELETE FROM Photo WHERE PhotoID=? AND UserID = ?`, photo.ID, photo.UserID)
	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrPhotoDoesNotExist
	}
	return nil
}
