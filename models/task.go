package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB // main.goで初期化する

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
	var tasks = []Task{}                          // 変数宣言時点で空のメモリを割り当てるため、データが0件でも空の配列を返す（var tasks []Taskだとデータがないときにnilを返す）
	if err := DB.Find(&tasks).Error; err != nil { // tasksが示す型Taskのテーブルから全てのレコードを取得し[]tasksに直接格納
		return nil, err
	}
	return tasks, nil
}

func FindTaskById(ID uint) (*Task, error) {
	var task = Task{} // 明示的な初期化を表す。（var task TasでもGolangの変数宣言時に自動的にゼロ値で初期化される特性上、機能的差異はない）
	if err := DB.First(&task, ID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *Task) SaveTask(updateTask Task) (*Task, error) {
	if err := DB.Save(&updateTask).Error; err != nil {
		return nil, err
	}
	return &updateTask, nil
}
