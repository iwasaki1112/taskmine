package http

import (
	"net/http"
	"taskmine/application"
	"taskmine/domain/entity"

	"github.com/gin-gonic/gin"
)

type TaskInteractor interface {
	CreateTask(input application.CreateTaskInput) (*entity.Task, error)
	UpdateTask(input application.UpdateTaskInput) (*entity.Task, error)
	DeleteTask(input application.DeleteTaskInput) (*entity.Task, error)
}

type TaskHandler struct {
	interactor TaskInteractor
}

func NewTaskHandler(interactor TaskInteractor) *TaskHandler {
	return &TaskHandler{
		interactor: interactor,
	}
}

func (h TaskHandler) CreateTask(c *gin.Context) {
	var input application.CreateTaskInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.interactor.CreateTask(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h TaskHandler) UpdateTask(c *gin.Context) {
	var input application.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.interactor.UpdateTask(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h TaskHandler) DeleteTask(c *gin.Context) {
	var input application.DeleteTaskInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.interactor.DeleteTask(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
