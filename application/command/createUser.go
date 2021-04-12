package command

import (
	"errors"

	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type CreateUserRequest struct {
	mediator.Request
	UserInput       models.CreateUserInput      `json:"input"`
}

type CreateUserRequestHandler struct {
	Repository      repository.Repository   `inject:"repository"`
}

var ErrEmailAlreadyInUse = errors.New("email already in use")

func (h *CreateUserRequestHandler) Handle(r *CreateUserRequest) (*models.User, error) {
	var user = models.User{
		Name: r.UserInput.Name,
		Email: r.UserInput.Email,
	}
	values, err := h.Repository.SetModel(&models.User{}).
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
	err = h.Repository.SetModel(&models.User{}).Create(&user)
	if err != nil {
		return nil, err
	}
    return &user, nil
}

func init() {
    mediator.RegisterRequest(&CreateUserRequest{}, &CreateUserRequestHandler{})
}
