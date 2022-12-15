package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddLike(like Like) (err error) {
	res, err := db.c.Exec(`INSERT INTO LIKE(UserIdPutting) VALUES (?)`,
		like.UserID)

	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	like.UserID = int(lastInsertID)
	return nil

}
