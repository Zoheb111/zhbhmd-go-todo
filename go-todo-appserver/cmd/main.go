package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zhbhmd/go-todo/pkg/routes"
)

func main() {

	r := mux.NewRouter()

	routes.RegisterUserRoutes(r)
	routes.RegisterTodoRoutes(r)

	http.Handle("/", r)

	log.Println("Server starting....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
