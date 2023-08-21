package database

func (db *appdbimpl) DeleteComment(commentid int) error {

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
