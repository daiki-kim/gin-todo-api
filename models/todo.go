package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content     string `json:"content"`
	IsCompleted bool   `json:"is_completed" default:"false"`
}

var TodoDB *gorm.DB

func InitTodoDB() {
	var err error
	TodoDB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	TodoDB.AutoMigrate(&Todo{})
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	result := TodoDB.Find(&todos)
	return todos, result.Error
}

func GetTodoById(ID uint) (Todo, error) {
	var todo Todo
	result := TodoDB.First(&todo, ID)
	return todo, result.Error
}

func CreateNewTodo(todo Todo) (Todo, error) {
	result := TodoDB.Create(&todo)
	return todo, result.Error
}

func UpdateTodo(todo Todo) (Todo, error) {
	result := TodoDB.Save(&todo)
	return todo, result.Error
}

func DeleteTodo(ID uint) error {
	result := TodoDB.Delete(&Todo{}, ID)
	return result.Error
}
