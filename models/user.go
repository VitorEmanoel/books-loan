package models

import "time"

// User model
type User struct {
	ID            int64         `json:"id" gorm:"primary_key;"`
	Name          string        `json:"name" valid:"required"`
	Email         string        `json:"email" gorm:"index:idx_email,unique" valid:"required,email"`
	CreatedAt     time.Time     `json:"createdAt" gorm:"default:current_timestamp"`
}

// AllowedFields for user model
func (u *User) AllowedFields() []string{
	return Fields{"id", "name", "email", "createdAt"}
}

