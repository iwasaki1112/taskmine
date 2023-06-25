package main

import (
	"log"
	"taskmine/application"
	"taskmine/config"
	"taskmine/domain/entity"
	"taskmine/infrastructure"
	"taskmine/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("No get db info")
	}

	DB, err := infrastructure.ConnectMysql(dbConfig)
	if err != nil {
		log.Fatal("faild to connect Mysql")
	}

	err = DB.AutoMigrate(&entity.Task{})
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := infrastructure.NewMysqlTaskRepository(DB)
	taskInteractor := application.NewTaskInteractor(taskRepository)
	taskHandler := http.NewTaskHandler(taskInteractor)
	r := gin.Default()

	taskGroup := r.Group("/task")
	taskGroup.POST("/", taskHandler.CreateTask)
	taskGroup.PATCH("/", taskHandler.UpdateTask)
	taskGroup.DELETE("/", taskHandler.DeleteTask)
	r.Run(":8080")
}
