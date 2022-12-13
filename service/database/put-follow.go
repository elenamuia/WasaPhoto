package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) PutFollow(u User) (User, error) {
	res, err := db.c.Exec(`INSERT INTO followers (UserId) VALUES (?)`,
		u.ID)

	if err != nil {
		return u, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return u, err
	}

	u.ID = string(lastInsertID)
	return u, nil

}
