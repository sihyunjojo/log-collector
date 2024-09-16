package main

import (
	"fmt"
	"log"
	"os"

	_ "log-collector/docs"
	"log-collector/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Log Collector API
// @version 1.0
// @description <h2>회원별 키워드 로깅을 위한 간단한 로그 수집 API입니다.</h2>

// @contact.name 조시현
// @contact.email si4018@naver.com

// @schemes http https

func main() {
	err := runApplication()
	if err != nil {
		log.Fatalf("Application failed with error: %v", err)
	}
}

func runApplication() error {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf(".env 파일을 로드하는데 실패했습니다: %w", err)
	}

	// Fiber 인스턴스 생성
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// 라우터 설정
	router.SetupRoutes(app)

	// 서버 실행
	port := os.Getenv("SERVER_PORT")
	return app.Listen(":" + port)
}
