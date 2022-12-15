package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) UnbanUser(banningUser User, bannedUser User) (err error) {
	res, err := db.c.Exec(`DELETE FROM Banned WHERE BannedID=? AND BanningID = ?`, bannedUser.ID, banningUser.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrUserDoesNotExist
	}
	return nil
}
