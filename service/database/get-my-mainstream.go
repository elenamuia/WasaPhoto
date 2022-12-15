package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetMyMainstream(user User) (ArrayofPhotos []Photo, err error) {
	var ret []Photo

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT PhotoID, UserID, Photo, NumComment, NumLike, DataPost FROM Photo WHERE UserID = ?`, user.ID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.UserID, &photo.numComm, &photo.numLikes, &photo.datapost, &photo.PhotoStructure)
		if err != nil {
			return nil, err
		}

		ret = append(ret, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ret, nil
}
