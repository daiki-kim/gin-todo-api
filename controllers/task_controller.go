package controllers

import (
	"gin-todo-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (t *TaskHandler) GetTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64) // str->uint64に変換
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	getTask, err := models.FindTaskById(uint(taskID)) // uint64->uintに変換
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
	}

	c.JSON(http.StatusOK, getTask)
}

func (t *TaskHandler) UpdateTask(c *gin.Context) {
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 64) // str->uint64に変換
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	targetTask, err := models.FindTaskById(uint(taskID))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if updatedTask.ID != targetTask.ID {
		targetTask.ID = updatedTask.ID
	}
	if updatedTask.Content != targetTask.Content {
		targetTask.Content = updatedTask.Content
	}
	if updatedTask.DueDate != targetTask.DueDate {
		targetTask.DueDate = updatedTask.DueDate
	}
	if updatedTask.IsCompleted != targetTask.IsCompleted {
		targetTask.IsCompleted = updatedTask.IsCompleted
	}
	targetTask.SaveTask(updatedTask)
	c.JSON(http.StatusOK, targetTask)
}
