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
	NumLikes       int
	NumComm        int
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
	Follower []int
	Followed []int
}

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrPhotoDoesNotExist = errors.New("Photo does not exist")
var ErrCommentDoesNotExist = errors.New("Comment does not exist")
var ErrLikeDoesNotExist = errors.New("Like does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	DeletePhoto(int) error
	AddComment(comment Comment) error
	AddLike(like Like) error
	BanUser(ban Banned) error
	DeleteComment(int) (err error)
	DeleteFollow(follow Follow) (err error)
	DeleteLike(like Like) (err error)
	DeleteProfile(string) (err error)
	GetPhoto(int) (Photo, error)
	GetProfile(name string) (p Profile, err error)
	LoginUser(l Login) (name string, isNew bool, err error)
	CheckAuthToken(AuthToken string) (bool, error)
	PostPhoto(photo Photo) (Photo, error)
	PutFollow(follow Follow) error
	UnbanUser(ban Banned) (err error)
	GetMyMainstream() (ArrayofPhotos []Photo, err error)
	Updateusername(User) (string, error)
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
			AuthToken string NOT NULL PRIMARY KEY,
			Name string NOT NULL UNIQUE
				
			) WITHOUT ROWID;
			
			CREATE TABLE IF NOT EXISTS Follower (
				Follower string ,
				Followed string,
				PRIMARY KEY(Follower, Followed)
				FOREIGN KEY (Follower) 
				  REFERENCES Users (Name) 
					 ON DELETE CASCADE 
					 ON UPDATE CASCADE
				FOREIGN KEY (Followed) 
				  REFERENCES Users (Name) 
					 ON DELETE CASCADE 
					 ON UPDATE CASCADE
			) WITHOUT ROWID;

			CREATE TABLE IF NOT EXISTS Banned (
				Banned string,
				Banning string,
				PRIMARY KEY (Banned, Banning)
				FOREIGN KEY (Banned) 
				  REFERENCES Users (Name) 
					ON DELETE CASCADE 
					ON UPDATE CASCADE
				FOREIGN KEY (Banning) 
				  REFERENCES Users (Name) 
					 ON DELETE CASCADE 
					 ON UPDATE CASCADE
			) WITHOUT ROWID;


			CREATE TABLE IF NOT EXISTS Photo (
					PhotoID INTEGER PRIMARY KEY AUTOINCREMENT,
				    User string,
				    Photo string NOT NULL,
				    DataPost string NOT NULL,
				    FOREIGN KEY (User) 
				      REFERENCES Users (Name) 
				        ON DELETE CASCADE 
				         ON UPDATE CASCADE
				);

				CREATE TABLE IF NOT EXISTS Comments (
						PhotoID int,
					    UserReceiving string NOT NULL,
					    CommentID INTEGER PRIMARY KEY AUTOINCREMENT,
						CommentMessage string NOT NULL,
						UserPutting string NOT NULL,
					    DataPost string NOT NULL,
					    FOREIGN KEY (UserReceiving) 
					      REFERENCES Users (Name) 
					       ON DELETE CASCADE 
					         ON UPDATE CASCADE
					    FOREIGN KEY (UserPutting) 
					      REFERENCES Users (Name) 
					       ON DELETE CASCADE 
					         ON UPDATE CASCADE
					    FOREIGN KEY (PhotoID) 
					      REFERENCES Photo (PhotoID) 
					       ON DELETE CASCADE 
					         ON UPDATE CASCADE
					);

				CREATE TABLE IF NOT EXISTS Like (

						UserPutting string NOT NULL,
						PhotoID string NOT NULL,
						UserReceiving string NOT NULL,
						DataPost string NOT NULL,
						PRIMARY KEY (PhotoID, UserPutting)
						FOREIGN KEY (UserReceiving) 
						REFERENCES Users (Name) 
							ON DELETE CASCADE 
							ON UPDATE CASCADE
						FOREIGN KEY (UserPutting) 
						REFERENCES Users (Name) 
							ON DELETE CASCADE 
							ON UPDATE CASCADE
						FOREIGN KEY (PhotoID) 
						REFERENCES Photo (PhotoID) 
							ON DELETE CASCADE 
							ON UPDATE CASCADE
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
