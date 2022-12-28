package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) Updateusername(user User) error {
	res, err1 := db.c.Exec(`UPDATE Users SET username = ? WHERE UserID=?`,
		user.Name, user.ID)
	if err1 != nil {
		return err1
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
