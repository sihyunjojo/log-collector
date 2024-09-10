package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/joho/godotenv" // godotenv 패키지
)

var port string

func main() {
    // .env 파일 로드
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 포트 번호 환경 변수로 관리
    port = getPort()

    // 경로 설정: /log/click/board 경로에 대한 핸들러 등록
    http.HandleFunc("/log/click/board", handleClickBoardLog)

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
