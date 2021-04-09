package models

import "time"

type User struct {
	ID            int64      `json:"id" gorm:"primary_key;"`
	Name          string      `json:"name" valid:"required"`
	Email         string      `json:"email" gorm:"index:idx_email,unique" valid:"required,email"`
	CreatedAt     time.Time   `json:"createdAt" gorm:"default:current_timestamp"`
}

