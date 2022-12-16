package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

type Users struct {
	ID       int    `json:"UserID"`
	Username string `json:"Username"`
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

type Like struct {
	LikeID   int    `json:"LikeID"`
	PhotoID  int    `json:"PhotoID"`
	Datapost string `json:"Datapost"`
	UserRec  int    `json:"UserRec"`
}

type Follow struct {
	FollowerID int `json:"FollowerID"`
	FollowedID int `json:"FollowedID"`
}

type Photo struct {
	ID             int       `json:"PhotoID"`
	PhotoStructure string    `json:"PhotoStructure"`
	NumLikes       int       `json:"numLikes"`
	NumComm        int       `json:"numComm"`
	ArrayofLike    []Like    `json:"ArrayofLike"`
	ArrayofComment []Comment `json:"ArrayofComment"`
	Datapost       string    `json:"DataPost"`
	UserID         int       `json:"UserID"`
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

func (l *Like) ToDatabaseLike() database.Like {
	return database.Like{
		LikeID:   l.LikeID,
		PhotoID:  l.PhotoID,
		DataPost: l.Datapost,
		UserRec:  l.UserRec,
	}
}

func (f *Follow) ToDatabaseFollow() database.Follow {
	return database.Follow{
		FollowerID: f.FollowerID,
		FollowedID: f.FollowedID,
	}
}

func (p *Photo) ToDatabasePhoto() database.Photo {
	return database.Photo{
		ID:             p.ID,
		PhotoStructure: p.PhotoStructure,
		NumLikes:       p.NumLikes,
		NumComm:        p.NumComm,
		ArrayofLike:    p.ToDatabasePhoto().ArrayofLike,
		ArrayofComment: p.ToDatabasePhoto().ArrayofComment,
		Datapost:       p.Datapost,
		UserID:         p.UserID,
	}
}

func (u *Users) ToDatabaseUser() database.User {
	return database.User{
		ID:   u.ID,
		Name: u.Username,
	}
}
