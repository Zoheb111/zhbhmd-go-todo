package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3307)/go-todo?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db

}
