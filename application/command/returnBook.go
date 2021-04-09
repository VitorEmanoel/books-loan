package command

import (
	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/models"
)

type ReturnBookRequest struct {
	application.BaseRequest
	LoggedUserId        int64
	BookId              int64
}

type ReturnBookRequestHandler struct {

}

func (h *ReturnBookRequestHandler) Handle(request ReturnBookRequest) (*models.BookLoan, error) {
    return nil, nil
}

func init() {
    mediator.RegisterRequest(&ReturnBookRequest{}, &ReturnBookRequestHandler{})
}
