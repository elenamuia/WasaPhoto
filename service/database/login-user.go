package database

import (
	"fmt"
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

func (db *appdbimpl) LoginUser(l Login) (User string, isNew bool, err error) {

	var AuthToken = RandStringBytes(15)

	rows, err1 := db.c.Query(`SELECT * from Users where Name = ?`, l.UsernameLog)

	if err1 != nil {
		return "", false, err
	}

	defer func() { _ = rows.Close() }()
	if err = rows.Err(); err != nil {
		return "", false, err
	}

	if !rows.Next() {

		_, err = db.c.Exec(`INSERT into Users (Name, AuthToken) VALUES (?, ?)`,
			l.UsernameLog, AuthToken)

		if err != nil {
			fmt.Println("err: ", err)
			return "", true, err
		}

		return l.UsernameLog, true, nil
	}

	return l.UsernameLog, false, nil

}
