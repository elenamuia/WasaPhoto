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
	User           string
	PhotoStructure string
	Datapost       time.Time
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
	AddComment(comment Comment) (int, error)
	AddLike(like Like) (Like, error)
	BanUser(string, string) (string, error)
	CheckIfHasBannedMe(string, string) (bool, error)
	CheckIfFollow(string, string) (bool, error)
	CheckIfIBanned(string, string) (bool, error)
	DeleteComment(int) error
	DeleteFollow(follow Follow) (err error)
	DeleteLike(string, int) (err error)
	DeleteProfile(string) (err error)
	GetPagePhoto(userid string, offset int, perPage int) (photos []Photo, err error)
	GetProfile(name string) (p Profile, err error)
	LoginUser(l Login) (User, error)
	CheckAuthToken(AuthToken string) (bool, error)
	PostPhoto(photo Photo) (Photo, error)
	PutFollow(follow Follow) (string, error)
	UnbanUser(string, string) error
	GetMyMainstream(string, int, int) (ArrayofPhotos []Photo, err error)
	CheckILiked(string, int) (bool, error)
	Updateusername(string, string) (string, error)
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

	err2 := db.Ping()
	if err2 != nil {
		return nil, fmt.Errorf("error pinging database: %w", err2)
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
				Follower string NOT NULL,
				Followed string NOT NULL,
				PRIMARY KEY(Follower, Followed),
				FOREIGN KEY (Follower) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (Followed) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE
				
			) WITHOUT ROWID;

			CREATE TABLE IF NOT EXISTS Banned (
				Banned string NOT NULL,
				Banning string NOT NULL, 
				PRIMARY KEY (Banned, Banning),
				FOREIGN KEY (Banned) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (Banning) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE
			) WITHOUT ROWID;


			CREATE TABLE IF NOT EXISTS Photo (
				PhotoID INTEGER PRIMARY KEY AUTOINCREMENT,
				User string NOT NULL,
				Photo string NOT NULL,
				DataPost TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (User) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE
				    
			);

			CREATE TABLE IF NOT EXISTS Comments (
				PhotoID int NOT NULL ,
				UserReceiving string NOT NULL,
				CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
				CommentMessage string NOT NULL,
				UserPutting string NOT NULL,
				FOREIGN KEY (UserReceiving) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (UserPutting) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (PhotoID) REFERENCES Photo(PhotoID) ON DELETE CASCADE ON UPDATE CASCADE					
			);

			CREATE TABLE IF NOT EXISTS Like (
				UserPutting string NOT NULL,
				PhotoID string NOT NULL,
				UserReceiving string NOT NULL,
				PRIMARY KEY (PhotoID, UserPutting)
				FOREIGN KEY (UserPutting) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (UserReceiving) REFERENCES Users(Name) ON DELETE CASCADE ON UPDATE CASCADE,
				FOREIGN KEY (PhotoID) REFERENCES Photo(PhotoID) ON DELETE CASCADE ON UPDATE CASCADE					
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
