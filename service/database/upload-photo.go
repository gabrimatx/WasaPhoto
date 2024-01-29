package database

import (
	"time"

	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) UploadPhoto(photo components.Photo, PublisherId uint64) (uint64, error) {
	currentTime := time.Now().UTC()
	res, err := db.c.Exec(
		"INSERT INTO Photos(ReleaseDate, Caption, PublisherId, Likes) VALUES (?, ?, ?, 0)",
		currentTime.Format("2006-01-02 15:04:05"), photo.Caption, PublisherId)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil

}
