package handler

import (
	"github.com/gofiber/fiber/v2"
)

// HomeHandler는 "/" 라우팅에서 요청을 처리하는 핸들러
func HomeHandler(c *fiber.Ctx) error {
	// 간단한 예시로 로그 생성
	return c.SendString("Hello, World!")
}
