package query

import (
	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/graphql"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type FindUserRequest struct {
	mediator.Request
	UserId              int64
}

type FindUserRequestHandler struct {
	Repository      repository.Repository       `inject:"repository"`
}

func (h *FindUserRequestHandler) Handle(r *FindUserRequest) (interface{}, error) {
	var user = models.User{}
	var allowedFields = user.AllowedFields()
	var fields = graphql.GetFields(r.Context(), allowedFields...)
	fields = append(fields, "id")
	finalUser, err := h.Repository.SetModel(&models.User{}).Find(r.UserId, repository.Select(fields...))
	if err != nil {
		return nil, UserNotFound
	}
	return finalUser, nil
}

func init() {
    mediator.RegisterRequest(&FindUserRequest{}, &FindUserRequestHandler{})
}
