package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func ConnectToDB() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Content     string    `json:"content" gorm:"not null"`
	DueDate     time.Time `json:"due_date,omitempty"`
	IsCompleted bool      `json:"is_completed,omitempty" gorm:"default:false"`
}

func (t *Task) Create() (*Task, error) {
	if err := DB.Create(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func FindAllTasks() ([]Task, error) {
	var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil { // tasksが示す型Taskのテーブルから全てのレコードを取得し[]tasksに直接格納
		return nil, err
	}
	return tasks, nil
}
