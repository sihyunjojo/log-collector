package main

import (
	"log"
	"log-collector/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"log-collector/router" // 라우터 임포트
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

	// 라우터 설정
	router.SetupRoutes(app)

	// 서버 실행
	port := os.Getenv("SERVER_PORT")
	log.Fatal(app.Listen(":" + port))
}

//package main
//
//import (
//	"fmt"
//	"log"
//	"log-collector/config"
//	"net/http"
//
//	"github.com/gofiber/fiber/v2"
//
//	customerrors "log-collector/errors"
//	"log-collector/writer"
//	"os"
//
//	// godotenv는 .env 파일에 정의된 키-값 쌍을 환경 변수로 설정하는 라이브러리입니다.
//	"github.com/joho/godotenv" // godotenv 패키지
//)
//
//var port string
//
//func main() {
//	// .env 파일 로드
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//
//	// Fiber 인스턴스 생성
//	app := fiber.New()
//
//	userLogger := config.SetupLogger()
//
//	// 포트 번호 환경 변수로 관리
//	port, err = getPort()
//	if err != nil {
//		log.Fatal(err) // 에러가 발생하면 프로그램을 중단하고 로그 출력
//	}
//
//	// 자정마다 로그 파일을 변경하는 작업을 백그라운드에서 실행
//	go writer.ScheduleLogRotation()
//
//	// 경로 설정: /log/watch 경로에 대한 핸들러 등록
//	http.HandleFunc("/log/watch", handlers.HandleWatchLog) // 패키지 이름을 명시적으로 사용
//
//	// 서버 시작
//	fmt.Printf("Starting server on :%s...\n", port)
//	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
//}
//
//// getPort 함수: 환경 변수가 없으면 커스텀 에러 발생
//func getPort() (string, error) {
//	port := os.Getenv("LOG_SERVER_PORT")
//	if port == "" {
//		return "", customerrors.NewPortNotFoundError("LOG_SERVER_PORT")
//	}
//	return port, nil
//}
