package query

import (
	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/application"
	"github.com/VitorEmanoel/books-loan/graphql"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
)

type FindUserRequest struct {
	application.BaseRequest
	UserId              int64
}

type FindUserRequestHandler struct {

}

func (h *FindUserRequestHandler) Handle(r *FindUserRequest) (interface{}, error) {
	var fields = graphql.GetFields(r.Context())
	return r.Repository.SetModel(&models.User{}).Find(r.UserId, repository.Select(fields...))
}

func init() {
    mediator.RegisterRequest(&FindUserRequest{}, &FindUserRequestHandler{})
}
