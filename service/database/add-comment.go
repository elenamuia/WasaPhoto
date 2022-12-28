package database

import "time"

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(comment Comment) error {
	res, err1 := db.c.Exec(`INSERT INTO Comments(CommentID, PhotoID, UserIDReceiving, CommentID, UserIDPutting, UserIDReceiving, DataPost) VALUES (?,?,?,?,?,?,?)`,
		comment.CommentID, comment.PhotoID, comment.UserIDPut, comment.UserIDRec, time.Now)

	if err1 != nil {
		return err1
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {

		return err
	}

	comment.CommentID = int(lastInsertID)
	return nil
}
