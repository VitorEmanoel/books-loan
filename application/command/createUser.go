package command

import (
	"errors"

	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type CreateUserRequest struct {
	application.BaseRequest
	UserInput       models.CreateUserInput      `json:"input"`
}

type CreateUserRequestHandler struct {

}

var ErrEmailAlreadyInUse = errors.New("email already in use")

func (h *CreateUserRequestHandler) Handle(r *CreateUserRequest) (*models.User, error) {
	var user = models.User{
		Name: r.UserInput.Name,
		Email: r.UserInput.Email,
	}
	values, err := r.Repository.SetModel(&models.User{}).
		FindAll(
			repository.Select("id"),
			repository.Where("email = ?", r.UserInput.Email),
		)
	if err != nil {
		return nil, err
	}
	users, ok := values.([]*models.User)
	if ok && len(users) > 0 {
		return nil, ErrEmailAlreadyInUse
	}
	err = r.Repository.SetModel(&models.User{}).Create(&user)
	if err != nil {
		return nil, err
	}
    return &user, nil
}

func init() {
    mediator.RegisterRequest(&CreateUserRequest{}, &CreateUserRequestHandler{})
}
