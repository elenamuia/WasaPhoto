package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(Banned User, Banning User) (err error) {
	_, err = db.c.Exec(`INSERT INTO Banned(BannedID, BanningID) VALUES (?,?)`,
		Banned.ID, Banning.ID)

	if err != nil {
		return err
	}

	return nil
}
