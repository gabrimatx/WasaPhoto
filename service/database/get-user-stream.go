package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetUserStream(UserId uint64) (components.PhotoList, error) {
	rows, err := db.c.Query(
		`
		SELECT * 
		FROM Photos 
		WHERE PublisherId IN (
			SELECT FollowedId
			FROM UserId JOIN Follows 
			WHERE UserId = FollowerId AND UserId = ?
		)
		ORDER BY ReleaseDate
		`, UserId)
	defer rows.Close()
	var ToReturn components.PhotoList
	for rows.Next() {
		var TempPhoto components.Photo
		var placeholderId uint64
		var dateHolder string
		var likes_holder int
		if err := rows.Scan(&placeholderId, &dateHolder, &TempPhoto.Caption, &TempPhoto.PublisherId, &likes_holder); err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
	}
	return ToReturn, err
}
