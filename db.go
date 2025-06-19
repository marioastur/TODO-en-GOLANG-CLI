package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func openDB() (*sql.DB, error) {
	return sql.Open("sqlite", "todo.db")
}
