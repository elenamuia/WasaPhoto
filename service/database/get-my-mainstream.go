package database

// GetName is an example that shows you how to query data
func (db *appdbimpl) GetMyMainstream() ([]Photo, error) {
	var ArrayofPhotos []Photo

	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT PhotoID, UserID, Photo, NumComment, NumLike, DataPost FROM Photo ORDER BY DataPost DESC `)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	for rows.Next() {
		var photo Photo
		err = rows.Scan(&photo.ID, &photo.UserID, &photo.NumComm, &photo.NumLikes, &photo.Datapost, &photo.PhotoStructure)
		if err != nil {
			return nil, err
		}

		ArrayofPhotos = append(ArrayofPhotos, photo)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return (ArrayofPhotos), nil
}
