package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

type TaskHandler interface {
	CreateTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
}

type Router struct {
	API         *gin.Engine
	taskHandler TaskHandler
}

func NewRouter(taskHandler TaskHandler) *Router {
	api := gin.Default()
	taskGroup := api.Group("/task")
	taskGroup.POST("/", taskHandler.CreateTask)
	taskGroup.PATCH("/", taskHandler.UpdateTask)
	taskGroup.DELETE("/", taskHandler.DeleteTask)
	return &Router{
		API:         api,
		taskHandler: taskHandler,
	}
}

func (r *Router) StartServer() {
	err := r.API.Run(":8080")
	if err != nil {
		log.Fatal("faild to run server")
	}
}
