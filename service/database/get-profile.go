package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetProfile(userid int) (p Profile, err error) {

	rows, err1 := db.c.Query("SELECT p.PhotoID, p.UserID, count(f.FollowerID) as NumFollowed,  count(f.FollowedID) as NumFollower FROM Photo as p, Follower as f WHERE UserID=  ? ", userid)
	if err1 != nil {
		return p, err
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		var pr Profile
		err2 := rows.Scan(&pr.UserID, &pr.Photos, &pr.Followed, &pr.Follower)
		if err2 != nil {
			return p, err
		}

	}
	if err3 := rows.Err(); err3 != nil {
		return p, err
	}

	return p, nil
}
