package command

import "errors"

var LoggedUserNotfound = errors.New("logged user not found")

var BookNotFound = errors.New("book not found")

var ToUserNotFound = errors.New("to user not found")

var BookAlreadyBorrowed = errors.New("this book already borrowed")

var BookNotBorrowed = errors.New("you not already borrowed this book")

var BookAlreadyReturn = errors.New("this book already returned to the owner")