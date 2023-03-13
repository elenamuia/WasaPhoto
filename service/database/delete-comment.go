package database

func (db *appdbimpl) DeleteComment(commentid int) error {
	/*
		var c Comment
		err2 := db.c.QueryRow(`SELECT * FROM Comments WHERE CommentID = ?`, commentid).Scan(c.PhotoID, c.UserRec, c.CommentID, c.CommMessage, c.UserPut, c.Datapost)
		var userput string
		userput = c.UserPut

		if err2 != nil {
			return userput, err2
		}
	*/
	res, err1 := db.c.Exec(`DELETE FROM Comments WHERE CommentID=? `, commentid)
	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrCommentDoesNotExist
	}
	return nil
}
