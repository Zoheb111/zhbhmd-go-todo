package models

import (
	"time"

	"github.com/zhbhmd/go-todo/pkg/config"
	"gorm.io/gorm"
)

type Todo struct {
	Id          int64     `gorm:"column:id;primaryKey" json:"id"`
	Heading     string    `gorm:"column:heading" json:"heading"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	CreatedBy   string    `gorm:"column:created_by" json:"created_by"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (t *Todo) CreateTodo() *Todo {
	db.Create(&t)
	return t
}

func GetAllTodos() []Todo {
	var Todos []Todo
	db.Find(&Todos)
	return Todos
}

func GetTodoById(Id int64) (*Todo, *gorm.DB) {
	var todo Todo
	db := db.Where("ID=?", Id).Find(&todo)
	return &todo, db
}

func GetTodoByUserId(Id int64) (*Todo, *gorm.DB) {
	var todo Todo
	db := db.Where("Created_By=?", Id).Find(&todo)
	return &todo, db
}

func DeleteTodo(ID int64) Todo {
	var todo Todo
	db.Where("ID=?", ID).Delete(todo)
	return todo
}
