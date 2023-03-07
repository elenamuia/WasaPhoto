package database

func (db *appdbimpl) UnbanUser(ban Banned) error {
	res, err1 := db.c.Exec(`DELETE FROM Banned WHERE Banned=? AND Banning = ?`, ban.Banned, ban.Banning)
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
