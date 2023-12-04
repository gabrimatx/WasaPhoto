package database

func (db *appdbimpl) SetUsername(UserId int, new_username string) error {
	_, err := db.c.Exec(
		`UPDATE Users 
	     SET Name = ?
		 WHERE Id = ?`, new_username, UserId)
	return err
}
