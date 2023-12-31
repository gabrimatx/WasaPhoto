package database

func (db *appdbimpl) InsertUser(newUsername string) (uint64, error) {
	res, err := db.c.Exec(
		`INSERT INTO Users(Name)
         VALUES (?)`, newUsername)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}
