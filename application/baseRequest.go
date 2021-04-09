package application

import (
	mediator "github.com/VitorEmanoel/gMediator"

	"github.com/VitorEmanoel/books-loan/repository"
)

type BaseRequest struct {
	mediator.Request
	Repository          repository.Repository
}

func NewRequest(repository repository.Repository) BaseRequest {
	return BaseRequest{
		Repository: repository,
	}
}
