package database

func (db *appdbimpl) BanUser(ban Banned) (err error) {
	_, err = db.c.Exec(`INSERT INTO Banned(BannedID, BanningID) VALUES (?,?)`,
		ban.BannedID, ban.BanningID)

	if err != nil {
		return err
	}

	return nil
}
