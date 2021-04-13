package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/VitorEmanoel/books-loan/api/graphql"
)

// NewAPI create routers of GraphQL API and /health router
func NewAPI(app *fiber.App) {
	app.Get("/health", Health)
	var api = app.Group("/api")
	{
		graphql.NewGraphQL(api)
	}
}
