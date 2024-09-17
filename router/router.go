package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log-collector/config"
	"log-collector/handler"
)

// SetupRoutes는 라우팅과 미들웨어 설정을 담당
func SetupRoutes(app *fiber.App) {
	// 로거 미들웨어 설정
	// 이 라인은 Fiber의 로거 미들웨어(logger.New)를 설정하고, 모든 요청에 대한 로그를 기록하는 역할을 합니다.
	// writer.GetLogger() 함수는 로그 파일을 생성하고, 해당 로그 파일에 로그가 기록되도록 설정합니다.

	// Fiber의 logger 미들웨어는 클라이언트 요청이 들어올 때마다 요청에 대한 정보를 자동으로 기록합니다.
	// 이 정보는 요청 시간, HTTP 메서드, 경로, 응답 상태 코드, 응답 시간이 포함됩니다.

	// 기본적으로 이 미들웨어는 로그를 stdout(콘솔)에 출력하지만,
	// 여기서는 writer.GetLogger()를 사용하여 로그 파일에 기록하도록 출력 대상을 변경한 것입니다.

	// 모든 요청에 GetLogger()를 사용함.
	// logger.New()는 Fiber의 로깅 미들웨어를 설정하는 메서드입니다. 이 메서드를 통해 서버에 들어오는 모든 HTTP 요청을 자동으로 기록
	app.Use(logger.New(config.GetLogger()))

	// 각 요청 마다 해당하는 handler 로 이동
	app.Get("/", handler.HomeHandler)
	app.Post("/log/keyword", handler.HandleKeywordLog)
	app.Post("/log/keyword/:memberId", handler.HandleKeywordLogByMember)
}
