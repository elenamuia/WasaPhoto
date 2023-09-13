package database

func (db *appdbimpl) AddLike(like Like) (Like, error) {
	var LIKE Like
	_, err1 := db.c.Exec(`INSERT INTO LIKE(UserPutting, PhotoID, UserReceiving) VALUES (?,?,?)`,
		like.LikeID, like.PhotoID, like.UserRec)

	LIKE.LikeID = like.LikeID
	LIKE.PhotoID = like.PhotoID
	LIKE.UserRec = like.UserRec

	if err1 != nil {

		return LIKE, err1
	}

	return LIKE, nil

}
