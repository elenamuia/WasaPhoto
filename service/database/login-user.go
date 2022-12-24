package database

import "math/rand"

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, err error) {
	var AuthToken int = rand.Int()

	_, err = db.c.Exec(`INSERT OR IGNORE into Users(UserID, username, AuthToken) ?, ?, ?)`,
		l.IDlog, l.UsernameLog, AuthToken)
	if err != nil {
		return 0, err
	}
	rows, err := db.c.Query(`SELECT * from Users where username = ?`, l.UsernameLog)
	if err != nil {
		return 0, err
	}
	var id int
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.AuthToken)
		if err != nil {
			return 0, err
		}
		id = u.ID
	}

	return id, nil
}
