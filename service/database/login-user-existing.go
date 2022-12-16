package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUserExisting(user User) (err error) {
	res, err := db.c.Exec(`SELECT username FROM Users where UserID =?`,
		user.ID)
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
