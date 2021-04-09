package command

import (
	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/models"
)

type LendBookRequest struct {
	application.BaseRequest
	LoggedUserId        int64
	LendBookInput       models.LendBookInput
}

type LendBookRequestHandler struct {

}

func (h *LendBookRequestHandler) Handle(r *LendBookRequest) (*models.BookLoan, error) {
	var bookLoan = models.BookLoan{
		BookId:     r.LendBookInput.BookID,
		FromUser:   r.LoggedUserId,
		ToUser:     r.LendBookInput.ToUserID,
	}
	err := r.Repository.SetModel(&models.BookLoan{}).Create(&bookLoan)
	if err != nil {
		return nil, err
	}
    return &bookLoan, nil
}

func init() {
    mediator.RegisterRequest(&LendBookRequest{}, &LendBookRequestHandler{})
}
