package main

import (
	"html/template"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zhbhmd/go-todo/database"
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

	router := mux.NewRouter()

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/user.html"))
		users := database.FindAllUser()
		data := UserPageData{
			Title:    "Users",
			UserList: users,
		}
		tmpl.Execute(w, data)
	})

	router.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("web/todo.html"))
		users := database.FindAllUser()
		data := UserPageData{
			Title:    "Users",
			UserList: users,
		}
		tmpl.Execute(w, data)
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html")
	})

	http.ListenAndServe(":8000", router)

}
