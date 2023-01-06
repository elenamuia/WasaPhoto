package database

import "fmt"

func (db *appdbimpl) BanUser(ban Banned) (err error) {
	_, err = db.c.Exec(`INSERT INTO Banned(BannedID, BanningID) VALUES (?,?)`,
		ban.BannedID, ban.BanningID)

	if err != nil {
		fmt.Println("message3")
		return err
	}

	return nil
}
