package database

func (db *appdbimpl) PutFollow(follow Follow) (err error) {
	_, err = db.c.Exec(`INSERT INTO Follower(FollowerID, FollowedID) VALUES (?,?)`,
		follow.FollowerID, follow.FollowedID)

	if err != nil {
		return err
	}

	return nil

}
