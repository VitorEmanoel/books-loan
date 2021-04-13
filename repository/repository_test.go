package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id      int64       `json:"id"`
	Name    string      `json:"name"`
}

type Rows []string

func TestRepositoryFind(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1").AddRow(2, "user 2"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	user, err := repository.Find(1)
	assert.Nil(t, err)
	assert.Equal(t, user, &User{Id: 1, Name: "user 1"})
}

func TestRepositoryFindAll(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1").AddRow(2, "user 2"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	users, err := repository.FindAll()
	assert.Len(t, users, 2)
	assert.Contains(t, users, &User{Id: 1, Name: "user 1"})
	assert.Contains(t, users, &User{Id: 2, Name: "user 2"})
}

func TestRepositoryFindWithSelect(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "name" FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(Rows{"name"}).AddRow("user 1"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	user, err := repository.Find(1, Select("name"))
	assert.Nil(t, err)
	assert.Equal(t, &User{Name: "user 1"}, user)
}

func TestRepositoryFindAllWithLimit(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" LIMIT 1`)).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	users, err := repository.FindAll(Limit(1))
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}

func TestRepositoryFindAllWithWhere(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE name = ($1)`)).
		WithArgs("user 1").
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	users, err := repository.FindAll(Where("name = ?", "user 1"))
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}

func TestRepositoryFindAllWithOffset(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" OFFSET 1`)).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	users, err := repository.FindAll(Offset(1))
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}

func TestRepositoryFindAllWithOrder(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" ORDER BY name`)).
		WillReturnRows(sqlmock.NewRows(Rows{"id", "name"}).AddRow(1, "user 1"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	users, err := repository.FindAll(Order("name"))
	assert.Nil(t, err)
	assert.Len(t, users, 1)
}

func TestRepositoryTotal(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(1) FROM "users"`)).
		WillReturnRows(sqlmock.NewRows(Rows{"count"}).AddRow(2))

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	total, err := repository.Total()
	assert.Nil(t, err)
	assert.Equal(t, int64(2), total)
}

func TestRepositoryCreate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users" ("name","id") VALUES ($1,$2) RETURNING "id"`)).
		WithArgs("user 1", 1).
		WillReturnRows(sqlmock.NewRows(Rows{"id"}).AddRow(1))
	mock.ExpectCommit()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	var newUser = User{
		Id:   1,
		Name: "user 1",
	}
	err = repository.Create(&newUser)
	assert.Nil(t, err)
}


func TestRepositoryUpdate(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "name"=$1 WHERE id = $2`)).
		WithArgs("user 1 1", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	var updateUser = User{
		Name: "user 1 1",
	}
	err = repository.Update(&updateUser, 1)
	assert.Nil(t, err)
}


func TestRepositoryDelete(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	assert.NotNil(t, mock)
	assert.NotNil(t, mockDB)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:  mockDB,
	}), &gorm.Config{})
	assert.Nil(t, err)
	assert.NotNil(t, db)
	var repository = NewRepositoryModel(&User{}, db)
	err = repository.Delete(1)
	assert.Nil(t, err)
}