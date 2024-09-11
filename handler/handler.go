package handler

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary 초기 메서드
// @Description Hello, World! 메시지를 반환합니다.
// @Tags 기본
// @Produce  plain
// @Success 200 {string} string "Hello, World!"
// @Router / [get]
func HomeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
