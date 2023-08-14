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
	Name      string
	AuthToken string
}

type Photo struct {
	ID             int
	PhotoStructure []byte
	Datapost       time.Time
	User           string
}

type Like struct {
	LikeID   string
	PhotoID  int
	DataPost time.Time
	UserRec  string
}

type Comment struct {
	CommentID   int
	CommMessage string
	PhotoID     int
	UserPut     string
	UserRec     string
	Datapost    time.Time
}
type Banned struct {
	Banned  string
	Banning string
}

type Follow struct {
	Follower string
	Followed string
}

type Login struct {
	UsernameLog string
}
type Profile struct {
	User     string
	Photos   []int
	Follower []string
	Followed []string
}

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetListComments(int) ([]Comment, error)
	GetListLike(int) ([]Like, error)
	DeletePhoto(int) error
	AddComment(comment Comment) error
	AddLike(like Like) error
	BanUser(string, string) error
	CheckIfBanned(string, string) (bool, error)
	DeleteComment(int) error
	DeleteFollow(follow Follow) (err error)
	DeleteLike(string, int) (err error)
	DeleteProfile(string) (err error)
	GetPhoto(int) (Photo, error)
	GetProfile(name string) (p Profile, err error)
	LoginUser(l Login) (name string, isNew bool, err error)
	CheckAuthToken(AuthToken string) (bool, error)
	PostPhoto(photo Photo) (Photo, error)
	PutFollow(follow Follow) error
	UnbanUser(ban Banned) (err error)
	GetMyMainstream(userid string) (ArrayofPhotos []Photo, err error)
	Updateusername(string, string) (string, error)
	GetBannedList(string) ([]string, error)
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
		sqlStmt := `
			PRAGMA foreign_keys = ON;
			CREATE TABLE  IF NOT EXISTS Users (
			Name string NOT NULL PRIMARY KEY,
			AuthToken string NOT NULL UNIQUE	
			) WITHOUT ROWID;
			
			CREATE TABLE IF NOT EXISTS Follower (
				Follower string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				Followed string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				PRIMARY KEY(Follower, Followed)
				
			) WITHOUT ROWID;

			CREATE TABLE IF NOT EXISTS Banned (
				Banned string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE, 
				Banning string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE, 
				PRIMARY KEY (Banned, Banning)
			) WITHOUT ROWID;


			CREATE TABLE IF NOT EXISTS Photo (
					PhotoID INTEGER PRIMARY KEY AUTOINCREMENT,
				    User string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				    Photo string NOT NULL,
				    DataPost string NOT NULL
				    
				);

				CREATE TABLE IF NOT EXISTS Comments (
						PhotoID int NOT NULL REFERENCES Photo(PhotoID) ON DELETE CASCADE ON UPDATE CASCADE,
					    UserReceiving string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
					    CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
						CommentMessage string NOT NULL,
						UserPutting string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
					    DataPost string NOT NULL
						
					);

			CREATE TABLE IF NOT EXISTS Like (

					UserPutting string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
					PhotoID string NOT NULL REFERENCES Photo(PhotoID) ON DELETE CASCADE ON UPDATE CASCADE,
					UserReceiving string NOT NULL REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
					DataPost string NOT NULL,
					
					PRIMARY KEY (PhotoID, UserPutting)
					
				);
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
