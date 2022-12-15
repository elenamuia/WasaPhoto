package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(comment Comment, photo Photo, userrec User, usersend User) (err error) {
	res, err := db.c.Exec(`INSERT INTO Comment(CommentID, PhotoID, UserIDReceiving, CommentID, UserIDSending, DataPost) VALUES (?,?,?,?,?,?)`,
		comment.CommentID, photo.ID, userrec.ID, usersend.ID, comment.datapost)

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
