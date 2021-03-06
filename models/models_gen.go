// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AddBookInput struct {
	Title string `json:"title"`
	Pages int    `json:"pages"`
}

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LendBookInput struct {
	BookID   int64 `json:"bookId"`
	ToUserID int64 `json:"toUserId"`
}
