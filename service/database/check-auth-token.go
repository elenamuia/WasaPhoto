package database

func (db *appdbimpl) CheckAuthToken(userId int, AuthToken string) (bool, error) {

	var newAuthToken string
	err := db.c.QueryRow(`SELECT AuthToken FROM Users WHERE UserID=  ? `, userId).Scan(&newAuthToken)
	if err != nil {
		return false, err

	}

	if newAuthToken != AuthToken {
		return false, nil
	}
	return true, nil
}
