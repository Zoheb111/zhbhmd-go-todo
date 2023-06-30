package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zhbhmd/go-todo/pkg/config"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Id        int64     `json:"id"`
	Username  string    `gorm:"" json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var user User
	db := db.Where("ID=?", Id).Find(&user)
	return &user, db
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
