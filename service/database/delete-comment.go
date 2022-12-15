package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteComment(comment Comment) (err error) {
	res, err := db.c.Exec(`DELETE FROM Comments WHERE CommentID=? AND PhotoID = ?`, comment.CommentID, comment.PhotoID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrCommentDoesNotExist
	}
	return nil
}
