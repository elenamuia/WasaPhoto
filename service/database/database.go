/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

type Photo struct {
	ID             int
	PhotoStructure string
	numLikes       int
	numComm        int
	ArrayofLike    []Like
	ArrayofComment []Comment
}

type Like struct {
	LikeID   int
	PhotoID  Photo
	datapost string
}

type Comment struct {
	CommentID   int
	CommMessage string
	PhotoID     int
	UserID      int
	datapost    string
}

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	DeletePhoto(PhotoID int, u User) error
	AddComment(comment Comment, photo Photo, userrec User, usersend User) error
	AddLike(like Like, photo Photo, userrec User) error
	BanUser(User) error
	DeleteComment(comment Comment, photo Photo) (err error)
	DeleteFollow(User) (err error)
	DeleteLike(sendUser User, photo Photo) (err error)
	DeleteProfile(ArrayPhotoID []int, ArrayCommentID []int, ArrayLikeID []int) (err error)
	GetPhoto() (Photo, error)
	GetProfile() (UserList []User, PhotoID []int, NumFollower []int, NumFollowed []int, err error)
	LoginUser(UserName string) (UserID int, err error)
	PostPhoto(PhotoStructure string) (PhotoID int, err error)
	PutFollow(u User) error
	UnbanUser(User) (err error)
	GetMyMainstream(ArrayofPhotos []Photo) (err error)
	Updateusername(User) (err error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='database.db';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (
			UserID int PRIMARY KEY,
			username string NOT NULL	
			) WITHOUT ROWID;
			
			CREATE TABLE IF NOT EXISTS Follower (
				FollowerID int ,
				FollowedID int,
				PRIMARY KEY(FollowerID, FOllowedID)
				FOREIGN KEY (FollowerID) 
				  REFERENCES Users (UserID) 
					 ON DELETE CASCADE 
					 ON UPDATE NO ACTION
				FOREIGN KEY (FollowedID) 
				  REFERENCES Users (UserID) 
					 ON DELETE CASCADE 
					 ON UPDATE NO ACTION
			) WITHOUT ROWID;

			CREATE TABLE IF NOT EXISTS Banned (
				BannedID int,
				BanningID int,
				PRIMARY KEY (BannedID, BanningID)
				FOREIGN KEY (BannedID) 
				  REFERENCES Users (UserID) 
					ON DELETE CASCADE 
					ON UPDATE NO ACTION
				FOREIGN KEY (BanningID) 
				  REFERENCES Users (UserID) 
					 ON DELETE CASCADE 
					 ON UPDATE NO ACTION
			) WITHOUT ROWID;


			CREATE TABLE IF NOT EXISTS Photo (
					PhotoID int,
				    UserID int,
				    Photo string NOT NULL,
				    DataPost string NOT NULL,
					PRIMARY KEY (PhotoID, UserID)
				    FOREIGN KEY (UserID) 
				      REFERENCES Users (UserID) 
				        ON DELETE CASCADE 
				         ON UPDATE NO ACTION
				) WITHOUT ROWID;

				CREATE TABLE IF NOT EXISTS Comment (
						PhotoID int,
					    UserIDReceiving int NOT NULL,
					    CommentID int,
					    UserIDSending int NOT NULL,
					    DataPost string NOT NULL,
						PRIMARY KEY (PhotoID, CommentID)
					    FOREIGN KEY (UserIDReceiving) 
					      REFERENCES Users (UserID) 
					       ON DELETE CASCADE 
					         ON UPDATE NO ACTION
					    FOREIGN KEY (UserIDSending) 
					      REFERENCES Users (UserID) 
					       ON DELETE CASCADE 
					         ON UPDATE NO ACTION
					    FOREIGN KEY (PhotoID) 
					      REFERENCES Users (UserID) 
					       ON DELETE CASCADE 
					         ON UPDATE NO ACTION
					) WITHOUT ROWID;

					 CREATE TABLE IF NOT EXISTS Like (
 
						     UserIDPutting int,
						     PhotoID string NOT NULL,
						     UserIDReceiving int NOT NULL,
						     DataPost string NOT NULL,
						     PRIMARY KEY (PhotoID, UserIDPutting)
						     FOREIGN KEY (UserIDReceiving) 
						       REFERENCES Users (UserID) 
					             ON DELETE CASCADE 
					              ON UPDATE NO ACTION
					         FOREIGN KEY (UserIDPutting) 
					           REFERENCES Users (UserID) 
					             ON DELETE CASCADE 
					              ON UPDATE NO ACTION
					         FOREIGN KEY (PhotoID) 
					           REFERENCES Photo (PhotoID) 
					             ON DELETE CASCADE 
					              ON UPDATE NO ACTION
					     ) WITHOUT ROWID;
`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
