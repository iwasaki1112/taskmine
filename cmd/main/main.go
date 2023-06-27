package main

import (
	"log"
	"taskmine/application"
	"taskmine/config"
	"taskmine/domain/entity"
	"taskmine/infra/database"
	"taskmine/infra/notifier"
	"taskmine/infra/router"
	"taskmine/interface/http"
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
	router := router.NewRouter(taskHandler)
	router.StartServer()

}
