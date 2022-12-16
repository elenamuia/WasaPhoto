package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) PutFollow(follow Follow) (err error) {
	_, err = db.c.Exec(`INSERT INTO Follower(FollowerID, FollowedID) VALUES (?,?)`,
		follow.FollowerID, follow.FollowedID)

	if err != nil {
		return err
	}

	return nil

}
