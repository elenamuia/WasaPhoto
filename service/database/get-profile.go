package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetProfile(userid int) (p Profile, err error) {

	rows, err := db.c.Query("SELECT p.PhotoID, p.UserID, count(f.FollowerID) as NumFollowed,  count(f.FollowedID) as NumFollower FROM Photo as p, Follower as f WHERE UserID=  ? ", userid)
	if err != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var pr Profile
		err = rows.Scan(&pr.UserID, &pr.Photos, &pr.Followed, &pr.Follower)
		if err != nil {
			return p, err
		}

	}
	if err = rows.Err(); err != nil {
		return p, err
	}

	return p, nil
}
