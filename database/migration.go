package database

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {

	query := `
	DROP TABLE IF EXISTS todo;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	query = `
		DROP TABLE IF EXISTS user;`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	query = `
		CREATE TABLE user (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	query = `

	CREATE TABLE todo (
		id INT AUTO_INCREMENT,
		heading TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at DATETIME,
		created_by TEXT NOT NULL,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

}
