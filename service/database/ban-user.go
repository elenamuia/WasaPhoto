package database

import "fmt"

func (db *appdbimpl) BanUser(banning string, banned string) (err error) {

	_, err = db.c.Exec(`INSERT INTO Banned(Banned, Banning) VALUES (?,?)`, banned, banning)

	if err != nil {
		fmt.Println("err:", err)
		return err
	}

	return nil
}
