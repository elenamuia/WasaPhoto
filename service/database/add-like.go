package database

import "time"

func (db *appdbimpl) AddLike(like Like) error {
	_, err1 := db.c.Exec(`INSERT INTO LIKE(UserIdPutting, PhotoID, UserIDReceiving, Datapost) VALUES (?,?,?,?)`,
		like.LikeID, like.PhotoID, like.UserRec, time.Now())

	if err1 != nil {
		return err1
	}

	return nil

}
