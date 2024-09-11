package handler

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary Home page
// @Description Returns a simple Hello, World! message
// @Tags home
// @Produce  plain
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
