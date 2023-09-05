package database

func (db *appdbimpl) AddComment(comment Comment) (commentid int, err error) {
	_, err1 := db.c.Exec(`INSERT INTO Comments( PhotoID, UserPutting, UserReceiving,CommentMessage) VALUES (?,?,?,?)`,
		comment.PhotoID, comment.UserPut, comment.UserRec, comment.CommMessage)

	if err1 != nil {
		return 0, err
	}

	err2 := db.c.QueryRow(`SELECT CommentID from Comments where UserPutting = ? ORDER BY CommentID DESC LIMIT 1`, comment.UserPut).Scan(&commentid)
	if err2 != nil {
		return 0, err
	}

	return commentid, nil
}
