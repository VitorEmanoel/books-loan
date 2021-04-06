package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/VitorEmanoel/books-loan/api/graphql/generated"
	entities "github.com/VitorEmanoel/books-loan/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input entities.CreateUserInput) (*entities.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddBookToMyCollection(ctx context.Context, loggedUserID string, input entities.AddBookInput) (*entities.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LendBook(ctx context.Context, loggedUserID string, input entities.LendBookInput) (*entities.BookLoan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReturnBook(ctx context.Context, loggedUserID string, bookID string) (*entities.BookLoan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*entities.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
