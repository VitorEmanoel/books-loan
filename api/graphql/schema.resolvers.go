package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/api/graphql/generated"
	"github.com/VitorEmanoel/books-loan/application/command"
	"github.com/VitorEmanoel/books-loan/application/query"
	"github.com/VitorEmanoel/books-loan/models"
)

func (r *bookResolver) CreatedAt(ctx context.Context, obj *models.Book) (string, error) {
	var formattedValue = obj.CreatedAt.Format(time.RFC3339Nano)
	return formattedValue, nil
}

func (r *bookLoanResolver) Book(ctx context.Context, obj *models.BookLoan) (*models.Book, error) {
	book, err := mediator.Send(&query.FindBookRequest{
		BookId:  obj.BookId,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return book.(*models.Book), nil
}

func (r *bookLoanResolver) LentAt(ctx context.Context, obj *models.BookLoan) (string, error) {
	var formattedValue = obj.LentAt.Format(time.RFC3339Nano)
	return formattedValue, nil
}

func (r *bookLoanResolver) ReturnedAt(ctx context.Context, obj *models.BookLoan) (*string, error) {
	if obj.ReturnedAt != nil {
		var formattedValue = obj.ReturnedAt.Format(time.RFC3339Nano)
		return &formattedValue, nil
	}
	return nil, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	user, err := mediator.Send(&command.CreateUserRequest{
		UserInput: input,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (r *mutationResolver) AddBookToMyCollection(ctx context.Context, loggedUserID int64, input models.AddBookInput) (*models.Book, error) {
	book, err := mediator.Send(&command.AddBookRequest{
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
		UserId:  id,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *models.User) (string, error) {
	var formattedValue = obj.CreatedAt.Format(time.RFC3339Nano)
	return formattedValue, nil
}

func (r *userResolver) Collection(ctx context.Context, obj *models.User) ([]*models.Book, error) {
	books, err := mediator.Send(&query.FindBooksRequest{
		UserId:  obj.ID,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return books.([]*models.Book), nil
}

func (r *userResolver) LentBooks(ctx context.Context, obj *models.User) ([]*models.BookLoan, error) {
	booksLoan, err := mediator.Send(&query.FindBookLoansRequest{
		FromUser: obj.ID,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return booksLoan.([]*models.BookLoan), nil
}

func (r *userResolver) BorrowedBooks(ctx context.Context, obj *models.User) ([]*models.BookLoan, error) {
	booksLoan, err := mediator.Send(&query.FindBookLoansRequest{
		ToUser: obj.ID,
	}, mediator.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	return booksLoan.([]*models.BookLoan), nil
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
