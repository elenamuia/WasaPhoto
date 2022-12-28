package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteLike(like Like) error {
	res, err1 := db.c.Exec(`DELETE FROM Like WHERE UserIDPutting=? AND PhotoID = ?`, like.LikeID, like.PhotoID)
	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrLikeDoesNotExist
	}
	return nil
}
