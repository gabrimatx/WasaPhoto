package database

import (
	"time"

	components "github.com/gabrimatx/WasaPhoto/service"
)

// GetName is an example that shows you how to query data
func (db *appdbimpl) UploadPhoto(photo components.Photo) (components.Photo, error) {
	currentTime := time.Now()
	_, err := db.c.Exec(
		"INSERT INTO Photos(Id, ReleaseDate, Caption, PublisherId, Likes) VALUES (?, ?, ?, ?, 0)",
		photo.Id, currentTime.Format("01-02-2006"), photo.Caption, photo.PublisherId)
	return photo, err
}
