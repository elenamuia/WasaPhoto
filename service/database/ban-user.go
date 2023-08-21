package database

import "fmt"

func (db *appdbimpl) BanUser(banning string, banned string) (err error) {

	res, err := db.c.Exec(`INSERT INTO Banned(Banned, Banning) VALUES (?,?)`, banned, banning)

	if err != nil {
		fmt.Println("err:", err)
		return err
	}
	_, err2 := res.RowsAffected()
	if err2 != nil {
		fmt.Println("err:", err2)
		return err
	}

	return nil
}
