package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) DeleteFollow(follower User) (err error) {
	res, err := db.c.Exec(`DELETE FROM Follower WHERE FollowerID=? `, follower.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrUserDoesNotExist
	}
	return nil
}
