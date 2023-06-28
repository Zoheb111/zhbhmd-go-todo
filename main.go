package main

import (
	"html/template"

	"github.com/Zoheb111/zhbhmd-go-todo/database"
	_ "github.com/go-sql-driver/mysql"

	"net/http"
)

func init() {
	database.ConnectDB()
}

func main() {

	database.Migrate(database.DB)
	database.Seed(database.DB)

	type UserPageData struct {
		Title    string
		UserList []database.User
	}

	tmpl := template.Must(template.ParseFiles("web/layout.html"))
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := database.FindAllUser()
		data := UserPageData{
			Title:    "Users",
			UserList: users,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})
	http.ListenAndServe(":8000", nil)

}
