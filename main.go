//go:generate gqlgen

package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {

	var environment = LoadEnvironment()
	var app = NewApp(environment)
	if err := app.Listen(fmt.Sprintf(":%s", environment.Port)); err != nil {
		logrus.Error("Error in starting server. Error: ", err.Error())
	}
}
