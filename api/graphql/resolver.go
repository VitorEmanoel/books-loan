package graphql

import (
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB              *gorm.DB
	Repository      repository.Repository
}
