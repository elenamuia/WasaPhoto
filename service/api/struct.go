package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

type Users struct {
	ID       int    `json:"id"`
	username string `json:"username"`
}

type Banned struct {
	BannedID  int `json:"BannedID"`
	BanningID int `json:"BanningID"`
}

type Comment struct {
	CommentID      int    `json:"CommentID"`
	CommentContent string `json:"CommentContent"`
	PhotoID        int    `json:"PhotoID"`
	UserIDPut      int    `json:"UserIDPut"`
	DataPost       string `json:"datapost"`
	UserIDRec      int    `json:"UserIDRec"`
}

//func (b *Banned) FromDatabaseBanned(ban database.Banned) {
//	b.BannedID = ban.BannedID
//	b.BanningID = ban.BanningID

//}

// ToDatabase returns the fountain in a database-compatible representation
func (b *Banned) ToDatabaseBanned() database.Banned {
	return database.Banned{
		BannedID:  b.BannedID,
		BanningID: b.BanningID,
	}
}

func (c *Comment) ToDatabaseComment() database.Comment {
	return database.Comment{
		CommentID:   c.CommentID,
		CommMessage: c.CommentContent,
		UserIDPut:   c.UserIDPut,
		PhotoID:     c.PhotoID,
		Datapost:    c.DataPost,
		UserIDRec:   c.UserIDRec,
	}
}
