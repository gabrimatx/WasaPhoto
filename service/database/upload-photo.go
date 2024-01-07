package database

import (
	"time"

	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) UploadPhoto(photo components.Photo, PublisherId uint64) (uint64, error) {
	currentTime := time.Now()
	res, err := db.c.Exec(
		"INSERT INTO Photos(ReleaseDate, Caption, PublisherId, Likes) VALUES (?, ?, ?, 0)",
		currentTime.Format("01-02-2006"), photo.Caption, PublisherId)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil

}
