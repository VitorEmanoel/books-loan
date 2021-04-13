package repository

import (
	"reflect"

	"gorm.io/gorm"
)

// Repository interface for search, create, update and delete data of models in database
type Repository interface {
	FindAll(options ...Options) (interface{}, error)
	Find(id int64, options ...Options) (interface{}, error)
	Create(entity interface{}, options ...Options) error
	Delete(id int64, options ...Options) error
	Update(entity interface{}, id int64, options ...Options) error
	Total(options ...Options) (int64, error)
	SetModel(model interface{}) Repository
	GetDatabase() *gorm.DB
}

// RepositoryContext context for Repository
type RepositoryContext struct {
	DB        *gorm.DB
	ModelType reflect.Type
	Model     interface{}
}

func (r *RepositoryContext) GetDatabase() *gorm.DB {
	return r.DB
}

// FindAll findall items in database
func (r *RepositoryContext) FindAll(options ...Options) (interface{}, error) {
	var db = r.applyOptions(options)
	var items = reflect.New(reflect.SliceOf(reflect.TypeOf(r.Model))).Interface()
	err := db.Find(items).Error
	if err != nil {
		return nil, err
	}
	return reflect.ValueOf(items).Elem().Interface(), nil
}

// Total get total items of search
func (r *RepositoryContext) Total(options... Options) (int64, error) {
	var db = r.applyOptions(options)
	var total int64
	err := db.Model(r.Model).Count(&total).Error
	if err != nil {
		return 0, err
	}
	return total, nil

}

// Find find specify registry
func (r *RepositoryContext) Find(id int64, options ...Options) (interface{}, error) {
	var db = r.applyOptions(options)
	var item = reflect.New(r.ModelType).Interface()
	err := db.First(item, id).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Create create item in database
func (r *RepositoryContext) Create(entity interface{}, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Create(entity).Error
	if err != nil {
		return err
	}
	return nil
}

// Delete delete item in database
func (r *RepositoryContext) Delete(id int64, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Delete(r.Model, id).Error
	if err != nil {
		return err
	}
	return nil
}

// Update item in database
func (r *RepositoryContext) Update(entity interface{}, id int64, options ...Options) error {
	var db = r.applyOptions(options)
	err := db.Where("id = ?", id).Updates(entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryContext) applyOptions(options []Options) *gorm.DB {
	var db = r.DB.Model(r.Model)
	for _, option := range options {
		option.Apply(db)
	}
	return db
}

// SetModel set model of repository
func (r *RepositoryContext) SetModel(model interface{}) Repository {
	var modelType = reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}
	r.Model = model
	r.ModelType = modelType
	return r
}

// NewRepository create repository
func NewRepositoryModel(model interface{}, db *gorm.DB) Repository {
	var context = &RepositoryContext{DB: db}
	context.SetModel(model)
	return context
}

func NewRepository(db *gorm.DB) Repository {
	return &RepositoryContext{DB: db}
}
