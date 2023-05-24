package database

func (db *appdbimpl) CheckAuthToken(AuthToken string) (bool, error) {

	rows, err := db.c.Query(`SELECT AuthToken FROM Users WHERE AuthToken = ? `, AuthToken)
	if err != nil {

		return false, err

	}

	if !rows.Next() {
		return false, nil
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		return false, err
	}
	return true, nil
}
