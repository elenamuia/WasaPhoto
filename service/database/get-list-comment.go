package database

func (db *appdbimpl) GetListComments(photoid int) (arrayofcom []Comment, err error) {

	rows, err1 := db.c.Query("SELECT * FROM Comments WHERE Photo = ? ", photoid)
	if err1 != nil {
		return arrayofcom, err
	}

	for rows.Next() {
		var comm Comment
		err2 := rows.Scan(&comm.PhotoID, &comm.UserRec, &comm.UserRec, &comm.CommentID, &comm.CommMessage, &comm.UserPut, &comm.Datapost)
		if err2 != nil {

			return arrayofcom, err
		}

		arrayofcom = append(arrayofcom, comm)
	}

	return arrayofcom, nil
}
