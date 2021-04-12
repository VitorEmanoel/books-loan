package query

import (
	"github.com/VitorEmanoel/books-loan/graphql"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
)

type FindBooksRequest struct {
    mediator.Request
	UserId      int64
}

type FindBooksRequestHandler struct {
	Repository          repository.Repository       `inject:"repository"`
}

func (h *FindBooksRequestHandler) Handle(r *FindBooksRequest) (interface{}, error) {
	var book models.Book
	var allowedFields = book.AllowedFields()
	var fields = graphql.GetFields(r.Context(), allowedFields...)
	fields = append(fields, "id")
	return h.Repository.SetModel(&models.Book{}).FindAll(repository.Where("user_id = ?", r.UserId), repository.Select(fields...))
}

func init() {
    mediator.RegisterRequest(&FindBooksRequest{}, &FindBooksRequestHandler{})
}
