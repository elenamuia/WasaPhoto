package database

func (db *appdbimpl) BanUser(ban Banned) (err error) {
	_, err = db.c.Exec(`INSERT INTO Banned(Banned, Banning) VALUES (?,?)`,
		ban.Banned, ban.Banning)

	if err != nil {

		return err
	}

	return nil
}
