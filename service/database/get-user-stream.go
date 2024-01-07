package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetUserStream(UserId uint64) (components.PhotoStreamList, error) {
	var ToReturn components.PhotoStreamList
	rows, err := db.c.Query(
		`
		SELECT * 
		FROM Photos 
		WHERE PublisherId IN (
			SELECT FollowedId
			FROM Follows 
			WHERE FollowerId = ?
		)
		ORDER BY ReleaseDate
		LIMIT 20
		`, UserId)
	if err != nil {
		return ToReturn, err
	}
	defer rows.Close()

	for rows.Next() {
		var TempPhoto components.PhotoStreamListElement
		var userId uint64
		if err := rows.Scan(&TempPhoto.Id, &TempPhoto.ReleaseDate, &TempPhoto.Caption, &userId, &TempPhoto.Likes); err != nil {
			return ToReturn, err
		}
		TempPhoto.PublisherName, err = db.GetUserName(userId)
		if err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
	}

	if err := rows.Err(); err != nil {
		return ToReturn, err
	}

	return ToReturn, err
}
