package query

import (
	"github.com/VitorEmanoel/books-loan/graphql"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
)

type FindBookRequest struct {
    mediator.Request
    BookId      int64
}

type FindBookRequestHandler struct {
	Repository      repository.Repository       `inject:"repository"`
}

func (h *FindBookRequestHandler) Handle(r *FindBookRequest) (interface{}, error) {
	var book models.Book
	var allowedFields = book.AllowedFields()
	var fields = graphql.GetFields(r.Context(), allowedFields...)
	fields = append(fields, "id")
	return h.Repository.SetModel(&models.Book{}).Find(r.BookId, repository.Select(fields...))
}

func init() {
    mediator.RegisterRequest(&FindBookRequest{}, &FindBookRequestHandler{})
}
