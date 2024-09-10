package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"log/logwriter" // 로그 저장 로직을 처리하는 logwriter 패키지를 임포트
)

// LogRequest 클릭 로그 요청의 데이터를 담을 구조체
type LogRequest struct {
	Subject string `json:"subject"` // 클라이언트가 보내는 JSON 데이터의 "subject" 필드에 매핑되는 구조체 필드
}

// HandleWatchLog 클릭 로그 핸들러 함수
func HandleWatchLog(w http.ResponseWriter, r *http.Request) {
	// 요청이 POST가 아니면 오류 반환
	if r.Method != http.MethodPost {
		// HTTP 메서드가 POST가 아닌 경우 "Method not allowed" 메시지와 함께 405 상태 코드 반환
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 요청의 Body 데이터를 읽음
	body, err := ioutil.ReadAll(r.Body)
	// Body 데이터를 읽는 과정에서 오류가 발생하면 400 상태 코드와 함께 "Error reading request body" 메시지를 반환
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// JSON 데이터를 구조체로 변환
	var logRequest LogRequest
	// 요청 본문을 JSON에서 LogRequest 구조체로 변환. 변환 중 오류가 발생하면 400 상태 코드와 "Invalid JSON format" 메시지를 반환
	err = json.Unmarshal(body, &logRequest)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 로그 파일에 기록할 데이터 생성. 현재 시각과 요청에서 받은 subject 데이터를 기록
	logData := fmt.Sprintf("Timestamp: %s, Subject: %s\n", time.Now().Format(time.RFC3339), logRequest.Subject)

	// 로그 파일 경로 설정. logs 디렉터리에 click_board_log.txt 파일에 저장할 것을 지정
	logFilePath := "./watch_log.txt"

	// 로그 저장. logwriter 패키지의 SaveLog 함수를 호출하여 로그를 파일에 저장
	if err := logwriter.SaveLog(logData, logFilePath); err != nil {
		// 로그 저장 중 오류가 발생하면 500 상태 코드와 함께 "Error writing log" 메시지를 반환
		http.Error(w, "Error writing log", http.StatusInternalServerError)
		return
	}

	// 성공적으로 로그가 저장되면 200 상태 코드와 함께 "Log saved successfully" 메시지를 클라이언트에 전송
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log saved successfully"))
}
