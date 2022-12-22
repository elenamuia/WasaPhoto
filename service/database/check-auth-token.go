package database

func (db *appdbimpl) CheckAuthToken(userId int, authToken string) (bool, error) {

	var newAuthToken string
	err := db.c.QueryRow("SELECT authToken FROM Users WHERE UserID=  ? ", userId).Scan(&newAuthToken)
	if err != nil {
		return false, err
	}
	if newAuthToken != authToken {
		return false, nil
	}
	return true, nil
}
