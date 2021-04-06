package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/VitorEmanoel/books-loan/api/graphql"
)

func NewAPI(app *fiber.App) {
	var api = app.Group("/api")
	{
		graphql.NewGraphQL(api)
	}
}
