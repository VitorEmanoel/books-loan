package main

import (
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/VitorEmanoel/books-loan/api"
	"github.com/VitorEmanoel/books-loan/database"
)

func NewApp(environment *Environment) *fiber.App {
	var app = fiber.New()
	mediator.NewMediator()
	db := database.Open(environment.GetDatabaseInfo())
	err := database.Migrate(db)
	if err != nil {
		logrus.Fatalln("Error in migrate database. Error: ", err.Error())
	}
	app.Use(database.UseDatabase(db))
	api.NewAPI(app)
	return app
}
