package query

import (
	"github.com/VitorEmanoel/books-loan/graphql"
	"github.com/VitorEmanoel/books-loan/models"
	"github.com/VitorEmanoel/books-loan/repository"
	mediator "github.com/VitorEmanoel/gMediator"
)

type FindBookLoansRequest struct {
    mediator.Request
	FromUser        int64
    ToUser          int64
}

type FindBookLoansRequestHandler struct {
	Repository          repository.Repository       `inject:"repository"`
}

func (h *FindBookLoansRequestHandler) Handle(r *FindBookLoansRequest) (interface{}, error) {
	var options []repository.Options
	if r.ToUser > 0 {
		options = append(options, repository.Where("to_user = ?", r.ToUser))
	}
	if r.FromUser > 0 {
		options = append(options, repository.Where("from_user = ?", r.FromUser))
	}
	var bookLoan models.BookLoan
	var allowedFields = bookLoan.AllowedFields()
	var fields = graphql.GetFields(r.Context(), allowedFields...)
	fields = append(fields, "id", "book_id")
	options = append(options, repository.Select(fields...))
	return h.Repository.SetModel(&models.BookLoan{}).FindAll(options...)
}

func init() {
    mediator.RegisterRequest(&FindBookLoansRequest{}, &FindBookLoansRequestHandler{})
}
