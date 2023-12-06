package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) FollowUser(IdUserToFollow int, IdFollowingUser int) error {
	_, err := db.c.Exec(
		"INSERT INTO Follows(FollowerId, FollowedId) VALUES (?, ?)", IdFollowingUser, IdUserToFollow)
	return err
}
