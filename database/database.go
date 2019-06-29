package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dbHost string

func SetDB() {
	dbHost = os.Getenv("DATABASE_URL")
	if len(dbHost) == 0 {
		log.Fatal("DATABASE_URL is empty")
	}
	db, err := sql.Open("postgres", dbHost)
	if err != nil {
		log.Fatal("Can't connect db", err.Error())
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS customer(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table fatal error", err.Error())
	}

}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	return db, err
}
