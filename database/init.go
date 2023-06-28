package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func ConnectDB() {

	DB, err = sql.Open("mysql", "root:root@(127.0.0.1:3307)/go-todo?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}

}
