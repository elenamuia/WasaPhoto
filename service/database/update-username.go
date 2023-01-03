package database

func (db *appdbimpl) Updateusername(user User) (string, error) {
	var username string
	res, err1 := db.c.Exec(`UPDATE Users SET username = ? WHERE UserID=?`,
		user.Name, user.ID)
	if err1 != nil {
		return "", err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return username, err
	} else if affected == 0 {

		return "", ErrUserDoesNotExist
	}
	username = user.Name
	return username, nil
}
