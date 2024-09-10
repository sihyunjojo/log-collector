package logwriter

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// SaveLog 공통적인 로그 저장 함수
func SaveLog(logData, logFilePath string) error {
	// .env 파일 로드
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 로그 폴더 경로 추출 (환경 변수 LOG_FOLDER 사용)
	logFolder := os.Getenv("LOG_FOLDER")
	if logFolder == "" {
		logFolder = "./logs" // 기본 폴더 경로 설정
	}

	// 로그 폴더가 없으면 생성
	if _, err := os.Stat(logFolder); os.IsNotExist(err) {
		err := os.Mkdir(logFolder, 0755)
		if err != nil {
			return err
		}
	}

	// 로그 파일에 데이터 기록 (추가 모드)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}(f)

	// 로그 데이터를 파일에 저장
	if _, err := f.WriteString(logData); err != nil {
		return err
	}
	return nil
}
