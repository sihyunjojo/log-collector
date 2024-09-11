package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	// Fiber 앱 생성
	app := fiber.New()

	// 테스트할 라우팅 등록
	app.Get("/", HomeHandler)

	// 테스트 요청 생성
	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)

	// 오류가 없는지 확인
	assert.Nil(t, err)

	// 상태 코드가 200 OK 인지 확인
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	/**
	ioutil.ReadAll() 사용: 응답의 Body를 읽기 위해 ioutil.ReadAll(resp.Body)를 사용하여 응답 바디를 문자열로 변환합니다.
	assert.Equal()로 문자열 비교: 변환된 응답 바디를 기대값 "Hello, World!"와 비교합니다.
	*/
	// 응답 바디 읽기
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	// 응답 바디가 "Hello, World!" 인지 확인
	assert.Equal(t, "Hello, World!", string(body))
}

// 이건 왜 안됫지?

// 응답 바디가 "Hello, World!" 인지 확인
//body := httptest.NewRecorder()
//app.Test(req, -1)
//
//assert.Equal(t, "Hello, World!", body.Body.String())
