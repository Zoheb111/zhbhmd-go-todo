package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func Seed(db *sql.DB) { // Insert a new user
	username := "admin"
	password := "admin"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO user (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)

	username = "zhbhmd"
	password = "zhbhmd"
	createdAt = time.Now()

	result, err = db.Exec(`INSERT INTO user (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err = result.LastInsertId()
	fmt.Println(id)

	heading := "Learn Go language"
	description := "Spend 40 mins everyday"
	createdBy := "zhbhmd"
	createdAt = time.Now()

	result, err = db.Exec(`INSERT INTO todo (heading, description, created_at, created_by) VALUES (?, ?, ?, ?)`, heading, description, createdAt, createdBy)
	if err != nil {
		log.Fatal(err)
	}

	id, err = result.LastInsertId()
	fmt.Println(id)

}
