package database

import (
	"database/sql"
)

func (db *appdbimpl) GetUserIdFromPhotoId(photoId uint64) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT PublisherId FROM Photos WHERE Id = ?", photoId).Scan(&Id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return Id, err
}

func (db *appdbimpl) GetUserIdFromCommentId(commentId uint64) (uint64, error) {
	var Id uint64
	err := db.c.QueryRow("SELECT UserId FROM Comments WHERE Id = ?", commentId).Scan(&Id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return Id, err
}
