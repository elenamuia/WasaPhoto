package database

import "fmt"

func (db *appdbimpl) GetBanningList(userid string) (listban []string, err error) {

	rows, err1 := db.c.Query("SELECT b.Banning FROM Banned as b WHERE Banned = ? ", userid)
	fmt.Println("rows")

	if err1 != nil {
		return listban, err
	}

	for rows.Next() {
		var ban string
		err2 := rows.Scan(&ban)
		if err2 != nil {
			fmt.Println("ban: " + ban)
			return listban, err
		}

		listban = append(listban, ban)
		fmt.Println(listban)
	}
	rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return listban, nil
}
