package command

import (
	mediator "github.com/VitorEmanoel/gMediator"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/models"
	. "github.com/VitorEmanoel/books-loan/repository"
)

type LendBookRequest struct {
	mediator.Request
	LoggedUserId        int64
	LendBookInput       models.LendBookInput
}

type LendBookRequestHandler struct {
	Repository      Repository   `inject:"repository"`
}


func (h *LendBookRequestHandler) Handle(r *LendBookRequest) (*models.BookLoan, error) {
	var bookLoan = models.BookLoan{
		BookId:     r.LendBookInput.BookID,
		FromUser:   r.LoggedUserId,
		ToUser:     r.LendBookInput.ToUserID,
	}

	_, err := h.Repository.SetModel(&models.User{}).Find(r.LoggedUserId, Select("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, LoggedUserNotfound
		}
		return nil, err
	}
	_, err = h.Repository.SetModel(&models.Book{}).Find(r.LendBookInput.BookID, Select("id"), Where("user_id = ?", r.LoggedUserId))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, BookNotFound
		}
		return nil, err
	}
	_, err = h.Repository.SetModel(&models.User{}).Find(r.LendBookInput.ToUserID, Select("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ToUserNotFound
		}
		return nil, err
	}
	booksValue, err := h.Repository.SetModel(&models.BookLoan{}).FindAll(Select("id"), Where("book_id = ? AND returned_at IS NULL", r.LendBookInput.BookID))
	if err != nil {
		return nil, err
	}
	books := booksValue.([]*models.BookLoan)
	if len(books) > 0 {
		return nil, BookAlreadyBorrowed
	}
	err = h.Repository.SetModel(&models.BookLoan{}).Create(&bookLoan)
	if err != nil {
		return nil, err
	}
    return &bookLoan, nil
}

func init() {
    mediator.RegisterRequest(&LendBookRequest{}, &LendBookRequestHandler{})
}
