package database

import (
	"database/sql"
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

func (db *appdbimpl) LoginUser(l Login) (user User, err error) {

	var AuthToken = RandStringBytes(15)

	err1 := db.c.QueryRow(`SELECT * from Users where Name = ?`, l.UsernameLog).Scan(&user.Name, &user.AuthToken)

	if err1 != nil {
		if err1 == sql.ErrNoRows {
			_, err = db.c.Exec(`INSERT into Users (Name, AuthToken) VALUES (?, ?)`,
				l.UsernameLog, AuthToken)

			if err != nil {

				return user, err
			}
			user.Name = l.UsernameLog
			user.AuthToken = AuthToken

			return user, nil
		}
	} else {
		return user, nil
	}
	return user, nil

}
