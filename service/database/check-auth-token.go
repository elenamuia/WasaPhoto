package database

import "fmt"

func (db *appdbimpl) CheckAuthToken(userId int, AuthToken string) (bool, error) {

	var newAuthToken string
	err := db.c.QueryRow(`SELECT AuthToken FROM Users WHERE UserID=  ? `, userId).Scan(&newAuthToken)
	if err != nil {
		return false, err

	}

	fmt.Println(newAuthToken)

	if newAuthToken != AuthToken {
		return false, nil
	}
	return true, nil
}
