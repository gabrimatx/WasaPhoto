package database

import (
	"time"

	"github.com/gabrimatx/WasaPhoto/service/api"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) UploadPhoto(photo api.Photo) (api.Photo, error) {
	currentTime := time.Now()
	_, err := db.c.Exec(
		"INSERT INTO Photos(Id, File, ReleaseDate, Caption, PublisherId, Likes) VALUES (?, ?, ?, ?, ?, 0)",
		photo.Id, photo.File, currentTime.Format("01-02-2006"), photo.Caption, photo.PublisherId)
	return photo, err
}
