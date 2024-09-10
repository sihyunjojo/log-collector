package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "./logwriter" 
)

// 클릭 로그 요청의 데이터를 담을 구조체
type LogRequest struct {
    Subject string `json:"subject"`
}

// 클릭 로그 핸들러
func handleClickBoardLog(w http.ResponseWriter, r *http.Request) {
    // 요청이 POST가 아니면 오류 반환
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 요청의 Body에서 데이터를 읽음
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusBadRequest)
        return
    }

    // JSON 데이터를 구조체로 변환
    var logRequest LogRequest
    err = json.Unmarshal(body, &logRequest)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    // 로그 파일에 기록할 데이터 생성
    logData := fmt.Sprintf("Timestamp: %s, Subject: %s\n", time.Now().Format(time.RFC3339), logRequest.Subject)

    // 로그 파일 경로 설정
    logFilePath := fmt.Sprintf("%s/click_board_log.txt", getLogFolder())

    // 로그 저장
    if err := logwriter.SaveLog(logData, logFilePath); err != nil {
        http.Error(w, "Error writing log", http.StatusInternalServerError)
        return
    }

    // 응답 전송
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Log saved successfully"))
}
