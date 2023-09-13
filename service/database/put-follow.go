package database

func (db *appdbimpl) PutFollow(follow Follow) (string, error) {
	_, err := db.c.Exec(`INSERT INTO Follower(Follower, Followed) VALUES (?,?)`,
		follow.Follower, follow.Followed)

	if err != nil {
		return "", err
	}

	return follow.Followed, nil

}
