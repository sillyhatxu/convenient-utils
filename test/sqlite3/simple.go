package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
CREATE TABLE IF NOT EXISTS user_info
(
  id                 INTEGER PRIMARY KEY AUTOINCREMENT,
  login_name         TEXT NOT NULL,
  password           TEXT NOT NULL,
  status             INTEGER  DEFAULT 1,
  created_time       datetime default current_timestamp,
  last_modified_time datetime default current_timestamp
)
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
