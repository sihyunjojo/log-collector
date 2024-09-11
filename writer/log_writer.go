package writer

import (
	"github.com/joho/godotenv" // .env 파일을 로드하는 라이브러리
	"log"                      // 기본 로깅 기능을 제공하는 패키지
	"os"                       // 파일 시스템 및 환경 변수 관리를 위한 패키지
	"path/filepath"            // 파일 경로를 처리하는 패키지
	"sync"                     // 동시성 처리를 위한 뮤텍스 패키지
	"time"                     // 시간 처리 및 시간대 설정을 위한 패키지
)

// 파일 접근 동기화를 위한 뮤텍스
var fileMutex sync.Mutex // 여러 고루틴에서 파일에 동시 접근하는 것을 막기 위한 뮤텍스

// MoveLogAtMidnight 자정에 로그 파일 이름을 변경하는 함수
func MoveLogAtMidnight(logFileName string) error {
	fileMutex.Lock()         // 파일에 접근할 때 동시성 문제를 방지하기 위해 뮤텍스 잠금
	defer fileMutex.Unlock() // 함수가 끝나면 잠금을 해제하여 다른 작업이 접근 가능하도록 함

	// .env 파일 로드
	// .env 파일을 로드하여 환경 변수를 사용 가능하게 함. 실패하면 프로그램 종료
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 한국 시간대(Asia/Seoul)를 설정
	// 한국 시간을 기준으로 로그 파일 이름을 변경하기 위해 시간대를 설정
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatal("Error loading time zone")
	}

	// 현재 한국 시간 가져오기
	// 현재 시간을 한국 시간으로 가져옴
	now := time.Now().In(location)

	// 어제 날짜 계산
	// 어제 날짜를 구하기 위해 현재 시간에서 하루를 뺌
	yesterday := now.AddDate(0, 0, -1)
	// 어제 날짜를 yyyy-mm-dd 형식으로 변환하여 파일 이름에 추가
	yesterdayFileName := yesterday.Format("2006-01-02") + "_" + logFileName

	// 로그 폴더 경로 추출 (환경 변수 LOG_FOLDER 사용)
	// 환경 변수 LOG_FOLDER에서 로그 폴더 경로를 가져옴. 없으면 기본값 "./logs" 사용
	logFolder := os.Getenv("LOG_FOLDER")
	if logFolder == "" {
		logFolder = "./logs" // 기본 폴더 경로 설정
	}

	// logFolder와 logFileName을 결합하여 전체 파일 경로 생성
	// 현재 로그 파일 경로와 어제 날짜로 변경될 파일 경로 생성
	currentLogFilePath := filepath.Join(logFolder, logFileName)
	yesterdayLogFilePath := filepath.Join(logFolder, yesterdayFileName)

	// 현재 로그 파일이 존재하는지 확인하고, 존재하면 어제 날짜로 파일 이름 변경
	// 현재 로그 파일이 있는 경우, 어제 날짜가 포함된 파일 이름으로 변경
	if _, err := os.Stat(currentLogFilePath); err == nil {
		// 파일 이름을 어제 날짜로 변경
		err := os.Rename(currentLogFilePath, yesterdayLogFilePath)
		if err != nil {
			return err // 이름 변경 중 오류가 발생하면 해당 오류를 반환
		}
	}

	// 새로운 로그 파일(watch_log.txt)을 생성
	// 파일이 없으면 새로 생성하고, 있으면 기존 파일에 덧붙임
	f, err := os.OpenFile(currentLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err // 파일 생성 또는 열기 실패 시 에러 반환
	}
	defer f.Close() // 파일 사용 후 반드시 닫음

	return nil // 모든 작업이 성공적으로 완료되면 nil 반환
}

// SaveLog 공통적인 로그 저장 함수
func SaveLog(logData, logFileName string) error {
	fileMutex.Lock()         // 파일에 접근할 때 동시성 문제를 방지하기 위해 뮤텍스 잠금
	defer fileMutex.Unlock() // 함수가 끝나면 잠금을 해제하여 다른 작업이 접근 가능하도록 함

	// 로그 폴더 경로 추출 (환경 변수 LOG_FOLDER 사용)
	// 환경 변수에서 로그 폴더 경로를 가져옴. 설정되지 않은 경우 기본값 "./logs" 사용
	logFolder := os.Getenv("LOG_FOLDER")
	if logFolder == "" {
		logFolder = "./logs" // 기본 폴더 경로 설정
	}

	// logFolder와 logFileName을 결합하여 전체 파일 경로 생성
	// 로그를 저장할 파일의 전체 경로를 생성
	logFilePath := filepath.Join(logFolder, logFileName)

	// 로그 파일에 데이터 기록 (추가 모드)
	// 파일을 열어 로그 데이터를 추가. 파일이 없으면 새로 생성
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err // 파일 열기 또는 생성 중 에러가 발생하면 반환
	}
	defer func(f *os.File) {
		err := f.Close() // 파일 사용 후 닫는 작업 수행
		if err != nil {
			log.Printf("Error closing file: %v", err) // 파일 닫기 중 에러가 발생하면 로그로 출력
		}
	}(f)

	// 로그 데이터를 파일에 저장
	// 로그 데이터를 파일에 추가. 에러가 발생하면 해당 에러를 반환
	if _, err := f.WriteString(logData); err != nil {
		return err // 로그 기록 중 오류 발생 시 반환
	}
	return nil // 성공적으로 로그가 기록되면 nil 반환
}
