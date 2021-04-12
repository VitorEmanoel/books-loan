package main

import (
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/api"
	"github.com/VitorEmanoel/books-loan/database"
)

func Inject(container mediator.Container, db *gorm.DB) {
	container.Inject("db", db)
	container.Inject("repository", repository.NewRepository(db))
}

func NewApp(environment *Environment) *fiber.App {
	var app = fiber.New()
	var med = mediator.NewMediator()
	db := database.Open(environment.GetDatabaseInfo())
	err := database.Migrate(db)
	if err != nil {
		logrus.Fatalln("Error in migrate database. Error: ", err.Error())
	}
	Inject(med.GetContainer(), db)
	api.NewAPI(app)
	return app
}
