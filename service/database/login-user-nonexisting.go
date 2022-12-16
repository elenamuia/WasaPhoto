package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) LoginUserNonExisting(user User) (err error) {
	res, err := db.c.Exec(`INSERT INSERT INTO Users(UserID, username) VALUES (?,?)`,
		user.ID, user.Name)
	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	user.ID = int(lastInsertID)
	return nil
}
