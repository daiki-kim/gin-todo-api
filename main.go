package main

import (
	"gin-todo-api/controllers"
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	models.InitDB()

	router := gin.Default()

	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	return router
}

func main() {

	router := setupRouter()
	router.Run(":8080")
}
