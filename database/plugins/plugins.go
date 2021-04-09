package plugins

import "gorm.io/gorm"

func SetupPlugins(db *gorm.DB) error {
	if err := db.Use(NewValidator()); err != nil {
		return err
	}
	return nil
}
