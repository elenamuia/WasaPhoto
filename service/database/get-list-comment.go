package database

import "fmt"

func (db *appdbimpl) GetListComments(photoid int) (arrayofcom []Comment, err error) {

	rows, err1 := db.c.Query("SELECT PhotoID, UserReceiving, CommentID, CommentMessage, UserPutting FROM Comments WHERE PhotoID = ? ", photoid)
	if err1 != nil {
		fmt.Println("err1: ", err1)
		return arrayofcom, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var comm Comment
		err2 := rows.Scan(&comm.PhotoID, &comm.UserRec, &comm.CommentID, &comm.CommMessage, &comm.UserPut)
		if err2 != nil {
			fmt.Println("err2: ", err2)
			return arrayofcom, err
		}

		arrayofcom = append(arrayofcom, comm)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("err: ", err1)
		return nil, err
	}

	return arrayofcom, nil
}
