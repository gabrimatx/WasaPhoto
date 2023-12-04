package database

import (
	"github.com/gabrimatx/WasaPhoto/service/api"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) AddComment(PhotoId int, Commnt api.Comment) error {
	_, err := db.c.Exec(
		"INSERT INTO Comments(Id, File, ReleaseDate, Caption, PublisherId, Likes) VALUES (?, ?, ?, ?, ?, 0)",
		photo.Id, photo.File, currentTime.Format("01-02-2006"), photo.Caption, photo.PublisherId)
	return photo, err
}
