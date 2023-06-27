package main

import (
	"log"
	"taskmine/application"
	"taskmine/config"
	"taskmine/domain/entity"
	"taskmine/infra/database"
	"taskmine/infra/notifier"
	"taskmine/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("No get db info")
	}

	DB, err := database.ConnectMysql(config)
	if err != nil {
		log.Fatal("faild to connect Mysql")
	}

	err = DB.AutoMigrate(&entity.Task{})
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := database.NewMysqlTaskRepository(DB)
	slackNotifier := notifier.NewSlackNotifier(config.SlackWebHookURL)
	taskInteractor := application.NewTaskInteractor(taskRepository, slackNotifier)
	taskHandler := http.NewTaskHandler(taskInteractor)
	r := gin.Default()

	taskGroup := r.Group("/task")
	taskGroup.POST("/", taskHandler.CreateTask)
	taskGroup.PATCH("/", taskHandler.UpdateTask)
	taskGroup.DELETE("/", taskHandler.DeleteTask)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("faild to run server")
	}
}
