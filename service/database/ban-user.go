package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) BanUser(banned Banned, banning Banned) (err error) {
	_, err = db.c.Exec(`INSERT INTO Banned(BannedID, BanningID) VALUES (?,?)`,
		banned.BannedID, banning.BanningID)

	if err != nil {
		return err
	}

	return nil
}
