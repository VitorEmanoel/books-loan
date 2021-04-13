package command

import (
	"time"

	. "github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/models"
)

type ReturnBookRequest struct {
	mediator.Request
	LoggedUserId        int64
	BookId              int64
	Date                *time.Time
}

type ReturnBookRequestHandler struct {
	Repository Repository `inject:"repository"`
}

func (h *ReturnBookRequestHandler) Handle(r *ReturnBookRequest) (*models.BookLoan, error) {
	_, err := h.Repository.SetModel(&models.User{}).Find(r.LoggedUserId, Select("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, LoggedUserNotfound
		}
		return nil, err
	}
	_, err = h.Repository.SetModel(&models.Book{}).Find(r.BookId, Select("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, BookNotFound
		}
		return nil, err
	}
	borrowedBooks, err := h.Repository.SetModel(&models.BookLoan{}).
		FindAll(
			Where("book_id = ?", r.BookId),
			Where("to_user = ?", r.LoggedUserId),
			Order("returned_at desc"),
			Limit(1),
		)
	if err != nil {
		return nil, err
	}
	books := borrowedBooks.([]*models.BookLoan)
	if len(books) == 0 {
		return nil, BookNotBorrowed
	}
	var borrowedBook = books[0]
	if borrowedBook.ReturnedAt != nil {
		return nil, BookAlreadyReturn
	}
	if r.Date == nil {
		var now = time.Now()
		r.Date = &now
	}
	borrowedBook.ReturnedAt = r.Date
	err = h.Repository.Update(borrowedBook, borrowedBook.ID, Select("returned_at"))
	if err != nil {
		return nil, err
	}
	return borrowedBook, nil
}

func init() {
    mediator.RegisterRequest(&ReturnBookRequest{}, &ReturnBookRequestHandler{})
}
