package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/VitorEmanoel/books-loan/api/graphql/generated"
)

func NewGraphQL(app fiber.Router) {
	playgroundServer := playground.Handler("GraphQL", "/api/graphql")

	app.Get("/playground", func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundServer)(ctx.Context())
		return nil
	})

	app.Post("/graphql", func(ctx *fiber.Ctx) error {
		graphqlServer := handler.NewDefaultServer(
			generated.NewExecutableSchema(generated.Config{
				Resolvers: &Resolver{},
			}))
		fasthttpadaptor.NewFastHTTPHandler(graphqlServer)(ctx.Context())
		return nil
	})
}
