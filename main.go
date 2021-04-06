//go:generate gqlgen

package main

import "github.com/sirupsen/logrus"

func main() {
	var app = NewApp()
	if err := app.Listen(":8080"); err != nil {
		logrus.Error("Error in starting server. Error: ", err.Error())
	}
}
