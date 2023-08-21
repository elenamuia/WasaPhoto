package database

func (db *appdbimpl) UnbanUser(userid string, banned string) error {
	res, err1 := db.c.Exec(`DELETE FROM Banned WHERE Banned=? AND Banning = ?`, banned, userid)

	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return err
	} else if affected == 0 {

		return ErrUserDoesNotExist
	}

	return nil
}
