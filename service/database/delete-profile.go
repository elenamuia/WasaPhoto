package database

func (db *appdbimpl) DeleteProfile(user User) error {
	res, err1 := db.c.Exec(`DELETE FROM Users WHERE UserID = ?`, user.ID)
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
