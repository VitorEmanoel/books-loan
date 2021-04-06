package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/VitorEmanoel/books-loan/api"
)

func NewApp() *fiber.App {
	var app = fiber.New()
	api.NewAPI(app)
	return app
}
