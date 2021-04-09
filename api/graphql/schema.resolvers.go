package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/api/graphql/generated"
	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/application/command"
	"github.com/VitorEmanoel/books-loan/application/query"
	"github.com/VitorEmanoel/books-loan/models"
)

func (r *bookResolver) CreatedAt(ctx context.Context, obj *models.Book) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookLoanResolver) Book(ctx context.Context, obj *models.BookLoan) (*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookLoanResolver) LentAt(ctx context.Context, obj *models.BookLoan) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookLoanResolver) ReturnedAt(ctx context.Context, obj *models.BookLoan) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	user, err := mediator.Send(&command.CreateUserRequest{
		BaseRequest: application.NewRequest(r.Repository),
		UserInput: input,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (r *mutationResolver) AddBookToMyCollection(ctx context.Context, loggedUserID int64, input models.AddBookInput) (*models.Book, error) {
	book, err := mediator.Send(&command.AddBookRequest{
		BaseRequest: application.NewRequest(r.Repository),
		LoggedUserId: loggedUserID,
		BookInput:    input,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return book.(*models.Book), nil
}

func (r *mutationResolver) LendBook(ctx context.Context, loggedUserID int64, input models.LendBookInput) (*models.BookLoan, error) {
	bookLoan, err := mediator.Send(&command.LendBookRequest{
		BaseRequest: application.NewRequest(r.Repository),
		LoggedUserId:  loggedUserID,
		LendBookInput: input,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return bookLoan.(*models.BookLoan), nil
}

func (r *mutationResolver) ReturnBook(ctx context.Context, loggedUserID int64, bookID int64) (*models.BookLoan, error) {
	bookLoan, err := mediator.Send(&command.ReturnBookRequest{
		BaseRequest: application.NewRequest(r.Repository),
		LoggedUserId: loggedUserID,
		BookId:       bookID,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return bookLoan.(*models.BookLoan), nil
}

func (r *queryResolver) User(ctx context.Context, id int64) (*models.User, error) {
	user, err := mediator.Send(&query.FindUserRequest{
		BaseRequest: application.NewRequest(r.Repository),
		UserId:  id,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Collection(ctx context.Context, obj *models.User) ([]*models.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LentBooks(ctx context.Context, obj *models.User) ([]*models.BookLoan, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) BorrowedBooks(ctx context.Context, obj *models.User) ([]*models.BookLoan, error) {
	panic(fmt.Errorf("not implemented"))
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// BookLoan returns generated.BookLoanResolver implementation.
func (r *Resolver) BookLoan() generated.BookLoanResolver { return &bookLoanResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type bookResolver struct{ *Resolver }
type bookLoanResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
