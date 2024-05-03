package main

import (
	"gin-todo-api/controllers"
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	models.InitTodoDB()
	models.InitUserDB()

	router := gin.Default()

	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos", controllers.GetTodos)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.DeleteTodo)

	router.POST("/users/register", controllers.Register)
	router.POST("/users/login", controllers.LoginUser)

	return router
}

func main() {

	router := setupRouter()
	router.Run(":8080")
}
