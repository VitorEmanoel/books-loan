package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
var returned = time.Now().Add(24 * time.Hour)
var bookOne = BookLoan{
	ID:         1,
	BookId:     1,
	FromUser:   1,
	ToUser:     2,
	LentAt:     time.Now(),
	ReturnedAt: &returned,
}
var bookTwo = BookLoan{
	ID:         2,
	BookId:     2,
	FromUser:   2,
	ToUser:     1,
	LentAt:     time.Now(),
}

var booksLoan = BooksLoan([]*BookLoan{&bookOne, &bookTwo})

func TestGetBookLoan(t *testing.T) {
	findBook := booksLoan.GetBookLoan(2)
	assert.Equal(t, &bookTwo, findBook)
}

func TestBorrowedBooks(t *testing.T) {
	borrowed := booksLoan.BorrowedBooks()
	assert.Len(t, borrowed, 1)
	assert.Contains(t, borrowed, &bookTwo)
}

func TestReturnedBooks(t *testing.T) {
	returned := booksLoan.ReturnedBooks()
	assert.Len(t, returned, 1)
	assert.Contains(t, returned, &bookOne)
}