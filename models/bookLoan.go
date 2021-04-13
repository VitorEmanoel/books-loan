package models

import (
	"time"
)

//BookLoan model
type BookLoan struct {
	ID              int64               `json:"id" gorm:"primaryKey;"`
	BookId          int64               `json:"bookId"`
	FromUser        int64               `json:"fromUser"`
	ToUser          int64               `json:"toUser"`
	LentAt          time.Time           `json:"lentAt" gorm:"default:current_timestamp"`
	ReturnedAt      *time.Time          `json:"returnedAt"`
}

// AllowedFields for BookLoan
func (b *BookLoan) AllowedFields() []string {
	return Fields{"id", "bookId", "fromUser", "toUser", "lentAt", "returnedAt"}
}