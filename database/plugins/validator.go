package plugins

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Validator interface {
	Validate(db *gorm.DB) error
}

type context struct {

}

func Validate(db *gorm.DB) {
	var model = db.Statement.Dest
	if model == nil {
		return
	}
	_, err := govalidator.ValidateStruct(model)
	if  err != nil {
		_ = db.AddError(err)
		return
	}
	validator, ok := model.(Validator)
	if !ok {
		return
	}
	err = validator.Validate(db)
	if err != nil {
		_ = db.AddError(err)
		return
	}
}

func (v *context) Name() string {
	return "validator"
}

func (v *context) Initialize(db *gorm.DB) error {
	var createCallback = db.Callback().Create()
	var updateCallback = db.Callback().Update()
	err := createCallback.Before("gorm:create").Register("validator:on_create", Validate)
	if err != nil {
		return err
	}
	err = updateCallback.Before("gorm:update").Register("validator:on_update", Validate)
	if err != nil {
		return err
	}
	return nil
}

func NewValidator() gorm.Plugin {
	return &context{}
}
