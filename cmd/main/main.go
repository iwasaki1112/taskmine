package main

import (
	"log"
	"taskmine/application"
	"taskmine/config"
	"taskmine/domain/entity"
	"taskmine/domain/repository"
	"taskmine/domain/service"
	"taskmine/infra/database"
	"taskmine/infra/database/mysql"
	"taskmine/infra/notifier"
	"taskmine/infra/router"
	"taskmine/interface/http"
)

func main() {
	// initialize environemnet variables.
	if err := config.LoadEnvironmentVariables(); err != nil {
		log.Fatal(err)
	}

	// initialize db variables.
	DBConfig, err := config.LoadDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	// initialize google AOuth variables.
	googleAOuthConfig, err := config.LoadGoogleOAuthConfig()
	if err != nil {
		log.Fatal(err)
	}

	// initialize database connection.
	var dbConnector database.DbConnector = mysql.NewConnector()
	if err = dbConnector.Connect(DBConfig); err != nil {
		log.Fatal(err)
	}
	if err = dbConnector.AutoMigration(&entity.Task{}); err != nil {
		log.Fatal(err)
	}

	var taskRepository repository.TaskRepository = mysql.NewTaskRepository(dbConnector.GetDB())
	var slackNotifier service.WebhookNotifier = notifier.NewSlackNotifier(DBConfig.SlackWebHookURL)

	var taskInteractor = application.NewTaskInteractor(taskRepository, slackNotifier)
	var taskHandler = http.NewTaskHandler(taskInteractor)
	var authHandler = http.NewAuthHandler(googleAOuthConfig)
	var router *router.Router = router.NewRouter(taskHandler, authHandler)
	router.StartServer()

}
