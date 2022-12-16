package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(comment Comment) (err error) {
	res, err := db.c.Exec(`INSERT INTO Comments(CommentID, PhotoID, UserIDReceiving, CommentID, UserIDPutting, UserIDReceiving, DataPost) VALUES (?,?,?,?,?,?,?)`,
		comment.CommentID, comment.PhotoID, comment.UserIDPut, comment.UserIDRec, comment.Datapost)

	if err != nil {
		return err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	comment.CommentID = int(lastInsertID)
	return nil
}
