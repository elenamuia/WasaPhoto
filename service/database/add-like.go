package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddLike(like Like, photo Photo, userrec User) (err error) {
	res, err := db.c.Exec(`INSERT INTO LIKE(UserIdPutting, PhotoID, UserIDReceiving, Datapost) VALUES (?,?,?,?)`,
		like.LikeID, photo.ID, userrec.ID, like.datapost)

	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	like.LikeID = int(lastInsertID)
	return nil

}
