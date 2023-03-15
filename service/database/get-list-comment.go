func (db *appdbimpl) GetListComments(photoid int) (arrayofcom []Comments, error err) {

	rows, rr1 := db.c.Query("SELECT * FROM Comments WHERE Photo = ? ", photoid)
	if err1 != nil {
		return arrayofcomm, err
	}
	
	
	for rows.Next() {
		var comm Comments
		err2 := rows.Scan(&comm.PhotoID, &comm.UserRec,&comm.UserRec, &comm.CommentID, &comm.CommMessage, &comm.UserPut, &comm.Datapost)
		if err2 != nil {

			return p, err
		}

		arrayofcom = append(arrayofcom, comm)

	return arrayofcom, nil
}
