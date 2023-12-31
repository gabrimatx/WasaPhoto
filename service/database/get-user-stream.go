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
		if err := rows.Scan(&placeholderId, &TempPhoto.ReleaseDate, &TempPhoto.Caption, &TempPhoto.PublisherId, &TempPhoto.Likes); err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
	}
	return ToReturn, err
}
