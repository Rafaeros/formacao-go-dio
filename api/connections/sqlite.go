package connections

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbFile = "client.db"

func Connect() (*sql.DB, error) {

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("Database file does not exist. Creating...")
		file, err := os.Create(dbFile)
		if err != nil {
			log.Println("Error creating database file:", err)
			return nil, err
		}
		file.Close()
		log.Println("Database file created successfully.")
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Println("Error opening database:", err)
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		phone INTEGER
	);`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Println("Error creating table:", err)
		return nil, err
	}

	log.Println("Database connected successfully.")
	return db, nil
}

func Disconnect(db *sql.DB) {
	db.Close()
}