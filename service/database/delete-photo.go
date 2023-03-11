package database

import "fmt"

func (db *appdbimpl) DeletePhoto(phId int) error {
	fmt.Println(phId)
	res, err1 := db.c.Exec(`DELETE FROM Photo WHERE PhotoID=? `, phId)

	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrPhotoDoesNotExist
	}
	return nil
}
