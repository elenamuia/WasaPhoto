package database

func (db *appdbimpl) DeleteProfile(user string) error {
	res, err1 := db.c.Exec(`DELETE FROM Users WHERE Name = ?`, user)
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
