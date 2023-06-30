package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zhbhmd/go-todo/pkg/models"
	"github.com/zhbhmd/go-todo/pkg/utils"
)

var Todo models.Todo

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()
	res, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodoByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todos, _ := models.GetTodoByUserId(ID)
	res, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, _ := models.GetTodoById(ID)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}
	utils.ParseBody(r, todo)
	b := todo.CreateTodo()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo := models.DeleteTodo(ID)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var todo = &models.Todo{}
	utils.ParseBody(r, todo)
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	ID, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todoDetails, db := models.GetTodoById(ID)
	if todo.Heading != "" {
		todoDetails.Heading = todo.Heading
	}

	if todo.Description != "" {
		todoDetails.Description = todo.Description
	}

	db.Save(&todoDetails)
	res, _ := json.Marshal(todoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
