package models

import "time"

type Book struct {
	ID          int64       `json:"id" gorm:"primaryKey"`
	UserId      int64       `json:"-" valid:"required"`
	Title       string      `json:"title" valid:"required"`
	Pages       int         `json:"pages" valid:"required"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"default:current_timestamp"`
}

