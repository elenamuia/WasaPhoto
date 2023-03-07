package api

import (
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

type Users struct {
	Name      string `json:"Username"`
	AuthToken string `json:"AuthToken"`
}

type Banned struct {
	Banning string `json:"Banning"`
	Banned  string `json:"Banned"`
}

type Comment struct {
	CommentID      int       `json:"CommentID"`
	CommentContent string    `json:"CommentContent"`
	PhotoID        int       `json:"PhotoID"`
	UserPut        string    `json:"UserPut"`
	DataPost       time.Time `json:"datapost"`
	UserRec        string    `json:"UserRec"`
}

type Like struct {
	LikeID   string    `json:"LikeID"`
	PhotoID  int       `json:"PhotoID"`
	Datapost time.Time `json:"Datapost"`
	UserRec  string    `json:"UserRec"`
}

type Follow struct {
	Follower string `json:"Follower"`
	Followed string `json:"Followed"`
}

type Photo struct {
	ID             int                `json:"PhotoID"`
	PhotoStructure []byte             `json:"PhotoStructure"`
	NumLikes       int                `json:"numLikes"`
	NumComm        int                `json:"numComm"`
	Arrayoflike    []database.Like    `json:"Arrayoflike"`
	Arrayofcomment []database.Comment `json:"Arrayofcomment"`
	Datapost       time.Time          `json:"DataPost"`
	User           string             `json:"UserID"`
}

type Login struct {
	LoginName string `json:"LoginName"`
}

type Profile struct {
	User     int   `json:"User"`
	Photos   []int `json:"Photos"`
	Follower []int `json:"Follower"`
	Followed []int `json:"Followed"`
}

func (l *Login) FromDatabase(log string) {
	l.LoginName = log

}

// ToDatabase returns the fountain in a database-compatible representation
func (b *Banned) ToDatabase(banned string, banning string) database.Banned {
	return database.Banned{

		Banning: banning,
		Banned:  banned,
	}
}

func (c *Comment) ToDatabaseComment(userid string, userputting string, photoid int, comcont string) database.Comment {
	return database.Comment{
		CommentID:   c.CommentID,
		CommMessage: comcont,
		UserPut:     userputting,
		PhotoID:     photoid,
		Datapost:    c.DataPost,
		UserRec:     userid,
	}
}

func (l *Like) ToDatabaseLike(userrec string, userput string, photoid int) database.Like {
	return database.Like{
		LikeID:   userput,
		PhotoID:  photoid,
		DataPost: l.Datapost,
		UserRec:  userrec,
	}
}

func (f *Follow) ToDatabaseFollow(follower string, followed string) database.Follow {
	return database.Follow{
		Follower: follower,
		Followed: followed,
	}
}

func (p *Photo) ToDatabasePhoto(userid string, photoStruct []byte) database.Photo {
	return database.Photo{
		ID:             p.ID,
		PhotoStructure: photoStruct,
		NumLikes:       p.NumLikes,
		NumComm:        p.NumComm,
		Datapost:       p.Datapost,
		User:           p.User,
	}
}

func (u *Users) ToDatabaseUser() database.User {
	return database.User{

		Name:      u.Name,
		AuthToken: u.AuthToken,
	}
}

func (u *Users) FromDatabase(username string) {
	u.Name = username
}

func (log *Login) ToDatabaseLogin() database.Login {
	return database.Login{
		UsernameLog: log.LoginName,
	}
}
