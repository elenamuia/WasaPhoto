package database

import "time"

func (db *appdbimpl) AddComment(comment Comment) error {
	_, err1 := db.c.Exec(`INSERT INTO Comments( PhotoID, UserIDPutting, UserIDReceiving,CommentMessage, DataPost) VALUES (?,?,?,?,?)`,
		comment.PhotoID, comment.UserIDPut, comment.UserIDRec, comment.CommMessage, time.Now())

	if err1 != nil {
		return err1
	}

	return nil
}
