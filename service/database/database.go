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
	"time"
)

type User struct {
	ID        int
	Name      string
	AuthToken string
}

type Photo struct {
	ID             int
	PhotoStructure string
	NumLikes       int
	NumComm        int
	ArrayofLike    []Like
	ArrayofComment []Comment
	Datapost       time.Time
	UserID         int
}

type Like struct {
	LikeID   int
	PhotoID  int
	DataPost time.Time
	UserRec  int
}

type Comment struct {
	CommentID   int
	CommMessage string
	PhotoID     int
	UserIDPut   int
	UserIDRec   int
	Datapost    time.Time
}
type Banned struct {
	BannedID  int
	BanningID int
}

type Follow struct {
	FollowerID int
	FollowedID int
}

type Login struct {
	IDlog       int
	UsernameLog string
}
type Profile struct {
	UserID   int
	Photos   []int
	Follower []int
	Followed []int
}

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	DeletePhoto(photo Photo) error
	AddComment(comment Comment) error
	AddLike(like Like) error
	BanUser(ban Banned) error
	DeleteComment(comment Comment) (err error)
	DeleteFollow(follow Follow) (err error)
	DeleteLike(like Like) (err error)
	DeleteProfile(user User) (err error)
	GetPhoto() (Photo, error)
	GetProfile(userid int) (p Profile, err error)
	LoginUser(l Login) (UserID int, err error)
	CheckAuthToken(userId int, authToken string) (bool, error)
	PostPhoto(photo Photo) (err error)
	PutFollow(follow Follow) error
	UnbanUser(ban Banned) (err error)
	GetMyMainstream() (ArrayofPhotos []Photo, err error)
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
		sqlStmt := `CREATE TABLE  IF NOT EXISTS Users (
			UserID int PRIMARY KEY,
			username string NOT NULL
			authToken string NOT NULL	
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
					NumComment int NOT NULL,
    				NumLike int NOT NULL,
				    DataPost string NOT NULL,
					PRIMARY KEY (PhotoID, UserID)
				    FOREIGN KEY (UserID) 
				      REFERENCES Users (UserID) 
				        ON DELETE CASCADE 
				         ON UPDATE NO ACTION
				) WITHOUT ROWID;

				CREATE TABLE IF NOT EXISTS Comments (
						PhotoID int,
					    UserIDReceiving int NOT NULL,
					    CommentID int,
					    UserIDPutting int NOT NULL,
					    DataPost string NOT NULL,
						PRIMARY KEY (PhotoID, CommentID)
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
