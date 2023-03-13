package database

func (db *appdbimpl) DeleteLike(likeid string, photoid int) error {
	res, err1 := db.c.Exec(`DELETE FROM Like WHERE UserPutting=? AND PhotoID = ?`, likeid, photoid)
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
