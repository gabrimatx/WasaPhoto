package database

func (db *appdbimpl) FollowUser(IdUserToFollow uint64, IdFollowingUser uint64) error {
	_, err := db.c.Exec(
		"INSERT INTO Follows(FollowerId, FollowedId) VALUES (?, ?)", IdFollowingUser, IdUserToFollow)
	return err
}
