package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/VitorEmanoel/books-loan/api/graphql/generated"
	"github.com/VitorEmanoel/books-loan/database"
	"github.com/VitorEmanoel/books-loan/repository"
)

func NewGraphQL(app fiber.Router) {
	playgroundServer := playground.Handler("GraphQL", "/api/graphql")

	app.Get("/playground", func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(playgroundServer)(ctx.Context())
		return nil
	})

	app.Post("/graphql", func(ctx *fiber.Ctx) error {
		db := database.GetDatabase(ctx)
		graphqlServer := handler.NewDefaultServer(
			generated.NewExecutableSchema(generated.Config{
				Resolvers: &Resolver{
					DB: db,
					Repository: repository.NewRepository(db),
				},
			}))
		fasthttpadaptor.NewFastHTTPHandler(graphqlServer)(ctx.Context())
		return nil
	})
}
