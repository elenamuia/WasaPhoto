package database

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUser(l Login) (UserID int, isNew bool, err error) {

	var AuthToken = RandStringBytes(15)

	rows, err1 := db.c.Query(`SELECT * from Users where UserID = ?`, l.IDlog)

	if err1 != nil {
		return 0, false, err
	}

	if err = rows.Err(); err != nil {
		return 0, false, err
	}

	if !rows.Next() {
		_, err = db.c.Exec(`INSERT into Users (UserID, username, AuthToken) VALUES (?, ?, ?)`,
			l.IDlog, l.UsernameLog, AuthToken)
		if err != nil {
			return 0, true, err
		}
		return l.IDlog, true, nil
	}
	/*
		var u User
		err2 := rows.Scan(&u.ID, &u.Name, &u.AuthToken)
		if err2 != nil {
			return 0, false, err
		}
		fmt.Println(u.ID)
	*/
	return l.IDlog, false, nil

}
