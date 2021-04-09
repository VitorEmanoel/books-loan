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

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/database/plugins"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type Rows []string

func TestCreateUser(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	var id int64 = 1
	var now = time.Now()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("test@mail.com").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("name","email") VALUES ($1,$2)`)).
		WithArgs("Test User", "test@mail.com").
		WillReturnRows(sqlmock.NewRows(Rows{"created_at", "id"}).AddRow(now, id))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "Test User",
		Email: "test@mail.com",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	var expectedUser = models.User{
		ID:             id,
		Name:          "Test User",
		Email:         "test@mail.com",
		CreatedAt:     now,
	}
	assert.Nil(t, err)
	assert.Equal(t, &expectedUser, user)
}

func TestCreateUserWithEmailAlreadyExists(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("test@mail.com").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(1))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "Test User",
		Email: "test@mail.com",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.ErrorIs(t, ErrEmailAlreadyInUse, err)
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("test_invalidmail.com").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "Test User",
		Email: "test_invalidmail.com",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Len(t, err, 1)
	errs := err.(govalidator.Errors)
	oneError := errs[0].(govalidator.Error)
	assert.Equal(t, "email", oneError.Name)
	assert.Equal(t, "email", oneError.Validator)
}

func TestCreateUserWithEmptyName(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("test@mail.com").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "",
		Email: "test@mail.com",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Len(t, err, 1)
	errs := err.(govalidator.Errors)
	oneError := errs[0].(govalidator.Error)
	assert.Equal(t, "required", oneError.Validator)
	assert.Equal(t, "name", oneError.Name)
}

func TestCreateUserWithEmptyEmail(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "Test user",
		Email: "",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Len(t, err, 1)
	errs := err.(govalidator.Errors)
	oneError := errs[0].(govalidator.Error)
	assert.Equal(t, "required", oneError.Validator)
	assert.Equal(t, "email", oneError.Name)
}

func TestCreateUserWithEmptyEmailAndName(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "id" FROM "users" WHERE email = ($1)`)).
		WithArgs("").
		WillReturnRows(sqlmock.NewRows(Rows{"id"}))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	db = db.Debug()
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	var userInput = models.CreateUserInput{
		Name:  "",
		Email: "",
	}
	user, err := mediator.Send(&CreateUserRequest{
		BaseRequest: application.NewRequest(repo),
		UserInput:   userInput,
	})
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Len(t, err, 2)
	errs := err.(govalidator.Errors)
	oneError := errs[0].(govalidator.Error)
	twoError := errs[1].(govalidator.Error)
	assert.Equal(t, "required", oneError.Validator)
	assert.Equal(t, "name", oneError.Name)
	assert.Equal(t, "required", twoError.Validator)
	assert.Equal(t, "email", twoError.Name)
}