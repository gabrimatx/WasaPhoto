package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(Commnt components.Comment) error {
	_, err := db.c.Exec(
		"INSERT INTO Comments(PhotoId, UserId, Text_Comment) VALUES (?, ?, ?)",
		Commnt.PhotoId, Commnt.UserId, Commnt.Text_Comment)
	return err
}
