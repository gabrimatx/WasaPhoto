package database

import "fmt"

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteLike(IdPhoto uint64, UserLikeId uint64) error {
	result, err := db.c.Exec("DELETE FROM Likes WHERE PhotoId = ? AND UserId = ?", IdPhoto, UserLikeId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Check if any rows were affected
	if rowsAffected == 0 {
		return fmt.Errorf("Like not found for PhotoId: %d", IdPhoto)
	}

	return nil
}
