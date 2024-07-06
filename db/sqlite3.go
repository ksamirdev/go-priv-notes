package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Load() {
	os.Remove("./main.db")

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatalf("[database] Failed to connect database: %s\n", err.Error())
		return
	}

	query := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT PRIMARY KEY,
		pin TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS notes (
		id TEXT PRIMARY KEY,
		content TEXT NOT NULL,
		username TEXT,
		
		FOREIGN KEY (username) REFERENCES users(username)
	);
`
	DB = db

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("[database] error while seeding tables: %s", err.Error())
		return
	}

	log.Println("[database] connected!")

}
