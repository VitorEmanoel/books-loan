package main

import (
	"fmt"

	"github.com/VitorEmanoel/menv"
)

type Environment struct {
	Port                string      `env:"PORT" default:"8080" file:".env"`

	DatabaseHost        string      `env:"DB_HOST" required:"true"`
	DatabasePort        string      `env:"DB_PORT" required:"true"`
	DatabaseUsername    string      `env:"DB_USERNAME" required:"true"`
	DatabasePassword    string      `env:"DB_PASSWORD" required:"true"`
	DatabaseDatabase    string      `env:"DB_DATABASE" required:"true"`
}
// GetDatabaseInfo transform environment variable struct to database connection string
func (e *Environment) GetDatabaseInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", e.DatabaseHost, e.DatabasePort, e.DatabaseUsername, e.DatabasePassword, e.DatabaseDatabase)
}

func LoadEnvironment() *Environment{
	var environment = Environment{}
	err := menv.LoadEnvironment(&environment)
	if err != nil {
		panic(err.Error())
	}
	return &environment
}
