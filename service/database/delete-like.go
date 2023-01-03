package database

func (db *appdbimpl) DeleteLike(like Like) error {
	res, err1 := db.c.Exec(`DELETE FROM Like WHERE UserIDPutting=? AND PhotoID = ?`, like.LikeID, like.PhotoID)
	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrLikeDoesNotExist
	}
	return nil
}
