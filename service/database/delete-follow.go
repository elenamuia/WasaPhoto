package database

func (db *appdbimpl) DeleteFollow(follow Follow) error {
	res, err1 := db.c.Exec(`DELETE FROM Follower WHERE FollowerID=? AND FollowedID = ?`, follow.FollowerID, follow.FollowedID)
	if err1 != nil {
		return err1
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrUserDoesNotExist
	}
	return nil
}
