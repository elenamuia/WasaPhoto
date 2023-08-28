package database

import "fmt"

func (db *appdbimpl) Updateusername(oldname string, newname string) (string, error) {

	res, err1 := db.c.Exec(`UPDATE Users SET Name = ? WHERE Name = ?`, newname, oldname)
	if err1 != nil {
		fmt.Println("err1: ", err1)

		return "", err1
	}

	affected, err := res.RowsAffected()

	if err != nil {
		fmt.Println("err: ", err)
		return newname, err
	} else if affected == 0 {

		return "", ErrUserDoesNotExist
	}

	return newname, nil
}
