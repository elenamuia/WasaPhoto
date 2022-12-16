package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

type Users struct {
	ID       int    `json:"id"`
	username string `json:"username"`
}

type Banned struct {
	BannedID  int `json:"BannedID"`
	BanningID int `json:"BanningID"`
}

func (b *Banned) FromDatabaseBanned(ban database.Banned) {
	b.BannedID = ban.BannedID
	b.BanningID = ban.BanningID

}

// ToDatabase returns the fountain in a database-compatible representation
func (user *Banned) ToDatabaseBanned() database.Banned {
	return database.Banned{
		BannedID:  user.BannedID,
		BanningID: user.BanningID,
	}
}
