package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DBController struct {
	Path       string
	connection *sql.DB
}

func (db *DBController) Init() (err error) {
	// Test that the db file can be opened
	driver, err := sql.Open("sqlite3", db.Path)
	if err != nil {
		return err
	}

	defer driver.Close()
	return nil
}

func (db *DBController) Query(sql string) (err error) {
	return nil
}
