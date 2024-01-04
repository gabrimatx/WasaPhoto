package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetUserStream(UserId uint64) (components.PhotoList, error) {
	var ToReturn components.PhotoList
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
		`, UserId)
	if err != nil {
		return ToReturn, err
	}
	defer rows.Close()

	if !rows.Next() {
		return ToReturn, err
	}

	for rows.Next() {
		var TempPhoto components.PhotoListElement
		if err := rows.Scan(&TempPhoto.Id, &TempPhoto.ReleaseDate, &TempPhoto.Caption, &TempPhoto.PublisherId, &TempPhoto.Likes); err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
	}

	if err := rows.Err(); err != nil {
		return ToReturn, err
	}

	return ToReturn, err
}
