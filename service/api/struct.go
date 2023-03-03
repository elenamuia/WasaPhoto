package api

import (
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type Users struct {
	ID        int    `json:"UserID"`
	Username  string `json:"Username"`
	AuthToken string `json:"AuthToken"`
}

type Banned struct {
	BanningID int `json:"BanningID"`
	BannedID  int `json:"BannedID"`
}

type Comment struct {
	CommentID      int       `json:"CommentID"`
	CommentContent string    `json:"CommentContent"`
	PhotoID        int       `json:"PhotoID"`
	UserIDPut      int       `json:"UserIDPut"`
	DataPost       time.Time `json:"datapost"`
	UserIDRec      int       `json:"UserIDRec"`
}

type Like struct {
	LikeID   int       `json:"LikeID"`
	PhotoID  int       `json:"PhotoID"`
	Datapost time.Time `json:"Datapost"`
	UserRec  int       `json:"UserRec"`
}

type Follow struct {
	FollowerID int `json:"FollowerID"`
	FollowedID int `json:"FollowedID"`
}

type Photo struct {
	ID             int                `json:"PhotoID"`
	PhotoStructure []byte             `json:"PhotoStructure"`
	NumLikes       int                `json:"numLikes"`
	NumComm        int                `json:"numComm"`
	Arrayoflike    []database.Like    `json:"Arrayoflike"`
	Arrayofcomment []database.Comment `json:"Arrayofcomment"`
	Datapost       time.Time          `json:"DataPost"`
	UserID         int                `json:"UserID"`
}

type Login struct {
	LoginID   int    `json:"LoginID"`
	LoginName string `json:"LoginName"`
}

type Profile struct {
	UserID   int   `json:"UserID"`
	Photos   []int `json:"Photos"`
	Follower []int `json:"Follower"`
	Followed []int `json:"Followed"`
}

func (l *Login) FromDatabase(log int) {
	l.LoginID = log

}

// ToDatabase returns the fountain in a database-compatible representation
func (b *Banned) ToDatabase() database.Banned {
	return database.Banned{

		BanningID: b.BanningID,
		BannedID:  b.BannedID,
	}
}

func (c *Comment) ToDatabaseComment(userid int, userputting int, photoid int, comcont string) database.Comment {
	return database.Comment{
		CommentID:   c.CommentID,
		CommMessage: comcont,
		UserIDPut:   userputting,
		PhotoID:     photoid,
		Datapost:    c.DataPost,
		UserIDRec:   userid,
	}
}

func (l *Like) ToDatabaseLike(userrecid int, userputid int, photoid int) database.Like {
	return database.Like{
		LikeID:   userputid,
		PhotoID:  photoid,
		DataPost: l.Datapost,
		UserRec:  userrecid,
	}
}

func (f *Follow) ToDatabaseFollow() database.Follow {
	return database.Follow{
		FollowerID: f.FollowerID,
		FollowedID: f.FollowedID,
	}
}

func (p *Photo) ToDatabasePhoto(userid int, photoStruct []byte) database.Photo {
	return database.Photo{
		ID:             p.ID,
		PhotoStructure: photoStruct,
		NumLikes:       p.NumLikes,
		NumComm:        p.NumComm,
		Datapost:       p.Datapost,
		UserID:         userid,
	}
}

func (u *Users) ToDatabaseUser() database.User {
	return database.User{
		ID:        u.ID,
		Name:      u.Username,
		AuthToken: u.AuthToken,
	}
}

func (u *Users) FromDatabase(username string) {
	u.Username = username
}

func (log *Login) ToDatabaseLogin() database.Login {
	return database.Login{
		IDlog:       log.LoginID,
		UsernameLog: log.LoginName,
	}
}
