package database

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/VitorEmanoel/books-loan/database/plugins"
	"github.com/VitorEmanoel/books-loan/models"
)

func UseDatabase(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("database", db)
		return ctx.Next()
	}
}

func GetDatabase(ctx *fiber.Ctx) *gorm.DB {
	value := ctx.Locals("database")
	db, ok := value.(*gorm.DB)
	if !ok {
		return nil
	}
	return db
}

func Open(connection string) *gorm.DB {
	// Open connection
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: connection}), &gorm.Config{})
	if err != nil {
		logrus.Panic("Error in connect with database. Error: ", err.Error())
	}
	// Setup plugins
	err = plugins.SetupPlugins(db)
	if err != nil {
		logrus.Error("Failed in setup gorm plugins. Error: ", err.Error())
	}
	return db
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(models.Models...)
	if err != nil {
		return err
	}
	return nil
}