package command

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VitorEmanoel/books-loan/database/plugins"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestReturnBook(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var loggedUser int64 = 1
	var bookId int64 = 1
	var bookLoanId = 1
	var fromUser = 2
	var lentAt = time.Now().AddDate(0, -1, 0)
	var returnAt = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(bookId))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "book_loans" WHERE book_id = ($1) AND to_user = ($2) ORDER BY returned_at desc LIMIT 1`)).
		WithArgs(bookId, loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "book_id", "from_user", "to_user", "lent_at", "returned_at"}).
			AddRow(bookLoanId, bookId, fromUser, loggedUser, lentAt, nil),
		)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "book_loans" SET "returned_at"=$1 WHERE id = $2`)).
		WithArgs(returnAt, bookLoanId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	bookLoan, err := mediator.Send(&ReturnBookRequest{
		LoggedUserId: loggedUser,
		BookId:       bookId,
		Date: &returnAt,
	})
	book := bookLoan.(*models.BookLoan)
	assert.Nil(t, err)
	assert.NotNil(t, bookLoan)
	assert.Equal(t, &returnAt, book.ReturnedAt)
}

func TestReturnBookWithNotFoundUser(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var loggedUser int64 = 1
	var bookId int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	bookLoan, err := mediator.Send(&ReturnBookRequest{
		LoggedUserId: loggedUser,
		BookId:       bookId,
	})
	assert.Nil(t, bookLoan)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, LoggedUserNotfound)
}

func TestReturnBookNotFound(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var loggedUser int64 = 1
	var bookId int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	bookLoan, err := mediator.Send(&ReturnBookRequest{
		LoggedUserId: loggedUser,
		BookId:       bookId,
	})
	assert.Nil(t, bookLoan)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, BookNotFound)
}

func TestReturnBookNotBorrowed(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var loggedUser int64 = 1
	var bookId int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(bookId))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "book_loans" WHERE book_id = ($1) AND to_user = ($2) ORDER BY returned_at desc LIMIT 1`)).
		WithArgs(bookId, loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "book_id", "from_user", "to_user", "lent_at", "returned_at"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	bookLoan, err := mediator.Send(&ReturnBookRequest{
		LoggedUserId: loggedUser,
		BookId:       bookId,
	})
	assert.Nil(t, bookLoan)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, BookNotBorrowed)
}

func TestReturnBookAlreadyReturned(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var loggedUser int64 = 1
	var bookId int64 = 1
	var bookLoanId = 1
	var fromUser = 2
	var lentAt = time.Now().AddDate(0, -1, 0)
	var returnAt = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(bookId))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "book_loans" WHERE book_id = ($1) AND to_user = ($2) ORDER BY returned_at desc LIMIT 1`)).
		WithArgs(bookId, loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "book_id", "from_user", "to_user", "lent_at", "returned_at"}).
			AddRow(bookLoanId, bookId, fromUser, loggedUser, lentAt, returnAt),
		)


	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	bookLoan, err := mediator.Send(&ReturnBookRequest{
		LoggedUserId: loggedUser,
		BookId:       bookId,
	})
	assert.Nil(t, bookLoan)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, BookAlreadyReturn)
}