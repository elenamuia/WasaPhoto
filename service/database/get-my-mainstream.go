package database

func (db *appdbimpl) GetMyMainstream() ([]Photo, error) {
	var ArrayofPhotos []Photo

	rows, err := db.c.Query(`SELECT PhotoID, UserID, Photo, NumComment, NumLike, DataPost FROM Photo ORDER BY DataPost DESC `)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

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
