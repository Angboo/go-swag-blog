package outbound

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewMemoryDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("db open failed, err: %q", err)
	}
	return db
}
