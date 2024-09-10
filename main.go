package main

import (
	"fmt"
	"log"
	"log/logwriter"
	"net/http"
	"os"

	"github.com/joho/godotenv" // godotenv 패키지
	"log/errors"               // 패키지 경로 수정
	"log/handlers"             // 패키지 경로 수정
)

var port string

func main() {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 포트 번호 환경 변수로 관리
	port, err = getPort()
	if err != nil {
		log.Fatal(err) // 에러가 발생하면 프로그램을 중단하고 로그 출력
	}

	// 자정마다 로그 파일을 변경하는 작업을 백그라운드에서 실행
	go logwriter.ScheduleLogRotation()

	// 경로 설정: /log/watch 경로에 대한 핸들러 등록
	http.HandleFunc("/log/watch", handlers.HandleWatchLog) // 패키지 이름을 명시적으로 사용

	// 서버 시작
	fmt.Printf("Starting server on :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// getPort 함수: 환경 변수가 없으면 커스텀 에러 발생
func getPort() (string, error) {
	port := os.Getenv("LOG_SERVER_PORT")
	if port == "" {
		return "", customerrors.NewPortNotFoundError("LOG_SERVER_PORT")
	}
	return port, nil
}
