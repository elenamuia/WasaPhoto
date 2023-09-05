package database

func (db *appdbimpl) AddLike(like Like) error {
	_, err1 := db.c.Exec(`INSERT INTO LIKE(UserPutting, PhotoID, UserReceiving) VALUES (?,?,?)`,
		like.LikeID, like.PhotoID, like.UserRec)

	if err1 != nil {

		return err1
	}

	return nil

}
