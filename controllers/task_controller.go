package controllers

import (
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTodos(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	todoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	gotTodo, err := models.GetTodoById(uint(todoID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
	}

	c.JSON(http.StatusOK, gotTodo)
}

func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTodo, err := models.CreateNewTodo(newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTodo)
}

func UpdateTodo(c *gin.Context) {
	updateTodoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updateTodo models.Todo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateTodo.ID = uint(updateTodoID)
	updatedTodo, err := models.UpdateTodo(updateTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	deleteTodoID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"eeror": "invalid id"})
		return
	}
	if err := models.DeleteTodo(uint(deleteTodoID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
