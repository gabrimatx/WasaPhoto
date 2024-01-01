package database

import (
	"database/sql"
	"errors"
	"fmt"

	components "github.com/gabrimatx/WasaPhoto/service"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// photos
	UploadPhoto(photo components.Photo) (uint64, error)
	DeletePhoto(id uint64) error

	// users
	SetUsername(UserId uint64, new_username string) error
	InsertUser(newUsername string) (uint64, error)
	DeleteUser(UserId uint64) error
	GetUser(Username string) (uint64, error)
	GetUserStream(UserId uint64) (components.PhotoList, error)

	// comments
	AddComment(Commnt components.Comment) error
	DeleteComment(commentId uint64) error

	// Likes
	LikePhoto(IdPhoto uint64, UserLikeId uint64) error
	DeleteLike(IdPhoto uint64, UserLikeId uint64) error

	// follows
	FollowUser(IdUserToFollow uint64, IdFollowingUser uint64) error
	DeleteFollow(IdUserToNotFollow uint64, IdFollowingUser uint64) error

	// bans
	BanUser(IdUserToBan uint64, IdUser uint64) error
	DeleteBan(IdUserToUnban uint64, IdUser uint64) error

	// Utils
	GetUserIdFromCommentId(commentId uint64) (uint64, error)
	GetUserIdFromPhotoId(photoId uint64) (uint64, error)
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
					UserId INTEGER PRIMARY KEY AUTOINCREMENT,
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
