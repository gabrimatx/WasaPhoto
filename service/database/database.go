/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webcomponents` executable for the database data source name (add it to the
main.WebcomponentsConfiguration structure):

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

Then you can initialize the AppDatabase and pass it to the components package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"

	components "github.com/gabrimatx/WasaPhoto/service"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	//photos
	UploadPhoto(photo components.Photo) (components.Photo, error)
	DeletePhoto(id int) error

	//users
	SetUsername(UserId int, new_username string) error
	InsertUser(newUsername string) (int, error)
	DeleteUser(UserId int) error
	GetUser(Username string) (int, error)
	GetUserStream(UserId int) (components.PhotoList, error)

	//comments
	AddComment(Commnt components.Comment) error
	DeleteComment(commentId int) error

	//Likes
	LikePhoto(IdPhoto int, UserLikeId int) error
	DeleteLike(IdPhoto int, UserLikeId int) error

	//follows
	FollowUser(IdUserToFollow int, IdFollowingUser int) error
	DeleteFollow(IdUserToNotFollow int, IdFollowingUser int) error

	//bans
	BanUser(IdUserToBan int, IdUser int) error
	DeleteBan(IdUserToUnban int, IdUser int) error
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

	tableName = "Photos"
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Photos (
					Id INTEGER PRIMARY KEY AUTOINCREMENT,
					ReleaseDate VARCHAR(10),
					Caption TEXT,
					PublisherId INTEGER,
					Likes INTEGER,
					FOREIGN KEY (PublisherId) REFERENCES Users(Id)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Users"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (
					Id INTEGER PRIMARY KEY AUTOINCREMENT,
					Name VARCHAR(100)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Comments"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Comments (
					Id INTEGER PRIMARY KEY AUTOINCREMENT,
					PhotoId INTEGER,
					UserId INTEGER,
					Text_Comment TEXT,
					FOREIGN KEY (UserId) REFERENCES Users(Id),
					FOREIGN KEY (PhotoId) REFERENCES Photos(Id)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Likes"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Likes (
					PhotoId INTEGER,
					UserId INTEGER,
					PRIMARY KEY (PhotoId, UserId),
					FOREIGN KEY (UserId) REFERENCES Users(Id),
					FOREIGN KEY (PhotoId) REFERENCES Photos(Id)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Follows"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Follows (
					FollowerId INTEGER,
					FollowedId INTEGER,
					PRIMARY KEY (FollowerId, FollowedId),
					FOREIGN KEY (FollowerId) REFERENCES Users(Id),
					FOREIGN KEY (FollowedId) REFERENCES Users(Id)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Bans"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Bans (
					BannerId INTEGER,
					BannedId INTEGER,
					PRIMARY KEY (BannerId, BannedId),
					FOREIGN KEY (BannerId) REFERENCES Users(Id),
					FOREIGN KEY (BannedId) REFERENCES Users(Id)
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}
	return &appdbimpl{
		c: db,
	}, nil
}
