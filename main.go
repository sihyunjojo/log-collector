package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

var port string

func main() {
    // 포트 번호 환경 변수로 관리
    port = getPort()

    // 경로 설정: /log/click/board 경로에 대한 핸들러 등록
    http.HandleFunc("/log/click/board", handleClickBoardLog)

    // 서버 시작
    fmt.Printf("Starting server on :%s...\n", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// 포트 번호를 환경 변수로부터 가져오거나 기본값(8080)을 설정
func getPort() string {
    port := os.Getenv("LOG_SERVER_PORT")
    if port == "" {
        port = "8089" // 기본 포트 번호 설정
    }
    return port
}
