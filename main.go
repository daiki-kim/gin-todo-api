package main

import (
	"gin-todo-api/controllers"
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	taskHandler := controllers.TaskHandler{}
	router.POST("/tasks", taskHandler.CreateTask)
	router.GET("/tasks", taskHandler.GetTasks)

	return router
}

func main() {
	db := models.ConnectToDB() // ConnectToDB()でエラーハンドリングしているからここでは不要？

	models.DB = db // これがないと*gorm.DBに初期化が反映されず、task.goでのdbがnilのままでerrorになる

	models.DB.AutoMigrate(&models.Task{})

	router := setupRouter()
	router.Run(":8080")
}
