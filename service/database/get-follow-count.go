package database

func (db *appdbimpl) GetFollowedUsers(UserId uint64) (int, error) {
	var id int
	err := db.c.QueryRow(
		`
		SELECT COUNT(*) 
		FROM Follows 
		WHERE FollowerId = ?
		`, UserId).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, err
}

func (db *appdbimpl) GetFollowingUsers(UserId uint64) (int, error) {
	var id int
	err := db.c.QueryRow(
		`
		SELECT COUNT(*) 
		FROM Follows 
		WHERE FollowedId = ?
		`, UserId).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, err
}
