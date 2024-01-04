package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetProfilePhotos(UserId uint64) (components.PhotoList, error) {
	var ToReturn components.PhotoList
	rows, err := db.c.Query(
		`
		SELECT * 
		FROM Photos 
		WHERE PublisherId = ?
		ORDER BY ReleaseDate
		LIMIT 20
		`, UserId)
	if err != nil {
		return ToReturn, err
	}
	defer rows.Close()

	if !rows.Next() {
		return ToReturn, err
	} else {
		var TempPhoto components.PhotoListElement
		if err := rows.Scan(&TempPhoto.Id, &TempPhoto.ReleaseDate, &TempPhoto.Caption, &TempPhoto.PublisherId, &TempPhoto.Likes); err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
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
