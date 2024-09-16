package handler

import (
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"log-collector/config" // 로그 파일에 쓰기 위한 패키지
)

type KeywordLogRequest struct {
	Keyword string `json:"keyword"`
}

var logger io.Writer

// @Param <name> <location> <type> <required> <description>

// @Summary 키워드 로그 수집
// @Description 키워드를 받아와서 파일에 기록합니다.
// @Tags 로그
// @Accept  json
// @Produce  json
// @Param  keyword body string true "로그로 사용할 키워드"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /log/keyword [post]
func HandleKeywordLog(c *fiber.Ctx) error {
	// 요청 Body에서 JSON 데이터 추출
	var req KeywordLogRequest
	if err := c.BodyParser(&req); err != nil {
		// 오류가 발생하면 400 Bad Request 응답
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 로그 파일에 keyword 저장
	userLogger := config.SetupLogger("keyword", "keyword")                        // 로그 파일을 생성하고 설정
	config.RotateLogger(userLogger, "keyword", "keyword")                         // 자정마다 로그 파일을 롤링
	log.SetOutput(userLogger)                                                     // 기본 log 패키지에서 로그 출력 대상을 변경
	log.Printf("[%s] Received keyword: %s\n", config.GetSeoulTime(), req.Keyword) //  로그 메시지를 생성하고 기록

	//fmt.Fprintf(logger, "[%s] Received keyword: %s\n", config.GetSeoulTime(), req.Keyword)

	// 성공적으로 처리되었음을 응답
	return c.JSON(fiber.Map{
		"keyword": req.Keyword,
	})
}

// @Summary 사용자별 키워드 로그 수집
// @Description 사용자별 키워드를 받아와서 파일에 기록합니다.
// @Tags 로그
// @Accept  json
// @Produce  json
// @Param  memberId  path  string  true  "사용자 아이디(사용자 유니크 값)"
// @Param  keyword body string true "로그로 사용할 키워드"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /log/keyword/{memberId} [post]
func HandleKeywordLogByMember(c *fiber.Ctx) error {
	// 경로에서 userId 추출
	memberId := c.Params("memberId")

	// 요청 Body에서 JSON 데이터 추출
	var req KeywordLogRequest
	if err := c.BodyParser(&req); err != nil {
		// 오류가 발생하면 400 Bad Request 응답
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// 로그 파일에 keyword 저장
	userLogger := config.SetupLogger("keywordByMember", "keywordByMember")
	config.RotateLogger(userLogger, "keywordByMember", "keywordByMember") // 자정마다 로그 파일 롤링
	log.SetOutput(userLogger)
	log.Printf("[%s] MemberId: %s, Received keyword: %s\n", config.GetSeoulTime(), memberId, req.Keyword)

	//fmt.Fprintf(logger, "[%s] Received keyword: %s\n", config.GetSeoulTime(), req.Keyword)

	// 성공적으로 처리되었음을 응답
	return c.JSON(fiber.Map{
		"memberId": memberId,
		"keyword":  req.Keyword,
	})
}
