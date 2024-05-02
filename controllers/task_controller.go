package controllers

import (
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskHandler struct{}

func (t *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	createdTask, err := task.Create()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, *createdTask)
}

func (t *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := models.FindAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, tasks)
}
