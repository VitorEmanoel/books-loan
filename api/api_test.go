package api

import (
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/VitorEmanoel/books-loan/common"
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func App() http.HandlerFunc {
	var app = fiber.New()
	NewAPI(app)
	return common.FiberToHandlerFunc(app)
}

type Rows []string

func TestSimplesUserQuery(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var id = 1
	var name = "Test user"
	var email = "test@mail.com"
	var createdAt = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","name","email","created_at","id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name", "email", "created_at"}).AddRow(id, name, email, createdAt))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)

	apitest.New().HandlerFunc(App()).
		Post("/api/graphql").
		GraphQLQuery(`
			query {
				user(id: 1) {
					id,
					name,
					email,
					createdAt
				}
			}
		`).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Root("$.data.user").
			Equal("id", float64(id)).
			Equal("name", name).
			Equal("email", email).
			Equal("createdAt", createdAt.Format(time.RFC3339Nano)).
			End(),
		).
		End()
}

func TestUserCollectionQuery(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var id = 1
	var name = "Test user"
	var email = "test@mail.com"
	var createdAt = time.Now()

	var bookId = 1
	var bookTitle = "Test book"

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name", "email", "created_at"}).AddRow(id, name, email, createdAt))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","title","id" FROM "books" WHERE user_id = ($1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "title"}).AddRow(bookId, bookTitle))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)

	apitest.New().HandlerFunc(App()).
		Post("/api/graphql").
		GraphQLQuery(`
			query {
				user(id: 1) {
					collection {
						id,
						title
					}
				}
			}
		`).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$.data.user.collection", 1)).
		Assert(jsonpath.Root("$.data.user.collection[0]").
			Equal("id", float64(bookId)).
			Equal("title", bookTitle).
			End(),
		).
		End()
}

func TestUserBorrowedQuery(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var id = 1

	var bookId = 1
	var bookLoanId = 1
	var bookTitle = "Test book"
	var bookPages = 100

	var fromUser = 1
	var toUser = 1
	var lentAt = time.Now().AddDate(0, -1, 0)
	var returnedAt = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(id))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "from_user","to_user","lent_at","returned_at","id","book_id" FROM "book_loans" WHERE to_user = ($1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"from_user", "to_user", "lent_at", "returned_at", "id", "book_id"}).
			AddRow(fromUser, toUser, lentAt, &returnedAt, bookLoanId, bookId))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","title","pages","id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "title", "pages"}).AddRow(bookId, bookTitle, bookPages))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)

	apitest.New().HandlerFunc(App()).
		Post("/api/graphql").
		GraphQLQuery(`
			query {
				user(id: 1) {
					borrowedBooks {
						fromUser,
						toUser,
						lentAt,
						returnedAt,
						book {
							id,
							title,
							pages
						}
					}
				}
			}
		`).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$.data.user.borrowedBooks", 1)).
		Assert(jsonpath.Root("$.data.user.borrowedBooks[0]").
			Equal("fromUser", float64(fromUser)).
			Equal("toUser", float64(toUser)).
			Equal("lentAt", lentAt.Format(time.RFC3339Nano)).
			Equal("returnedAt", returnedAt.Format(time.RFC3339Nano)).
			End(),
		).
		Assert(jsonpath.Root("$.data.user.borrowedBooks[0].book").
			Equal("id", float64(bookId)).
			Equal("title", bookTitle).
			Equal("pages", float64(bookPages)).
			End(),
		).
		End()
}

func TestUserLentQuery(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var id = 1

	var bookId = 1
	var bookLoanId = 1
	var bookTitle = "Test book"
	var bookPages = 100


	var fromUser = 1
	var toUser = 1
	var lentAt = time.Now().AddDate(0, -1, 0)
	var returnedAt = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(id))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "from_user","to_user","lent_at","returned_at","id","book_id" FROM "book_loans" WHERE from_user = ($1)`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows(Rows{"from_user", "to_user", "lent_at", "returned_at", "id", "book_id"}).
			AddRow(fromUser, toUser, lentAt, &returnedAt, bookLoanId, bookId))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id","title","pages","id" FROM "books" WHERE "books"."id" = $1 ORDER BY "books"."id" LIMIT 1`)).
		WithArgs(bookId).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "title", "pages"}).AddRow(bookId, bookTitle, bookPages))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)

	apitest.New().HandlerFunc(App()).
		Post("/api/graphql").
		GraphQLQuery(`
			query {
				user(id: 1) {
					lentBooks {
						fromUser,
						toUser,
						lentAt,
						returnedAt
						book {
							id,
							title,
							pages
						}
					}
				}
			}
		`).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$.data.user.lentBooks", 1)).
		Assert(jsonpath.Root("$.data.user.lentBooks[0]").
			Equal("fromUser", float64(fromUser)).
			Equal("toUser", float64(toUser)).
			Equal("lentAt", lentAt.Format(time.RFC3339Nano)).
			Equal("returnedAt", returnedAt.Format(time.RFC3339Nano)).
			End(),
		).
		Assert(jsonpath.Root("$.data.user.lentBooks[0].book").
			Equal("id", float64(bookId)).
			Equal("title", bookTitle).
			Equal("pages", float64(bookPages)).
			End(),
		).
		End()
}