package command

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	mediator "github.com/VitorEmanoel/gMediator"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/database/plugins"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

func TestLendBook(t *testing.T) {
	mediator.NewMediator()
	mockDB, mock, err := sqlmock.New()
	assert.Nil(t, err)

	var id int64 = 1
	var now = time.Now()
	var fromUser int64 = 1
	var toUser int64 = 2
	var bookId int64 = 1

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "book_loans" ("book_id","from_user","to_user","returned_at") VALUES ($1,$2,$3,$4) RETURNING "lent_at","id"`)).
		WithArgs(bookId, fromUser, toUser, nil).
		WillReturnRows(sqlmock.NewRows(Rows{"created_at", "id"}).AddRow(now, id))

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})
	assert.Nil(t, err)
	err = plugins.SetupPlugins(db)
	assert.Nil(t, err)
	var repo = repository.NewRepository(db)
	bookLoan, err := mediator.Send(&LendBookRequest{
		BaseRequest:   application.NewRequest(repo),
		LoggedUserId:  fromUser,
		LendBookInput: models.LendBookInput{
			BookID:   bookId,
			ToUserID: toUser,
		},
	})
	var expectedBookLoan = models.BookLoan{
		ID:         id,
		BookId:     bookId,
		FromUser:   fromUser,
		ToUser:     toUser,
		LentAt:     now,
	}
	assert.Nil(t, err)
	assert.NotNil(t, bookLoan)
	assert.Equal(t, &expectedBookLoan, bookLoan)
}
