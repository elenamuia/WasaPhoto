package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) PutFollow(follower User, followed User) (err error) {
	_, err = db.c.Exec(`INSERT INTO Follower(FollowerID, FollowedID) VALUES (?,?)`,
		follower.ID, followed.ID)

	if err != nil {
		return err
	}

	return nil

}
