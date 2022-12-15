package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) Updateusername(user User) (err error) {
	res, err := db.c.Exec(`UPDATE Users SET username = ? WHERE UserID=?`,
		user.Name, user.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		//return ErrFountainDoesNotExist
	}
	return nil
}
