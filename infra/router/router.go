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

type AuthHandler interface {
	Auth(c *gin.Context)
}

type Router struct {
	API         *gin.Engine
	taskHandler TaskHandler
	authHandler AuthHandler
}

func NewRouter(taskHandler TaskHandler, authHandler AuthHandler) *Router {
	api := gin.Default()

	taskGroup := api.Group("/task")
	taskGroup.POST("/", taskHandler.CreateTask)
	taskGroup.PATCH("/", taskHandler.UpdateTask)
	taskGroup.DELETE("/", taskHandler.DeleteTask)

	authGroup := api.Group("/auth")
	authGroup.GET("/", authHandler.Auth)

	return &Router{
		API:         api,
		taskHandler: taskHandler,
		authHandler: authHandler,
	}
}

func (r *Router) StartServer() {
	err := r.API.Run(":8080")
	if err != nil {
		log.Fatal("faild to run server")
	}
}
