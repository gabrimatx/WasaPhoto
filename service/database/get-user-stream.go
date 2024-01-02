package database

import (
	"fmt"
	"log"

	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetUserStream(UserId uint64) (components.PhotoList, error) {
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
	defer rows.Close()
	var ToReturn components.PhotoList
	if err != nil {
		log.Fatal(err)
		return ToReturn, err
	}

	if !rows.Next() {
		fmt.Println("No photos found for the given criteria.")
		return ToReturn, err
	}

	for rows.Next() {
		var TempPhoto components.PhotoListElement
		if err := rows.Scan(&TempPhoto.Id, &TempPhoto.ReleaseDate, &TempPhoto.Caption, &TempPhoto.PublisherId, &TempPhoto.Likes); err != nil {
			return ToReturn, err
		}
		ToReturn.PList = append(ToReturn.PList, TempPhoto)
	}
	return ToReturn, err
}
