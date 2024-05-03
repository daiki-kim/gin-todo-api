package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Todo{})
	log.Println("Database migrated")
}

type Todo struct {
	gorm.Model
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed" default:"false"`
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	result := DB.Find(&todos)
	return todos, result.Error
}

func GetTodoById(ID uint) (Todo, error) {
	var todo Todo
	result := DB.First(&todo, ID)
	return todo, result.Error
}

func CreateNewTodo(todo Todo) (Todo, error) {
	result := DB.Create(&todo)
	return todo, result.Error
}

func UpdateTodo(todo Todo) (Todo, error) {
	result := DB.Save(&todo)
	return todo, result.Error
}

func DeleteTodo(ID uint) error {
	result := DB.Delete(&Todo{}, ID)
	return result.Error
}
