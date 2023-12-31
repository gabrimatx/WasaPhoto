package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteFollow(IdUserToNotFollow uint64, IdFollowingUser uint64) error {
	_, err := db.c.Exec(
		"DELETE FROM Follows WHERE FollowerId = ? AND FollowedId = ?", IdFollowingUser, IdUserToNotFollow)
	return err
}
