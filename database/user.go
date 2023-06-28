package database

import (
	"log"
	"time"
)

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func FindAllUser() []User {
	rows, err := DB.Query(`SELECT id, username, password, created_at FROM user`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User

		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func FindUserById(id string) {

}

func SaveUser(user *User) {

}

func UpdateUser(id string, user *User) {

}

func DeleteUser(id string) {

}
