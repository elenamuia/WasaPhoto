package database

func (db *appdbimpl) UnbanUser(ban Banned) error {
	res, err1 := db.c.Exec(`DELETE FROM Banned WHERE BannedID=? AND BanningID = ?`, ban.BannedID, ban.BanningID)
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
