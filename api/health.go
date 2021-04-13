package api

import (
	"github.com/VitorEmanoel/books-loan/common"
	"github.com/gofiber/fiber/v2"
)

func Health(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(common.JSON{
		"status": "OK",
	})
}
