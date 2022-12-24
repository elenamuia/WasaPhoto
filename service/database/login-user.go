package database

import "math/rand"

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, err error) {
	var AuthToken int = rand.Int()

	res, err := db.c.Exec(`INSERT OR IGNORE into Users(UserID, username, AuthToken) ((select max(id) from Users)+1,?, ?)`,
		l.UsernameLog, AuthToken)
	if err != nil {
		return 0, err
	}
	_, err = db.c.Query(`SELECT * from Users where username = ?`, l.UsernameLog)
	if err != nil {
		return 0, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	UserID = int(lastInsertID)

	return UserID, nil
}
