package database

func (db *appdbimpl) BanUser(banning string, banned string) (err error) {

	_, err = db.c.Exec(`INSERT INTO Banned(Banned, Banning) VALUES (?,?)`,
		banned, banning)

	if err != nil {

		return err
	}

	return nil
}
