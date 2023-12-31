package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) FollowUser(IdUserToFollow uint64, IdFollowingUser uint64) error {
	_, err := db.c.Exec(
		"INSERT INTO Follows(FollowerId, FollowedId) VALUES (?, ?)", IdFollowingUser, IdUserToFollow)
	return err
}
