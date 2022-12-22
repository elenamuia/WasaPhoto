package database

import "time"

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddLike(like Like) (err error) {
	res, err := db.c.Exec(`INSERT INTO LIKE(UserIdPutting, PhotoID, UserIDReceiving, Datapost) VALUES (?,?,?,?)`,
		like.LikeID, like.PhotoID, like.UserRec, time.Now())

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
