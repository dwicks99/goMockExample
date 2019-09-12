package util

import (
	"github.com/go-pg/pg"
)

// DbConnect ...
func DbConnect() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "wheaties",
		Database: "postgres",
	})

	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	if err != nil {
		return nil, err
	}
	return db, nil
}
