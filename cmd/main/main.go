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
	err := config.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	DBConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	googleAOuthConfig, err := config.LoadGoogleOAuthConfig()
	if err != nil {
		log.Fatal(err)
	}

	DB, err := database.ConnectMysql(DBConfig)
	if err != nil {
		log.Fatal("faild to connect Mysql")
	}

	err = DB.AutoMigrate(&entity.Task{})
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := database.NewMysqlTaskRepository(DB)
	slackNotifier := notifier.NewSlackNotifier(DBConfig.SlackWebHookURL)
	taskInteractor := application.NewTaskInteractor(taskRepository, slackNotifier)
	taskHandler := http.NewTaskHandler(taskInteractor)
	authHandler := http.NewAuthHandler(*googleAOuthConfig)
	router := router.NewRouter(taskHandler, authHandler)
	router.StartServer()

}
