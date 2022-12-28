package database

import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, err error) {

	var AuthToken = RandStringBytes(15)

	_, err = db.c.Exec(`INSERT OR IGNORE into Users (UserID, username, AuthToken) VALUES (?, ?, ?))`,
		l.IDlog, l.UsernameLog, AuthToken)
	if err != nil {
		return 0, err
	}
	rows, err := db.c.Query(`SELECT * from Users where username = ?`, l.UsernameLog)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.AuthToken)
		if err != nil {
			return 0, err
		}
		UserID = u.ID
	}

	return UserID, nil
}
