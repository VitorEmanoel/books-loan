package command

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/database/plugins"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

func TestAddBook(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var loggedUser int64 = 1
	var now = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "books" ("user_id","title","pages") VALUES ($1,$2,$3)`)).
		WithArgs(1, "My Book", 10).
		WillReturnRows(sqlmock.NewRows(Rows{"created_at", "id"}).AddRow(now, 1))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	book, err := mediator.Send(&AddBookRequest{
		LoggedUserId: loggedUser,
		BookInput:    models.AddBookInput{
			Title: "My Book",
			Pages: 10,
		},
	})
	var expectedBook = models.Book{
		ID:        1,
		UserId:    loggedUser,
		Title:     "My Book",
		Pages:     10,
		CreatedAt: now,
	}
	assert.Nil(t, err)
	assert.NotNil(t, book)
	assert.Equal(t, &expectedBook, book)
}

func TestAddBookWithNotFoundUser(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var loggedUser int64 = 1
	var now = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "books" ("user_id","title","pages") VALUES ($1,$2,$3)`)).
		WithArgs(1, "My Book", 10).
		WillReturnRows(sqlmock.NewRows(Rows{"created_at", "id"}).AddRow(now, 1))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	book, err := mediator.Send(&AddBookRequest{
		LoggedUserId: loggedUser,
		BookInput:    models.AddBookInput{
			Title: "My Book",
			Pages: 10,
		},
	})
	assert.Nil(t, book)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, LoggedUserNotfound)
}

func TestAddBookWithEmptyTitle(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var loggedUser int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	book, err := mediator.Send(&AddBookRequest{
		LoggedUserId: loggedUser,
		BookInput:    models.AddBookInput{
			Title: "",
			Pages: 10,
		},
	})
	assert.Nil(t, book)
	assert.NotNil(t, err)
	errs := err.(govalidator.Errors)
	oneError := errs[0].(govalidator.Error)
	assert.Equal(t, "title", oneError.Name)
	assert.Equal(t, "required", oneError.Validator)
}

func TestAddBookWithPageZero(t *testing.T) {
	var med = mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var loggedUser int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE "users"."id" = $1 ORDER BY "users"."id" LIMIT 1`)).
		WithArgs(loggedUser).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(loggedUser))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	med.GetContainer().Inject("repository", repo)
	book, err := mediator.Send(&AddBookRequest{
		LoggedUserId: loggedUser,
		BookInput:    models.AddBookInput{
			Title: "My book",
			Pages: 0,
		},
	})
	assert.Nil(t, book)
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, PageGreaterZero)
}