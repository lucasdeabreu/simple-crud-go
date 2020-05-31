package database

import (
	"database/sql"
	"log"

	// Database Driver
	_ "github.com/mattn/go-sqlite3"
)

// Db is a database handler
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./infra/user.db")
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = createTables(); err != nil {
		log.Fatal(err)
	}
}

func createTables() error {
	createSql := `
		create table if not exists User (
			id integer primary key autoincrement,
			name text,
			email text,
			document text
		)
	`
	stmt, err := Db.Prepare(createSql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
