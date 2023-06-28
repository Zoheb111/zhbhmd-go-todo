package main

import (
	_ "github.com/go-sql-driver/mysql"

	"net/http"
)

func main() {

	db := Connect()

	Migrate(db)
	Seed(db)

	// { // Query all users
	// 	type user struct {
	// 		id        int
	// 		username  string
	// 		password  string
	// 		createdAt time.Time
	// 	}

	// 	rows, err := db.Query(`SELECT id, username, password, created_at FROM user`)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer rows.Close()

	// 	var users []user
	// 	for rows.Next() {
	// 		var u user

	// 		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		users = append(users, u)
	// 	}
	// 	if err := rows.Err(); err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Printf("%#v", users)
	// }

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	http.ListenAndServe(":8000", nil)

}
