package command

import (
	"errors"

	mediator "github.com/VitorEmanoel/gMediator"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type AddBookRequest struct {
	application.BaseRequest
	LoggedUserId        int64
	BookInput           models.AddBookInput
}

var PageGreaterZero = errors.New("page value greater than zero")

type AddBookRequestHandler struct {
}

var UserNotFound = errors.New("user not exists")

func (h *AddBookRequestHandler) Handle(r *AddBookRequest) (*models.Book, error) {
	if r.BookInput.Pages <= 0{
		return nil, PageGreaterZero
	}
	_, err := r.Repository.SetModel(&models.User{}).Find(r.LoggedUserId, repository.Select("id"))
	if err == gorm.ErrRecordNotFound {
		return nil, UserNotFound
	}
	var book = models.Book{
		UserId:    r.LoggedUserId,
		Title:     r.BookInput.Title,
		Pages:     r.BookInput.Pages,
	}
	err = r.Repository.SetModel(&models.Book{}).Create(&book)
	if err != nil {
		return nil, err
	}
    return &book, nil
}

func init() {
    mediator.RegisterRequest(&AddBookRequest{}, &AddBookRequestHandler{})
}
