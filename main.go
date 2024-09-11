package main

import (
	"log"
	"os"

	"log-collector/config"
	// _는 패키지를 가져오지만, 그 패키지 내의 어떤 기능도 직접 사용하지 않는다는 의미입니다.
	_ "log-collector/docs" // 자동 생성된 Swagger 문서를 import
	"log-collector/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env 파일을 로드하는데 실패했습니다: %v", err)
	}
	// 로그 설정
	logger := config.SetupLogger("logs", "logs")
	config.RotateLogger(logger, "logs") // 로그 파일 롤링 작업 (날짜 기반)

	// 로그 출력 설정
	log.SetOutput(logger) // 모든 로그를 logger로 설정

	// Fiber 인스턴스 생성
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// 라우터 설정
	router.SetupRoutes(app)

	// 서버 실행
	port := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + port))
}

//	// 자정마다 로그 파일을 변경하는 작업을 백그라운드에서 실행
//	go writer.ScheduleLogRotation()
