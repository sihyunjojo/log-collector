package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"log-collector/config" // 로그 파일에 쓰기 위한 패키지
)

type WatchLogRequest struct {
	Keyword string `json:"keyword"`
}

// HandleWatchLog는 요청의 Body에서 keyword를 추출하고 로그로 저장하는 핸들러
// @Summary Log a keyword from request body
// @Description Extracts a keyword from the request body and logs it to a file
// @Tags logs
// @Accept  json
// @Produce  json
// @Param   keyword  body WatchLogRequest true "Keyword to log"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /log/keyword [post]
func HandleWatchLog(c *fiber.Ctx) error {
	// 요청 Body에서 JSON 데이터 추출
	var req WatchLogRequest
	if err := c.BodyParser(&req); err != nil {
		// 오류가 발생하면 400 Bad Request 응답
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 로그 파일에 keyword 저장
	userLogger := config.SetupLogger("keyword", "keyword") // 로그 파일을 생성하고 설정
	log.SetOutput(userLogger)                              // 기본 log 패키지에서 로그 출력 대상을 변경
	log.Printf("Received keyword: %s", req.Keyword)        //  로그 메시지를 생성하고 기록

	// 성공적으로 처리되었음을 응답
	return c.JSON(fiber.Map{
		"keyword": req.Keyword,
	})
}
