package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) PutFollow(u User) (err error) {
	res, err := db.c.Exec(`INSERT INTO Follower(UserId) VALUES (?)`,
		u.ID)

	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	u.ID = int(lastInsertID)
	return nil

}
