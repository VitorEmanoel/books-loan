package repository

import (
	"gorm.io/gorm"
)

// Options base options for apply mutable data in database connection
type Options interface {
	Apply(db *gorm.DB)
}

// DatabaseOptions database options builder
type DatabaseOptions struct {
	ApplyFunc func(db *gorm.DB)
}

// Apply apply database customization
func (d *DatabaseOptions) Apply(db *gorm.DB) {
	d.ApplyFunc(db)
}

// Preloads add preload in database connection
func Preloads(preloads ...string) Options {
	return &DatabaseOptions{
		ApplyFunc: func(db *gorm.DB) {
			for _, preload := range preloads {
				db.Preload(preload)
			}
		},
	}
}

// Limit apply limit for database query
func Limit(limit int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) {
		db.Limit(limit)
	}}
}

// Where apply where for database query
func Where(condition string, args ...interface{}) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) {
		db.Where(condition, args)
	}}
}

// Offset apply offset in database query
func Offset(offset int) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) {
		db.Offset(offset)
	}}
}

// Order apply ordering in database query
func Order(order string) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) {
		db.Order(order)
	}}
}

type FilterFields map[string][]interface{}

// Select apply filters in
func Select(fields... string) Options {
	return &DatabaseOptions{ApplyFunc: func(db *gorm.DB) {
		db.Select(fields)
	}}
}