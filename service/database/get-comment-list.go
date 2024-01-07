package database

import (
	components "github.com/gabrimatx/WasaPhoto/service"
)

func (db *appdbimpl) GetPhotoComments(photoId uint64) (components.CommentList, error) {
	var ToReturn components.CommentList
	rows, err := db.c.Query(
		`
		SELECT * 
		FROM Comments 
		WHERE PhotoId = ?
		ORDER BY Id
		LIMIT 40
		`, photoId)
	if err != nil {
		return ToReturn, err
	}
	defer rows.Close()

	for rows.Next() {
		var TempComment components.CommentListElement
		if err := rows.Scan(&TempComment.Id, &TempComment.PhotoId, &TempComment.PublisherId, &TempComment.CommentText); err != nil {
			return ToReturn, err
		}
		TempComment.PublisherName, err = db.GetUserName(TempComment.PublisherId)
		if err != nil {
			return ToReturn, err
		}
		ToReturn.CList = append(ToReturn.CList, TempComment)
	}

	if err := rows.Err(); err != nil {
		return ToReturn, err
	}

	return ToReturn, err
}
