package models

import "time"

// Book model
type Book struct {
	ID          int64       `json:"id" gorm:"primaryKey"`
	UserId      int64       `json:"-" valid:"required"`
	Title       string      `json:"title" valid:"required"`
	Pages       int         `json:"pages" valid:"required"`
	CreatedAt   time.Time   `json:"createdAt" gorm:"default:current_timestamp"`
}

// AllowedFields for book
func (b *Book) AllowedFields() []string {
	return Fields{"id", "title", "pages", "createdAt"}
}

