package database

import "github.com/gabrimatx/WasaPhoto/service/api"

func (db *appdbimpl) GetUserStream(UserId int) (api.PhotoList, error) {
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
	var to_return api.PhotoList
	for rows.Next() {

	}
	return photo, err
}
