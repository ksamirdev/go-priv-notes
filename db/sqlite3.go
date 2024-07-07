package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/samocodes/go-priv-notes/env"

	_ "github.com/mattn/go-sqlite3"
)

func Load() *sql.DB {
	if env.DefaultConfig.ENVIRONMENT == "dev" {
		os.Remove("./main.db")
	}

	db, err := sql.Open("sqlite3", "./main.db")
	if err != nil {
		log.Fatalf("[database] Failed to connect database: %s\n", err.Error())
		return nil
	}

	query := `
		CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY,
			pin TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS notes (
			id TEXT PRIMARY KEY,
			content TEXT NOT NULL,
			username TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

			FOREIGN KEY (username) REFERENCES users(username)
		);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("[database] error while seeding tables: %s", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("[database] error while pinging database: %s", err.Error())
		return nil
	}

	log.Println("[database] connected!")

	return db
}
